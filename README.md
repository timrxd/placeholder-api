# placeholder-api
Basic Go API that serves placeholder data

## Run
```
go build
./placeholder-api
```
- hosted on localhost:8080

## Endpoints
```
 GET    /items
    - gets all items 

 GET    /item/{id}
    - gets a specific item
 
 POST   /item
    - request body contains item
    - creates an item
    - id cannot already exist

 PUT    /item/{id}
    - request body contains item
    - updates an item in the database
    - item must exist

 DELETE  /item/{id}
    - deletes an item in the database
    - item must exist

 ```

 ## Item Spec
 ```
Item {
    int id
    int userId
    string title
    string body
}
 ```