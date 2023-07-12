# simplego

This repository aims to reduce boiler plate code to simplify writing 
Golang applications
The project is a side hobby , PR's are welcome but please follow the developers guideline below.

The project is built as a mono repo multi module in order to enable quick plug and play of different interfaces for the personal usage.
Each module will contain teskit module in order to encourage dependency injection while building application and allowing code testability.

## contents
### App
The application module is a simple but the root of all modules.
The module aims to reduce code around Start and Stop sequences and supports injection
of other simplego modules

### Logger
The logger module is a simple but the root of all modules.
The module aims to provide consistent usage of a logger and allowes plug and play of different implemenations
aswell as providing a testable logger module


## roadmap
The current forseeable roadmap for the project 
- [x] logger interface
- [x] default logger implementation - zerolog
- [ ] logger testkit
- [x] application struct and start sequence
- [x] application shutdown sequence
- [ ] metrics interface
- [ ] default metrics implementation - prometheus
- [ ] metrics testkit
- [ ] application start / stop sequences support metrics
- [ ] default application start scenarios - job
- [ ] server interface
- [ ] default server implementation - TBD (chi probably)
- [ ] server testkit
- [ ] application start / stop scenarios - server
- [ ] publisher / consumer interface
- [ ] default publisher / consumer implemenatation - nats
- [ ] publisher / consimer testkits
- [ ] application start / stop scenarios - publisher / consumer
- [ ] default cache implementation - redis
- [ ] cache testkit
- [ ] application start / stop sequences support cache


## contribution

### build
For convinience i provided a Makefile that simplifies actions for specific module.
An example of downloading module deps and building the logger module
```
make all DIR=logger
```

In the same format , is supported 
- lint
- build
- tidy
- test

### developer guidelines
PR's are welcomed , plz note im all alone at this, so patience please :)