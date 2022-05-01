package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"io"
	"os"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	storage := "/var/local/"

	r := gin.New()
	r.Use(gin.Logger())
	r.GET("/", func(c *gin.Context) {
		c.String(200, "curl --upload-file ./example http://%s\n", c.Request.Host)
	})
	r.GET("/:file", func(c *gin.Context) {
		name := c.Param("file")
		c.File(storage + name)
		go func(name string) {
			time.AfterFunc(time.Minute, func() {
				_ = os.Remove(storage + name)
			})
		}(name)
	})
	r.PUT("/:file", func(c *gin.Context) {
		var b bytes.Buffer
		_, err := io.Copy(&b, c.Request.Body)
		if err != nil {
			c.String(200, "%s\n", err.Error())
			return
		}
		name := uuid.NewV4().String()
		file, err := os.OpenFile(storage+name, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
		if err != nil {
			c.String(200, "%s\n", err.Error())
			return
		}
		defer file.Close()
		_, err = b.WriteTo(file)
		if err != nil {
			c.String(200, "%s\n", err.Error())
			return
		}
		c.String(200, "http://%s/%s\n", c.Request.Host, name)
	})

	err := r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
