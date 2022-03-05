# User Comments

## Get User profile by user ID 
Endpoint -POST- http://localhost:8080/v1/users/getProfile/1 <br>
Return data associated with user profile   

Response
```javascript
// if the user with given id exists:
{
{
    "data": {
        "ID": 1,
        "CreatedAt": "2022-03-04T18:29:35.563751Z",
        "UpdatedAt": "2022-03-04T18:29:35.563751Z",
        "DeletedAt": null,
        "userName": "johndoe",
        "firstName": "John",
        "lastName": "Doe",
        "Email": "johndoe@gatorshare.com",
        "zipcode": 0,
        "avatar": "",
        "password": "",
        "bookmark": ""
    }
}

// if the user isn't found:
{
  "error": "record not found"
}
```

## Authenticate user (Login)
Endpoint -Post- http://localhost:8080/v1/users/login/ <br>
Return all comments associated with a post with given id

### Request params
```javascript
# Request
{
    "username": "johndoe",
    "password": "password123"
}

# Response
// If user is succesfully authenticated
{
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDY0NDQ0NjYsInVzZXJfaWQiOjF9.-XzJYiPF7U-zYQ525kUQEiqkCtJM6foxqspjL73lprk"
}

// If user authentication fails
{
    "error": "unable to authenticate user"
}
```

## Register User
Endpoint -POST- http://localhost:8080/v1/users/register <br>
Register a user profile

### Request params
	username  string  required
	firstname string  required
	lastname  string  required
	email     string  required
	password  string  required
	zipcode	  string 

```javascript
# Request
{
    "Username": "johndoe",
    "Firstname": "John",
    "Lastname": "Doe",
    "Email": "johndoe@gatorshare.com",
    "Password": "password123"
}

# Response
// if user is succesfully registered 
{
    "data": 1
}

// if failed to register user
{
    "error": error message
}
```

## Delete user profile
Endpoint -DELETE- http://localhost:8080/v1/users/delete/ <br>
Delete user and his associated data with given id

```javascript
// If user profile get Successfully deleted
{
  "data": {
    "ID": 1,
    "CreatedAt": "2022-03-04T20:21:40.922675-05:00",
    "UpdatedAt": "2022-03-04T20:21:40.922675-05:00",
    "DeletedAt": null,
    "userName": "johndoe",
    "firstName": "John",
    "lastName": "Doe",
    "Email": "johndoe@gatorshare.com",
    "zipcode": 0,
    "avatar": "",
    "password": "",
    "bookmark": ""
  }
}

// If the record is already deleted or doesn't exist
{
    "error": "record not found"
}
```

## Update user profile
Endpoint -Patch- http://localhost:8080/v1/users/updateProfile/1 <br>
Update user details if user exists

```javascript
// If user profile get Successfully updated
{
    "data": {
        "ID": 1,
        "CreatedAt": "2022-03-04T20:26:01.0697293-05:00",
        "UpdatedAt": "2022-03-04T20:42:11.7414594-05:00",
        "DeletedAt": null,
        "userName": "johndoe",
        "firstName": "John",
        "lastName": "Smith",
        "Email": "Smith123@gatorshare.com",
        "zipcode": 0,
        "avatar": "",
        "password": "",
        "bookmark": ""
    }
}

// If unable to update user profile
{
    "error": "record not found"
}
```
