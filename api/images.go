package api

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/color"
	"log"
	"net/http"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/mjthecoder65/image-processing-service/db/sqlc"
	"github.com/mjthecoder65/image-processing-service/pkg/token"
	"github.com/mjthecoder65/image-processing-service/pkg/utils"
)

const (
	MAX_FILE_SIZE = 10 << 20
)

var SUPPORTED_IMAGE_TYPES = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

func (server *Server) uploadImage(ctx *gin.Context) {
	claims, exists := ctx.Get("auth_payload")

	if !exists {
		err := errors.New("user not authenticated")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}
	log.Println(claims, exists)

	file, err := ctx.FormFile("image")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if file.Size > MAX_FILE_SIZE {
		err := errors.New("file size exeeds 10MB")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ext := filepath.Ext(file.Filename)

	if !SUPPORTED_IMAGE_TYPES[ext] {
		err := errors.New("only JGEP and PNG files are allowed")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	src, err := file.Open()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	defer src.Close()

	fileName := file.Filename
	imageURL, err := server.storageClient.UploadFile(src, fileName)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateImageParams{
		Name:   fileName,
		UserID: claims.(*token.Claims).UserID,
		Url:    imageURL,
	}

	image, err := server.queries.CreateImage(context.Background(), arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, image)
}

type getImageRequest struct {
	ID pgtype.UUID `uri:"id" binding:"required"`
}

func (server *Server) getImage(ctx *gin.Context) {
	var req getImageRequest

	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	image, err := server.queries.GetImage(context.Background(), req.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, image)
}

func (server *Server) listImages(ctx *gin.Context) {
	claims, exists := ctx.Get("auth_payload")

	if !exists {
		err := errors.New("user unauthenticated ")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.GetUserImagesParams{
		UserID: claims.(*token.Claims).UserID,
		Offset: 0,
		Limit:  10,
	}

	images, err := server.queries.GetUserImages(context.Background(), arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, images)
}

type TransformRequest struct {
	Transformations struct {
		Resize *struct {
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"resize"`
		Crop *struct {
			Width  int `json:"width"`
			Height int `json:"height"`
			X      int `json:"x"`
			Y      int `json:"y"`
		} `json:"crop"`
		Rotate  *float64 `json:"rotate"`
		Format  *string  `json:"format"`
		Filters *struct {
			Grayscale bool `json:"grayscale"`
			Sepia     bool `json:"sepia"`
		} `json:"filters"`
	} `json:"transformations"`
}

func (server *Server) transformImage(ctx *gin.Context) {

	claims, exists := ctx.Get("auth_payload")

	if !exists {
		err := errors.New("user not authenticated")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	imageID := ctx.Param("id")

	if len(imageID) == 0 {
		err := errors.New("image id is required")
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	var req TransformRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	imageUUID, err := StringToUUID(imageID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	imageMetadata, err := server.queries.GetImage(context.Background(), imageUUID)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if imageMetadata.UserID != claims.(*token.Claims).UserID {
		err := errors.New("unauthorized to transform this image")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	// Download original image from s3 bucket.
	imageData, err := server.storageClient.GetImage(imageMetadata.Name)

	if err != nil {
		err := errors.New("failed to retrieve image")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	img, err := imaging.Decode(bytes.NewReader(imageData))

	if err != nil {
		err := errors.New("failed to decode downloaded image")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if req.Transformations.Resize != nil {
		img = imaging.Resize(img, req.Transformations.Resize.Width, req.Transformations.Resize.Height, imaging.Lanczos)
	}

	if req.Transformations.Crop != nil {
		rect := image.Rect(
			req.Transformations.Crop.X, req.Transformations.Crop.Y,
			req.Transformations.Crop.X+req.Transformations.Crop.Width,
			req.Transformations.Crop.Y+req.Transformations.Crop.Height,
		)

		img = imaging.Crop(img, rect)
	}

	if req.Transformations.Rotate != nil {
		img = imaging.Rotate(img, *req.Transformations.Rotate, nil)
	}

	if req.Transformations.Filters != nil {
		if req.Transformations.Filters.Grayscale {
			img = imaging.Grayscale(img)
		}

		if req.Transformations.Filters.Sepia {
			img = ApplySepia(img)
		}
	}

	format := imaging.JPEG

	if req.Transformations.Format != nil {
		switch *req.Transformations.Format {
		case "jpeg", "jpg":
			format = imaging.JPEG
		case "png":
			format = imaging.PNG
		default:
			err := errors.New("unsupported format")
			ctx.JSON(http.StatusBadRequest, errorResponse(err))
			return
		}
	}

	buf := new(bytes.Buffer)

	if err := imaging.Encode(buf, img, format); err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	filename := imageMetadata.Name

	newFilename := fmt.Sprintf("%s_transformed_%s%s",
		filename[:len(filename)-len(filepath.Ext(filename))],
		utils.RandomString(6), filepath.Ext(filename))

	transformedURL, err := server.storageClient.UploadFile(bytes.NewReader(buf.Bytes()), newFilename)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	transformation, err := json.Marshal(req.Transformations)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateImageTransformationParams{
		ImageID:        imageMetadata.ID,
		Url:            transformedURL,
		Transformation: transformation,
	}

	imageTransformation, err := server.queries.CreateImageTransformation(context.Background(), arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, imageTransformation)

}

func (server *Server) generateImage(ctx *gin.Context) {
	// TODO: use generative AI to generate images based on the prompt and save the image to s3 bucket.
}

func StringToUUID(value string) (pgtype.UUID, error) {
	u, err := uuid.Parse(value)

	if err != nil {
		return pgtype.UUID{}, nil
	}

	return pgtype.UUID{
		Bytes: u,
		Valid: true,
	}, nil
}

func ApplySepia(img image.Image) image.Image {
	bounds := img.Bounds()
	result := image.NewNRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()

			r8, g8, b8 := float64(r>>8), float64(g>>8), float64(b>>8)

			tr := 0.393*r8 + 0.769*g8 + 0.189*b8
			tg := 0.349*r8 + 0.686*g8 + 0.168*b8
			tb := 0.272*r8 + 0.534*g8 + 0.131*b8

			if tr > 255 {
				tr = 255
			}
			if tg > 255 {
				tg = 255
			}
			if tb > 255 {
				tb = 255
			}

			result.SetNRGBA(x, y, color.NRGBA{
				R: uint8(tr),
				G: uint8(tg),
				B: uint8(tb),
				A: uint8(a >> 8),
			})
		}
	}

	return result
}
