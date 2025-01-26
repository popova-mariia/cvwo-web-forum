# CVWO Web Forum
A web-based movie forum application built with React and Go.

## Features

- User authentication with JWT
- Create, view, and manage threads
- Post and reply to comments
- Responsive design

## Technologies Used
Frontend: React, Axios, Material-UI
Backend: Go, Gin, JWT
Database: PostgreSQL

## Setup

You need to have the following installed:
- Node.js and npm
- Go
- PostgreSQL

1. Clone the repository:

   ```bash
   git clone https://github.com/popova-mariia/cvwo-web-forum.git
   cd cvwo-web-forum
   ```

2. Set up the backend and run the server:

   - Navigate to the backend directory:
     ```bash
     cd backend
     ```
   - Install Go dependencies:
     ```bash
     go mod tidy
     ```
   - Set up your PostgreSQL database and update the connection settings in `db.go`.

   - Run the server:
   ```bash
   cd server
   go run main.go
   ```

4. Set up the frontend and run the development server:

   - Navigate to the frontend from root directory:
     ```bash
     cd frontend
     ```
   - Install npm dependencies:
     ```bash
     npm install
     ```
   - Run the development server:
   ```bash
   npm run dev
   ```

6. Access the application:

   - Open your browser and go to `http://localhost:5173`.

## Usage

- Register or log in to start creating and participating in threads.
- Use the navigation bar to explore different sections of the forum.
