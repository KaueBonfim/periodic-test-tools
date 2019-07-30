package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func (app *App) getTools(w http.ResponseWriter, r *http.Request) {

	var tool []Tool
	app.instance.Find(&tool)
	json.NewEncoder(w).Encode(tool)

}

func (app *App) getTool(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var tool Tool
	app.instance.Where("name = ?", vars["name"]).Find(&tool)
	if tool.ID == 0 {
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	} else {
		json.NewEncoder(w).Encode(tool)
	}
}

func (app *App) createTool(w http.ResponseWriter, r *http.Request) {
	var tool Tool
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&tool); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	dbc := app.instance.Create(&tool)
	if dbc.Error != nil {
		http.Error(w, dbc.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("OK"))
}

func (app *App) updateTool(w http.ResponseWriter, r *http.Request) {
	var last Tool
	var now Tool
	vars := mux.Vars(r)
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&now); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	app.instance.Where("name = ?", vars["name"]).Find(&last)
	dbc := app.instance.Model(&last).Updates(&now)
	if dbc.Error != nil {
		http.Error(w, dbc.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (app *App) deleteTool(w http.ResponseWriter, r *http.Request) {
	var tool Tool
	vars := mux.Vars(r)
	dbc := app.instance.Where("name = ?", vars["name"]).Delete(tool)
	if dbc.Error != nil {
		http.Error(w, dbc.Error.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}
