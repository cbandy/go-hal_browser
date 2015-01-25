
hal-browser.go: hal-browser/ hal-browser/.git
	go-bindata -o $@ -prefix=$< -ignore=[.]git -ignore=[.]md $<...

hal-browser/.git:
	git submodule update --init hal-browser

clean:
	rm -f hal-browser.go

setup:
	go get github.com/jteeuwen/go-bindata/...

.PHONY: clean setup
