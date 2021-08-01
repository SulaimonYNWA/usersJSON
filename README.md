### Run server
go run server.go

***Default address***: `127.0.0.1:3001`


The project follows the [MVC](https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93controllerhttps:/) (Model View Controller) pattern (excluding V-View).

- **Model** is where the entity and its structure is created, it is the layer where CRUD (Create Read Update Delete) operations are performed.
- **View** as the name suggests is responsible for displaying the data (not considered in this project)
- **Controller** uses model or multiple data models to perform a task. It may accept data from user and may return processed data.

### Dependencies:

go get github.com/gorilla/mux
go get github.com/go-sql-driver/mysql


