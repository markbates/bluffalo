#!/bin/bash

set -ex
go install ./cmd/bluffalo

cd coke

bluffalo -h

bluffalo fix -h
bluffalo fix
bluffalo fix pop
bluffalo fix plush

bluffalo generate -h
bluffalo generate goth facebook twitter
