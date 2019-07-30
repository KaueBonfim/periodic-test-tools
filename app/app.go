package app

import (
	"log"
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

type App struct {
	instance *gorm.DB
	router   *mux.Router
}


func (app *App) Init(Type string) {
	var Uri string
	if Type == "mysql" {
		Uri = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("USER_BANK"), os.Getenv("PASS_BANK"), os.Getenv("HOST_BANK"), os.Getenv("PORT_BANK"), os.Getenv("DB_BANK"))
	} else if Type == "sqlite3" {
		Uri = os.Getenv("FILE_BANK")
	} else {
		panic("Error: Type bank is Invalid")
	}
	db, err := gorm.Open(Type, Uri)
	if err != nil {
		panic(err.Error())
	}

	app.instance = db
	app.router = mux.NewRouter()
}


func (app *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, app.router))
}

func (app *App) InitRouters() {
	app.router.HandleFunc("/tools", app.getTools).Methods("GET")
	app.router.HandleFunc("/tools/{name}", app.getTool).Methods("GET")
	app.router.HandleFunc("/tools", app.createTool).Methods("POST")
	app.router.HandleFunc("/tools/{name}", app.updateTool).Methods("PUT", "PATCH")
	app.router.HandleFunc("/tools/{name}", app.deleteTool).Methods("DELETE")
}
