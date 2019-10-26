# Go GraphQL Server Template

Scalable web server written in Go that uses JWT's for authentication and MongoDB to persist data. The GraphQL endpoint is `/graphql`. The schemas for the implemented mutations and query can be found below. Since certain features like enforcing a certain password strength are not implemented yet, please use it with caution.

## Usage

Add an `.env` file according to the `.example.env` file. Install dependencies with the command `go get` and start the server with `go run main.go` in the root folder.

### SignUp Mutation

```graphql
mutation {
  signUp(
    input: {
      email: "name@example.com"
      firstName: "Some"
      lastName: "User"
      password: "example"
    }
  ) {
    token
  }
}
```

### Login Mutation

```graphql
mutation {
  Login(
    input: {
      email: "name@example.com"
      password: "example"
    }
  ) {
    token
  }
}
```

### Access Protected Resolvers

Add an `Authorization` header with a valid JWT as bearer token to the request. The following protected query can only be accessed if a valid JWT is provided.

```graphql
query {
  protected {
    message
  }
}
```
