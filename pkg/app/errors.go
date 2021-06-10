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
)

var canUseHoneybadger = false

func configureErrorReporting(e env.FizzEnv, honeybadgerApiKey string) (startMonitoring func()) {
	// Bypass honeybadger for development.
	if e.IsDevelopment() {
		return func() {

		}
	}

	honeybadger.Configure(honeybadger.Configuration{
		APIKey: honeybadgerApiKey,
		Env: string(e.Deployment.Type),
	})

	canUseHoneybadger = true

	return func() {
		honeybadger.Monitor()
	}
}

func Configure(
	e env.FizzEnv,
	appName string,
	honeybadgerApiKey string,
	sanitizeAppEnv func(),
) {
	sanitizeAppEnv()
	log.Init(e, appName)
	monitor := configureErrorReporting(e, honeybadgerApiKey)
	defer monitor()
}

func Notify(str string) {
	if !canUseHoneybadger {
		return
	}
	_, _ = honeybadger.Notify(str)
}
