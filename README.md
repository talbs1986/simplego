# simplego

This repository aims to reduce boiler plate code to simplify writing 
Golang applications
The project is a side hobby , PR's are welcome but please follow the developers guideline below.

The project is built as a mono repo multi module in order to enable quick plug and play of different interfaces for the personal usage.
Each module will contain teskit module in order to encourage dependency injection while building application and allowing code testability.

## roadmap
The current forseeable roadmap for the project 
- [x] logger interface
- [x] default logger implementation - zerolog
- [ ] logger testkit
- [x] application struct and start sequence
- [x] application shutdown sequence
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

### developer guidelines
TBD