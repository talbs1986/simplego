[![Build generate push](https://github.com/talbs1986/simplego/actions/workflows/build.yml/badge.svg?branch=main)](https://github.com/talbs1986/simplego/actions/workflows/build.yml)
# simplego

This repository aims to reduce boiler plate code to simplify writing 
Golang applications. <br>
The project is a side hobby , PR's are welcome but please follow the developers guideline below.


The project is built as a mono repo multi module in order to enable quick plug and play of different interfaces for the personal usage.
Each module will contain teskit module in order to encourage dependency injection while building application and allowing code testability.

## contents
### App
The application module is the root object which will provide
the developer with the basic functionalities 
and will be extended by each simplego module 
 
### Scenarios
#### job
a scenario of a job that starts by initializing metrics pusher stops after the executed code

#### service
a scenario of a service that starts by initializing a server which will log, observe and handle incoming requests
the scenario will stop at sigterm

#### publisher
TBD

#### consumer
TBD

### Logger
The [logger module](app/pkg/logger) is a simple asbstraction of a logger interface.
The module aims to provide consistent usage of a logger of different impls
#### impls
* [FMT](app/pkg/fmt)
* [Zerolog](zerolog-logger)
* [TestKit](testkit-logger)

### Configs
The [configs module](configs) is a simple lib to provide app configuraiton object injection.
The module aims to provide simple configuration objects usage in an app of different impls
#### impls
* [GoEnv](goenv-configs)

### Metrics
The [metrics module](metrics) is a simple lib to provide app metrics.
The module aims to provide simple api and models to support a push / get metrics objects of different impls
#### impls
* [Prometheus](prom-metrics)
* [TestKit](testkit-metrics)

### Server
The [server module](server) is a simple lib to provide a web server
The module aims to provide simple api and models to registering routes and starting to listen for request by different impls
#### impls
* [Chi](chi-server)
* [TestKit](testkit-server)

### Trace
TBD

### Cache
TBD

### Publisher
TBD

### Consumer
TBD

## roadmap
The current forseeable roadmap for the project 
- [x] logger module
- [x] default logger implementation - zerolog
- [x] logger testkit module
- [x] application struct and start sequence
- [x] application shutdown sequence
- [x] configuration module
- [x] default configuration injector implementation - go-envconfig
- [x] metrics module
- [x] default metrics implementation - prometheus
- [x] metrics testkit module
- [x] application start / stop sequences support metrics
- [x] default application start scenarios - job
- [x] server module
- [x] default server implementation - chi
- [x] server testkit module
- [x] application start / stop scenarios - service
- [ ] publisher / consumer modules
- [ ] default publisher / consumer implemenatation - nats
- [ ] publisher / consimer testkit modules
- [ ] application start / stop scenarios - publisher / consumer
- [ ] cache module
- [ ] default cache implementation - redis
- [ ] cache testkit module
- [ ] application start / stop sequences support cache
- [ ] trace module
- [ ] default trace implementation - jager
- [ ] trace testkit module
- [ ] application start / stop sequences support trace

## wishlist
- [ ] dynamic config object update - vault impl
- [ ] server middlewares
  - [ ] auth
  - [ ] context
  - [ ] twirp

## contribution

### build
For convinience i provided a Makefile that simplifies actions for specific module.
An example of downloading module deps and building the logger module
```
make all DIR=logger
```

the following command will work on the same pattern
- lint
- build
- tidy
- test

## build all
The following command will run the above `make all` on every module
```
make dev_all
```



### developer guidelines
PR's are welcomed , plz note im all alone at this, so patience please :)

**Issues:** <br>
Found an issue while using this library ? please open an Issue with the details 
1. Current library version
2. Logs if existing
3. Expected behaviour
4. Actual behaviour


**Feature Requests:** <br>
In case u find a feature is missing, please dont hesitate to open a FR (Feature Request), while keeping ine mind the following
1. What is the main use case ?
2. How is it missing from the current impl ?
3. Please share a code snippet describing the recommended usage for the FR

**Roadmap Features:** <br>
When you want to submit a new feature out of the roadmap. please take the following into considuraiton
1. Modules should be Testable
2. Tests should cover most of the cases to give assurance & example of usage
3. Modules should be isolated to reduce dependencies and thrive for more encapsulated code structure
4. Modules should provide a testkit
