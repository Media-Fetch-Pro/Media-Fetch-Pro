package server

import (
	"net/http"

	store "github.com/CorrectRoadH/video-tools-for-nas/backend/store"
	"github.com/CorrectRoadH/video-tools-for-nas/backend/store/db"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *gin.Engine

	Store *store.Store
}

func NewServer() (*Server, error) {
	httpServer := gin.Default()

	s := &Server{
		httpServer: httpServer,
		// store:      store,
	}

	httpServer.NoRoute(gin.WrapH(http.FileServer(GetFileSystem("dist"))))

	db := db.NewDB()
	storeInstance := store.NewStore(db.DBInstance)
	s.Store = storeInstance
	s.Store.LoadGlobalVideoInfo()

	apiGroup := httpServer.Group("/api")
	s.registerVideoRoutes(apiGroup)
	s.registerSystemRoutes(apiGroup)

	return s, nil
}

func (s *Server) Run(addr string) error {
	// such as ":7789"
	return s.httpServer.Run(addr)
}
