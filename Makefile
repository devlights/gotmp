default: install

install: prepare download extract
	'echo 'export PATH=~/bin/go:$$PATH' >> ~/.bashrc
	'echo 'export GOROOT=~/bin/go' >> ~/.bashrc

prepare:
	rm -rf ~/download
	mkdir -p ~/download
	mkdir -p ~/bin

download:
	wget -O ~/download/go116.tar.gz https://golang.org/dl/go1.16.linux-amd64.tar.gz

extract:
	tar -C ~/bin/ -zxf ~/download/go116.tar.gz