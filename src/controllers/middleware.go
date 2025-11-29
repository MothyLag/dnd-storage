package controllers

import (
	"dnd-storage/src/application/ports"
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(appRepo ports.AppRepository,keyService ports.KeyService,allowedRoles []string) gin.HandlerFunc{
	return func(c *gin.Context){
		apiKey := c.GetHeader("API-KEY")
		apiSecret := c.GetHeader("API-SECRET")

		if (apiKey == "" || apiSecret ==""){
			c.JSON(http.StatusUnauthorized,gin.H{"error":"App not Authorized"})
			c.Abort()
			return
		}

		appClient,err := appRepo.FindAppByApiKey(apiKey)	
		if err != nil{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid key pairs"})
			c.Abort()
			return
		}
		if keyService.ValidateKeyPair(apiSecret,appClient.AppSecret){
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid key pairs"})
			c.Abort()
			return
		}
		authorized := slices.Contains(allowedRoles,appClient.Rol)

		if !authorized {
			c.JSON(http.StatusForbidden,gin.H{"error":"Insufficient permissions"})
			c.Abort()
			return
		}

		c.Set("appClient",appClient)
		c.Next()
	}
}
