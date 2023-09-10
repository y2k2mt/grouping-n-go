# Grouping-N Go lang version

A tiny API: Just shuffles the given names.

### Setting up the project

1. Install go lang
1. Install [task](https://taskfile.dev/installation/)
1. Install [air](https://github.com/cosmtrek/air)
1. Install [migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
1. Run postgresql database server (e.g. `docker run --rm -d -p 5432:5432 -e POSTGRES_HOST_AUTH_METHOD=trust postgres:14`)
1. Run `task migrate` to migrate local database
1. Run `task dev` to start the app
1. Request 
```
curl localhost:1323/grouping -H 'Content-Type: application/json' -d '{"n": 2,"members": ["1","2","3","4","5","6","7","8"]}'
```
5. And response
```
< HTTP/1.1 200 OK
< Content-Type: application/json
< Content-Length: 126
{
  "groups": [
    {
      "members": [
        "8",
        "7",
        "3",
        "2"
      ]
    },
    {
      "members": [
        "5",
        "1",
        "4",
        "6"
      ]
    }
  ],
  "id": "08d6277f-89aa-4da2-b2d6-627cd91fc749"
}

```
