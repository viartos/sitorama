package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/art-sitedesign/sitorama/app/handlers"
	"github.com/art-sitedesign/sitorama/app/utils"
)

const (
	appDefaultPort = 8085
)

//todo: годоки дописать
//todo: логи сервисов прокидывать наружу или в stdout?

func main() {
	tmpl := template.Must(template.ParseFiles(
		"app/templates/html/index.html",
		"app/templates/html/project/create.html",
		"app/templates/html/project/confirm.html",
		"app/templates/html/settings-app.html",
		"app/templates/html/error.html",
	))

	http.HandleFunc("/", handlers.Index(tmpl))
	http.HandleFunc("/init", handlers.Init(tmpl))

	// экшены найтроек
	http.HandleFunc("/settings/app", handlers.SettingsApp(tmpl))

	// экшены проектов
	http.HandleFunc("/project/create", handlers.ProjectCreate(tmpl))
	http.HandleFunc("/project/create-confirm", handlers.ProjectCreateConfirm(tmpl))
	http.HandleFunc("/project/start", handlers.ProjectStart(tmpl))
	http.HandleFunc("/project/stop", handlers.ProjectStop(tmpl))
	http.HandleFunc("/project/remove", handlers.ProjectRemove(tmpl))

	// экшены контейнеров
	http.HandleFunc("/container/restart", handlers.ContainerRestart(tmpl))
	http.HandleFunc("/container/stop", handlers.ContainerStop(tmpl))
	http.HandleFunc("/container/start", handlers.ContainerStart(tmpl))
	http.HandleFunc("/container/remove", handlers.ContainerRemove(tmpl))

	appPort := utils.FindNearPort(appDefaultPort)

	fmt.Printf("Open GUI: http://127.0.0.1:%d\n", appPort)
	err := http.ListenAndServe(fmt.Sprintf(":%d", appPort), nil)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return
}
