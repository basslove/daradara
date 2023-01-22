# daradara

***Please execute the following command at RepoRoot***

## openapi
```
# # openapi gen -> type, spec, server
$ make oapi_gen

# # swagger up(ui, editor)
$ make oapi_up

# access swagger-ui-url
http://localhost:8189/docs/

# access swagger-editor-url
http://localhost:8188
```


## migrate

```
# # db-server run
$ docker-compose up -d 

# # install for Mac
$ brew install golang-migrate

# # generate migrate file name
$ make db_create_migration migration_name=xxxx

# # execute dev migrations(up / down)
$ make dev_db_migrate_up
$ make dev_db_migrate_down

# # execute test migrations(up / down)
$ make test_db_migrate_up
$ make test_db_migrate_down

# # case dev migrate fail
$ make dev_db_migrate_force

# # case test migrate fail
$ make test_db_migrate_force
```

## fixtures
```
$ make seed
```

## server run
```
$ make run
```
