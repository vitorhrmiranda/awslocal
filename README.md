# Localstack management with Serverless Framework

<h1 align="center">
    <img src="https://user-images.githubusercontent.com/76954948/136490297-35560c8a-d111-42eb-8d2a-0c88a6599f87.jpg" />
</h1>

This repository is a proof of concept (PoC) of using the `AWS Local` to manage the infrastructure dependencies for `AWS` in a development environment with `Localstack`

### Dependencies
- [Docker](https://www.docker.com/get-started)
- [AWS Local CLI](https://github.com/localstack/awscli-local)


### Install
```bash
docker-compose up -d
make setup
```

### Send message to SNS
```bash
make notify
```

## Metrics
### Read Dynamo
```bash
make scan
```

### Local Stack Logs
```bash
docker logs --tail 1000 -f awslocal
```

> Output
![image](https://user-images.githubusercontent.com/76954948/136490871-7b224b2b-8a7a-4db2-90fa-e2959c4fe22a.png)
