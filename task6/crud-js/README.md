
## `curl` commands for testing

Get all users

```curl
curl http://localhost:3000/users
```

Create new user
```curl
curl -X POST -d "name=John&age=30" http://localhost:3000/users 
```

Get user with id=1
```curl
curl http://localhost:3000/users/1
```

Update data of user with id=1
```curl
curl -X PUT -d "name=Jane&age=35" http://localhost:3000/users/1
```

Delete data of user with id=1
```curl
curl -X DELETE http://localhost:3000/users/1
```
