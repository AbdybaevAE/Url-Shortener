# Url shortener service 
### Technologies:
- Golang 
- Grpc
- Postgres

### Featres:
- Concurrency safe
- Low latency
- Default <code>BASE62</code> algorithm
- Customizable algorithm
- Multiple algorithms support

### Services:
- Link service 
- Key Generation service
- Algorithm Factory
- Number service

### Link service allow you shorten links
### Key generation service pre generate unique keys, depending on algorithm 
### Algorithm factory register all algorithms 
### Number service allow you concurently generate increment value in database



