
asset/hal-browser.go: hal-browser/ hal-browser/.git
	go-bindata -o $@ -pkg=asset -prefix=$< -ignore=[.]git -ignore=[.]md $<...

hal-browser/.git:
	git submodule update --init hal-browser

clean:
	rm -f asset/hal-browser.go

setup:
	go get github.com/jteeuwen/go-bindata/...
	go get -d github.com/elazarl/go-bindata-assetfs

.PHONY: clean setup
