#!/bin/bash

docker run -p 8920:8920 -p 3000:3000 -v $GOPATH/src/dienlanhphongvan-cdn:/go/src/dienlanhphongvan-cdn -it imgx
