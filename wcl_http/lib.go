package wcl_http

import "net/http"

func ListenAndServeTLS(port, svrCrt, svrKey string, handler http.Handler) error {
	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
	server.SetKeepAlivesEnabled(false)
	return server.ListenAndServeTLS(svrCrt, svrKey)
}
