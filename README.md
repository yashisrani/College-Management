# College Management System
A simple and efficient system to manage students, courses, faculties, and departments. Supports CRUD operations, authentication, and role-based access control. Designed for scalability and easy integration.​

## Features 🪄🌟
- **Student Management**: Add, update, view, and delete student records.

- **Course Management**: Manage course details and enrollments.

- **Faculty Management**: Handle faculty profiles and assignments.

- **Department Management**: Organize departments and associated data.

- **Authentication**: Secure login system with JWT-based authentication.

- **Role-Based Access Control**: Assign roles and permissions to users.

- **API Documentation**: Integrated Swagger UI for API exploration.

- **Monitoring**: Prometheus metrics and Grafana dashboards for system monitoring

## Technologies Used 👨🏻‍💻
**Backend**: Go (Golang) with Gin framework

**Database**: PostgreSQL

**Authentication**: JWT (JSON Web Tokens)

**API Documentation**: Swagger UI

**Monitoring**: Prometheus and Grafana

**Containerization**: Docker and Docker Compose

**CI/CD**: GitHub Actions​

## ScreenShots

<img width="1440" alt="Screen Shot 2025-05-01 at 12 57 43 PM" src="https://github.com/user-attachments/assets/f3ca557e-c3a3-4727-a75e-57592c5b2d37" />
<img width="1440" alt="Screen Shot 2025-05-01 at 12 57 49 PM" src="https://github.com/user-attachments/assets/387b2bab-cab6-4e8e-a4e8-8e68832fe9fa" />
<img width="1440" alt="Screen Shot 2025-05-01 at 12 57 59 PM" src="https://github.com/user-attachments/assets/afda27b9-b583-4efa-9622-17d0db687e58" />


<img width="1440" alt="Screen Shot 2025-03-15 at 10 22 04 PM" src="https://github.com/user-attachments/assets/28cdb19d-8d71-41ba-85ec-9a627c8fa605" />

<img width="1434" alt="Screen Shot 2025-03-16 at 2 24 48 PM" src="https://github.com/user-attachments/assets/3a0c9232-37d7-4cc5-a596-6bd49b89697d" />
<img width="1440" alt="Screen Shot 2025-03-15 at 10 18 48 PM" src="https://github.com/user-attachments/assets/5cec9644-f157-40c3-8b5a-5b9067c335b6" />

## Folder Structure
`api/` - API route definitions

`controllers/` - Request handlers and business logic

`model/` - Database models and schemas

`store/` - Database interactions and queries

`middleware/` - Custom middleware functions

`utils/` - Utility functions and helpers

`deployment/` - Deployment configurations

`docs/`- API documentation and related files

## Getting Started
### Prerequisites
- Go installed on your machine

- Docker and Docker Compose installed

- PostgreSQL database setup​ (Use Postgres Docker Image)

#### Installation
1) Clone the repository:

```
git clone https://github.com/yashisrani/College-Management.git

cd College-Management

go run cmd/main.go
```

2) Set up environment variables:

- Create a .env file in the root directory and configure the necessary environment variables as per your setup.

3) Run the application using Docker Compose:

```
docker-compose up --build
```

4) Access the application:

- API endpoints will be available at `http://localhost:8080/api`

- Swagger UI can be accessed at` http://localhost:8080/swagger/index.html`

- Grafana dashboards (if configured) will be available at `http://localhost:3000​`

