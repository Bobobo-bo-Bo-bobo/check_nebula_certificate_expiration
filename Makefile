GOPATH	= $(CURDIR)
BINDIR	= $(CURDIR)/bin

PROGRAMS = check_nebula_certificate_expiration

depend:
	env GOPATH=$(GOPATH) go get -u github.com/hako/durafmt
	env GOPATH=$(GOPATH) go get -u github.com/slackhq/nebula/cert

build:
	env GOPATH=$(GOPATH) go install $(PROGRAMS)

destdirs:
	mkdir -p -m 0755 $(DESTDIR)/usr/lib64/nagios/plugins

strip: build
	strip --strip-all $(BINDIR)/check_nebula_certificate_expiration

install: strip destdirs install-bin

install-bin:
	install -m 0755 $(BINDIR)/check_nebula_certificate_expiration $(DESTDIR)/usr/lib64/nagios/plugins

clean:
	/bin/rm -f bin/check_nebula_certificate_expiration

distclean: clean
	/bin/rm -rf src/github.com

uninstall:
	/bin/rm -f $(DESTDIR)/usr/lib64/nagios/plugins/check_nebula_certificate_expiration

all: build strip install

