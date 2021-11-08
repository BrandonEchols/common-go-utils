# A Collection Of Comms Integration Go Utilities

## Theory
This should allow us to import the utils to each of our services, and push changes to them as needed
### Versions
If you wish to check out a specific version, in your glide.yaml add something like:
```
- package: ci-go-utils
  version: <version>
  repo:    ssh://git@stash.xant.tech:7999/ci/ci-go-utils.git
  vcs:     git
```
For version you can use either a tag (^1.2.0), or a specific hash (f89effe81c1ece9c5b0fda359ebd9cf65f169a51)

To update a mock or create a new mock:
```
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
mockgen -source=routing/JwtAuthenticator.go -destination=mocks/MockJwtAuthenticator.go -package=mocks
```
