/*
 * @Author: tobyn
 * @LastEditors: tobyn
 * @Description:
 */
package handlers

import "net/http"

// livez is a liveness probe.
func livez(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}
