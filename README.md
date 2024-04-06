# bae.ly
Command line URL shortner using the [URLBae developer API](https://urlbae.com/developers).
Compile the package using `go build` and you can specify a file a directory to write to by using the `-o` flag.
An example for Windows systems: `go build -o bae.exe`. Running the executable generates the following message which outlines all the commands and flags available:

```
Usage:
  bae [command]

Available Commands:
  auth        Provide your API key to make requests
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        print a list of links you've shortend
  shorten     shorten links
  
Flags:
  -h, --help   help for bae

Use "bae [command] --help" for more information about a command.
```
