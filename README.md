# TodoBerry Server
TodoBerry is an opensource TodoList Application. 

This Repository belongs to TodoBerry's Server side application that serves our
RestFul API.

## Getting Start

| Make sure your redis server is running and configuration inside `config/conf/redis.toml` is up-to-date with your server configuration.

```
 $ go get github.com/TameshkCloud/todoberry-server
```

Change your current directory to project directory located inside your $GOPATH
and start coding.

## Start Development

We use [Glide Package Manager](https://github.com/Masterminds/glide) to manage
our project dependencies. Make sure Glide is installed on your system:

```bash
$ curl https://glide.sh/get | sh
```


Install dependencies:

```bash
$ glide up
```

Glide will install all dependencies listed in `glide.yml` file located in root directory
of project. All this project vendors are **patched** version and stable.

Check out `router/gin.go` to follow-up code execution

 - [ ] Documentation about project dependencies (English/Farsi)

## Generating Swagger documentations
simply run:

```
$ swag init
```

now you can run `go run main.go`  and check [http://127.0.0.1:3000/swagger/index.html](http://127.0.0.1:3000/swagger/index.html) to see documentations generated automatically by swagger!




