package main

import (
	"encoding/json"
	"strconv"

	"github.com/YangzhenZhao/hulai-predict/storage"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

func main() {
	storage.InitData()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowMethods:     []string{"POST", "GET", "OPTIONS", "PUT", "DELETE", "UPDATE"},
		AllowAllOrigins:  true,
		AllowCredentials: true,
	}))
	router.GET("/", func(c *gin.Context) {
		res := generateRes()
		dumpsRes, _ := json.Marshal(res)
		c.JSON(200, string(dumpsRes))
	})
	router.POST("/upload-gaochou", uploadGaochou)
	router.POST("/upload-ange", uploadAnge)

	// router.Use(TlsHandler(8000))
	router.Run(":8000")
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
