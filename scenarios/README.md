# simplego - scenarios

this repository provides an example code of how to use the simplego app framework
in order to run a simple job code that provides the basics
* env configurations
* logs
* metrics

## contents
### Job
An application which processes some functionality and then suppose to shutdown

### Service
An application which runs until sigterm and provides a web server for health checks and API

### Messaging Consumer
An application which runs until sigterm and provides server for health checks and messaging consumer
