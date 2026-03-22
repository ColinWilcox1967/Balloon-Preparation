@echo off
if exist balloon.exe del balloon.exe >nul
go build -o balloon.exe main.go 

