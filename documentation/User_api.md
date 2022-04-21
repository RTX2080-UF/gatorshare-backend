# User Comments

## Get User profile by user ID 
Endpoint -POST- http://localhost:8080/v1/users/getUserProfile/1 <br>
Return data associated with user profile   

Response
```json
// if the user with given id exists:
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
```json
# Request
{
    "username": "johndoe",
    "password": "password123"
}

# Response
// If user is succesfully authenticated
{
    "data": {
        "userDetails": {
            "ID": 3,
            "CreatedAt": "2022-04-09T23:26:06.4215922-04:00",
            "UpdatedAt": "2022-04-15T17:19:56.1959388-04:00",
            "DeletedAt": null,
            "userName": "johndoe",
            "firstName": "John",
            "lastName": "Doe",
            "Email": "johndoe@gatorshare.com",
            "zipcode": 0,
            "avatar": "",
            "password": "",
            "bookmark": ""
        },
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NTAzNDMxOTYsInVzZXJfaWQiOjN9.lC33tz1Cg1_SLT6R7SZgBWuQuIcUdawM0iqqibwT3aM"
    }
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

```json
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
    "error": "error message"
}
```

## Delete user profile
Endpoint -DELETE- http://localhost:8080/v1/users/delete/ <br>
Delete user and his associated data with given id

```json
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

```json
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
## Refresh user auth token 
Endpoint -POST- http://localhost:8080/v1/users/refreshToken <br>
Return new authentication with validity of 15 min   

Response
```json
// if the user with given id exists:
{
    "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJleHAiOjE2NDg4NjM2ODEsInVzZXJfaWQiOjJ9.dqRxryaxfSy6ix6EgKdhbJ1GUHI78V95BI9RIQRYQYk"
}

// if current token is invalid:
{
    "error": "Token expired, please login again",
    "errorDetails": "Token expired, please login again"
}
```
## Get current user profile 
Endpoint -POST- http://localhost:8080/v1/users/getProfile <br>
Return profile data associated with currently logged in user    

Response
```json
// if the user with given id exists:
{
    "data": {
        "ID": 2,
        "CreatedAt": "2022-03-22T12:33:50.3682646-04:00",
        "UpdatedAt": "2022-03-22T12:34:56.6777649-04:00",
        "DeletedAt": null,
        "userName": "alicesmith",
        "firstName": "John",
        "lastName": "Smith",
        "Email": "alicesmith@gatorshare.com",
        "zipcode": 0,
        "avatar": "",
        "password": "$2a$14$kH1wAUcCzmJrzwppLzJveOCRp..Riei0urFAdPwBflovtykr2.Woq",
        "bookmark": ""
    }
}

// if the user isn't found:
{
  "error": "record not found"
}
```

## Follow another user  
Endpoint -POST- http://localhost:8080/v1/users/follow/1 <br>
Return created record id if succesfully able to follow else error message   

Response
```json
// if succesfully able to follow user
{
    "data": 2
}

// if failed to follow user
{
    "error": "error message",
    "errorDetails": "error message"
}
```

## Follow another user  
Endpoint -POST- http://localhost:8080/v1/users/listFollowers/1 <br>
Return follower list for user with given id   

Response
```json
// if succesfully able to get user follower
{
    "data": [
        {
            "ID": 1,
            "CreatedAt": "2022-04-01T17:02:09.8321374-04:00",
            "UpdatedAt": "2022-04-01T17:02:09.8321374-04:00",
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
            "followerId": 2,
            "Follower": {
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
            }
        },
        {
            "ID": 2,
            "CreatedAt": "2022-04-01T17:29:28.2380571-04:00",
            "UpdatedAt": "2022-04-01T17:29:28.2380571-04:00",
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
            "followerId": 2,
            "Follower": {
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
            }
        }
    ]
}

// if failed to get followers list
{
    "error": "error message",
    "errorDetails": "error message"
}
```

## Request to reset forgotten password for given emailId 
Endpoint -GET- http://localhost:8080/v1/users/resetPassword?email=allensolley@gatorshare.com <br>
Send a password reset link to users email account   

Response
```json
// if the user with given id exists and password reset link sent successfully:
{
    "data": true
}

// if the user doesn't exist or not able to generate reset link:
{
    "error": "Unable to generate password reset link",
    "errorDetails": "Unable to generate password reset link"
}
```

## Update password using reset password link 
Endpoint -POST- http://localhost:8080/v1/users/updatePassword <br>
Update password to new password if correct reset token is provided    

this endpoint require  following x-www-form-urlencoded parameters
email = "allensolley@gatorshare.com" *users email address*
token = "$2a$14$aJia6YKg5cqBcl.r2sJVNO4mjzLj1FWxHlI6fEkbroDOUTCxr4TBq" *token from email*
password = "newpassword" *newpassword string*

Response
```json
// if the user with given id exists and password get reset succcesfully:
{
    "data": true
}

// if the user doesn't exist or not able to generate reset link:
{
    "error": "Unable to update password for user",
    "errorDetails": "Unable to update password for user"
}
```
