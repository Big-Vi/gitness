// Copyright 2021 Harness Inc. All rights reserved.
// Use of this source code is governed by the Polyform Free Trial License
// that can be found in the LICENSE.md file for this repository.

package user

import (
	"net/http"

	"github.com/harness/gitness/internal/api/render"
	"github.com/harness/gitness/internal/api/render/platform"
	"github.com/harness/gitness/internal/api/request"
)

// HandleFind returns an http.HandlerFunc that writes json-encoded
// account information to the http response body.
func HandleFind() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user, _ := request.UserFrom(ctx)
		render.JSON(w, user, 200)
	}
}

// func returns an http.HandlerFunc that writes json-encoded
// account information to the http response body in platform
// format.
func HandleCurrent() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user, _ := request.UserFrom(ctx)
		platform.RenderResource(w, user, 200)
	}
}
