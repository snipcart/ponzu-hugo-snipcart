#!/usr/bin/env bash

go build -o ./ponzuImport ./ponzuImport.go

./ponzuImport

hugo
