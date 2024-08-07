.SUFFIXES:
.PHONY: all not_all ALL



all:
	go run calculator
not_all:
	go run main.go
ALL:
	go build
