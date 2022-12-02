package interfacepollution

// Code is from this article: https://www.ardanlabs.com/blog/2016/10/avoid-interface-pollution.html

// ------------------------------------------------------------------------------------------------

// Here is the interface pollution smell list for the code :

// 1. The package declares an interface that matches the entire API of its own concrete type.
// 2. The factory function returns the interface value with the unexported concrete type value inside.
// 3. The interface can be removed and nothing changes for the user of the API.
// 4. The interface is not decoupling the API from change.

// ------------------------------------------------------------------------------------------------

// Server defines a contract for tcp servers.
type Server interface {
	Start() error
	Stop() error
	Wait() error
}

// server is our Server implementation.
type server struct {
	host string
}

// THIS IS A BAD CONSTRUCTOR:

// NewServer returns an interface value of type Server
// with an xServer implementation.
func NewServer(host string) Server {
	return &server{host}
}

// The next code listing shows how removing the interface changes nothing for the user:

// func NewServer(host string) *Server {
// 		return &Server{host}
// }
// Having the user work with the concrete type directly doesn’t change anything for the user or the API.

// Start allows the server to begin to accept requests.
func (s *server) Start() error {
	/* impl */
	return nil
}

// Stop shuts the server down.
func (s *server) Stop() error {
	/* impl */
	return nil
}

// Wait prevents the server from accepting new connections.
func (s *server) Wait() error {
	/* impl */
	return nil
}
