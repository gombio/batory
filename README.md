# Batory

## Backend

Golang + WebSocket + RethinkDB

## Frontend

ReactJS based app.

Created with `create-react-app`, so make sure to yarn add missing things.

### General instructions

```
Inside that directory, you can run several commands:

  yarn start
    Starts the development server.

  yarn build
    Bundles the app into static files for production.

  yarn test
    Starts the test runner.

  yarn eject
    Removes this tool and copies build dependencies, configuration files
    and scripts into the app directory. If you do this, you can’t go back!
```

### Test data / procedures:

#### Frontend / WebSocket API testing:

1. Go here http://websocket.org/echo.html
2. Use this as Location: ws://localhost:8081/ws

##### Logging in

Failed authentication:

```
{"type":"auth.login","data":{"login":"foo@example.com","password":"xxx"}}
```

Successful authentication:

```
{"type":"auth.login","data":{"login":"john@example.com","password":"test123"}}
```
