# User Posts

## Get posts by ID 
Endpoint -GET- http://localhost:8080/v1/posts/getOne/2 <br>
Return details of the post with given id if it exists else return error  

Response <br>
```json
// if the record exist:
{
  "data": {
    "ID": 2,
    "CreatedAt": "2022-02-22T03:56:47.6075532+05:30",
    "UpdatedAt": "2022-02-22T03:56:47.6075532+05:30",
    "DeletedAt": null,
    "userId": 1,
    "title": "Test post 2",
    "description": "My second post",
    "userLimit": 4,
    "participants": "1",
    "expiry": 24,
    "viewCount": 0,
    "status": 2,
    "Categories": "",
    "Tags": ""
  }
}

// if the record doesn't exist: 
{
  "error": "record not found"
}
```

## All Posts 
Endpoint -GET- http://localhost:8080/v1/posts/getAll/1 <br>
Retrieve all posts created by a user using user id

```json
// If user have any post
{
  "data": [
    {
      "ID": 2,
      "CreatedAt": "2022-02-22T03:56:47.6075532+05:30",
      "UpdatedAt": "2022-02-22T03:56:47.6075532+05:30",
      "DeletedAt": null,
      "userId": 1,
      "title": "Test post 2",
      "description": "My second post",
      "userLimit": 4,
      "participants": "1",
      "expiry": 24,
      "viewCount": 0,
      "status": 2,
      "Categories": "",
      "Tags": ""
    },
    {
      "ID": 3,
      "CreatedAt": "2022-02-22T04:29:47.2209425+05:30",
      "UpdatedAt": "2022-02-22T04:29:47.2209425+05:30",
      "DeletedAt": null,
      "userId": 1,
      "title": "Test post 2",
      "description": "My second post",
      "userLimit": 4,
      "participants": "1",
      "expiry": 24,
      "viewCount": 0,
      "status": 2,
      "Categories": "",
      "Tags": ""
    },
    {
      "ID": 4,
      "CreatedAt": "2022-02-22T04:30:51.0203078+05:30",
      "UpdatedAt": "2022-02-22T04:30:51.0203078+05:30",
      "DeletedAt": null,
      "userId": 1,
      "title": "Test post 2",
      "description": "Testing Post Creation",
      "userLimit": 4,
      "participants": "1",
      "expiry": 24,
      "viewCount": 0,
      "status": 2,
      "Categories": "",
      "Tags": ""
    }
  ]
}

// If user have not created any post yet
{
  "error": "record not found"
}
``` 

## Create a Post 
Endpoint -POST- http://localhost:8080/v1/posts/create <br>
Create a post with following post params

### Request params
	userId       uint    required 
	title        string  required 
	description  string  
    participants uint  	 
    expiry       float32 
    viewCount    int64   
    userLimit    uint    required
	status       int     required
	Categories   string
	Tags         string

```json
# Request
{
    "userId" : 1,
    "title" : "Test post 2",
    "description" : "Testing Post Creation",
    "userLimit" : 4,
    "status" : 2
}

# Response
// if post get succsesfully created 
{
    "data": 8
}

// if failed to create post
{
    "error": "error message"
}
```

## Delete a post
Endpoint -DELETE- http://localhost:8080/v1/posts/delete/1 <br>
Delete a post with provided id

```json
// If record get Successfully deleted
{
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

// If there is no record with given id 
{
    "error": "record not found"
}
```

## Update user post
Endpoint -Patch- http://localhost:8080/v1/posts/update/1 <br>
Update post details if post exists else return error

```json
// If user profile get Successfully updated
{
    "data": {
        "ID": 1,
        "CreatedAt": "2022-03-04T20:44:03.8333128-05:00",
        "UpdatedAt": "2022-03-04T20:44:10.3285685-05:00",
        "DeletedAt": null,
        "userId": 1,
        "User": {
            "ID": 1,
            "CreatedAt": "2022-03-04T20:26:01.0697293-05:00",
            "UpdatedAt": "2022-03-04T20:42:29.9002396-05:00",
            "DeletedAt": null,
            "userName": "johndoe",
            "firstName": "John",
            "lastName": "Smith",
            "Email": "Smith123@gatorshare.com",
            "zipcode": 0,
            "avatar": "",
            "password": "",
            "bookmark": ""
        },
        "title": "Test post 1",
        "description": "My first post updated 2",
        "userLimit": 5,
        "participantNum": 1,
        "Expiry": 24,
        "viewCount": 0,
        "status": 2,
        "categories": "",
        "tags": ""
    }
}

// If unable to update user posts
{
    "error": "error message"
}
```

