package controller

import "net/http"

func (c *Controller) Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Application Running!"))
}
