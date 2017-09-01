## Task

Write an HTTP service that exposes an endpoint "/numbers". This endpoint receives a list of URLs 
though GET query parameters. The parameter in use is called "u". It can appear 
more than once.

```
http://yourserver:8080/numbers?u=http://example.com/primes&u=http://foobar.com/fibo
```

When the /numbers is called, your service shall retrieve each of these URLs if 
they turn out to be syntactically valid URLs. Each URL will return a JSON data 
structure that looks like this:

```
{ "numbers": [ 1, 2, 3, 5, 8, 13 ] }
```

The JSON data structure will contain an object with a key named "numbers", and 
a value that is a list of integers. After retrieving each of these URLs, the 
service shall merge the integers coming from all URLs, sort them in ascending 
order, and make sure that each integer only appears once in the result. The 
endpoint shall then return a JSON data structure like in the example above with 
the result as the list of integers.

The endpoint needs to return the result as quickly as possible, but always 
within 500 milliseconds. It needs to be able to deal with error conditions when 
retrieving the URLs. If a URL takes too long to respond, it must be ignored. It 
is valid to return an empty list as result only if all URLs returned errors or 
took too long to respond.

## Usage

Install dependencies:

```
make depend
```

Start **Mock** and **Application** server:
```
make run-mock
```

Start **Application** server:

```
Application Options:
 --listen=         Listen address (default: :8080)
 --reader-timeout= Timeout to read numbers (default: 490)
```

```
make run-app
```

Application will print information about server address and timeout for reader:

```
INFO[0000] Listening on the address :8080      listen=:8080 reader-timeout=490
```

Run tests:

```
$ make test
 
go test ./...
?   github.com/mstovicek/showcase-go-api-number-aggregator/app [no test files]
ok  github.com/mstovicek/showcase-go-api-number-aggregator/app/Numbers 0.007s
?   github.com/mstovicek/showcase-go-api-number-aggregator/app/api [no test files]
?   github.com/mstovicek/showcase-go-api-number-aggregator/app/api/handler [no test files]
?   github.com/mstovicek/showcase-go-api-number-aggregator/app/api/middleware [no test files]
ok  github.com/mstovicek/showcase-go-api-number-aggregator/app/loader 0.015s
?   github.com/mstovicek/showcase-go-api-number-aggregator/mock [no test files]
```

### Mock

See output of the mock:

```
$ curl 'http://localhost:8090/ten/'
 
{"numbers":[10,9,8,7,6,5,4,3,2,1]}
```

Mock endpoints:
- `/ten/` - number 1..10
- `/even/` -  even number smaller than 20
- `/odd/` - odd numbers smaller than 20
- `/fibo/` - fibonacci numbers smaller than 20
- `/primes/` - prime numbers smaller than 20 
- `/wrong/` - wrong key in the output JSON

```
$ curl 'http://localhost:8090/wrong/'
 
{"wrong":[1,2,3,4,5]}
```

Mock output can be defined number of milisecond:

```
$ time curl 'http://localhost:8090/ten/?sleep_ms=1000'
 
{"numbers":[10,9,8,7,6,5,4,3,2,1]}
curl 'http://localhost:8090/ten/?sleep_ms=1000'  0.00s user 0.00s system 0% cpu 1.021 total
```

Mock log prints information about request:

```
INFO[0408] Sleep before response         ms=1000 path=/ten/
``` 

### Application

```
$ time curl 'http://localhost:8080/numbers/?u=http://localhost:8090/fibo/&u=http://localhost:8090/ten/?sleep_ms=500&u=foo'
 
{"numbers":[1,2,3,5,8,13]}
curl   0.00s user 0.00s system 1% cpu 0.510 total
```

Application log print information about request:

```
WARN[0654] Ignore invalid url             url=foo
INFO[0654] Loading numbers                url=http://localhost:8090/ten/?sleep_ms=500
INFO[0654] Loading numbers                url=http://localhost:8090/fibo/
INFO[0654] Numbers loaded                 numbers=[13 8 5 3 2 1] time=1.305926ms url=http://localhost:8090/fibo/
WARN[0654] Request failed                 error=Get http://localhost:8090/ten/?sleep_ms=500: net/http: request canceled (Client.Timeout exceeded while awaiting headers) timeout=490ms url=http://localhost:8090/ten/?sleep_ms=500
INFO[0654] Request has been processed     method=GET path=/numbers/ response_time=493.841605ms
``` 


## Learning

- theory behind pointer and non-pointer method receivers ([nathanleclaire.com](https://nathanleclaire.com/blog/2014/08/09/dont-get-bitten-by-pointer-vs-non-pointer-method-receivers-in-golang/))
- avoid lint warnings by defining exported interface as return type

```
exported func *** returns unexported type ***, which can be annoying to use
```

- Basics about testing ([blog.golang.org](https://blog.golang.org/examples), [blog.alexellis.io](https://blog.alexellis.io/golang-writing-unit-tests/))
- Acronyms are upper case in the function and variable names (regarding to linter and conventions)   
