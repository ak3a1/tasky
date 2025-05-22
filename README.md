# Docker
A Dockerfile has been provided to run this application.  The default port exposed is 8080.

# Environment Variables
The following environment variables are needed.
|Variable|Purpose|example|
|---|---|---|
|`MONGODB_URI`|Address to mongo server|`mongodb://servername:27017` or `mongodb://username:password@hostname:port` or `mongodb+srv://` schema|
|`SECRET_KEY`|Secret key for JWT tokens|`secret123`|

Alternatively, you can create a `.env` file and load it up with the environment variables.

# Running with Go

Clone the repository into a directory of your choice Run the command `go mod tidy` to download the necessary packages.

You'll need to add a .env file and add a MongoDB connection string with the name `MONGODB_URI` to access your collection for task and user storage.
You'll also need to add `SECRET_KEY` to the .env file for JWT Authentication.

Run the command `go run main.go` and the project should run on `locahost:8080`

# Updates from source repo

Golang version and alpine version updated to remove some CVE's from the generated image

[Link to Grype scan results when using the previous image](https://github.com/ak3a1/tasky/blob/main/grype_scans/5_21_2025/previous_image_golang_1.19_alpine_3.17.0.txt)

[Link to Grype scan results when using the new updated image](https://github.com/ak3a1/tasky/blob/main/grype_scans/5_21_2025/new_image_golang_1.24_alpine_3.21.3.txt)

Fixed for the signup form to show errors on the form itself (similarly to what the login page already does)

[Link to commit](https://github.com/ak3a1/tasky/commit/17c7ec24c460390ee267bdb1c557ce747c4da4d9)

# License

This project is licensed under the terms of the MIT license.

Original project: https://github.com/dogukanozdemir/golang-todo-mongodb
