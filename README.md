# Halted
Pass the given submission time by Rakamin (1 month), but I still work on this anyway as side-project
## 🔄 Project Status

**Current completion: ~50%**

✅ Completed:
- Basic authentication
- Product management
- Regional integration
- File upload structure

🚧 In Progress:
- Store management
- Category system
- Transaction processing
- Product logging

# Evermos E-commerce API

A RESTful API service built with Go for Evermos, a Muslim-focused social commerce platform. This project implements clean architecture principles and provides essential e-commerce functionalities.

## 🌟 Features

### Core Services
- **Authentication**
  - User registration with automatic store creation
  - JWT-based authentication
  - Profile management
  
- **Product Management**
  - Product CRUD operations
  - Advanced filtering and pagination
  - Category-based organization
  - Price range filtering

- **Location Services**
  - Integration with Indonesian Regional API
  - Province and city data management
  - Address validation

- **File Management**
  - Product image upload
  - User avatar management

## 🏗 Architecture

The project follows Clean Architecture principles with the following layers:
```go
├── domain          # Enterprise business rules
├── repository      # Database operations
├── usecase         # Application business rules
└── delivery        # External interfaces (HTTP handlers)
```

## 🛠 Tech Stack

- **Backend:** Go 1.19+
- **Database:** MySQL
- **Authentication:** JWT
- **API Documentation:** Swagger/OpenAPI
- **Development Tools:** Git, Make

## 📋 Prerequisites

- Go 1.19 or higher
- MySQL 8.0+
- Make (optional, for using Makefile commands)

## 🚀 Getting Started

1. Clone the repository
```bash
git clone https://github.com/yourusername/evermos-api.git
```

2. Set up environment variables
```bash
cp .env.example .env
# Edit .env with your configurations
```

3. Install dependencies
```bash
go mod download
```

4. Run migrations
```bash
make migrate
```

5. Start the server
```bash
make run
```

## 📝 API Documentation

API documentation is available at `/swagger/index.html` when running the server.

Key endpoints:
- `POST /auth/register` - User registration
- `POST /auth/login` - User authentication
- `GET /products` - List products with filtering
- `POST /products` - Create new product
- `GET /regions` - Get regional data
