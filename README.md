# simplego

This repository aims to reduce boiler plate code to simplify writing 
Golang applications.
The project is a side hobby , PR's are welcome but please follow the developers guideline below.


The project is built as a mono repo multi module in order to enable quick plug and play of different interfaces for the personal usage.
Each module will contain teskit module in order to encourage dependency injection while building application and allowing code testability.

## contents
### App
The application module is a simple but the root of all modules.
The module aims to reduce code around Start and Stop sequences and supports injection
of other simplego modules

### Logger
The logger module is a simple asbstraction of a logger interface.
The module aims to provide consistent usage of a logger and allowes plug and play of different implemenations
aswell as providing a testable logger module

### Configs
The configs module is a simple lib to provide app configuraiton object injection.
The module aims to provide simple configuration objects usage in an app and allowes plug and play of different
configuration parsers for more complex solutions.
()


## roadmap
The current forseeable roadmap for the project 
- [x] logger module
- [x] default logger implementation - zerolog
- [x] logger testkit module
- [x] application struct and start sequence
- [x] application shutdown sequence
- [x] configuration module
- [x] default configuration injector implementation - go-envconfig
- [ ] metrics module
- [ ] default metrics implementation - prometheus
- [ ] metrics testkit module
- [ ] application start / stop sequences support metrics
- [ ] default application start scenarios - job
- [ ] server module
- [ ] default server implementation - TBD (chi probably)
- [ ] server testkit module
- [ ] application start / stop scenarios - server
- [ ] publisher / consumer modules
- [ ] default publisher / consumer implemenatation - nats
- [ ] publisher / consimer testkit modules
- [ ] application start / stop scenarios - publisher / consumer
- [ ] cache module
- [ ] default cache implementation - redis
- [ ] cache testkit module
- [ ] application start / stop sequences support cache

## wishlist
- [ ] dynamic config object update - vault impl

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

*Issues:*
Found an issue while using this library ? please open an Issue with the details 
1. Current library version
2. Logs if existing
3. Expected behaviour
4. Actual behaviour


*Feature Requests:*
In case u find a feature is missing, please dont hesitate to open a FR (Feature Request), while keeping ine mind the following
1. What is the main use case ?
2. How is it missing from the current impl ?
3. Please share a code snippet describing the recommended usage for the FR

*Roadmap Features:*
When you want to submit a new feature out of the roadmap. please take the following into considuraiton
1. Modules should be Tesatble
2. Tests should cover most of the cases to give assurance & example of usage
3. Modules should be isolated to reduce dependencies and thrive for more encapsulated code structure
4. Modules should provide a testkit