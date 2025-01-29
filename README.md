## HNG 12 Stage 0 Task

This is a simple backend API built with the [Go](https://go.dev) programming language.

### API Endpoints
There is only one endpoint in the API:

- `"/"`: which has the following JSON response
```json
{
  "email": "fawomath@gmail.com",
  "current_datetime": "2025-01-29T18:53:06Z",
  "github_url": "https://github.com/Abdulrasheed1729/hng12-stage0"
}
```

### Local Setup

To run this project locally following the following steps:
- Install the [Go](https://go.dev) programming language on your machine.
- Clone the repository

```sh
git clone https://github.com/Abdulrasheed1729/hng12-stage0

cd hng12-stage0/
```

- Build the project
```
go build .
```

- Run the project
```
./hng12-stage0
```

- Make a Sample request
```
curl -X GET http://localhost:8080
```
