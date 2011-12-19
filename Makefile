include $(GOROOT)/src/Make.inc

TARG=gsh
GOFILES=\
	keychain.go\
	main.go\

#DEPS=$(HOME)/devel/ssh
DEPS=$(GOROOT)/src/pkg/exp/ssh

include $(GOROOT)/src/Make.cmd
