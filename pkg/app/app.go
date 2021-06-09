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
	"github.com/honeybadger-io/honeybadger-go"
	"log"
	"net/http"
)

func ConfigureErrorReporting(honeybadgerApiKey string) (startMonitoring func()) {
	honeybadger.Configure(honeybadger.Configuration{APIKey: honeybadgerApiKey})
	return func() {
		honeybadger.Monitor()
	}
}

func ListenAndServe(port string, handler http.Handler) {
	log.Fatal(http.ListenAndServe(port, honeybadger.Handler(handler)))
}
