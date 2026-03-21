// SPDX-FileCopyrightText: 2026 Jarkad <jarkad@jarkad.net.eu.org>
//
// SPDX-License-Identifier: EUPL-1.2

package handlers

import (
	"context"
	"log"
	"net/http"

	"git.jarkad.net.eu.org/jarkad/pocket-expo/components"
)

// CounterService is expected to maintain a global counter.
//
// Get() should retrieve its value.
// Inc() and Dec() should increment or decrement it respectively.
type CounterService interface {
	Get(ctx context.Context) (int, error)
	Inc(ctx context.Context) error
	Dec(ctx context.Context) error
}

// HomepageHandler handles the homepage.
type HomepageHandler struct {
	Log     *log.Logger
	Counter CounterService
}

// NewHomepageHandler creates a new homepageHandler.
func NewHomepageHandler(log *log.Logger, counter CounterService) *HomepageHandler {
	return &HomepageHandler{
		Log:     log,
		Counter: counter,
	}
}

// MaxHomepageRequestSize specifies maximum request body size for the homepage.
const MaxHomepageRequestSize = 1 << 20 // 1MB

func (h *HomepageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, MaxHomepageRequestSize)

	defer func() {
		err := r.Body.Close()
		if err != nil {
			h.Log.Println("while handling homepage: while closing body:", err)
		}
	}()

	var errorMessage string

	switch r.FormValue("action") {
	case "": // no action
	case "+":
		err := h.Counter.Inc(r.Context())
		if err != nil {
			h.Log.Println("cannot increment count:", err)

			errorMessage = "could not increment"
		}
	case "-":
		err := h.Counter.Dec(r.Context())
		if err != nil {
			h.Log.Println("cannot decrement count:", err)

			errorMessage = "could not decrement"
		}
	default:
		h.Log.Println("unknown action:", r.FormValue("action"))

		errorMessage = "unknown action (oops!)"
	}

	count, err := h.Counter.Get(r.Context())
	if err != nil {
		h.Log.Println("cannot get count:", err)

		errorMessage = "could not query database"
		count = 42
	}

	view := components.Homepage(count, errorMessage)
	Render(w, r, view)
}
