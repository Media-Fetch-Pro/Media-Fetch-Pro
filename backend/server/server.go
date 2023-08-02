package server

import (
	"net/http"
	"time"

	store "github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/store"
	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/store/db"
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

	go func() {
		s.Store.SchedulerDownload()
		time.Sleep(2 * time.Second)
	}()
	return s, nil
}

func (s *Server) Run(addr string) error {
	// such as ":7789"
	return s.httpServer.Run(addr)
}
