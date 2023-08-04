package server

import (
	"net/http"
	"time"

	"github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/server/profile"
	store "github.com/Media-Fetch-Pro/Media-Fetch-Pro/backend/store"
	"github.com/gin-gonic/gin"
)

type Server struct {
	httpServer *gin.Engine
	Profile    *profile.Profile
	Store      *store.Store
}

func NewServer(profile *profile.Profile, store *store.Store) (*Server, error) {
	httpServer := gin.Default()

	s := &Server{
		httpServer: httpServer,
		Profile:    profile,
		Store:      store,
	}

	httpServer.NoRoute(gin.WrapH(http.FileServer(GetFileSystem("dist"))))

	s.Store.LoadGlobalVideoInfo()

	apiGroup := httpServer.Group("/api")
	s.registerVideoRoutes(apiGroup)
	s.registerSystemRoutes(apiGroup)
	s.registerSettingsRoutes(apiGroup)

	go func() {
		for {
			s.Store.SchedulerDownload()
			time.Sleep(2 * time.Second)
		}
	}()
	return s, nil
}

func (s *Server) Run(addr string) error {
	// such as ":7789"
	return s.httpServer.Run(addr)
}
