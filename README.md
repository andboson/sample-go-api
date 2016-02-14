###sample-go-api

```
    cd $GOPATH
    
    mkdir -p src/app && cd src/app
    
    git clone git@github.com:andboson/sample-go-api.git .
    
    glide install
    
    go test -v $(glide novendor)
    
    go build -v    
```

####glide

Glide is golang package manager. It is very similar to Composer (PHP package manager). 

You must have Glide. See [github glide](https://github.com/Masterminds/glide) for howto

#### package manager
https://github.com/Masterminds/glide