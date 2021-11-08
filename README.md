# A Collection Of Common Go Utilities

### Versions
For version you can use either a tag (^1.2.0), or a specific hash (f89effe81c1ece9c5b0fda359ebd9cf65f169a51)

To update a mock or create a new mock:
```
go get github.com/golang/mock/gomock
go get github.com/golang/mock/mockgen
mockgen -source=routing/JwtAuthenticator.go -destination=mocks/MockJwtAuthenticator.go -package=mocks
```
