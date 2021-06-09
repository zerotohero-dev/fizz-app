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

func Configure(
	appName string,
	deploymentType env.DeploymentType,
	honeybadgerApiKey string,
	sanitizeAppEnv func(),
) {
	sanitizeAppEnv()
	log.Init(appName)
	monitor := ConfigureErrorReporting(honeybadgerApiKey, deploymentType)
	defer monitor()
}

func Notify(str string) {
	_, _ = honeybadger.Notify(str)
}

func ListenAndServe(
	appName string,
	port string,
	deploymentType env.DeploymentType,
	handler http.Handler,
) {
	log.Info("🦄 Service '%s' will listen at port '%s'.", appName, port)

	Notify(fmt.Sprintf(
		"'%s' will listen at port '%s' on '%s'.", appName, port, deploymentType,
		),
	)

	log.Fatal(http.ListenAndServe(port, honeybadger.Handler(handler)))
}
