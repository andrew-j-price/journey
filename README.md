# journey
This is an my educational `journey` repo to `drive` (binary built) my learnings on Golang, GitHub Actions, and Argo tools.
* Some practices applied may not be ideal

## setup
```bash
cd ~\code\journey
go mod init journey

sudo ln -s ~/code/journey/drive /usr/local/bin/drive

```

## vscode
settings.json
```json
"gopls": {
    "experimentalWorkspaceModule": true,
}
```
commands
```bash
cd ~/code/tmp
go get -v golang.org/x/tools/gopls
go install golang.org/x/tools/gopls@latest

# NOTE: did not do anything with vscode, to analyze further
go get -v github.com/haya14busa/goplay/cmd/goplay
go install github.com/haya14busa/goplay/cmd/goplay@latest
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
