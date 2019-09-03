### Dockerize

#### Create Dockerfile

#### Create docker-compose

#### How to deploy in swarm mode

Build database 
```docker-compose -f db.docker-compose.yml build```

Deploy database service with stack dockerize
```docker stack deploy --compose-file=db.docker-compose.yml dockerize```

Build app   
```docker-compose -f service.docker-compose.yml build```

Deploy app with stack dockerize 
```docker stack deploy --compose-file=service.docker-compose.yml dockerize```