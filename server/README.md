# Asiap API sample app

## Event Storming
<img src="./docs/event-storming.png" width="721">

## Diagram
<img src="./docs/diagram.png" width="721">

## How to run

The whole environment can be run with:
```bash
docker-compose up
```

There is still an error when running application for the first time because of the connection for RabbitMQ is refused. The App has to be restarted to establish connection with RabbitMQ