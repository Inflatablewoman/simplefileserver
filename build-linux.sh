CGO_ENABLED=0 GOOS=linux go build -o ./bin/simplefileserver.linux -a -tags netgo -ldflags '-w' github.com/inflatablewoman/simplefileserver
