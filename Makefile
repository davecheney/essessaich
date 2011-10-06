include $(GOROOT)/src/Make.inc

TARG=essessaich
GOFILES=\
	main.go\

DEPS=$(GOROOT)/src/pkg/exp/ssh

include $(GOROOT)/src/Make.cmd
