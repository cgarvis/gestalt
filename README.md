# gestalt

Configuration for Golang

## Usage

Defaults to looking for the environment varable with the same name.
It expects the variable to be all caps
and converts all `.` to `_`.

```go
config := gestalt.New()

# Lookup for PORT
port := config.Int('port', 8080)

# Lookup for MONGO_URI
mongo_uri := config.String('mongo.uri', 'localhost:27072')
```
