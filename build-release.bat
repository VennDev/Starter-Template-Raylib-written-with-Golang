@echo off
title Building Game for Windows
go build -x -ldflags="-H=windowsgui"  -o ./build/GameRelease.exe
