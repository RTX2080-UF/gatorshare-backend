### Video:
#### Demo:
[![Sprint 2 Demo](https://img.youtube.com/vi/Uqa2b2MekAw/0.jpg)]()


#### Postman
[![Sprint 2 Postman](https://img.youtube.com/vi/QqsWVejhuFI/0.jpg)]()
# gatorshare-backend
Backend Code of GatorShare, Software Engineering Spring 2022 project.
Built by:
- Anuj Singh (Backend) {R3dI0}
- Dinesh Valasala (Frontend) {valasaladinesh}
- Ekleen Kaur (Backend) {eclairss17}
- Rishabh Tatiraju (Frontend) {tatirajurishabh}

Stack:
- Frontend: React.js
- Backend: GoLang
- Database: Postgres and sqlite

## Sprint 2 Deliverables

Sprint2 has functional testing and unit testing for the backend. Test cases for both postman and go have been included inside the scope of sprint2. 

## Issues Resolved
- Create CI/CD pipelines for devops.
- Migrate database to postgres from Sqlite
- Fix issue with foreign key dependency between tables
- Add Support for using environment variables 
- Add postman collection for relevant API requests.
- Create test for Posts, Comments and Registration. 
- Fix the backend bug for fetch post comments
- Add Documentation for Sprint 2

## Api Endpoints features
GET- http://localhost:8080/ <br>
Json Response received- <br>
    "data": "Welcome to Gatorshare made with the help of Go and Gin!"

## Api Details (Documentation)
- [User Api Endpoints](./documentation/User_api.md)
- [Post Api Endpoints](./documentation/Posts_api.md)
- [Comment Api Endpoints](./documentation/Comment_api.md)

# Api endpoints delivered in this sprint 

| Type | Api Endpoint | Description | Status |
| ----------- | ----------- | ----------- | ----------- |
| Patch | `/v1/posts/update/:id`   | Update an existing post | Active |
| Patch |  `/v1/comments/update/:id` | Update an existing comment | Active |
| POST | `/v1/users/register` | Register new users | Active |
| POST | `/v1/users/login` | Authenticate and create user session | Active |
| GET | `/v1/users/getProfile/:id` | Get user profile by Id | Active |
| DELETE | `/v1/users/deleteProfile/:id` | Delete user and associated resource | Active |
| PATCH | `/v1/users/updateProfile/:id` | Update user details | Active |
