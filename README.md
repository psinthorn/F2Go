# F2Go
# F2.co.th Web Application development by Golang
>
> On our official website www.f2.co.th we use as a development learning website that we're start to build full web application by golang  following MVC microserivces structure. And we will try to write every details on the way of learning maybe it can benefit to someone that start to learing golang web development like us. Go for golang and enjoys with Go :) 

# Structure MVC Microservices

# Getting Start
## Start Server

## We're use Getenv to check and get current server port running by using function below.

```golang
    func (s *server) PortRunning() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8089"
	}
	return port
}
```
## Go Module Init
 ```
 go mod init github.com/your_github_usernams/your_repos

 ```
# Folders Structure
    -----
        |- app
        |- 
  

# Templates 
## How to set glbal templates

## Template Nested Template

# Database Section
## Connect to databse 

*** What is dufferrent about PUT and PATCH ***
- PUT needs all fields data input to update otherwise will update no data input fileds to blank.
- PATCH will update only a field that have data input and remain other fiedls data as the same.
