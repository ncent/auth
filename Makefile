ifdef DOTENV
	DOTENV_TARGET=dotenv
else
	DOTENV_TARGET=./.env
endif

.PHONY: build clean


build: clean #test # generate_graphql test
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -a -tags netgo -installsuffix netgo -o bin/auth/create/secret handlers/auth/create/secret/main.go
	env GOOS=linux go build -ldflags="-s -w" -a -tags netgo -installsuffix netgo -o bin/auth/create/jwt handlers/auth/create/jwt/main.go
	env GOOS=linux go build -ldflags="-s -w" -a -tags netgo -installsuffix netgo -o bin/auth/validate/jwt handlers/auth/validate/jwt/main.go
	chmod +x bin/auth/create/secret
	chmod +x bin/auth/create/jwt
	chmod +x bin/auth/validate/jwt
	zip -j bin/auth/create/secret.zip bin/auth/create/secret
	zip -j bin/auth/create/jwt.zip bin/auth/create/jwt
	zip -j bin/auth/validate/jwt.zip bin/auth/validate/jwt
	
clean:
	-rm -rf ./bin

test: build
	go test -race $$(go list ./... | grep -v /vendor/) -v -coverprofile=coverage.out
	go tool cover -func=coverage.out
