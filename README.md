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

- **Go**: Backend implementation
- **Gin**: Web framework
- **PostgreSQL**: Database
- **SQLC**: Database query management
- **AWS S3**: Cloud storage for images
- **Docker & docker-compose**: Containerization
- **Imaging Library**: Image processing

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

- **Register**: `POST /register`
- **Login**: `POST /login`

### Image Management

- **Upload Image**: `POST /images`
- **Transform Image**: `POST /images/:id/transform`
- **Retrieve Image**: `GET /images/:id`
- **List Images**: `GET /images?page=1&limit=10`
