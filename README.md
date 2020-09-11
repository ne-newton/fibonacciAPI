# fibonacciAPI

API steps through the Fibonacci sequence.  
It runs on `localhost:8080` with three endpoints:
```
\current
\next
\previous
```

`\current` returns the current number in Fibonacci sequence, beginning on 0  
`\next` returns the next number in the Fibonacci sequence, after 1, then 1, 2, 3, 5 ... and sets current to that number  
`\previous` moves back throught the sequence, setting current to that number  
each endpoint returns it's name plus the number from the sequence, eg `current -> 0`  

fibonacci.go contains the fibonnaci struct and methods  
api.go contains the endpoint functions for the server

## Dependencies
Aside from the standard Go librarie, the API is built using the [gin-gonic](http://github.com/gin-gonic/gin) framework, and uses [endless](github.com/fvbock/endless) for graceful restarts.  

These can be retreived by running `go get github.com/gin-gonic/gin` and `go get github.com/fvbock/endless`

## Running

A macOS precompiled binary is available, which can be run `./fibonaccAPI` from the same directory. Then browse to `localhost:8080\{endpoint}` to use.  

Otherwise, the api can be built through `go build` from the project directory, or `go run main.go fibonacci.go api.go` to run without building
