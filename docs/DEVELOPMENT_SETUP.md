# dev setup

### local env
Please use mac machine 

### install docker
https://docs.docker.com/engine/install/

```
# # installed check
$ docker version
```

### install golang to local-machine

```
# # installed check
$ go version 
```

### install node

```
$ brew install nodebrew

# # add zshrc or bashrc
export PATH=$HOME/.nodebrew/current/bin:$PATH

# # source
$ source ~/.zshrc

# # set
$ nodebrew setup

$ nodebrew install v19.9.0

$ nodebrew use v19.9.0

$ node -v && npm -v
v19.9.0
9.6.3
```

### install npm tool
```
$ npm install -g swagger-cli@4.0.4
```

### clone
```
$ cd $HOME
$ mkdir hoge_projects && cd $_
$ git clone git@github.com:basslove/daradara.git
$ cd daradara
```

***Please execute the following command at RepoRoot***

-----------------------------------

### check pwd
```
$ pwd

/XXXXX/XXXXX/hoge_projects/daradara
```

### set git config
```
$ git config --local user.name "XXXX"
$ git config --local user.email XXXX@XXXXX.com

$ git config --local --list
```

### .env.sample copy & rename .env
```
$ cp .env.sample .env
```

### open-api start
```
$ make oapi_up

# # access swagger-ui-url
http://localhost:8189/docs/

# # access swagger-editor-url
http://localhost:8188
```

### db start
```
$ docker compose up [-d]
```

### login-check db
```
# # develop
$ docker container exec -it golang_dev_db bash
=# psql -h localhost -U dev01 -d dev01

# # test
$ docker container exec -it golang_test_db bash
=# psql -h localhost -U test01 -d test01
```

### db migrate
```
# # install for Mac
$ brew install golang-migrate

# # execute dev migrations
$ make dev_db_migrate_up

# # execute test migrations
$ make test_db_migrate_up
```

### db seed
```
$ make seed
```

### test run(optional)
```
$ go test -v ./...
```

### server start
```
$ make run
```
