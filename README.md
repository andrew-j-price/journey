# journey
This is an my educational `journey` repo to `drive` my learnings on Golang, GitHub Actions, and Argo tools.

## setup
```
cd ~\code\journey
go mod init journey

```

## vscode
```
cd ~/code/tmp
go get -v golang.org/x/tools/gopls

```

## multi-arch
* DockerHub - [Consume](https://hub.docker.com/r/andrewprice/journey/tags) or [Manage](https://hub.docker.com/repository/docker/andrewprice/journey/tags)
* NOTE: if not pushing, build result will only remain in the build cache
```
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
```
docker login images.linecas.com  # developer / ...
docker pull images.linecas.com/journey/journey:pluralsight

```
