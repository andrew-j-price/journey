# journey
This is an my educational `journey` repo to `drive` (binary built) my learnings on Golang, GitHub Actions, and Argo tools.
* Some practices applied may not be ideal


## go install and clean
```bash
# clean - commented out for manual protection
# sudo rm -rf /usr/local/go/
# sudo rm -rf ~/go/
mkdir -p ~/go

# install - https://go.dev/doc/install
cd ~/downloads
wget https://go.dev/dl/go1.17.3.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.17.3.linux-amd64.tar.gz
go version
```

## vscode
actions:
* ctrl-shift-p -> Go: Install/Update Tools -> Select all -> install
* output `ls -la ~/go/bin/`
  ```
  Tools environment: GOPATH=/home/andrew/go
  Installing 10 tools at /home/andrew/go/bin in module mode.
    gopkgs
    go-outline
    gotests
    gomodifytags
    impl
    goplay
    dlv
    dlv-dap
    staticcheck
    gopls
  ...
  ```

settings.json
```json
"gopls": {
    "experimentalWorkspaceModule": true,
}
```

## package management
```bash
# Global
go install package@ver
go install package@new_ver

# Module
go get package@ver
go get package@new_ver
go get package@none
go clean package
go clean -r package

# Samples
go install golang.org/x/tools/gopls@latest
go get golang.org/x/tools/gopls

go install github.com/haya14busa/goplay/cmd/goplay@latest
go get github.com/haya14busa/goplay/cmd/goplay
```

## multi-arch
* DockerHub - [Consume](https://hub.docker.com/r/andrewprice/journey/tags) or [Manage](https://hub.docker.com/repository/docker/andrewprice/journey/tags)
* NOTE: if not pushing, build result will only remain in the build cache
```bash
# go list
go tool dist list

# prep buildx
docker buildx create --name gojourneybuilder
docker buildx use gojourneybuilder
docker buildx inspect gojourneybuilder
# see if `Status: inactive` from above, if so run:
docker buildx inspect gojourneybuilder --bootstrap

# perform actual build
docker buildx build --platform linux/amd64,linux/arm64 -t andrewprice/journey . --push

```


## nexus
* Nexus - [HTML](https://nexus.linecas.com/service/rest/repository/browse/docker/v2/journey/journey/) or [Manage](https://nexus.linecas.com/#browse/browse:docker)
```bash
docker login images.linecas.com  # developer / ...
docker pull images.linecas.com/journey/journey:pluralsight

```
