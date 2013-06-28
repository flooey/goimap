#!/bin/sh

ragel -Z parser/parser.rgo
go run goimap/main.go 8000
