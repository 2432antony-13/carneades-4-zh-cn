package main

import (
	"flag"
	"github.com/carneades/carneades-4/src/web"
	"log"
	"os"
	"path/filepath"
)

const helpServer = `
usage: carneades server [-p port-number] [-t template-directory] [-l locale]

Starts the Carneades web service.

The -p flag ("port") specifies the port number to use (default: 8080)

The -t flag ("templates") specifies the full path to the directory with
the HTML templates used by the service 
(default: $GOPATH/src/github.com/carneades/carneades-4/src/web/templates/)

The -l flag ("locale") selects a localized template directory below the
template directory. Use "en" for the original templates or "zh-CN" for
Simplified Chinese or "zh-TW" for Traditional Chinese (default: en).
Users can also change the language in the browser.
`

func getGoPath() string {
	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		return os.Getenv("HOME") + "/go"
	} else {
		return goPath
	}
}

const defaultPort = "8080"

var defaultTemplatesDir = filepath.Join(getGoPath(), "/src/github.com/carneades/carneades-4/src/web/templates/")

func templatesDirForLocale(templatesDir string, locale string) string {
	if locale == "" || locale == "en" {
		return templatesDir
	}
	return filepath.Join(templatesDir, locale)
}

func webServerCmd() {
	webFlags := flag.NewFlagSet("webFlags", flag.ContinueOnError)
	port := webFlags.String("p", defaultPort, "the port number of the HTTP Server")
	templatesDir := webFlags.String("t", defaultTemplatesDir, "the path to the templates directory")
	locale := webFlags.String("l", "en", "the locale of the web interface")
	if err := webFlags.Parse(os.Args[2:]); err != nil {
		log.Fatal(err)
	}

	web.CarneadesServer(*port, templatesDirForLocale(*templatesDir, *locale))
}
