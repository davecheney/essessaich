include $(GOROOT)/src/Make.inc

TARG=essessaich
GOFILES=\
	keychain.go\
	main.go\

#DEPS=$(HOME)/devel/ssh
DEPS=$(GOROOT)/src/pkg/exp/ssh

include $(GOROOT)/src/Make.cmd
