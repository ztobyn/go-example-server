/*
 * @Author: tobyn
 * @LastEditors: tobyn
 * @Description:
 */
package handlers

import (
	"log"
	"sync/atomic"
	"time"

	"github.com/gorilla/mux"
)

// Router register necessary routes and returns an instance of a router.
func Router(buildTime, commit string) *mux.Router {
	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()

	r := mux.NewRouter()
	r.HandleFunc("/version", versionInfo(buildTime, commit)).Methods("GET")
	r.HandleFunc("/healthz", healthz)
	r.HandleFunc("/livez", livez)
	r.HandleFunc("/readyz", readyz(isReady))
	return r
}
