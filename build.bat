go version
set GO111MODULE=on
set GOPROXY=https://mirrors.aliyun.com/goproxy/

goversioninfo -platform-specific=true -icon=icon.ico -manifest=elevator.manifest

echo 'Building amd64'
set GOOS=windows
set GOARCH=amd64
go build -ldflags "-s -w -H=windowsgui" -o elevator-amd64.exe .


echo 'Building arm64'
set GOARCH=arm64
@REM if you want to build arm64, you need to install aarch64-w64-mingw32-gcc
set CC="D:/Program Files/llvm-mingw-20240518-ucrt-x86_64/bin/aarch64-w64-mingw32-gcc.exe"
go build -ldflags "-s -w -H=windowsgui" -o elevator-arm64.exe .
