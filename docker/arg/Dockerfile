FROM node:latest

COPY . $GOPATH/src/github.com/hengkyawijaya/dockerize/docker/arg

WORKDIR $GOPATH/src/github.com/hengkyawijaya/dockerize/docker/arg/react-app

RUN npm install --global yarn react-scripts serve

RUN yarn install

ARG REACT_APP_SERVICE_NAME=kidcon
ARG REACT_APP_LINK=https://www.unitypaloalto.org/wp-content/uploads/LoveLiftedUp.jpg

RUN yarn build

CMD serve -s build

EXPOSE 5000