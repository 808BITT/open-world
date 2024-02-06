@echo off

cd ..
go build -o bin\run.exe
cd bin
run.exe
cd ..\scripts
