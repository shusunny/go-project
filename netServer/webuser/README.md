
This program uses map as an alternative method to store user data in local. To use this, follow the following commands:

## POST users

Enter these commands at the terminal

```
curl -X POST -H "Content-Type: application/json" -d '{"name":"James Bond","gender":"male","age":32}' http://localhost:8080/user
```

```
curl -X POST -H "Content-Type: application/json" -d '{"name":"Miss Moneypenny","gender":"female","age":27}' http://localhost:8080/user
```

```
curl -X POST -H "Content-Type: application/json" -d '{"name":"Q","gender":"female","age":54}' http://localhost:8080/user
```

-X is short for --request
Specifies a custom request method to use when communicating with the HTTP server.

-H is short for --header

-d is short for --data

## GET a user

Enter this at the terminal

```
curl http://localhost:8080/user/<enter-user-id-here>
```

## DELETE a user

Enter this at the terminal

```
curl -X DELETE http://localhost:8080/user/<enter-user-id-here>
```