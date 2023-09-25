package api

import (
	db "com.wlq/simplebank/db/sqlc"
	"github.com/gin-gonic/gin"
)

// http请求服务
type Server struct {
	store  *db.Store
	router *gin.Engine
}

// 创建服务配置路由
func NewServer(store *db.Store) *Server {
	server := &Server{
		store: store,
	}
	router := gin.Default()

	// 路由
	router.POST("/accounts", server.createAccount)
	// url参数
	router.GET("/accounts/:id", server.getAccount)
	router.GET("/accounts", server.listAccount)

	server.router = router

	return server
}

// 启动
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

// 异常响应
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
