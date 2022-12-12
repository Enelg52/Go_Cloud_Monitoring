# Go_Cloud_Monitoring
J'espère qu'on trouvera un nom plus rigolo :)

## TODO
- [x] Faire le test pour (ListEc2)[/Go_Cloud_Monitoring/monitoring/list_ec2.go]

## Setup

### Golang version
last version (02.12.22)
```shell
$ go version
go version go1.19.1 windows/amd64
```
### JetBrains Goland version
last version (02.12.22)
help > about > version
```
GoLand 2022.2.5
Build #GO-224.4459.23, built on November 21, 2022
```
### aws sdk version
last version (02.12.22)
```shell
$ go get github.com/aws/aws-sdk-go
go: downloading github.com/aws/aws-sdk-go v1.44.151
go: added github.com/aws/aws-sdk-go v1.44.151
go: added github.com/jmespath/go-jmespath v0.4.0
```
### linux test environment
Je n'ai pas encore décidé comment j'allais tester mon code sur linux, mais j'ai plusieurs idées:
- utiliser un container docker
- utiliser un wsl2
- pipeline GitHub action

Mais normalement, le code qui tourne sur windows devrait tourner sur linux sans problème.
