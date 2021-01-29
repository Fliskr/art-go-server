# art-go-server
Golang server implementation with graphql 

This is first time i use GraphQL ;)

I used https://github.com/99designs/gqlgen instead of https://github.com/graphql-go/graphql which is more popular because of this:

https://codinglatte.com/posts/golang/golang-building-a-graphql-server-part-1/

To start app with sample data use:

``
    PORT=XXXX go run main.go -fill
``

DB is removed and recreated on every startup. 