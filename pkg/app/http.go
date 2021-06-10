/*
 *  \
 *  \\,
 *   \\\,^,.,,.                    “Zero to Hero”
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/honeybadger-io/honeybadger-go"
	"github.com/rs/cors"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"net/http"
)

func ListenAndServe(
	appName string,
	port string,
	deploymentType env.DeploymentType,
	handler http.Handler,
) {
	log.Info("🦄 Service '%s' will listen at port '%s'.", appName, port)

	Notify(fmt.Sprintf(
		"'%s' will listen at port '%s' on '%s'.", appName, port, deploymentType,
	))

	// Bypass Honeybadger for development.
	if isDevelopment(deploymentType) {
		log.Fatal(http.ListenAndServe(":"+port, handler))
		return
	}

	log.Fatal(http.ListenAndServe(":"+port, honeybadger.Handler(handler)))
}

func HandleCors(r *mux.Router) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"https://fizzbuzz.pro",
			"https://www.fizzbuzz.pro",
			"https://staging.fizzbuzz.pro",
			"https://staging.fizzbuzz.pro",
			// mapped in /etc/hosts to 127.0.0.1 for local development.
			"http://local.fizzbuzz.pro",
		},
		AllowCredentials: true,
		AllowedHeaders: []string{"Authorization", "Content-Type"},
		AllowedMethods: []string{"GET", "POST", "PUT", "HEAD", "OPTIONS"},
		// Debug: true,
	})

	return c.Handler(r)
}