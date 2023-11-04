# GoLang MusicApp Platform Assignment

This is a simple MusicApp platform built using GoLang, Gin framework, MySQL database, Unit 
testing and Microservice Architecture. 
It provides basic following operations for managing Songs.

- Insert a new Song(Name,Artists,Genre,Publish Year, Language) in App .
- Retieve songs by specific Genre.


## Prerequisites

Before running this application, make sure you have the following:

- GoLang installed on your system
- MySQL database server running
- Postman for testing the APIs (see Postman collection in postmanCollection.txt)

## Setup

1. Unzip the folder/Clone this repository:

   ```shell
   git clone "https://github.com/PATILSHUBHAM69/MyMusicApp.git"

2. Create a .env file in the project root directory with the following content, 
replacing the placeholders with your database credentials:
DB_USERNAME=your_db_username
DB_PASSWORD=your_db_password
DB_NAME=your_db_name

Also make mock database for testing.

3. Install the required Go packages:

-- go mod tidy

4. Run the application:
-- redirect to api-Gateway repository 
-- go run main.go


5. Check all api on postman mention in postmanCollection.txt file

6. Also run test file mentioned in InsertNewSong/controller and GetSongByGenre/controller using
 "go test" command
