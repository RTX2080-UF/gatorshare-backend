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

## User Features
### Ride Sharing
- Find people to share rides in their area. Our algorithm will automatically match people based on locality and their schedule flexibility.
- Plan rental cars together.
- Automatic email notification when a partner is found.
- You're under control: set limits on how many partners you are looking for. 

### Item Sharing
- Buy bulk items together like pizza, eggs, stationary and split among people.
- Subscribe to topics of your interest like perishables, gardening items, groceries, etc, and get notified when someone posts a share request.
- Order groceries together to bypass the minimum order limit. (We know you do it!)
- Share coupon codes to get additional discounts.

### Events
- Plan potlucks, group studies and much more!

## Enterprise features
KPIs available with a simple REST API:
- Riding patterns of people.
- Riding patterns of an area.

# Geek Zone
## Basic features
- User profiles: Login and registration.
- Community: Posts and comments, upvote and downvote users and posts.
- Notify: Push and email notifications.
- Follow: Create and subscribe to topics.

More features TBD...

# Setup and run instructions
1. First download and install go setup from *[here](https://go.dev/doc/install)*

2. If you have Make tool install run
```
$ make server 
```
else you can directly run
```
$ go run main.go
```
3. You can check server status in terminal and website will be up onm port `8080` by default.

# Api endpoints

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
