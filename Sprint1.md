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

## Sprint 1 Deliverables

Backend has three database tables, one table for users, one table for user posts and one table for user comments.
The issues resolved resemble the slow integration with the backend design, the template for Entity-Relationship diagram has been linked to one of the issues, and it also contains the design database for database architecture. We have used gin and gorm in the backend and SQlite for the database. The Api response has been tested through postman. 


### User Stories 
- User should be able to create a post (Split request).
- User should be able to comment on other posts.
- Create database design.

## Issues Resolved
- User should be able to create a post (Split request).
- User should be able to comment on other posts.
- Create database design.
- A Make file is needed to automate running server.
- Add postman collection for relevant API requests.

## Basic features
- User profiles: Login and registration.
- Community: Posts and comments, upvote and downvote users and posts.
- Notify: Push and email notifications.
- Follow: Create and subscribe to topics.


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
