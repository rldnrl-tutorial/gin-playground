# Gin Playground

## Gin이란?
Golang Web Framework입니다. httprouter보다 40초 빠릅니다.

## 사용 방법
### GET

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func main() {
    fmt.Println("Hello World")
    
    r := gin.Default()
    r.GET("/", HomePage)
    
    r.Run()
}
```

### GET (QueryStrings)

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
    fmt.Println("Hello World")
    
    r := gin.Default()
	r.GET("/query", QueryStrings)
    
    r.Run()
}
```

### GET (Path Parameter)

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func PathParameter(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/path/:name/:age", PathParameter)

	r.Run()
}
```

### POST

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func PostHomePage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.POST("/", PostHomePage)

	r.Run()
}
```