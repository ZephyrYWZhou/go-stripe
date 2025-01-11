package main

import "net/http"

func (app *application) VirtualTermianl(w http.ResponseWriter, r *http.Request) {
	app.infoLog.Println("Hit the Handler")
}
