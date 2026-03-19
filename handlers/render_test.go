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
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRender(t *testing.T) {
	t.Parallel()

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.Render(w, r, components.HelloWorld())
	}))
	defer ts.Close()

	req, err := http.NewRequestWithContext(t.Context(), "xxx", "asfd", nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "text/html; charset=utf-8", resp.Header.Get("Content-Type"))

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	err = resp.Body.Close()
	require.NoError(t, err)

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
