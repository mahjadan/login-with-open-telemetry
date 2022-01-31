### OpenTelemetry demo:
using two go apps we will try to demonstrate how to inject tracing for all the routes using middleware and for communication between services using httpclient

### How to Run:
`docker-compose up`

this will run both the apps with a jaeger container to receive the traces and visualize them

### Diagram:



### How to Test:
`curl http://localhost:8081/register_test`

`curl http://localhost:8081/login_test`

