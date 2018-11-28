## Gubrak

an experimental `Command Line performance testing tool` for your services. What this means, this `tool` will run concurrently againts your service.

### Gubrak Gopher

<p align="center">
  <img width="460" height="400" src="./files/gubrak.png">
</p>

### TODO
- Better code
- Writing test
- Better output

### Usage

- Install from `Homebrew`
```shell
$ brew install Bhinneka/tool/gubrak
```

- Install `binary` from source
```shell
$ go get github.com/Bhinneka/gubrak

$ go install github.com/Bhinneka/gubrak/cmd/gubrak

$ gubrak --version
```

- create `config.json` file with signature like this:
```json
{
    "url": "http://example.com",
    "headers": {
        "Authorization": "Basic exnfekeoeoeojsjalaljahhd",
        "Content-Type": "application/json",
        "Accept": "application/json"
    },
	"payload": {
		"from": "Bob",
		"content": {
			"header": "This is Message 3",
			"body": "Hello There"
		}
	}
}
```

- run `gubrak`
```shell
$ gubrak -m get -c /Users/wurianto/Documents/config.json -u https://jsonplaceholder.typicode.com/posts -r 100
```

### List flag and arguments
- `-m | --method` (default `GET`) HTTP method, example `-m POST` or `-m post`
- `-r` (default `10`) Size of Concurrent `request`, example `-r 1000`
- `-c | --config` `config.json` (default `config.json`) location, example `-c /Users/wurianto/Documents/config.json`
- `-u | --url` `URL full with path` (default in `config.json`), example `-u https://jsonplaceholder.typicode.com/posts`
- `-v | --version` show `gubrak` version, example `gubrak -v`
- `-h | --help` show `Help`, example `./gubrak -h`

##

### Author
Wuriyanto musobar https://github.com/wuriyanto48