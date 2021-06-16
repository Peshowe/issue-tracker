# Parvus JIRA (demo project)

A demo app for tracking issues (the goal is to explore building apps using a microservice architecture and Go)

## App architecture

The app consists of a couple of services:
- Tracker service:
    - responsible for mainly the CRUD aspect of the app 
    - stores projects and their respective issues
    - publishes domain events
- Mail service:
    - sends email notifications to users
    - responsible for users' notification preferences
    - consumes domain events
- Gateway services:
    - exposes a RESTful-ish API to the outer world
    - serves the frontend client
    - responsible for edge functions - authentication (uses Google's oAuth2 server for identity provider) 
- Infrastructure services:
    - MongoDB
    - Kafka (with ZooKeeper)
    - SMTP server
- Frontend client:
    - a ReactJS frontend for the whole app
    - communicates with the gateway's REST API

### Repo structure

Each service is in its own sub-directory.

The [cmd](cmd) directory contains CLI scripts for: 
- compiling and copying over the frontend app 
- compiling and copying over the Proto files for the gRPC contract 
- building, tagging and pushing the Docker images

The [k8-configs](k8-configs) directory contains the Kubernetes config files. Alternatively there is also a [docker-compose.yml](docker-compose.yml) file.

The [grpc-contract](grpc-contract) directory contains the Proto files for the gRPC communication.


![Alt text](app_structure.png?raw=true "App structure")