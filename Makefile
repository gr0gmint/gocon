# Copyright 2009 The Go Authors. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=main
GOFILES=\
	wanderer.go\
	broadcast.go\
	protohandlers.go\
	pwan.go\
	filters.go

default:
	8g $(GOFILES)
	8l wanderer.8

client: 
	8g wandererclient.go
	8l -o wan wandererclient.8 