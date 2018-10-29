## Gubrak

an experimental `Command Line` performance testing tool for your services. What this means, this `tool` will run concurrently againts your service

### Usage

- build from source
```shell
go get -U github.com/Bhinneka/gubrak/cmd
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
./gubrak -m get -c /Users/wurianto/Documents/config.json
```

### List flag and arguments
- `-m` HTTP method, example `-m POST` or `-m post`
- `-r` Size of Concurrent `request`, example `-r 1000`
- `-c` `config.json` location, example `-c /Users/wurianto/Documents/config.json`
- `-h` show `Help`, example `./gubrak -h`

##

### Author
wuriyanto musobar https://github.com/wuriyanto48