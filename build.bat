@echo off
if exist main.exe del main.exe >nul
if exist balloon.exe del balloon.exe >nul
go build -o balloon.exe main.go 

