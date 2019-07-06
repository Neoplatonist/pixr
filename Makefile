DIR = $(shell pwd)

devFront:
	@cd frontend; npm run dev

devBack:
	@realize start

build:
	@cd server; packr2
	@cd cmd/pixrsvc; go build -o ../../pixr
	@./pixr

buildFront:
	@cd frontend; npm run build

all: buildFront build
