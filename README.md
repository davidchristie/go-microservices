# Go Microservices

## Getting Started

To run this example on your local machine:

1. Verify [Docker](https://www.docker.com/) is installed by running:

```console
docker --version
```

2. Clone the repository:

```console
git clone https://github.com/davidchristie/go-microservices.git
cd go-microservices
```

3. Build Docker images for each of the services:

```console
cd environments/production
docker-compose build
```

4. Start the services running:

```console
docker-compose up -d
```

5. You should now be able to open the GraphQL interface in browser by clicking [here](http://localhost:5000).

## Services

- [Accounts](services/accounts): an internal service for managing user accounts
- [Gateway](services/gateway): public GraphQL interface
