package hal_browser

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestHandlerServesHALBrowser(t *testing.T) {
	for _, tt := range []struct{ path, mime string }{
		{"browser.html", "text/html"},
		{"styles.css", "text/css"},
		{"js/hal/browser.js", "application/javascript"},
		{"vendor/img/ajax-loader.gif", "image/gif"},
	} {
		r, _ := http.NewRequest("GET", tt.path, nil)
		w := httptest.NewRecorder()

		Handler.ServeHTTP(w, r)

		if w.Code != http.StatusOK {
			t.Errorf("Expected %s to return %v, got %v", tt.path, http.StatusOK, w.Code)
		}

		if content_type := w.HeaderMap["Content-Type"]; len(content_type) < 1 {
			t.Errorf("Expected %s to return a Content-Type, got %v", tt.path, w.HeaderMap)

		} else if !strings.HasPrefix(content_type[0], tt.mime) {
			t.Errorf("Expected %s to return %v, got %v", tt.path, tt.mime, content_type)
		}
	}
}

func TestHandlerReturnsNotFound(t *testing.T) {
	for _, path := range []string{
		"other.html",
		"nope.css",
	} {
		r, _ := http.NewRequest("GET", path, nil)
		w := httptest.NewRecorder()

		Handler.ServeHTTP(w, r)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected %s to return %v, got %v", path, http.StatusNotFound, w.Code)
		}
	}
}

func TestHandlerWorksWithStripPrefix(t *testing.T) {
	r, _ := http.NewRequest("GET", "browser.html", nil)
	expected := httptest.NewRecorder()

	Handler.ServeHTTP(expected, r)

	r, _ = http.NewRequest("GET", "configured/browser.html", nil)
	w := httptest.NewRecorder()

	http.StripPrefix("configured", Handler).ServeHTTP(w, r)

	if !reflect.DeepEqual(w, expected) {
		t.Errorf("Expected %v, got %v", expected, w)
	}
}

func TestAtReturnsBrowserHTML(t *testing.T) {
	r, _ := http.NewRequest("GET", "browser.html", nil)
	expected := httptest.NewRecorder()

	Handler.ServeHTTP(expected, r)

	r, _ = http.NewRequest("GET", "configured", nil)
	w := httptest.NewRecorder()

	At("configured").ServeHTTP(w, r)

	if !reflect.DeepEqual(w, expected) {
		t.Errorf("Expected %v, got %v", expected, w)
	}
}

func TestAtWorksWithStripPrefix(t *testing.T) {
	r, _ := http.NewRequest("GET", "browser.html", nil)
	expected := httptest.NewRecorder()

	Handler.ServeHTTP(expected, r)

	r, _ = http.NewRequest("GET", "configured/", nil)
	w := httptest.NewRecorder()

	http.StripPrefix("configured", At("/")).ServeHTTP(w, r)

	if !reflect.DeepEqual(w, expected) {
		t.Errorf("Expected %v, got %v", expected, w)
	}
}
