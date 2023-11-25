package main

import (
	"encoding/json"
	"strconv"

	"github.com/YangzhenZhao/hulai-predict/storage"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	storage.InitData()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		res := generateRes()
		dumpsRes, _ := json.Marshal(res)
		c.JSON(200, string(dumpsRes))
	})

	router.Use(TlsHandler(8000))
	router.RunTLS(":8000", "./public.pem", "./private.key")
}

func TlsHandler(port int) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + strconv.Itoa(port),
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		// If there was an error, do not continue.
		if err != nil {
			return
		}
		c.Next()
	}
}
