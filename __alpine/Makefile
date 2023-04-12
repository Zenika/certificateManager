all: 
	cd ../src
	go build . -o ./certificateManager.exe

prepare:
	sudo /opt/bin/install_golang.sh 1.19.7 amd64
	sudo apk update

install:
	test -d $(DESTDIR)$(PREFIX)/bin || mkdir -p $(DESTDIR)$(PREFIX)/bin
	install -pm 755 $(BUILDDIR)/certificateManager.exe $(DESTDIR)$(PREFIX)/bin/certificateManager
	strip $(DESTDIR)$(PREFIX)/bin/certificateManager
