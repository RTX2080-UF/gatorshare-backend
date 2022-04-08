# User Posts

## Get tag by ID 
Endpoint -GET- http://localhost:8080/v1/tags/getOne/2 <br>
Return details of the tags with given id if it exists else return error  

Response <br>
```json
// if the record exist:
{
    "data": {
        "ID": 1,
        "CreatedAt": "2022-04-01T20:12:26.5152944-04:00",
        "UpdatedAt": "2022-04-01T20:12:26.5152944-04:00",
        "DeletedAt": null,
        "name": "Test tag",
        "creatorId": 2,
        "Creator": {
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
        "votes": 0,
        "description": "Test tag"
    }
}

// if the record doesn't exist: 
{
  "error": "record not found"
}
```

## Create a tag 
Endpoint -POST- http://localhost:8080/v1/tags/create <br>
Create a tag with following post params

### Request params
	Name         string    required 
	description  string

```json
# Request
{
    "Name": "Test tag",
    "Description": "Test tag"
}

# Response
// if post get succsesfully created 
{
    "data": 1
}

// if failed to create post
{
    "error": "error message"
}
```

## Delete a tag
Endpoint -DELETE- http://localhost:8080/v1/tags/delete/1 <br>
Delete a tag with provided id

```json
// If record get Successfully deleted
{
    "data": {
        "ID": 1,
        "CreatedAt": "2022-04-01T20:12:26.5152944-04:00",
        "UpdatedAt": "2022-04-01T20:12:26.5152944-04:00",
        "DeletedAt": null,
        "name": "Test tag",
        "creatorId": 2,
        "Creator": {
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
        "votes": 0,
        "description": "Test tag"
    }
}

// If there is no record with given id 
{
    "error": "record not found"
}
```

## Update tags
Endpoint -Patch- http://localhost:8080/v1/tags/update/1 <br>
Update tag details if tag exists else return error

```json
# Request
{
    "Name": "Test tag updated",
    "Description": "Test tag updated"
}

// If tag get Successfully updated
{
    "data": {
        "ID": 2,
        "CreatedAt": "2022-04-01T20:45:52.9114396-04:00",
        "UpdatedAt": "2022-04-01T20:46:15.042908-04:00",
        "DeletedAt": null,
        "name": "Test tag update",
        "creatorId": 2,
        "Creator": {
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
        "votes": 0,
        "description": "Test tag update"
    }
}

// If unable to update user posts
{
    "error": "error message"
}
```


## Follow tags
Endpoint -GET- http://localhost:8080/v1/tags/update/1 <br>
Add tag to users tag subscription list if tag exists else return error

```json
// If user profile get Successfully updated
{
    "data": 1
}

// If unable to update user posts
{
    "error": "FOREIGN KEY constraint failed",
    "errorDetails": "unable to associate tag with given id"
}
```
