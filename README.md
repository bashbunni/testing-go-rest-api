## About

This repo is for practicing writing unit tests for REST APIs in Go.

I will pretend I'm doing API calls for an online chat as an example.

## Endpoints

GetAllMessages -> `/messages` -> GET all messages\
GetMessagesByUserID -> `/messages/:id` -> GET all messages for a given user\
NewMessage -> `/messages` -> POST new message\
DeleteMessage -> `/messages/:id` -> DELETE a message by ID\
\
GetAllUsers -> `/users` -> GET all users\
GetUserByID => `/users/:id` -> GET user by ID\
EditUser -> `/users/:id` -> PUT update user information by ID\
