# User Notifications

## Get new notifications for user 
Endpoint -GET- http://localhost:8080/v1/notifications/getNew   <br>
Return most recent notifications for user  

Response
```json
// if the user has any new notifications:
{
    "data": [
        {
            "ID": 0,
            "CreatedAt": "2022-04-19T11:10:12.8316666-04:00",
            "UpdatedAt": "2022-04-19T11:10:12.8316666-04:00",
            "DeletedAt": null,
            "userId": 5,
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
            "link": "",
            "seen": false,
            "description": "User AnujK followed you recently check html link below to see the notification"
        },
        {
            "ID": 0,
            "CreatedAt": "2022-04-19T11:21:57.8979087-04:00",
            "UpdatedAt": "2022-04-19T11:21:57.8979087-04:00",
            "DeletedAt": null,
            "userId": 5,
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
            "link": "",
            "seen": false,
            "description": "User allensolley reacted on your post"
        }
    ]
}

// If the there are no new notifications then 204 is returned:
// If unable to retrieve notifications then a error message is returned 
{
  "error": "error message"
}
```

## Update notifications status for user 
Endpoint -GET- http://localhost:8080/v1/notifications/updateStatus   <br>
Update read status of notifications to seen

Response
```json
// if succesfully update the status then true response is returned:
{
    "data": true
}

// If failed to update the status then error message is returned  
{
  "error": "error message"
}
```