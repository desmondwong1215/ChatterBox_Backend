# ChatterBox Backend

ChatterBox Backend is the server-side application for the ChatterBox project, a feature-rich forum platform. This repository provides the backend APIs and handles business logic, database interactions, and user authentication. It is built using **Go** and utilizes **GORM** for database management with **PostgreSQL**.

## Table of Contents
- [Features](#features)
- [Technologies Used](#technology-used)
- [Getting Started](#getting-started)
- [Running Locally](#running-locally)
- [API Endpoints](#api-endpoints)
- [Deployment](#deployment)
- [Troubleshooting](#troubleshooting)
- [Contributing](#contributing)
- [Contact](#contact)

## Features
- **User Authentication**: Secure JWT-based authentication.
- **CRUD Operations**: Comprehensive API for managing forums, comments, like, as well as editing and deletion of the forums and comments.
- **Database Integration**: PostgreSQL powered by GORM ORM.
- **CORS Configuration**: Supports cross-origin requests for seamless frontend-backend communication.
- **Error Handling**: Structured responses for client errors and internal issues.

## Technologies Used
- **Go**: Backend programming language.
- **GORM**: ORM for database interactions.
- **PostgreSQL**: Relational database system.
- **JWT**: JSON Web Token for user authentication.
- **Fiber**: Web framework for building APIs.

## Getting Started
This website is live at https://chatterbox-backend-mkjs.onrender.com/.

### Prerequisites
Before running the project, ensure you have the following installed:
- [Go](https://go.dev/dl/): Version 1.19 or higher.
- [PostgreSQL](https://www.postgresql.org/download/): Version 13 or higher.
- [Git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git): Version control system.

### Installation
1. Clone the repository:
   ```bash
   git clone https://github.com/desmondwong1215/ChatterBox-Backend.git
   ```

2. Navigate to the project directory:
   ```bash
   cd ChatterBox-Backend
   ```

3. Install dependencies:
   ```bash
   go mod tidy 
   ```

## Running locally
### Configure Environment Variables
Create `.env` file and add following environment variable
    ```bash
    DB_HOST=your-database-host (should be `localhost`)
    DB_PORT=your-database-port (should be `5432`)
    DB_USER=your-database-username
    DB_PASSWORD=your-database-password
    DB_NAME=your-database-name
    DB_SSLMODE=disable
    PORT=9090
    KEY=your-jwt-key
    ```
Replace placeholders with your actual database configuration and JWT secret. Go to `/storage/postgres.go`, uncomment line 24-28 and comment out line 31-34.

### Start the Server
Start the application:
    ```bash
    go run main.go
    ```
The backend server will run at http://localhost:9090

## API Endpoints
Below is the summary of key API endpoints:

### Authentication
- `POST /login`: Authenticate the user and create a access token (last for 15 minutes, stored in local storage) and a return token (last for 30 days, stored in cookies).

### Forum and Comment
- `GET /forum/:username/:keywords`: Return the forum according to the keywords from the user.
- `POST /forum/:username`: Create new post or comment
- `PUT /forum/:username`: Edit existing post or comment
- `PATCH /forum/:username/:id`: Handle the like of the post or comment according to the id.
- `DELETE /forum/:username/:type/:id`: Delete the post or comment according to the type and id.

## Deployment
### Deploy on Render
1. Log in to [Render](https://render.com/).
2. Create a new Web Service.
3. Connect your Github repository to Render.
4. Set the following settings:
    - **Build Command**: `go build -o main .`
    - **Start Command**: `./main`
5. Add environment variables in the Render dashboard.
6. Deploy the service.

## Troubleshooting
1. CORS Issues
Ensure that the backend CORS middleware is configured to allow requests from your frontend domain:
    ```bash
    app.Use(cors.New(cors.Config{
        AllowOrigins: "https://your-frontend-domain",
        AllowMethods: "GET,POST,PUT,DELETE",
        AllowHeaders: "Content-Type,Authorization",
    }))
    ```

2. Database Connection Issues
    - Verify that the database credentials in your `.env` file are correct.
    - Ensure the PostgreSQL server is running and accessible.

3. JWT Authentication Errors
Ensure the `KEY` in your `.env` matches the secret used to generate tokens.

## Contributing
Contributions are welcome! If you have ideas for improvements or encounter any issues, feel free to open an issue or submit a pull request.

### Steps to Contribute
1. Fork the repository.
2. Create a new branch:
   ```bash
   git checkout -b your-feature-name
   ```
3. Make your changes and commit them:
   ```bash
   git commit -m "Add your feature"
   ```
4. Push the branch to your fork:
   ```bash
   git push origin your-feature-name
   ```
5. Open a pull request to the `main` branch.

## Contact
1. For any inquiries or support, please contact:
   - **Name**: Desmond Wong Hui Sheng
   - **Email**: e1356940@u.nus.edu
   - **GitHub**: https://github.com/desmondwong1215

---
Feel free to explore, contribute, and build something amazing with ChatterBox!