// SPDX-FileCopyrightText: 2026 Jarkad <jarkad@jarkad.net.eu.org>
//
// SPDX-License-Identifier: EUPL-1.2

package handlers

import (
	"net/http"

	"github.com/a-h/templ"
)

// Render a given Templ component.
//
// If request has ?fragment=…, render only that fragment.
// Must call ParseForm() before calling this method.
func Render(w http.ResponseWriter, r *http.Request, component templ.Component) {
	var options []func(*templ.ComponentHandler)

	if fragments := r.Form["fragment"]; len(fragments) > 0 {
		fragmentsAny := make([]any, 0, len(fragments))
		for _, f := range fragments {
			fragmentsAny = append(fragmentsAny, f)
		}

		options = append(options, templ.WithFragments(fragmentsAny...))
	}

	templ.Handler(component, options...).ServeHTTP(w, r)
}
