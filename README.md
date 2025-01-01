# Go Echo Restful API Template

<hr>

## Description

This project is a template for building RESTful APIs using the Go programming language and the Echo framework. It provides a structured and scalable architecture, including layers for API handlers, domain logic, infrastructure, middleware, and services. The template also integrates various tools and libraries for database connections, message brokers, logging, configuration management, and more.

The project is designed to help developers quickly set up and start building robust and maintainable APIs with best practices in mind.

## Features

- **Clean Architecture**: Structured and scalable architecture with layers for API handlers, domain logic, infrastructure, middleware, and services.
- **Database Support**: Includes configurations for MySQL, MongoDB, and Redis.
- **Dependency Injection**: Uses Google Wire for dependency injection.
- **Configuration Management**: Viper for managing configuration files.
- **Logging**: Zap for structured logging.
- **API Documentation**: Swagger for API documentation.
- **Health Checks**: Endpoints for health, readiness, and liveness checks.
- **Metrics**: Prometheus metrics endpoint.
- **Docker Support**: Dockerfile and docker-compose for containerization.

```
Project structure  
│
└───cmd
│   │   main.go  -> entry point of the service
│   │
│   └───internal
│       │   
│       └───api     -> api handlers
│       │   
│       └───domain  -> domain layer
│       │
│       └───infrastructure   -> infrastructure layer
│       │   │         
│       │   └───adapter     -> http service
│       │   │
│       │   └───grpc        -> grpc service
│       │   │
│       │   └───repository  -> database
│       │   │
│       │   └───publisher   -> message broker
│       │   
│       └─── middleware  -> middleware
│       │
│       └───service      -> service layer
│       │
│       └───router       -> router
│       │
│       └───subscriber   -> listener
│       
└──────pkgs
│       │   
│       └───database     -> database connection
│       │   
│       └───heath        -> healthcheck
│       │
│       └───httpCaller   -> restify client (http)
│       │
│       └───rabbit_client -> rabbitmq client
│       │
│       └───ultis   -> ultis
│       
└───config
│    │   config.go -> config instance
│    └── config.yml -> config file
│
└───docs -> swagger docs
```

## Dependencies

- Go 1.22
- fiber 2.55
- gorm 1.25
- validator v10
- bytedance/sonic (json parser)
- go-redis v9
- mongodb
- mysql
- rabbitmq
- grpc 1.65
- swag (gen docs)
- zap (logger)
- wire (dependency injection)
- viper (config)
- prometheus

## Usage

Run docker-compose if you want to use the database and message broker

```bash
    docker-compose up -d
```

To build the service, run the following command in Makefile

```bash 
    make setup # to install dependencies
    make wire # to generate wire.go
    make build # to build the service
    make run # to run the service
```

## Endpoints

- GET /health - Check the health of the service
- GET /metrics - Get the metrics of the service
- GET /readiness - Check if the service is ready
- GET /liveness - Check if the service is alive
- GET /docs - Get the swagger documentation of the service


<hr>
<h4>Author: <a href="github.com/tdatIT">tdat.it</a></h4>