# Learn Go with tests

Code from the tutorial

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world


# Lets work on this again
- [Buffer Stuff in DI Chapter](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection#write-the-minimal-amount-of-code-for-the-test-to-run-and-check-the-failing-test-output)
- [Context](https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/context) 
# Useful things

### Code Coverage

```
go test -cover
```

### Benchmarking
```
go test -bench=. 
```
[Race Detector](https://go.dev/blog/race-detector)
```
go test -race . 
```


### Dependency Management
network update for named packages
```
go get -u  github.com/schmiddim/learn-go-with-tests/16-math 
```


### Wire up a new project
```
mkdir blogposts
cd blogposts
go mod init github.com/{your-name}/blogposts
touch blogposts_test.go
```


### Todos
- Different Modules in same folder ... -> https://stackoverflow.com/a/67694381
- For that reason, it is recommended that you research The [Test Pyramid](https://quii.gitbook.io/learn-go-with-tests/build-an-application/http-server#integration-tests)
