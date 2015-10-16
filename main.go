package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"flag"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

var buildFlag = flag.Bool("build", false, "just build the site now")
var repoFlag = flag.String("dir", "repo", "directory to clone repo to")

func main() {
	flag.Parse()

	viper.SetDefault("remote", "")
	viper.SetDefault("branch", "master")
	viper.SetDefault("build_script", "")
	viper.SetDefault("build_dir", "")

	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	autobuild := &AutoBuild{
		Git: &Git{
			Remote: viper.GetString("git.remote"),
			Branch: viper.GetString("git.branch"),
		},
		Build: &Build{
			Dir:     viper.GetString("build.dir"),
			Command: viper.GetString("build.command"),
			Args:    viper.GetString("build.args"),
		},
	}

	if *buildFlag {
		autobuild.Run(*repoFlag)
		return
	}

	webhook := &Webhook{
		HandlePush: func(branch string) {
			if branch == autobuild.Git.Branch {
				autobuild.Run(*repoFlag)
			}
		},
	}

	r := mux.NewRouter()
	r.Handle("/events", webhook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	fmt.Println("Starting server on port", port)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
