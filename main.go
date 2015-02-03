package hal_browser

import (
	"net/http"

	"github.com/cbandy/go-hal_browser/asset"
	"github.com/elazarl/go-bindata-assetfs"
)

var Handler = http.FileServer(&assetfs.AssetFS{
	Asset: asset.Asset, AssetDir: asset.AssetDir,
})

func At(path string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if path == r.URL.Path {
			r.URL.Path = "browser.html"
		}

		Handler.ServeHTTP(w, r)
	})
}
