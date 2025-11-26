go build -o bin\task-cli.exe .\main.go
copy bin\task-cli.exe "%GOPATH%\bin"
set PATH=%PATH%;"%GOPATH%\bin"