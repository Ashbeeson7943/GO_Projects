# StatelessAPI; Calculator service
### Endpoints
### add
Command:
``` curl -d '{"num1":2, "num2":3}' http://localhost:8080/add/ ```  
Returns:
```
Calculation: 2 + 3
Answer: 5
```
### subtract
Command:
``` curl -d '{"num1":2, "num2":3}' http://localhost:8080/subtract/ ```  
Returns:
```
Calculation: 2 - 3
Answer: -1
```
### divide
``` curl -d '{"num1":2, "num2":3}' http://localhost:8080/divide/ ```  
Returns:
```
Calculation: 2 / 3
Answer: 0.6666666666666666
```
### multiply
``` curl -d '{"num1":2, "num2":3}' http://localhost:8080/multiply/ ```  
Returns:
```
Calculation: 2 x 3
Answer: 6
```

## ToDo: 
- Add in rate limiter to prevent misuse of the API
- Add in a database to keep track of all of the calculations that have taken place
- Add in a middleware that adds a request ID to the http.Request object.