# Quasar Attack Operation
*Han solo has been recently named commander of the Rebel Alliance and he seeks for give a great blow against the Galactic Empire to revivate the resistance's call.*

*The rebel intelligent service has detected a helping call from the imperial load holder adrift in an asteroid field. The ship's manifest is high secret, but it's rumored that it's transporting food and weapons for a entire legion.*

## Heroku hosted

This project is being deployed and hosted on Heroku:

https://quasar-attack-operation.herokuapp.com

## Endpoint list:
_{{url}}_ Refers to the url where the project was deployed or is running

1. [POST] {{url}}/topsecret
2. [POST] {{url}}/topsecret/sato
3. [GET] {{url}}/topsecret/sato

**Examples**

Localhost:
- [POST]

http://localhost:9000/topsecret

Heroku: 
- [POST] 

https://quasar-attack-operation.herokuapp.com/topsecret/sato

- [GET] 

https://quasar-attack-operation.herokuapp.com/topsecret/sato

## How to run it
___
First, clone the repository
1. `git clone https://github.com/juancsr/quasar-attack-operation.git`
2. or use go package manager `go get -u github.com/juancsr/quasar-attack-operation`

**Note**: *The following commands I assume you're on the root folder project*

### Runing by Go
___
Once you're on the root project folder you need to download all of the dependecies in the mod file, type:

`$ go mod download`

Just in case, run the `tidy` option, just to be sure to remove all of unused packages:

`$ go mod tidy`

If the two upper commands were right, run the main file to start the gin server:

`go run main.go`

### Running by docker-compose
___
Make sure you have docker and docker compose installed on your machine.
 (You can verify this by simple running the version commands for the two of them):

 ```
 $ docker version
 Client: Docker Engine - Community
 Version:           20.10.5
 API version:       1.41
 Go version:        go1.13.15
 ....

 $ docker-compose version
 docker-compose version 1.28.6, build 5db8d86f
docker-py version: 4.4.4
CPython version: 3.7.10
...
 ```

On the root folder just type: 

`$ docker-compose up`

This will create two containers, one for the project and another one for the mongodb database connection.

### Running using the air development tool (Docker need it)
___
For more info about the air tool for live reloading check out the official repository: https://github.com/cosmtrek/air

**This one is the best approach for a development environment.**

On the project folder run: 

```
$ docker run -ti --rm  \
    -w $(pwd) -v \
    $(pwd):$(pwd) \
    -p 9000:8080 \
    -e GIN_MODE=debug \
    -e USER=<user> \
    -e PASSWORD=<password> \
    -e DB=<db_name> \
    -e CLUSTER=<db_host> \
    cosmtrek/air
```

**Note**: _This is project needs environment variables, so make sure you havo those into your machine_


### Env variables
___
The list of environment variables that should be on your machine are:

- **USER**

This is the user for the authentication database 

- **PASSWORD**

This is the user's password for the authentication database 

- **DB**

The name of the mongodb database to connect

- **CLUSTER**

The url for the mongodb database to connect

- **MODE**

The project mode, its refer when is in development or prod mode. The connection URI building depends of this var value
