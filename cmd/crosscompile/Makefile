all: \
	display \
	clean

build: \
	_build_linux_x86 \
	_build_linux_x64

display: \
	build \
	_display_od \
	_display_hexdump \
	_display_elf_format

_build_linux_x86:
	GOOS=linux GOARCH=386 go build -o linux_x86 main.go

_build_linux_x64:
	GOOS=linux GOARCH=amd64 go build -o linux_x64 main.go

_display_od:
	@echo ''
	@echo '[od     ]----------------------------------------------------'
	od -t x1 < linux_x86 2>/dev/null | head -n 1
	@echo '[od     ]----------------------------------------------------'
	od -t x1 < linux_x64 2>/dev/null | head -n 1
	@echo '[od     ]----------------------------------------------------'
	@echo ''

_display_hexdump:
	@echo ''
	@echo '[hexdump]----------------------------------------------------'
	hexdump -C < linux_x86 2>/dev/null | head -n 1
	@echo '[hexdump]----------------------------------------------------'
	hexdump -C < linux_x64 2>/dev/null | head -n 1
	@echo '[hexdump]----------------------------------------------------'
	@echo ''

_display_elf_format:
	@echo 'ELF形式は先頭 5バイト 目の値が 0x01 なら32bit, 0x02 なら64bit'
	@echo 'ref: https://en.wikipedia.org/wiki/Executable_and_Linkable_Format#File_header'


clean:
	@$(RM) linux_*
