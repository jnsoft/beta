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
```

```
beta http get <url> -o json
beta http get <url> -p <proxy_url>
```

```
beta key hex > key.out
key=$(<key.out)
beta hmac sha3 -k $key -f tmp.out > hmac.out
hmac=$(<hmac.out)
beta hmac verify sha3 -k $key --hmac $hmac -f tmp.out
``` 

```
key=$(./beta key hex -n 32)
./beta aes encrypt "Hello, World!" -k $key
./beta aes decrypt <encrypted_b64_string> -k $key

./beta aes encrypt "Hello, World!" -k $key

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

# Comments
Use fmt.Println() instead of cmd.Println() to print to stdout (to for example use > into a file)
But, this makes it hard to test the output from a command.

Rune is an int32, but by convention used to store charachters: []rune(str)
