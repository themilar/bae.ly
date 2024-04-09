# bae.ly
Command line URL shortner using the [URLBae developer API](https://urlbae.com/developers).
Compile the package using `go build` and you can specify a file a directory to write to by using the `-o` flag.
An example for Windows systems: `go build -o bae.exe`. Running the executable generates the following message which outlines all the commands and flags available:

```
This is a command line application for shortening and managing URLS.

Usage:
  bae [command]

Available Commands:
  auth        Authenticate requests using your API key
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  list        List links you've shortend
  shorten     Shorten links

Flags:
  -h, --help   help for bae

Use "bae [command] --help" for more information about a command.
```
