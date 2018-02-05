# Show your respect!

Show your respect to packages which you use in your projects

[![Build Status](https://travis-ci.org/iamsalnikov/gorespect.svg?branch=master)](https://travis-ci.org/iamsalnikov/gorespect)
[![Go Report Card](https://goreportcard.com/badge/github.com/iamsalnikov/gorespect)](https://goreportcard.com/report/github.com/iamsalnikov/gorespect)
[![Exago](https://api.exago.io/badge/tests/github.com/iamsalnikov/gorespect)](https://exago.io/project/github.com/iamsalnikov/gorespect)
[![Exago](https://api.exago.io/badge/cov/github.com/iamsalnikov/gorespect)](https://exago.io/project/github.com/iamsalnikov/gorespect)

*Installation*

```
go get -u github.com/iamsalnikov/gorespect
```

*Usage*

```
# you can run command inside your package
gorespect

# you can specify package directory. Here we will show respect
# to all dependencies of package
gorespect -dir=$GOPATH/src/github.com/vendor/package

# you can specify your own config file
# I store here your github username and token for working with API
gorespect -c=/path/to/your/custom/config/file.json
```

Program will give a link for generating access token in first run.
Then it will use your token for starring repos which you use.

![](https://jokideo.com/wp-content/uploads/meme/2014/06/Reaction-Pic---My-respect.jpg)