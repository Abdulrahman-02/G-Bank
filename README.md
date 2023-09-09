# G-Bank

This repository contains the codes of the [Backend Master Class](https://bit.ly/backendmaster) course.

This course is a step-by-step guide on how to design, develop, test and deploy a backend web service from scratch.

The backend service is a simple bank system that allows:

- Create and manage accounts and users.
- Record all balance changes to each account.
- Make transfers between accounts.

## Course Outline

### **Section 1**: Working with database [Postgres]

- Design the database schema
- Generate codes to talk to the DB
- Understand the DB isolation levels
- How to use Docker
- How to use Git and GitHub Actions to automate the CI/CD process

### **Section 2**: Building RESTful HTTP JSON API [Gin]

- Build a RESTful API using Gin
- Loading app configs
- Mocking DB for more robust unit tests
- Handling errors
- Authentication and authorization using JWT & PASETO

### **Section 3**: Deploying the application to production [Kubernetes + AWS]

- Build app with Docker and deploy to Kubernetes Cluster on AWS
- Setup free tier AWS account
- Create a production DB on AWS RDS
- Store and retrieve production secrets using AWS Secrets Manager
- Create a Kubernetes Cluster with EKS
- Use GitHub Actions to automatically build and deploy the image to the Kubernetes Cluster
- Buy a domain name and route traffics to the service
- Secure with HTTPS and auto-renew SSL certificates from Let's Encrypt

### **Section 4**: Advanced Backend Topics [Sessions + gRPC]

- Managing user sessions
- Building gRPC API
- Using gRPC Gateway to serve both gRPC and HTTP requests at the same time
- Embedding Swagger UI to document the API
- Updating a record using optional parameters
- Writing structured logger HTTP middlewares and gRPC interceptors

### **Section 5**: Asynchronous processing with background workers [Asynq + Redis]

- Asynchronous processing using background workers and Redis as its message queue
- Create and send emails to users via Gmail SMTP server
- Writing Unit Test for our gRPC services that might involve multiple dependencies at once

### **Section 6**: Improve the stability and security of the server

- Improve stability and security of the server
- Updating dependenby packages to the latest version
- Use Cookies to make the refresh token more secure
- How to gracefully shutdown the server to protect the processing ressources

## Docs

- Lecture #0: [Simple Bank](docs/0-%20A%20Simple%20Bank.md)
- Lecture #1: [Database Design](docs/1-%20DATABASE%20Design.md)
- Lecture #2: [Install & use Docker & Postgres](docs/2-%20Install%20&%20use%20Docker%20&%20Postgres.md)
- Lecture #3: [How to write & run DB migration in Golang](docs/3-%20How%20to%20write%20&%20run%20DB%20migration%20in%20Golang.md)
- Lecture #4: [Generate CRUD Golang code from SQL](docs/4-%20Generate%20CRUD%20Golang%20code%20from%20SQL.md)
- Lecture #5: [Unit Testing in Go db CRUD](docs/5-%20Unit%20Testing%20in%20Go%20db%20CRUD.md)
- Lecture #6: [DB Transaction in Golang](docs/6-%20DB%20Transaction%20in%20Golang.md)
- Lecture #7: [DB Transaction lock & how to handle deadlock](docs/7-%20DB%20Transaction%20lock%20&%20how%20to%20handle%20deadlock.md)
- Lecture #8: [How to avoid DB Deadlock](docs/8-%20How%20to%20avoid%20DB%20Deadlock.md)
- Lecture #9: [Transaction isolation level](docs/9-%20Transaction%20isolation%20level.md)
- Lecture #10: [Github actions](docs/10-%20Github%20actions.md)
- Lecture #11: [Implement RESTful API](docs/11-%20Implement%20RESTful%20API.md)
- Lecture #12: [Load Config from file & env vars](docs/12-%20Load%20Config%20from%20file%20&%20env%20vars.md)
- Lecture #13: [Mock DB for testing HTTP API](docs/13-%20Mock%20DB%20for%20testing%20HTTP%20API.md)
- Lecture #14: [Custom params validator](docs/14-%20Custom%20params%20validator.md)
- Lecture #15: [Add user table with unique & foreign key constraints](docs/15-%20Add%20user%20table%20with%20unique%20&%20foreign%20key%20constraints.md)
- Lecture #16: [handle db errors](docs/16-%20handle%20db%20erros.md)
- Lecture #17: [Hash Password](docs/17-%20Hash%20Password.md)
- Lecture #18: [Stronger Unit tests](docs/18-%20Stronger%20Unit%20tests.md)
- Lecture #19: [Why PASETO is better than JWT](docs/19-%20Why%20PASETO%20is%20better%20than%20JWT.md)
- Lecture #20: [Create & verify JWT & PASETO](docs/20-%20Create%20&%20verify%20JWT%20&%20PASETO.md)
- Lecture #21: [Implement login API with paseto & jwt](docs/21-%20Implement%20login%20API%20with%20paseto%20&%20jwt.md)
- Lecture #22: [How to implement authentication middleware](docs/22-%20How%20to%20implement%20authentication%20middleware.md)
- Lecture #23: [Build a minimal golang docker image](docs/22-%20Build%20a%20minimal%20golang%20docker%20image.md)
- Lecture #24: [How to connect containers in the same docker network](docs/24-%20How%20to%20connect%20containers%20in%20the%20same%20docker%20network.md)
- Lecture #25: [How to use Docker compose](docs/25-%20How%20to%20use%20Docker%20compose.md)
- Lecture #26: [How to create a free tier AWS account](docs/26-%20How%20to%20create%20a%20free%20tier%20AWS%20account.md)
- Lecture #27: [Build & push docker image to AWS ECR](docs/27-%20Build%20&%20push%20docker%20image%20to%20AWS%20ECR.md)
- Lecture #28: [How to create a production DB on AWS RDS](docs/28-%20How%20to%20create%20a%20production%20DB%20on%20AWS%20RDS.md)
- Lecture #29: [Store & retrieve production secrets with AWS secrets manager](docs/29-%20Store%20&%20retrieve%20production%20secrets%20with%20AWS%20secrets%20manager.md)
- Lecture #30: [How to create an AWS EKS Cluster](docs/30-%20How%20to%20create%20an%20AWS%20EKS%20Cluster.md)
- Lecture #31: [Kubectl & k9s](docs/31-%20Kubectl%20&%20k9s.md)
- Lecture #32: [Deploy a web app to Kubernetes on aws](docs/32-%20Deploy%20a%20web%20app%20to%20Kubernetes%20on%20aws.md)
- Lecture #33: [Register a domain & set up A-record using Route 53](docs/33-%20Register%20a%20domain%20&%20set%20up%20A-record%20using%20Route%2053.md)
- Lecture #34: [How to use Ingress to route traffics to different services in K8s](docs/34-%20How%20to%20use%20Ingress%20to%20route%20traffics%20to%20different%20services%20in%20K8s.md)
- Lecture #35: [Automatic issue TLS certificates in Kubernetes with Let's Encrypt](docs/35-%20Automatic%20issue%20TLS%20certificates%20in%20Kubernetes%20with%20Let's%20Encrypt.md)
- Lecture #36: [Automatic deploy to Kubernetes with Github action](docs/36-%20Automatic%20deploy%20to%20Kubernetes%20with%20Github%20action.md)
- Lecture #37: [how to manage user session with refresh token](docs/37-%20how%20to%20manage%20user%20session%20with%20refresh%20token.md)
- Lecture #38: [Generate DB docs](docs/38-%20Generate%20DB%20docs.md)
- Lecture #39: [Introduction to gRPC](docs/39-%20Introduction%20to%20gRPC.md)
- Lecture #40: [Define gRPC API and generate Go code with protobug](docs/40-%20Define%20gRPC%20API%20and%20generate%20Go%20code%20with%20protobug.md)
- Lecture #41: [How to run gRPC server](docs/41-%20How%20to%20run%20gRPC%20server.md)
- Lecture #42: [Implement gRPC API](docs/42-%20Implement%20gRPC%20API.md)
- Lecture #43: [Write code once, serve both gRPC & HTTP requests](docs/43-%20Write%20code%20once,%20serve%20both%20gRPC%20&%20HTTP%20requests.md)
- Lecture #44: [Meta data in gRPC](docs/44-%20Meta%20data%20in%20gRPC.md)
- Lecture #45: [Generate & serve Swagger with Go](docs/45-%20Generate%20%26%20serve%20Swaggre%20with%20Go.md)
- Lecture #46: [Embed static frontend files inside Golang backend](docs/46-%20Embed%20static%20frontend%20files%20inside%20Golang%20backend.md)
- Lecture #47: [gRPC params validation](docs/47-%20gRPC%20params%20validation.md)
- Lecture #48: [Run DB migration with Go](docs/48-%20Run%20DB%20migration%20with%20Go.md)
- Lecture #49: [Partial update & null params in go with sqlc](docs/49-%20Partial%20update%20%26%20null%20params%20in%20go%20with%20sqlc.md)
- Lecture #50: [Build gRPC update with optional parameters](docs/50-%20Build%20gRPC%20update%20with%20optional%20parameters.md)
- Lecture #51: [Add authorization to protect gRPC API](docs/51-%20Add%20authorization%20to%20protect%20gRPC%20API.md)
- Lecture #52: [Write structured logs for gRPC APIs](docs/52-%20Write%20structured%20logs%20for%20gRPC%20APIs.md)
- Lecture #53: [HTTP logger middleware](docs/53-%20HTTP%20logger%20middleware.md)
- Lecture #54: [Process async tasks with background workers](docs/54-%20Process%20async%20tasks%20with%20background%20workers.md)
- Lecture #55: [Integrate async worker to Go web server](docs/55-%20Integrate%20async%20worker%20to%20Go%20web%20server.md)
- Lecture #56: [Send async tasks to Redis within a DB transaction](docs/56-%20Send%20async%20tasks%20to%20Redis%20within%20a%20DB%20transaction.md)
- Lecture #57: [How to handle errors and print logs for Go Asynq workers](docs/57-%20How%20to%20handle%20errors%20and%20print%20logs%20for%20Go%20Asynq%20workers.md)
- Lecture #58: [The importance of delay in async tasks](docs/58-%20The%20importance%20of%20delay%20in%20async%20tasks.md)
- Lecture #59: [How to send emails in go](docs/59-%20How%20to%20send%20emails%20in%20go.md)
- Lecture #60: [How to skip test](docs/60-%20How%20to%20skip%20test.md)
- Lecture #61: [Email verification](docs/61-%20Email%20verification.md)
- Lecture #62: [Implement email verfication API](docs/62-%20Implement%20email%20verfication%20API.md)
