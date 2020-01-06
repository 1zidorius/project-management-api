### Requirements
* Any modern docker installed

###  To start application

`docker-compose up -d`

#### Usage
##### Populate dummy data using curl
`curl --request POST localhost:8080/api/v1/users --header 'Content-Type:application/json' --data-raw '{"username":"proficientuser","password":"supersafepass","email":"jonas@example.org","name":"jonas","surname":"jonatis"}'`

`GET localhost:8080/api/v1/users` returns all users
<br>
`GET localhost:8080/api/v1/users/{userId}` returns user specific information
<br>
`PUT localhost:8080/api/v1/users/{userId}` for updating fields
<br>
`DELETE localhost:8080/api/v1/users/{userId}` for deleting record

