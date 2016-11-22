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

#### Some packages:

- [github.com/julienschmidt/httprouter](github.com/julienschmidt/httprouter) - very fast mux\router
- [github.com/smartystreets/goconvey](github.com/smartystreets/goconvey) - useful library for testing
- [github.com/jinzhu/gorm](github.com/jinzhu/gorm) - beautiful orm-library
- [github.com/andboson/configlog](github.com/andboson/configlog) - library for config and use logfile

### Sample docs:

### model.find_by_name

Return model api-response

Params:

    - `name`  Name of model.

####sample request:

**request:**

```javascript
{
	"name": "test-name"
}
```

**response:**

```javascript
{
    "id": 3,
    "article": 1007455,
    "name": "test-name",   			
    "date": "2016-05-22"
}

```
