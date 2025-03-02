package api

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	db "github.com/mjthecoder65/image-processing-service/db/sqlc"
	"github.com/mjthecoder65/image-processing-service/pkg/utils"
)

type CreateUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type userResponse struct {
	ID        pgtype.UUID        `json:"id"`
	Email     string             `json:"email"`
	CreatedAt pgtype.Timestamptz `json:"createdAt"`
	UpdatedAt pgtype.Timestamptz `json:"updatedAt"`
}

func NewUserResponse(user db.User) userResponse {
	return userResponse{
		ID:        user.ID,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req CreateUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	passwordHash, err := utils.HashPassword(req.Password)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Email:        req.Email,
		PasswordHash: passwordHash,
	}

	user, err := server.queries.CreateUser(context.Background(), arg)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, NewUserResponse(user))
}

type loginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type logingUserResponse struct {
	AccessToken string       `json:"accessToken"`
	User        userResponse `json:"user"`
}

func (server *Server) login(ctx *gin.Context) {
	var req loginUserRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.queries.GetUserByEmail(context.Background(), req.Email)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.VerifyPassword(user.PasswordHash, req.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong email or Password"})
		return
	}

	token, err := server.maker.CreateToken(user.ID, 2*time.Minute) // TODO: Move token expiration duration to the configuration.

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, logingUserResponse{
		AccessToken: token,
		User:        NewUserResponse(user),
	})

}
