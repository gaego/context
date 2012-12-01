// Copyright 2012 AEGo Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Copyright 2012 The Gorilla Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine

package context

import (
	"appengine"
	"net/http"
)

type Context appengine.Context

// NewContext returns a new context for an in-flight HTTP request.
func NewContext(req *http.Request) appengine.Context {
	return appengine.NewContext(req)
}
