#!/bin/sh
# Copyright 2011 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

set -e
gcc $(go env GOGCCFLAGS) -shared -o libcgotest.so cgotest_c.c
go build myf.go
export LD_LIBRARY_PATH=. 
./myf
#rm -f libcgosotest.so main
