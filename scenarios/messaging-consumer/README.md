# simplego - scenario-messaging-consumer

this repository provides an example code of how to use the simplego app framework
in order to run a simple service process which ends on sigterm and consumes from config topics

## contents
* simplego app
* goenv-configs
* zerolog-logger
* prom-metrics
* chi-server
* messaging
* nats-messaging

### cmd
an exmaple code of running a service that runs with default health routes and metrics and consumes from a topic
/health
/ready
/metrics

### scenarios
api and models of providing the service scenario functionalities 