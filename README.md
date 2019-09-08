### Dockerize
#### Install portainer to manage docker easly from web based application
1. open utils directory
2. Deploy portainer service with stack portainer
```
docker stack deploy --compose-file=db.docker-compose.yml dockerize
```
3. open http://localhost:9000 from your browser
4. setup the username and password
5. login using username and password 

#### How to run a container from dockerfile
1. open docker directory
2. you will see many folders that implement dockerfile
3. choose a folder and open it
4. build the docker image from dockerfile
```
docker build --tag "react-app:latest"  .
```
5. start the container from docker image
```
docker run --name react-app -p 5000:5000 react-app:latest
```
6. open http://localhost:5000 from the browser

#### How to run multiple container from docker-compose
1. open docker-compose directory
2. build multiple services with docker-compose
```
docker-compose -f service.docker-compose.yml build
```
3. run multiple service containers with docker-compose
```
docker-compose -f serive.docker-compose.yml up
```
3. remove multiple service containers with docker-compose
```
docker-compose -f service.docker-compose.yml down
```

#### How to deploy in swarm mode

Build database 
```docker-compose -f db.docker-compose.yml build```

Deploy database service with stack dockerize
```docker stack deploy --compose-file=db.docker-compose.yml dockerize```

Build app   
```docker-compose -f service.docker-compose.yml build```

Deploy app with stack dockerize 
```docker stack deploy --compose-file=service.docker-compose.yml dockerize```