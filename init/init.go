package initialize

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"Society-Synergy/base/controllers"
	"Society-Synergy/base/routes"
	"Society-Synergy/base/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server *gin.Engine

	us services.ServiceUser
	ls services.ServiceLogs

	uc controllers.UserController
	lc controllers.LogsController

	userc    *mongo.Collection
	logc     *mongo.Collection
	clubc    *mongo.Collection
	memClubc *mongo.Collection
	adminc   *mongo.Collection
	// solc  *mongo.Collection
	// quec  *mongo.Collection

	rs routes.RouterService

	ctx         context.Context
	mongoclient *mongo.Client
	err         error
)

func InitializeSetup() {
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("err loading: %v", err)
	}

	ctx = context.TODO()

	mongoconn := options.Client().ApplyURI(os.Getenv("DATABASE"))
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal("error while connecting with mongo", err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal("error while trying to ping mongo", err)
	}

	fmt.Println("mongo connection established")

	logc = mongoclient.Database("Society-Synergy").Collection("auditlogs")
	userc = mongoclient.Database("Society-Synergy").Collection("users")
	clubc = mongoclient.Database("Society-Synergy").Collection("clubs")
	memClubc = mongoclient.Database("Society-Synergy").Collection("membersclubs")
	adminc = mongoclient.Database("Society-Synergy").Collection("adminsclubs")
	// solc = mongoclient.Database("Society-Synergy").Collection("solutions")
	// quec = mongoclient.Database("Society-Synergy").Collection("questions")

	us = services.NewServiceUser(userc, clubc, memClubc, adminc, ctx)
	ls = services.NewServiceLogs(logc, ctx)

	uc = controllers.NewUserController(us)
	lc = controllers.NewLogsController(ls)

	rs = routes.NewRouterService(uc, lc)

	server = gin.Default()
	// server.Use(cors.New(cors.Config{
	//     AllowOrigins:     []string{"*"},
	//     AllowMethods:     []string{"POST", "GET"},
	//     AllowCredentials: true,
	// }))
	// server.SetTrustedProxies(nil)
	// server.Use(cors.Default())
	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
}

func StartServer() {
	defer mongoclient.Disconnect(ctx)

	api := server.Group("/api")

	rs.RegisterRoutes(api)

	port := os.Getenv("PORT")

	log.Fatal(server.Run(":" + port))
}
