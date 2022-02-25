# prpcl
## How to use: 

### Run
```
sh start.sh
```
The docker-compose already contained the mysql and migration, it took about 5s for everything to be ready.


### Unit test
```
make test
```
The unit test is running directly without docker thus ensure that go is already installed.
### API Documentation
Swagger can be accessed via `localhost:8080/prpcl/swagger/index.html`