## Get reactions on a post by ID 
Endpoint -GET- http://localhost:8080/v1/posts/getReactions/2 <br>
Return reactions on post with given id if they exists else return error  

Response <br>
```json
// if reactions exist for the post:
{
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2022-04-13T18:38:24.5347992-04:00",
            "UpdatedAt": "2022-04-13T18:38:24.5347992-04:00",
            "DeletedAt": null,
            "postId": 2,
            "Post": {
                "ID": 0,
                "CreatedAt": "0001-01-01T00:00:00Z",
                "UpdatedAt": "0001-01-01T00:00:00Z",
                "DeletedAt": null,
                "userId": 0,
                "User": {
                    "ID": 0,
                    "CreatedAt": "0001-01-01T00:00:00Z",
                    "UpdatedAt": "0001-01-01T00:00:00Z",
                    "DeletedAt": null,
                    "userName": "",
                    "firstName": "",
                    "lastName": "",
                    "Email": "",
                    "zipcode": 0,
                    "avatar": "",
                    "password": "",
                    "bookmark": ""
                },
                "title": "",
                "description": "",
                "userLimit": 0,
                "participantNum": 0,
                "Expiry": 0,
                "viewCount": 0,
                "status": 0,
                "categories": "",
                "tags": ""
            },
            "tagId": 3,
            "User": {
                "ID": 3,
                "CreatedAt": "2022-04-09T23:26:06.4215922-04:00",
                "UpdatedAt": "2022-04-15T17:19:56.1959388-04:00",
                "DeletedAt": null,
                "userName": "AnujK",
                "firstName": "Anuj-1",
                "lastName": "Koli",
                "Email": "anuj.singh.koli@gmail.com",
                "zipcode": 0,
                "avatar": "",
                "password": "",
                "bookmark": ""
            },
            "Reaction": "INTERESTED"
        },
        {
            "ID": 2,
            "CreatedAt": "2022-04-13T18:38:26.0346405-04:00",
            "UpdatedAt": "2022-04-13T18:38:26.0346405-04:00",
            "DeletedAt": null,
            "postId": 2,
            "Post": {
                "ID": 0,
                "CreatedAt": "0001-01-01T00:00:00Z",
                "UpdatedAt": "0001-01-01T00:00:00Z",
                "DeletedAt": null,
                "userId": 0,
                "User": {
                    "ID": 0,
                    "CreatedAt": "0001-01-01T00:00:00Z",
                    "UpdatedAt": "0001-01-01T00:00:00Z",
                    "DeletedAt": null,
                    "userName": "",
                    "firstName": "",
                    "lastName": "",
                    "Email": "",
                    "zipcode": 0,
                    "avatar": "",
                    "password": "",
                    "bookmark": ""
                },
                "title": "",
                "description": "",
                "userLimit": 0,
                "participantNum": 0,
                "Expiry": 0,
                "viewCount": 0,
                "status": 0,
                "categories": "",
                "tags": ""
            },
            "tagId": 3,
            "User": {
                "ID": 3,
                "CreatedAt": "2022-04-09T23:26:06.4215922-04:00",
                "UpdatedAt": "2022-04-15T17:19:56.1959388-04:00",
                "DeletedAt": null,
                "userName": "AnujK",
                "firstName": "Anuj-1",
                "lastName": "Koli",
                "Email": "anuj.singh.koli@gmail.com",
                "zipcode": 0,
                "avatar": "",
                "password": "",
                "bookmark": ""
            },
            "Reaction": "INTERESTED"
        }
    ]
}

// if no reaction exist then
{
  "data": []
} 

// if error retrieving the reactions then: 
{
  "error": "record not found"
}
```

## POST reactions on a post by given ID 
Endpoint -POST- http://localhost:8080/v1/posts/reactToPost <br>
Add reaction/Interest to a post with given ID if it exists else return error  

this endpoint require following x-www-form-urlencoded parameters  <br>
postid = 18 *numeric ID* <br>
reaction = {"INTERESTED", "MAYBE", "NOTINTERESTED"} *anyone of the given strings*

```json
// If post reaction get Successfully added
{
    "data": 3
}

// If unable to update reaction on user posts
{
    "error": "error message"
}
```