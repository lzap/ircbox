#!/bin/bash
go build
strip ircbox
cp ./ircbox ./ircbox-linux64
git add ircbox-linux64
