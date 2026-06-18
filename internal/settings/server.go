package settings

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"sync"
)

type SettingsServer struct {
	logger  *slog.Logger
	server  *http.Server
	port    int
	mux     *http.ServeMux
	service *Service
	started bool
	mu      sync.Mutex
}

func NewServer(logger *slog.Logger, service *Service) *SettingsServer {
	mux := http.NewServeMux()
	return &SettingsServer{
		logger:  logger,
		mux:     mux,
		service: service,
	}
}

func (s *SettingsServer) Start() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.started {
		return fmt.Sprintf("http://127.0.0.1:%d", s.port), nil
	}

	port, err := findFreePort()
	if err != nil {
		return "", fmt.Errorf("find free port: %w", err)
	}
	s.port = port

	s.registerRoutes()

	addr := fmt.Sprintf("127.0.0.1:%d", port)
	s.server = &http.Server{
		Addr:    addr,
		Handler: s.mux,
	}

	go func() {
		s.logger.Info("settings server started", "port", port)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.logger.Error("settings server error", "error", err)
		}
	}()

	s.started = true
	return fmt.Sprintf("http://127.0.0.1:%d", port), nil
}

func (s *SettingsServer) Stop() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.server != nil {
		s.server.Shutdown(context.Background())
		s.logger.Info("settings server stopped")
		s.started = false
	}
}

func (s *SettingsServer) registerRoutes() {
	s.mux.HandleFunc("/", s.handleUI)
	s.mux.HandleFunc("/api/settings", s.handleSettings)
	s.mux.HandleFunc("/api/cameras", s.handleCameras)
	s.mux.HandleFunc("/api/mappings", s.handleMappings)
}

func findFreePort() (int, error) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer ln.Close()
	return ln.Addr().(*net.TCPAddr).Port, nil
}
