# Url shortener service 
### Technologies:
- Golang 
- Grpc
- Postgres

### Featres:
- Clean architecture
- Concurrency safe
- Low latency
- Default <code>BASE62</code> algorithm
- Customizable algorithm
- Multiple algorithms support

### Services:
- Link Service
- Key Generation Service
- Algorithm Factory
- Number Service


Link service allow you shorten links

Key generation service pre generate unique keys, depending on algorithm 

Algorithm factory register all algorithms 

Number service allow you concurently generate increment value in database



### Notes:
In base62 implementation every long url refference on unique integer, that converts to base62 (digits and alphabetical symbols). Key grows up to 62^n( 56 800 235 584 unique keys for 6 chars size)




