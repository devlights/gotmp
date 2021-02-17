default: installgo116

installgo116: \
	_prepare \
	_download \
	_extract \
	_updateenv

_prepare:
	rm -rf ~/download
	rm -rf ~/bin/go
	mkdir -p ~/download
	mkdir -p ~/bin

_download:
	wget --quiet -O ~/download/go116.tar.gz https://golang.org/dl/go1.16.linux-amd64.tar.gz

_extract:
	tar -C ~/bin/ -zxf ~/download/go116.tar.gz

_updateenv:
	echo 'export PATH=~/bin/go/bin:$$PATH' >> ~/.bashrc
	echo 'export GOROOT=~/bin/go' >> ~/.bashrc
