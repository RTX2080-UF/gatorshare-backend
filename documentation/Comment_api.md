# User Comments

## Get comment by comment ID 
Endpoint -GET- http://localhost:8080/v1/comments/getOne/2 <br>
Return comment with a given id if exists else return error  

Response
```javascript
// if the comment with given id exists:
{
  "data": {
    "ID": 2,
    "CreatedAt": "2022-02-23T23:25:56.5582548+05:30",
    "UpdatedAt": "2022-02-23T23:25:56.5582548+05:30",
    "DeletedAt": null,
    "userId": 2,
    "postId": 1,
    "message": "Very helpful post",
    "parentId": 0,
    "votes": 5
  }
}

// if the record doesn't exist:
{
  "error": "record not found"
}
```

## Get all Comment by post id
Endpoint -GET- http://localhost:8080/v1/comments/getAll/ <br>
Return all comments associated with a post with given id

Response
```javascript
// If post has comments 
{
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2022-02-22T03:39:52.2541628+05:30",
      "UpdatedAt": "2022-02-22T03:39:52.2541628+05:30",
      "DeletedAt": null,
      "userId": 1,
      "postId": 1,
      "message": "checkout this post",
      "parentId": 0,
      "votes": 5
    },
    {
      "ID": 2,
      "CreatedAt": "2022-02-23T23:25:56.5582548+05:30",
      "UpdatedAt": "2022-02-23T23:25:56.5582548+05:30",
      "DeletedAt": null,
      "userId": 2,
      "postId": 1,
      "message": "Nice article",
      "parentId": 0,
      "votes": 5
    },
    {
      "ID": 3,
      "CreatedAt": "2022-02-23T23:32:00.9310077+05:30",
      "UpdatedAt": "2022-02-23T23:32:00.9310077+05:30",
      "DeletedAt": null,
      "userId": 3,
      "postId": 1,
      "message": "Share me the link",
      "parentId": 0,
      "votes": 5
    },
    {
      "ID": 4,
      "CreatedAt": "2022-02-23T23:32:57.668267+05:30",
      "UpdatedAt": "2022-02-23T23:32:57.668267+05:30",
      "DeletedAt": null,
      "userId": 3,
      "postId": 1,
      "message": "DJ Snake is Savage",
      "parentId": 0,
      "votes": 5
    }
  ]
}

// If there are no comment on the post
{
  "error": "record not found"
}
```

## Create Comment on a post
Endpoint -POST- http://localhost:8080/v1/comments/create <br>
Create a comment on a post

### Request params
	userId   uint   required
	postId   uint   required
	message  string required
	parentId uint   
	votes    int    

```javascript
# Request
{
	"userId" : 1,
	"postId" : 1,
	"message" : "Sherlock rulez",
	"parentId" : 0,
	"votes" : 5 
}

# Response
// if comment get succsesfully created 
{
    "data": 7
}

// if failed to create comment
{
    "error": error message
}
```

## Delete comment
Endpoint -DELETE- http://localhost:8080/v1/posts/delete/ <br>
Delete a comment with given id

```javascript
// If record get Successfully deleted
{
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

// If the record is already deleted or doesn't exist
{
    "error": "record not found"
}
```


## Update user comment
Endpoint -Patch- http://localhost:8080/v1/comments/update/1 <br>
Update comment message if post exists and comment exist else return error

```javascript
// If user comment on a post get Successfully updated
{
    "data": {
        "ID": 1,
        "CreatedAt": "2022-03-04T20:45:44.6222994-05:00",
        "UpdatedAt": "2022-03-04T20:47:28.5251872-05:00",
        "DeletedAt": null,
        "userId": 1,
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
        "postId": 1,
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
        "message": "Working update",
        "parentId": 0,
        "votes": 5
    }
}
// If unable to update user comments
{
    "error": "error message"
}
```
