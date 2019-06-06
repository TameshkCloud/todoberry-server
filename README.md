# TodoBerry Server
TodoBerry is an opensource TodoList Application. 

This Repository belongs to TodoBerry's Server side application that serves our
RestFul API.

## Getting Start

```
 $ go get -u github.com/TameshkCloud/todoberry-server
```

> Make sure your redis server is running and configuration inside `config/conf/redis.toml` is up-to-date with your server configuration.


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

## Generating Swagger documentations
Install `gin-swagger`:

```bash
$ go get -u github.com/swaggo/swag/cmd/swag
```

and simply run:

```
$ swag init
```

now you can run `go run main.go`  and check [http://127.0.0.1:3000/swagger/index.html](http://127.0.0.1:3000/swagger/index.html) to see documentations generated automatically by swagger!

there is two additional health checker apis you can check them here:
 - [http://127.0.0.1:3000/health/welcome](http://127.0.0.1:3000/health/welcome)
 - [http://127.0.0.1:3000/health/check/redis](http://127.0.0.1:3000/health/check/redis)
 
 
 
## Todos
 - [ ] MongoDB driver
 - [ ] JWT Authentication
 - [ ] Writing unit tests
 - [ ] `etcd` integration with Viper
 - [ ] command line tools using Cobra
 - [ ] Dependency Documents (English/Persian)
 
## Contact Us
feel free to send your ideas in twitter for us: [@rezaseyf2013](https://twitter.com/rezaseyf2013)
