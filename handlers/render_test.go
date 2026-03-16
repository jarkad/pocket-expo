package handlers_test

import (
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"git.jarkad.net.eu.org/jarkad/pocket-expo/components"
	"git.jarkad.net.eu.org/jarkad/pocket-expo/handlers"
	"github.com/gkampitakis/go-snaps/snaps"
)

func TestRender(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.Render(w, r, components.HelloWorld())
	}))
	defer ts.Close()

	req, err := http.NewRequestWithContext(t.Context(), http.MethodGet, ts.URL, nil)
	if err != nil {
		t.Fatal("cannot contact test server: ", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("cannot contact test server: ", err)
	}

	snaps.MatchInlineSnapshot(t,
		resp.StatusCode,
		snaps.Inline("int(200)"),
	)

	snaps.MatchInlineSnapshot(t,
		resp.Header.Get("Content-Type"),
		snaps.Inline("text/html; charset=utf-8"),
	)

	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()

	if err != nil {
		t.Fatal("cannot read body:", err)
	}

	snaps.WithConfig(snaps.Ext(".html")).
		MatchStandaloneSnapshot(t, string(body))
}

func TestMain(m *testing.M) {
	exitCode := m.Run()

	_, err := snaps.Clean(m)
	if err != nil {
		log.Println("could not clean snapshots:", err)
	}

	os.Exit(exitCode)
}
