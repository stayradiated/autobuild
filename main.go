package main

import (
	"fmt"
	"log"
	"net/http"

	"flag"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

var buildFlag = flag.Bool("build", false, "just build the site now")

func main() {
	flag.Parse()

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	autobuild := &AutoBuild{
		Dir: viper.GetString("dir"),
		Git: &Git{
			Remote: viper.GetString("git.remote"),
			Branch: viper.GetString("git.branch"),
		},
		Build: &Build{
			Dir:     viper.GetString("build.dir"),
			Command: viper.GetString("build.command"),
			Args:    viper.GetString("build.args"),
		},
		Webhook: &Webhook{
			Secret: viper.GetString("webhook.secret"),
		},
	}

	if *buildFlag {
		autobuild.Run()
		return
	}

	r := mux.NewRouter()
	r.Handle("/events", autobuild)

	port := viper.GetString("webhook.port")
	fmt.Println("Starting server on port", port)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
