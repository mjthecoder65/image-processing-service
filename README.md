# Image Processing Service

A scalable backend service built with Go for uploading, transforming, and managing images, inspired by services like Cloudinary. Features include user authentication, image storage on AWS S3, and transformations such as resize, crop, rotate, grayscale, and sepia.

## Features

- **User Authentication**: Register and log in with JWT-based authentication.
- **Image Management**:
  - Upload images via `POST /api/v1/images`.
  - Retrieve images by ID or list all user images.
- **Image Transformations**: Apply transformations (resize, crop, rotate, grayscale, sepia) via `POST /api/v1/images/:id/transform`.
- **Storage**: Images stored in AWS S3 with database metadata in PostgreSQL.
- **Database**: PostgreSQL with SQLC for type-safe queries.

## Tech Stack

- **Language**: Go
- **Framework**: Gin (HTTP server)
- **Database**: PostgreSQL with SQLC
- **Storage**: AWS S3
- **Image Processing**: `github.com/disintegration/imaging`
- **Authentication**: JWT (JSON Web Tokens)
- **Deployment**: Docker with `docker-compose`

## Prerequisites

- Go 1.23+
- PostgreSQL 15+
- AWS S3 bucket and credentials
- Docker (optional, for containerized deployment)

## Setup

1. **Clone the Repository**:
   ```bash
   git clone https://github.com/yourusername/image-processing-service.git
   cd image-processing-service
   ```
