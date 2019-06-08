DIR = $(shell pwd)

build:
	@cd cmd/pixrsvc; go build -o ../pixr
	@./pixr

buildFront:
	@cd frontend; npm run build

all: buildFront build
