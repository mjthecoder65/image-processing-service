# Image Processing Service

This project is a backend service for image processing, similar to Cloudinary. It allows users to upload images, apply transformations, and retrieve images in different formats. The system includes user authentication, image management, and efficient retrieval mechanisms.

## Features

### User Authentication

- **Sign-Up**: Users can create an account.
- **Log-In**: Users can log in.
- **JWT Authentication**: Secure endpoints using JWTs for authenticated access.

### Image Management

- **Upload Image**: Users can upload images.
- **Transform Image**: Perform transformations such as resize, crop, rotate, watermark, etc.
- **Retrieve Image**: Fetch saved images in different formats.
- **List Images**: View all uploaded images with metadata.

### Image Transformations

- Resize
- Crop
- Rotate
- Watermark
- Flip
- Mirror
- Compress
- Change format (JPEG, PNG, etc.)
- Apply filters (grayscale, sepia, etc.)

## Tech Stack

- **Language**: Go
- **Framework**: Gin (HTTP server)
- **Database**: PostgreSQL with SQLC
- **Storage**: AWS S3
- **Image Processing**: `github.com/disintegration/imaging`
- **Authentication**: JWT (JSON Web Tokens)
- **Deployment**: Docker with `docker-compose`

## Setup Instructions

1. Clone the repository:

   ```sh
   git clone https://github.com/your-repo/image-processing-service.git
   cd image-processing-service
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Set up environment variables:
   Create a `.env` file and configure database, AWS credentials, and JWT settings.

4. Run database migrations:

   ```sh
   make migrate-up
   ```

5. Start the service:
   ```sh
   docker-compose up --build
   ```

## API Endpoints

### Authentication

- **Register**: `POST /api/v1/auth/register`
- **Login**: `POST /api/v1/auth/login`

### Image Management

- **Upload Image**: `POST /api/v1/images`
- **Transform Image**: `POST /api/v1/images/:id/transform`
- **Retrieve Image**: `GET /api/v1/images/:id`
- **List Images**: `GET /api/v1/images?page=1&limit=10`

## API Usage with cURL

```sh
#!/bin/bash

# Register
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email": "michael53161@gmail.com.com", "password": "coolhand"}'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email": "michael53161@gmail.com.com", "password": "coolhand"}'

ACCESS_TOKEN="your_jwt_token_here"

# Upload Image
curl -X POST http://localhost:8080/api/v1/images/ \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -F "image=@/Users/michael/Desktop/sky.jpg"

# Get Image
curl -X GET http://localhost:8080/api/v1/images/a80b97f6-c411-4093-adfa-e84682341e62 \
    -H "Authorization: Bearer $ACCESS_TOKEN"

# Get All Images
curl -X GET http://localhost:8080/api/v1/images/ \
    -H "Authorization: Bearer $ACCESS_TOKEN"

# Test Image Rotate
curl -X POST http://localhost:8080/api/v1/images/a80b97f6-c411-4093-adfa-e84682341e62/transform \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"transformations": {"rotate": 90}}'

# Test Grayscale
curl -X POST http://localhost:8080/api/v1/images/a80b97f6-c411-4093-adfa-e84682341e62/transform \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"transformations": {"filters": {"grayscale": true}}}'

# Test Sepia
curl -X POST http://localhost:8080/api/v1/images/a80b97f6-c411-4093-adfa-e84682341e62/transform \
  -H "Authorization: Bearer $ACCESS_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"transformations": {"filters": {"sepia": true}}}'
```
