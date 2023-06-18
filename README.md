# journey
This is an my educational `journey` repo to `drive` (binary built) my learnings on Golang, GitHub Actions, and anything else that interests me.
* Some practices applied may not be ideal
* Consumed by [companion](https://github.com/andrew-j-price/companion)
* Package docs on [pkg.go.dev](https://pkg.go.dev/github.com/andrew-j-price/journey)


## install golang
```bash
# clean (if necessary) - commented out for manual protection
# sudo rm -rf /usr/local/go/
# sudo rm -rf ~/go/
mkdir -p ~/go

# install - https://go.dev/doc/install
mkdir ~/downloads
cd ~/downloads
VERSION=1.20.5
wget https://go.dev/dl/go${VERSION}.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go${VERSION}.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin
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

launch.json for debugger
```json
        {
            "name": "go:journey",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "console": "integratedTerminal",
            "program": "${workspaceFolder}/journey",
            "args": [
                //"-api"
                "-debug", "-math", "add", "5", "7"
                //"-random"
            ],
            "env": {
                "DEBUGGER": "True"
            },
            "envFile": "${workspaceFolder}/.vscode/.env"
        },
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

## git workflow
```bash
# squash commits on feature branch
git checkout feature_branch
git reset --soft main

# standard vscode workflow for add, commit, pull-push
# on GitHub Pull Request set "Squash and merge" when accepting

# Releasing from `main` branch
# navigate to https://github.com/andrew-j-price/journey/releases/new
# create semantic tag starting with `v` like: vX.Y.Z
# do not give the release a name, it will simply be the tag version then
# click "Generate release notes"
# click "Publish release"

# tag
git tag  # to list tags, otherwise generating tags via releases in GitHub
```

## goreleaser manual
* NOTE: intended for usage if GitHub Action job fails.
* NOTE: does not cover items such as Docker image build, push.
```bash
# install
go install github.com/goreleaser/goreleaser@latest

# run
goreleaser init    # initially populates .goreleaser.yaml
goreleaser check   # checks file .goreleaser.yaml
goreleaser release --snapshot --clean
goreleaser build   --single-target  # local-development testing

# release
source ~/code/.vscode/.env   # equals: export GITHUB_TOKEN="YOUR_GH_TOKEN"
TAG=v0.0.0
echo $TAG
git tag -a $TAG -m "$TAG"
git push origin $TAG
goreleaser release --clean

# NOTE: if during CI problem, can re-tag as necessary
git fetch --tags
git tag -l
# git tag -d $TAG
# git push origin :refs/tags/$TAG
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

## grpc (boondocks)

```bash
# Install protoc binary
sudo apt install protobuf-compiler
sudo apt install golang-goprotobuf-dev
protoc --version

# Alternative to protoc (need to evaluate further)
# https://docs.buf.build/installation#binary / https://github.com/bufbuild/buf
BIN="/usr/local/bin" && VERSION="1.6.0" && sudo curl -sSL "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" -o "${BIN}/buf" && sudo chmod +x "${BIN}/buf"


# Configure repo for grpc
cd ~/code/journey
go install google.golang.org/grpc
go install google.golang.org/protobuf
go install google.golang.org/protobuf/cmd/protoc-gen-go
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
ls -la ~/go/bin/
go mod tidy

# Build protocol buffers
cd ~/code/journey/
export PATH=~/go/bin:$PATH
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./boondocks/messages/messages.proto

## grpc server (terminal 1)
cd ~/code/journey/
make build
./drive -boondocks-server

## grpc client (terminal 2)
cd ~/code/journey/
make build
./drive -boondocks-client -boondocks-rps
./drive -boondocks-client -boondocks-name Drew
./drive -boondocks-client -boondocks-rps -boondocks-name Drew
watch ./drive -boondocks-client -boondocks-rps

```


## nexus
* Nexus - [HTML](https://nexus.linecas.com/service/rest/repository/browse/docker/v2/journey/journey/) or [Manage](https://nexus.linecas.com/#browse/browse:docker)
```bash
docker login images.linecas.com  # developer / ...
docker pull images.linecas.com/journey/journey:latest

```


## doc update
If [pkg.go.dev](https://pkg.go.dev/github.com/andrew-j-price/journey) is not updated, perform:  
### UI method
If a new package was added like [fake-package](https://pkg.go.dev/github.com/andrew-j-price/journey/fake)
* Simply navigate to the URL and click: Request “github.com/andrew-j-price/journey/fake”

### manual method
Example `curl` command with updated tag (should move to CI-CD)
```
$ curl https://sum.golang.org/lookup/github.com/andrew-j-price/journey@v0.2.0
10680626
github.com/andrew-j-price/journey v0.2.0 h1:z35bClodzfCsh2psGoTtNsb4oHbZcnzD4pzY8gk6lGk=
github.com/andrew-j-price/journey v0.2.0/go.mod h1:ucycXl93RbqYoM3rGHjem1W0TDIYRLsU8prl9/zKx0g=

go.sum database tree
10680630
k1UQFwFXWQa7MF1BREWVQxIhScvUdgRepXuyRS3sRuE=

— sum.golang.org Az3grqgyWctiMcdUtITQpkGSxEoEvbaAq7J3FpbaVlW1Zo47LX/jmeeE7LiBktpmR2d/1RsdocLU9HMXYXxHCVyf/AQ=
```

From command line of [companion](https://github.com/andrew-j-price/companion) repo, run with updated tag (should also move to CI-CD):
```
go get -d github.com/andrew-j-price/journey@v0.2.0
```
