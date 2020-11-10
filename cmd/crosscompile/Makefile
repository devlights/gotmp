all: \
	build \
	_display \
	_clean

build: \
	_build_linux_x86 \
	_build_linux_x64

_build_linux_x86:
	GOOS=linux GOARCH=386 go build -o linux_x86 main.go

_build_linux_x64:
	GOOS=linux GOARCH=amd64 go build -o linux_x64 main.go

_display:
	@echo '-------------------------------------------------------------'
	hexdump -C linux_x86 | head -n 1
	@echo '-------------------------------------------------------------'
	hexdump -C linux_x64 | head -n 1
	@echo '-------------------------------------------------------------'
	@echo 'ELF形式は先頭 5バイト 目の値が 0x01 なら32bit, 0x02 なら64bit'
	@echo 'ref: https://en.wikipedia.org/wiki/Executable_and_Linkable_Format#File_header'

_clean:
	@$(RM) linux_*
