# Show your respect!

Show your respect to packages which you use in your projects

![](https://jokideo.com/wp-content/uploads/meme/2014/06/Reaction-Pic---My-respect.jpg)

*Installation*

```
go get github.com/iamsalnikov/go-respect
```

*Usage*

```
# you can run command inside your package
go-respect

# you can specify package directory. Here we will show respect
# to all dependencies of package
go-respect -dir=$GOPATH/src/github.com/vendor/package

# you can specify your own config file
# I store here your github username and token for working with API
go-respect -c=/path/to/your/custom/config/file.json
```

