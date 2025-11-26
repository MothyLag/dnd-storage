package main

import (
	"context"
	"dnd-storage/src/application/usecases"
	"dnd-storage/src/controllers"
	"dnd-storage/src/infrastructure/repository"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main(){
	mongoHOST := os.Getenv("MONGO_HOST")
	mongoDBName := os.Getenv("MONGO_DB_NAME")
	timeoutStr := os.Getenv("MONGO_TIMEOUT")
	mongoUser := os.Getenv("MONGO_USER")
    mongoPass := os.Getenv("MONGO_PASS")
	portStr := os.Getenv("APP_PORT")
    if mongoHOST == "" || mongoUser == "" || mongoPass == "" || mongoDBName == "" {
        log.Fatal("Debes definir MONGO_HOST, MONGO_USER, MONGO_PASS y MONGO_DB_NAME en variables de entorno")
    }

	timeout := 10 * time.Second
	if timeoutStr != ""{
		if t,err := time.ParseDuration(timeoutStr + "s"); err == nil{
			timeout = t
		}
	}
	port := "5000"
	if portStr != ""{
		port = portStr
	}
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s", mongoUser, mongoPass, mongoHOST)
	ctx, cancel := context.WithTimeout(context.Background(),timeout)
	defer cancel()

	client, err := mongo.Connect(ctx,options.Client().ApplyURI(mongoURI))	
	if err != nil{
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	db := client.Database(mongoDBName)
	userRepo := repository.NewUserMongoRepository(db)
	appRepo := repository.NewAppMongoRepository(db)

	createUser := usecases.NewCreateUser(userRepo)
	createApp := usecases.NewCreateApp(appRepo)

	userController := controllers.NewUserController(createUser)
	appController := controllers.NewAppController(createApp)
	r := gin.Default()
	//users
	userGroup := r.Group("user")
	userGroup.POST("/",userController.CreateUserHandler)
	//apps
	appsGroup := r.Group("app")
	appsGroup.POST("/",appController.CreateAppHandler)

	log.Println("Server running at http://localhost:",port)
	if err := r.Run(":"+port); err != nil{
		log.Fatal(err)
	}
}
