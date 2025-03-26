package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

var availableColors = map[string]string{
	"red":    "#FF0000",
	"yellow": "#FFD700",
	"orange": "#FFA500",
	"lime":   "#00FF00",
	"green":  "#008000",
	"blue":   "#0000FF",
	"navy":   "#000080",
	"purple": "#800080",
	"pink":   "#FF00FF",
	"brown":  "#A52A2A",
	"grey":   "#808080",
	"black":  "#000000",
}

func checkEnvColor(c string) string {
	// Check if var is empty
	_, exists := os.LookupEnv("APP_COLOR")
	if !exists {
		log.Fatalln("APP_COLOR env var should not be empty !")
	}
	// Check if color is available
	var cValue string
	if value, t := availableColors[c]; t {
		cValue = value
	} else {
		colorKeys := []string{}
		for k := range availableColors {
			colorKeys = append(colorKeys, k)
		}
		log.Fatalf("Color not supported ! get=%s\navailable=%v", c, colorKeys)
	}
	return cValue
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	// Get color code if value is set and exist
	cValue := checkEnvColor(os.Getenv("APP_COLOR"))

	// Get hostname
	h, _ := os.Hostname()

	// Populate struct
	data := struct {
		Hostname string
		Color    string
	}{
		Hostname: h,
		Color:    cValue,
	}

	// Html template rendering
	t, err := template.ParseFiles("hello.html")
	if err != nil {
		log.Fatalln(err)
	}
	err = t.Execute(w, data)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// COnfigure application port if env var is set
	appPort, exists := os.LookupEnv("APP_PORT")
	if exists {
		appPort = ":" + appPort
	} else {
		appPort = ":8080"
	}

	// Check APP_COLOR env var and start webserver
	checkEnvColor(os.Getenv("APP_COLOR"))
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(appPort, nil))
}
