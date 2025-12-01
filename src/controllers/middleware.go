package controllers

import (
	"dnd-storage/src/application/ports"
	"dnd-storage/src/domain/services"
	"net/http"
	"slices"
	"strings"

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

func JwtAuthMiddleware(jwtService *services.JWTService,jwtSecret string,allowedRoles []string) gin.HandlerFunc{
	return func(c *gin.Context){
		tokenStr := c.GetHeader("Authorization")
		if tokenStr == ""{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Authorization required"})
			c.Abort()
			return
		}
		tokenStr = strings.TrimSpace(strings.TrimPrefix(tokenStr,"Bearer"))

		claims,err := jwtService.ValidateToken(tokenStr,jwtSecret)

		if err != nil{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Invalid Token: "+tokenStr})
			c.Abort()
			return
		}
		userRole,ok := claims["role"].(string)
		if(!ok){
			c.JSON(http.StatusUnauthorized,gin.H{"error":"User not authorized"})
			c.Abort()
			return
		}
		authorized := slices.Contains(allowedRoles,userRole)
		if !authorized{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"User not authorized"})
			c.Abort()
			return
		}

		c.Set("userId",claims["id"])
		c.Set("userName",claims["username"])
		c.Set("userRole",claims["role"])

		c.Next()
	}
}
