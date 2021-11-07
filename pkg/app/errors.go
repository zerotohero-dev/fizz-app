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
	"github.com/honeybadger-io/honeybadger-go"
	"github.com/zerotohero-dev/fizz-logging/pkg/log"
)

var canUseHoneybadger = false

func configureErrorReporting(isDevEnv bool, deploymentType string, honeybadgerApiKey string) (startMonitoring func()) {
	// Bypass honeybadger for development.
	if isDevEnv {
		return func() {

		}
	}

	honeybadger.Configure(honeybadger.Configuration{
		APIKey: honeybadgerApiKey,
		Env:    string(deploymentType),
	})

	canUseHoneybadger = true

	return func() {
		honeybadger.Monitor()
	}
}

type ConfigureOptions struct {
	IsDevEnv          bool
	AppName           string
	DeploymentType    string
	HoneybadgerApiKey string
	LogDestination    string
	SanitizeFn        func()
}

func Configure(opts ConfigureOptions) {
	opts.SanitizeFn()
	log.Init(log.InitParams{
		IsDevEnv:       opts.IsDevEnv,
		LogDestination: opts.LogDestination,
		SanitizeFn:     opts.SanitizeFn,
		AppName:        opts.AppName,
	})

	monitor := configureErrorReporting(
		opts.IsDevEnv,
		opts.DeploymentType,
		opts.HoneybadgerApiKey,
	)
	defer monitor()
}

func Notify(str string) {
	if !canUseHoneybadger {
		return
	}
	_, _ = honeybadger.Notify(str)
}
