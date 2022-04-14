# placement-cracker-API

## Installation process

### For mux router
```
go get -u github.com/gorilla/mux
```
### Setting the environment port
```
PORT=3000
```
#### Note that the port has to be declared inside an .env file

### Setting up the router
```
router.HandleFunc()
http.ListenAndServe()
```
## How to run
- Open postman
- Type
```
http://localhost:3000/(your endpoint)
```
- Send the request 
- Before sending the request run the following command in the project
```
go run main.go
```
## Tech Stack used to build the app
- Flutter, Golang, SQL, Firebase
