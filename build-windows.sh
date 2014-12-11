CGO_ENABLED=0 GOOS=windows go build -o ./bin/simplefileserver.exe -a -tags netgo -ldflags '-w' github.com/inflatablewoman/simplefileserver
