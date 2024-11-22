# beta

## About
CLI tool build in go, using cobra to handle commands.

## Start project
go mod init github.com/jnsoft/beta
go get -u github.com/spf13/cobra@latest


## Usage
beta -h  
beta [command]
### commands
```(bash)
beta version  (show version)  
beta base64 from "test" (base 64 encode string)  
beta base64 to "dGVzdA==" (base64 decode string)  
beta base64 encode -i <input_filename> (base64 encode file)  
beta base64 encode -i <input_filename> -o <output_filename> (base64 encode file)  
beta base64 decode -i <input_filename> (base64 decode file)  
beta base64 decode -i <input_filename> -o <output_filename> (base64 decode file) 

beta http get <url> -o json
beta http get <url> -p <proxy_url>
``` 

# Build and Test
```
go mod tidy
go build -o beta ./cmd
./beta b64 to "test"
./beta b64 from "dGVzdA=="
./beta b64 encode -i hej.txt -o hej.b64
./beta b64 decode -i hej.b64 -o hej2.txt

go test ./cmd
go test ./cmd -v
go test ./util/aesutil/

```
