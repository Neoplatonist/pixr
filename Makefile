DIR = $(shell pwd)
VERSION = $(shell grep -E 'Version = ' ./cmd/pixrsvc/cli.go | cut -f2 -d'"')

tag:
	@git tag -a $(VERSION)

devFront:
	@cd frontend; npm run dev

devBack:
	@realize start

build:
	@cd cmd/pixrsvc; go build -o ../../pixr
	@./pixr

buildFront:
	@cd frontend; npm run build

all: buildFront build
