# fibonacciAPI

API steps through the Fibonacci sequence.  
It runs on `localhost:8080` with three endpoints:
```
\current
\next
\previous
```

`\current` returns the current number in Fibonacci sequence, beginning on 0  
`\next` returns the next number in the Fibonacci sequence, after 1, then 1, 2, 3, 5 ... and sets current to that number. Will return an error if int overflows, and will not move forward to next number in sequence.  
`\previous` moves back throught the sequence, setting current to that number. If trying to go to a number previous to 0, it will return an error.  
Each endpoint returns it's name plus the number from the sequence, eg `current -> 0`, `next -> 1`.

## files
main.go runs the server  
fibonacci.go contains the fibonnaci struct and methods  
api.go contains the endpoint functions for the server  
db.go sets up the database and loads any previous fibonacci position

## Dependencies
Aside from the standard Go libraries, the API is built using the [gin-gonic](http://github.com/gin-gonic/gin) framework, uses [endless](github.com/fvbock/endless) for graceful restarts, and [gorm](http://github.com/jinzhu/gorm) as the database ORM.

These can be retreived by running `go get github.com/gin-gonic/gin`, `go get github.com/fvbock/endless` and `go get -u github.com/jinzhu/gorm`

## Running

A macOS precompiled binary is available, which can be run `./fibonaccAPI` from the same directory. Then navigating to `localhost:8080\{endpoint}` to use.  

Otherwise, the api can be built through `go build` from the project directory, or `go run main.go fibonacci.go api.go db.go` to run without building
