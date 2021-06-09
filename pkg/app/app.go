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
	"github.com/zerotohero-dev/fizz-env/pkg/env"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
	"net/http"
)

func ConfigureErrorReporting(
	honeybadgerApiKey string, deploymentType env.DeploymentType,
) (startMonitoring func()) {
	if deploymentType == env.Development {
		return func() {

		}
	}

	honeybadger.Configure(honeybadger.Configuration{
		APIKey: honeybadgerApiKey,
		Env: string(deploymentType),
	})

	return func() {
		honeybadger.Monitor()
	}
}

func Configure(deploymentType env.DeploymentType, appName string, honeybadgerApiKey string, sanitizeAppEnv func()) {
	sanitizeAppEnv()
	log.Init(appName)
	monitor := ConfigureErrorReporting(honeybadgerApiKey, deploymentType)
	defer monitor()
}

func Notify(str string) {
	_, _ = honeybadger.Notify(str)
}

func ListenAndServe(port string, handler http.Handler) {
	log.Fatal(http.ListenAndServe(port, honeybadger.Handler(handler)))
}
