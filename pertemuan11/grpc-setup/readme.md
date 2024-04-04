# gRPC Environment Setup

## Install Protobuf Compiler
### Windows


1. Buka Terminal/Powershell/Cmd
2. Ketik Perintah dibawah :

```bash
cd ~
curl https://github.com/protocolbuffers/protobuf/releases/download/v26.1/protoc-26.1-win64.zip -o protoc.zip
mkdir protoc
tar -xf protoc.zip -C protoc
copy protoc\bin\protoc.exe ~\go\bin\
```

### Mac


1. Buka Terminal
2. Silahkan install homebrew terlebih dahulu jika belum terinstal, ketik perintah :
```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

3. Jika sudah, ketik Perintah dibawah :

```bash
brew install protobuf
```

### Linux


1. Buka Terminal
2. Ketik Perintah dibawah :

```bash
PROTOC_ZIP=protoc-26.1-linux-x86_64.zip
curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v26.1/$PROTOC_ZIP
sudo unzip -o $PROTOC_ZIP -d /usr/local bin/protoc
sudo unzip -o $PROTOC_ZIP -d /usr/local 'include/*'
rm -f $PROTOC_ZIP
```

## Install Protobuf GRPC Generator
Jika sudah menginstal protobuf compiler, di terminal/cmd/powershell, kita ketikkan perintah :
```bash
 go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
 go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```