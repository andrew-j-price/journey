# journey


## 
```
cd ~\code\journey
# setup
go mod init journey

# docker build
docker build -t journey .
docker run -it -p 8080:8080 journey

# local build
CGO_ENABLED=0 go build -o drive
./drive

# api call
curl -m 2 127.0.0.1:8080
```

## vscode
```
cd ~/code/tmp
go get -v golang.org/x/tools/gopls

```