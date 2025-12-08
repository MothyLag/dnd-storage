package main

import (
	"context"
	"dnd-storage/src/application/usecases"
	"dnd-storage/src/controllers"
	"dnd-storage/src/domain/services"
	"dnd-storage/src/infrastructure/repository"
	repoChar"dnd-storage/src/infrastructure/repository/character"
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
	jwtSecret := os.Getenv("JWT_SECRET")
	portStr := os.Getenv("API_PORT")
    if mongoHOST == "" || mongoUser == "" || mongoPass == "" || mongoDBName == "" || jwtSecret == "" {
        log.Fatal("Debes definir JWT_SECRET,MONGO_HOST, MONGO_USER, MONGO_PASS y MONGO_DB_NAME en variables de entorno")
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
	//repos
	db := client.Database(mongoDBName)
	userRepo := repository.NewUserMongoRepository(db)
	appRepo := repository.NewAppMongoRepository(db)
	characterRepo := repoChar.NewCharacterMongoRepository(db)
	//services
	userService := services.NewUserService()
	appService := services.NewAppService()
	keyService := services.NewKeyService()
	jwtService := services.NewJWTService()
	characterService := services.NewCharacterService()
	//usecases
	createUser := usecases.NewCreateUser(userRepo,userService)
	loginUser := usecases.NewLoginUser(jwtSecret,userRepo,userService,keyService,jwtService)
	createApp := usecases.NewCreateApp(appRepo,appService,keyService)
	updateApp := usecases.NewUpdateApp(appRepo,appService)
	createChar := usecases.NewCharacterUseCase(characterRepo,characterService)
	lookCharByUser := usecases.NewLookCharacterByUser(characterRepo)
	lookChar := usecases.NewLookCharacter(characterRepo)
	//controllers
	userController := controllers.NewUserController(createUser,loginUser)
	appController := controllers.NewAppController(createApp,updateApp)
	characterController := controllers.NewCharacterController(createChar,lookCharByUser,lookChar)
	r := gin.Default()
	//users
	userGroup := r.Group("user")
	userGroup.Use(controllers.AuthMiddleware(appRepo,keyService,[]string{"system","admin"}))
	userGroup.POST("/",userController.CreateUserHandler)
	userGroup.POST("/login",userController.LoginUserHandler)
	//apps
	appsGroup := r.Group("app")
	appsGroup.Use(controllers.AuthMiddleware(appRepo,keyService,[]string{"system"}))
	appsGroup.POST("/",appController.CreateAppHandler)
	appsGroup.PUT("/:id",appController.UpdateAppHandler)
	//characters
	charGroup := r.Group("character")
	charGroup.Use(controllers.AuthMiddleware(appRepo,keyService,[]string{"system","admin"}),
		controllers.JwtAuthMiddleware(jwtService,jwtSecret,[]string{"admin","dm","player"}))
	charGroup.POST("/",characterController.PostCharacter)
	charGroup.GET("/",characterController.GetCharactersByUserId)
	charGroup.GET("/:id",characterController.GetCharacterById)
	log.Println("Server running at http://localhost:",port)
	if err := r.Run(":"+port); err != nil{
		log.Fatal(err)
	}
}
