DIR = $(shell pwd)

build:
	@cd cmd; go build -o ../pixr
	@./pixr

buildFront:
	@cd frontend; npm run build

all: buildFront build
