# User Home

## Get relevant posts for user 
Endpoint -GET- http://localhost:8080/v1/home/user   <br>
Return latest posts matching user interests to be displayed on homepage  

Response
```json
// if the user has any new post:
{
    "data": [
        {
            "ID": 16,
            "CreatedAt": "2022-04-17T14:00:03.4459524-04:00",
            "UpdatedAt": "2022-04-17T14:00:03.4459524-04:00",
            "DeletedAt": null,
            "userId": 3,
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
            "title": "Pizza party at Niche on Sunday",
            "description": "Come and join us for pizza party on this Sunday",
            "userLimit": 4,
            "participantNum": 1,
            "Expiry": 24,
            "viewCount": 0,
            "status": 2,
            "categories": "",
            "tags": ""
        },
        {
            "ID": 16,
            "CreatedAt": "2022-04-17T14:00:03.4459524-04:00",
            "UpdatedAt": "2022-04-17T14:00:03.4459524-04:00",
            "DeletedAt": null,
            "userId": 3,
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
            "title": "Beach trip this Sunday",
            "description": "Let's go to beach this Sunday",
            "userLimit": 4,
            "participantNum": 1,
            "Expiry": 24,
            "viewCount": 0,
            "status": 2,
            "categories": "",
            "tags": ""
        }
    ]
}

// If the there are no new post then 204 is returned:
// If unable to retrieve posts then a error message is returned 
{
  "error": "error message"
}
```

## Get generic posts 
Endpoint -GET- http://localhost:8080/latest?page=2&page_size=10   <br>
Return latest generic posts for homepage  
Pagination is implemented on this endpoint and could be used to retrieve post in batches
page defines the offset and page_size determines the number of post to fetch

Response
```json
// if the user has any new post:
{
    "data": [
        {
            "ID": 3,
            "CreatedAt": "2022-04-13T18:32:28.4820424-04:00",
            "UpdatedAt": "2022-04-13T18:32:28.4820424-04:00",
            "DeletedAt": null,
            "userId": 3,
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
            "title": "Test post 1",
            "description": "My first post",
            "userLimit": 4,
            "participantNum": 1,
            "Expiry": 24,
            "viewCount": 0,
            "status": 2,
            "categories": "",
            "tags": ""
        },
        {
            "ID": 2,
            "CreatedAt": "2022-03-21T00:27:09.5201269-04:00",
            "UpdatedAt": "2022-03-21T00:27:09.5201269-04:00",
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
            "title": "Test post 2",
            "description": "My second post",
            "userLimit": 4,
            "participantNum": 1,
            "Expiry": 24,
            "viewCount": 0,
            "status": 2,
            "categories": "",
            "tags": ""
        },
        {
            "ID": 1,
            "CreatedAt": "2022-03-20T23:16:37.1693305-04:00",
            "UpdatedAt": "2022-03-20T23:16:37.1693305-04:00",
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
            "title": "Test post 3",
            "description": "My third post",
            "userLimit": 4,
            "participantNum": 1,
            "Expiry": 24,
            "viewCount": 0,
            "status": 2,
            "categories": "",
            "tags": ""
        }
    ]
}

// If the there are no new post then 204 is returned:
// If unable to retrieve posts then a error message is returned 
{
  "error": "error message"
}
```