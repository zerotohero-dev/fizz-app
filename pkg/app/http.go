/*
 *  \
 *  \\,
 *   \\\,^,.,,.                     Zero to Hero
 *   ,;7~((\))`;;,,               <zerotohero.dev>
 *   ,(@') ;)`))\;;',    stay up to date, be curious: learn
 *    )  . ),((  ))\;,
 *   /;`,,/7),)) )) )\,,
 *  (& )`   (,((,((;( ))\,
 */

package app

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/honeybadger-io/honeybadger-go"
	"github.com/rs/cors"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"net/http"
	"reflect"
)

func ListenAndServe(
	e env.FizzEnv,
	appName string,
	port string,
	handler http.Handler,
) {
	log.Info("🦄 Service '%s' will listen at port '%s'.", appName, port)

	Notify(fmt.Sprintf(
		"'%s' will listen at port '%s' on '%s'.", appName, port, e.Deployment.Type,
	))

	// Bypass Honeybadger for development.
	if e.IsDevelopment() {
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

func ToErrorString(e interface{}) string {
	v := reflect.ValueOf(e)
	f := v.FieldByName("Err")
	return f.String()
}

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	responseErr := ToErrorString(response)

	if responseErr != "" {
		log.Err("EncodeResponse: error encoding response: %s", responseErr)

		res := reqres.GenericResponse{
			Err: "There is a problem in your request.",
		}

		w.WriteHeader(http.StatusBadRequest)

		return json.NewEncoder(w).Encode(res)
	}

	return json.NewEncoder(w).Encode(response)
}