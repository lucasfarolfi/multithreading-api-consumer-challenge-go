# multithreading-api-consumer-challenge-go
 This is a challenge from GoExpert course.

## Challenge statement
In this challenge you must to use that we have learned of Multithreading and APIs to search the fastest result between two distinct APIs.

Both requests will be done simultaneously for the following APIs:

"https://brasilapi.com.br/api/cep/v1/{cep}"
"http://viacep.com.br/ws/{cep}/json/"

The requirements for this challenge are:
- Accept the API that delivers the fastest response and discard the slowest response.
- The request result should be printed on command line with the response body from API.
- Limit response time in 1 second. Otherwise, the timeout error should be printed.

## How to run?
Just open the terminal in the project directory and run the command ```go run main.go```.