// file: web/middleware/basicauth.go

package middleware

import "github.com/8treenet/iris/v12/middleware/basicauth"

// BasicAuth middleware sample.
var BasicAuth = basicauth.New(basicauth.Config{
	Users: map[string]string{
		"admin": "password",
	},
})
