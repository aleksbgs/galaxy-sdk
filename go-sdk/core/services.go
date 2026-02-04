package core

import "sync"

// ServiceRegistry keeps msg servers by route.
type ServiceRegistry struct {
	mu     sync.RWMutex
	routes map[string]MsgServer
}

func NewServiceRegistry() *ServiceRegistry {
	return &ServiceRegistry{routes: make(map[string]MsgServer)}
}

func (r *ServiceRegistry) Register(server MsgServer) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.routes[server.Route()] = server
}

func (r *ServiceRegistry) Get(route string) (MsgServer, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	server, ok := r.routes[route]
	return server, ok
}
