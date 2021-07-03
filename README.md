# aws-golang-lambda

### This project demonstrates a basic golang application following the clean architecture principles and running on aws lambda 
### The project demonstrates the following:
1. Clean Architecture design.
2. Unit testing using golang's stdlib testing package.
3. Use of Golang's interfaces.
4. Use of Golang's Gin for http routing and middleware.
5. Use of AWS API Gateway for serving the API requests
6. Use of AWS custom lambda authorizer for working with JWT tokens.
7. Use of AWS SAM for quick and easy deployment.


### Architecture
![alt text](https://cdn-images-1.medium.com/max/800/1*7TOeidIPEOMR4Sns7Uvy8w.png)



### Installation:
There are two lambdas: an api lambda and a lambda custom authorizer lambda.
To deploy:

make sure to replace the S3 bucket in the Makefile
```
make build && make deploy
```


after deploying both lambdas, the final step is to deploy our API Gateway:
```
make deploy-apig
```


# API

### Sign up
```
curl --location --request POST 'https://<API ADDRESS>/deploy/users/signup' --headtion/json' --data-raw '{"username": "someuser", "password": "somepass", "address": "foobatz", "first_name": "foo", "last_name": "bar"}'
```

### Sign in
```
curl --location --request POST 'https://<API ADDRESS>/deploy/users/signin' --header 'Content-Type: application/json' --data-raw '{"username": "someuser", "password": "somepass"}'
```

### Say Hello
```
curl --location --request GET 'https://<API ADDRESS>/deploy/hello' --header 'Content-Type: application/json' --header 'Authorization: Bearer <TOKEN>'
```
