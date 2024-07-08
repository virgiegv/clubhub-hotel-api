# clubhub-hotel-api
Hotel franchises management project. 

## Building the project: 

Running the project with ***docker-compose***: 

```docker-compose up --build ```

The server will be accessible on localhost:8080. 

Alternatively, you can build it separately with these steps: 
1. Install postgresql 16
2. Create database 'clubhub'
3. Replace database credentials in .env with yours
4. Install go 1.22.4
5. Set GOPATH to go installation folder
6. Run ```go mod download ```
7. Install Swagger: ```go install github.com/swaggo/swag/cmd/swag@latest```
8. Add go/bin to PATH variable
9. Run ```swag init```
10. Run ```go build main```

### Debug mode 
If you run the project with docker-compose, I've left commented in the Dockerfile some instructions to install Delve debugger and run it in port 2345. If needed, a remote debugger can be configured to connect to this instance. 

### Swagger

Access the Swagger once the server is running: https://localhost:8080/clubhub/api/v1/swagger/index.html

## Documentation

The documentation for the project can be found in this repository's Wiki: https://github.com/virgiegv/clubhub-hotel-api/wiki

