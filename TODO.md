# Distributed calculator TODO list

this .md file represents TODO list of DC project with list of tasks for each microservice and subtasks for each task

## Map

- [Microservices](#microservices)
    1. [Microservices/sso](#sso)
    2. [Microservices/orchestrator](#orchestrator)
    3. [Microservices/agent](#agent)

## Microservices

### SSO

- [ ] Proper authorization & authentication:
    - [ ] access & refresh tokens based on JWT
- [ ] Proper user management:
    - [ ] create, fetch, update, delete users & user's data
    - [ ] roles & permissions management
- [ ] gRPC API
- [ ] Saving data in Postgres
- [ ] Dockerization

### Orchestrator

- [ ] Create, fetch, update, delete calculation tasks, status & results
- [ ] Communication with agents via Kafka & gRPC:
    - [ ] Task distribution logic, i.e load balancer based on round-robin algorithm
    - [ ] Healthchecks for agents
    - [ ] Implement resilience patterns
- [ ] gRPC API
- [ ] Saving data in Postgres
- [ ] Dockerization

### Agent

- [ ] Listen to task queue topic(s) and handle grpc calls:
    - [ ] Calculation logic:
        - [ ] Proper error handling and logging in calculation process
        - [ ] Implement resilience patterns
        - [ ] Save task results in Redis:
            - [ ] Temporal data storage
            - [ ] Remove after successful result reporting
        - [ ] Result reporting
- [ ] Healthchecks for orchestrator
- [ ] Dockerization

#### Misc

- [ ] Redis Caching
- [ ] Proper logging, code commenting & error handling
- [ ] Unit tests
- [ ] CI/CD
- [ ] API Documentation
