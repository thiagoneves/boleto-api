package api

import (
	"fmt"

	"bitbucket.org/mundipagg/boletoapi/log"
	gin "gopkg.in/gin-gonic/gin.v1"
)

//RequestLogger faz o log no SEQ para toda request
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Ola")
		log.Info("Teste SEQ")
		c.Next()
	}
}

// ReturnHeaders 'seta' os headers padrões de resposta
func ReturnHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Info("Registrando Headers")
		c.Header("Content-Type", "application/json")
		c.Next()
	}
}
