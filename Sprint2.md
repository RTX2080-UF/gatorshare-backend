### Video:
#### Demo:
[![Sprint 2 Demo](https://img.youtube.com/vi/Uqa2b2MekAw/0.jpg)]()


#### Postman
[![Sprint 2 Postman](https://img.youtube.com/vi/QqsWVejhuFI/0.jpg)]()
# gatorshare-backend
Backend Code of GatorShare, Software Engineering Spring 2022 project.
Built by:
- Anuj Singh (Backend)
- Dinesh Valasala (Frontend)
- Ekleen Kaur (Backend)
- Rishabh Tatiraju (Frontend)

Stack:
- Frontend: React.js
- Backend: GoLang
- Database: TBD

# GatorShare

College life is difficult, and expensive (unless you are Bill Gates' child). So, obviously, you'd want to save some money - let's say by sharing that cab ride to the airport (why would you even go there, though?), or maybe buy a large pizza and split the bill (and hopefully the pizza too) with different people. Whatever you want to do, GatorShare has you covered. Find people who want to share or split stuff with you and save money in the process. Wait, but its not just that, check out all that we offer below:

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
GET- http://localhost:8080/
Json Response received- 
    "data": "Welcome to Gatorshare made with the help of Go and Gin!"
## User Posts
By ID
-GET- http://localhost:8080/v1/posts/getOne/2
-Json Response received if the record exists:
-{"data":{"ID":2,"CreatedAt":"2022-02-22T03:56:47.6075532+05:30","UpdatedAt":"2022-02-22T03:56:47.6075532+05:30","DeletedAt":null,"userId":1,"title":"Test post 2","description":"My second post","userLimit":4,"participants":"1","expiry":24,"viewCount":0,"status":2,"Categories":"","Tags":""}}
-if the record doesn't exist:
-{"error":"record not found"}
-All Posts
-GET- http://localhost:8080/v1/posts/getAll/1
-Json Received if the record exists:
-{"data":[{"ID":2,"CreatedAt":"2022-02-22T03:56:47.6075532+05:30","UpdatedAt":"2022-02-22T03:56:47.6075532+05:30","DeletedAt":null,"userId":1,"title":"Test post 2","description":"My second post","userLimit":4,"participants":"1","expiry":24,"viewCount":0,"status":2,"Categories":"","Tags":""},{"ID":3,"CreatedAt":"2022-02-22T04:29:47.2209425+05:30","UpdatedAt":"2022-02-22T04:29:47.2209425+05:30","DeletedAt":null,"userId":1,"title":"Test post 2","description":"My second post","userLimit":4,"participants":"1","expiry":24,"viewCount":0,"status":2,"Categories":"","Tags":""},{"ID":4,"CreatedAt":"2022-02-22T04:30:51.0203078+05:30","UpdatedAt":"2022-02-22T04:30:51.0203078+05:30","DeletedAt":null,"userId":1,"title":"Test post 2","description":"Testing Post Creation","userLimit":4,"participants":"1","expiry":24,"viewCount":0,"status":2,"Categories":"","Tags":""}]}
-If the record doesn't exist:
-{"error":"record not found"}
-Create User Post 
-POST- http://localhost:8080/v1/posts/create
-Json Body Sent-
-{
    "userId" : 1,
    "title" : "Test post 2",
    "description" : "Testing Post Creation",
    "userLimit" : 4,
    "status" : 2
}
-Json response received- 
-{
    "data": 8
}
-DELETE- http://localhost:8080/v1/posts/delete/1
-Json Response Received- 
-{
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "userId": 0,
        "title": "",
        "description": "",
        "userLimit": 0,
        "participants": "",
        "expiry": 0,
        "viewCount": 0,
        "status": 0,
        "Categories": "",
        "Tags": ""
    }
}
-If the record is already deleted- 
-{
    "error": "record not found"
}

## User Comments
-By ID
-GET- http://localhost:8080/v1/comments/getOne/2
-Json Response received if the record exists:
-{"data":{"ID":2,"CreatedAt":"2022-02-23T23:25:56.5582548+05:30","UpdatedAt":"2022-02-23T23:25:56.5582548+05:30","DeletedAt":null,"userId":2,"postId":1,"message":"Satoshi is Life","parentId":0,"votes":5}}
-if the record doesn't exist:
-{"error":"record not found"}
-All Comments
-GET-  http://localhost:8080/v1/comments/getAll/1
-Json Received if the record exists:
-{"data":[{"ID":1,"CreatedAt":"2022-02-22T03:39:52.2541628+05:30","UpdatedAt":"2022-02-22T03:39:52.2541628+05:30","DeletedAt":null,"userId":1,"postId":1,"message":"Vitalik is God","parentId":0,"votes":5},{"ID":2,"CreatedAt":"2022-02-23T23:25:56.5582548+05:30","UpdatedAt":"2022-02-23T23:25:56.5582548+05:30","DeletedAt":null,"userId":2,"postId":1,"message":"Satoshi is Life","parentId":0,"votes":5},{"ID":3,"CreatedAt":"2022-02-23T23:32:00.9310077+05:30","UpdatedAt":"2022-02-23T23:32:00.9310077+05:30","DeletedAt":null,"userId":3,"postId":1,"message":"Shitcoins are growing","parentId":0,"votes":5},{"ID":4,"CreatedAt":"2022-02-23T23:32:57.668267+05:30","UpdatedAt":"2022-02-23T23:32:57.668267+05:30","DeletedAt":null,"userId":3,"postId":1,"message":"DJ Snake is Savage","parentId":0,"votes":5}]}
-If the record doesn't exist:
-{"error":"record not found"}
-Create User Comment 
-POST- http://localhost:8080/v1/comments/create
-Json Body Sent-
-{
	"userId" : 1,
	"postId" : 1,
	"message" : "Vitalik is God",
	"parentId" : 0,
	"votes" : 5 
}
-Json response received- 
-{
    "data": 8
}
-DELETE- http://localhost:8080/v1/posts/delete/1
-Json Response Received- 
-{
    "data": {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": null,
        "userId": 0,
        "postId": 0,
        "message": "",
        "parentId": 0,
        "votes": 0
    }
}
-If the record is already deleted- 
-{
    "error": "record not found"
}




# Api endpoints delivered in this sprint 

| Type | Api Endpoint | Description | Status |
| ----------- | ----------- | ----------- | ----------- |
| Get | `/` | Home (Provide Server Status) | Active |
| Get | `/v1/posts/getAll/:userId` | Get all user post from user post ID | Active |
| Get | `/v1/posts/getOne/:id` | Get a single post by post ID | Active |
| Post | `/v1/posts/create` | Create a user post | Active |
| Patch | `/v1/posts/update/:id`   | Update an existing post | In-progress |
| Delete | `/v1/posts/delete/:id`  | Delete an existing post by post ID | Active |
| Get | `/v1/comments/getAll/:userId` | Get all user comments from user comment Id | Active |
| Get | `/v1/comments/getOne/:id` | Get a single comment from user by Id | Active |
| Post | `/v1/comments/create` | Create a user comment | Active |
| Patch |  `/v1/comments/update/:id` | Update an existing comment | In-progress |
| Delete | `/v1/comments/delete/:id` | Delete a user comment | Active |


 