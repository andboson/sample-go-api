###sample-go-api

```
    cd $GOPATH
    
    mkdir -p src/app cd src/app
    
    git clone git@github.com:andboson/sample-go-api.git .
    
    glide install
    
    go test -v $(glide novendor)
    
    go build -v    
```

####glide

Glide is package golang manager. It is very similar to Composer (PHP package manager). 

See [github glide](https://github.com/Masterminds/glide)

#### package manager
https://github.com/Masterminds/glide