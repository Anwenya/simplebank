package api

import (
	"fmt"

	"com.wlq/simplebank/token"
	"com.wlq/simplebank/util"

	db "com.wlq/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// http请求服务
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// 创建服务配置路由
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker:%w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// 绑定自定义参数校验标签
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	// 路由
	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	// 默认路由
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)
	router.POST("/token/refresh", server.refreshAccessToken)

	// 添加认证中间件的路由组
	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount)
	authRoutes.GET("/accounts", server.listAccount)
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// 启动
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// 异常响应
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
