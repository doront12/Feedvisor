# Installation
  - clone the repository
  -  make sure  $GOPATH is configured
  - ** install easyjson from https://github.com/mailru/easyjson​
  - run ' go build'
  - you will find an executable file created after the dir name
  - run it as ./[dirname]

# Sending Requests
server is running under port 8080

  - Get all stored domains: GET localhost:8080
  - validate Domain + Path existance: POST localhost:8080/validate

payload example: 
{
	"domain": "www.google.com",
	"path": "/index"
}

# Notes
- ** easyJson compiled files are bundled in the repository, no need to compile the structs.
If you wish to compile run:  easyjson -all [StructName.go] for each struct in the 'api' dir for each struct

