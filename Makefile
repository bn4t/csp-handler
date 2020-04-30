SHELL = /bin/bash

# use bash strict mode
.SHELLFLAGS := -eu -o pipefail -c

.ONESHELL:
.DELETE_ON_ERROR:

# check if recipeprefix is supported
ifeq ($(origin .RECIPEPREFIX), undefined)
  $(error This Make does not support .RECIPEPREFIX. Please use GNU Make 4.0 or later)
endif
.RECIPEPREFIX = >

.SUFFIXES:      # delete the default suffixe
.SUFFIXES: .go  # add .go as suffix

PREFIX?=/usr/local
_INSTDIR=$(DESTDIR)$(PREFIX)
BINDIR?=$(_INSTDIR)/bin
GO?=go
GOFLAGS?=
RM?=rm -f # Exists in GNUMake but not in NetBSD make and others.


build:
> $(GO) build $(GOFLAGS) -o csp-handler .

run: build
> ./csp-handler

install: build
> useradd -MUr csp-handler
> install -m755 -gcsp-handler -ocsp-handler csp-handler $(BINDIR)/csp-handler
> mkdir -p /etc/csp-handler
> chown -R csp-handler:csp-handler /etc/csp-handler
> if [ ! -f "/etc/csp-handler/config.toml" ]; then
> 	install -m600 -gcsp-handler -ocsp-handler configs/config.example.toml /etc/csp-handler/config.toml
> fi

install-systemd:
> install -m644 -groot -oroot init/csp-handler.service /etc/systemd/system/csp-handler.service
> systemctl daemon-reload

clean:
> $(RM) csp-handler


RMDIR_IF_EMPTY:=sh -c '\
if test -d $$0 && ! ls -1qA $$0 | grep -q . ; then \
	rmdir $$0; \
fi'


uninstall:
> $(RM) $(BINDIR)/csp-handler
> $(RMDIR_IF_EMPTY) /etc/csp-handler

uninstall-systemd:
> $(RM) /etc/systemd/system/csp-handler.service
> systemctl daemon-reload

.DEFAULT_GOAL = build
.PHONY: all build install uninstall clean install-systemd

