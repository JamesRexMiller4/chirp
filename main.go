package main

import (
	"net/http"
)

type Handler struct{}

func (Handler) ServeHTTP(h http.ResponseWriter, r *http.Request) {
	h.Header().Set("Content-Type", "text/plain; charset=utf-8")
	h.WriteHeader(http.StatusOK)
	h.Write([]byte("OK"))

}

// type Server struct {
// 	Addr                         string
// 	Handler                      http.ServeMux
// 	DisableGeneralOptionsHandler bool
// 	TLSConfig                    *tls.Config
// 	ReadTimeout                  time.Duration
// 	ReadHeaderTimeout            time.Duration
// 	WriteTimeout                 time.Duration
// 	IdleTimeout                  time.Duration
// 	MaxHeaderBytes               int
// 	TLSNextProto                 map[string]func(*Server, *tls.Conn, Handler)
// 	ConnState                    func(net.Conn, ConnState)
// 	ErrorLog                     *log.Logger
// 	BaseContext                  func(net.Listener) context.Context
// 	ConnContext                  func(ctx context.Context, c net.Conn) context.Context
// }

func main() {
	mux := http.NewServeMux()
	mux.Handle("/app/", http.StripPrefix("/app/", http.FileServer(http.Dir("."))))
	mux.Handle("/assets/", http.StripPrefix("/app/assets/", http.FileServer(http.Dir("/assets"))))

	handler := Handler{}

	mux.HandleFunc("/healthz", handler.ServeHTTP)

	server := http.Server{}
	server.Handler = mux
	server.Addr = ":8080"
	server.ListenAndServe()
}
