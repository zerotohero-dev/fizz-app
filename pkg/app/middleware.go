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
	"context"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/zerotohero-dev/fizz-entity/pkg/reqres"
	"net/http"
	"strings"
)

func ContentTypeValidatingMiddleware(
	f httptransport.DecodeRequestFunc,
) httptransport.DecodeRequestFunc {
	return func(c context.Context, r *http.Request) (interface{}, error) {
		contentType := strings.ToLower(r.Header.Get("content-type"))

		if strings.Index(contentType, "application/json") != 0 {
			return reqres.ContentTypeProblemRequest{
				Err: "contentTypeValidatingMiddleware: Invalid content type.",
			}, nil
		}

		return f(c, r)
	}
}
