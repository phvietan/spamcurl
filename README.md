# spamcurl
A package to spam 'curl' command. Mainly used for checking if a single request causes server DOS or not !

# Installation

`$ go install github.com/phvietan/spamcurl@v1.0.1`

# Usage

`$ spamcurl -h`

```
Usage: spamcurl [options...] <url>
Will spam 'curl' command until receive SIGINT
Below is curl help
==============================================
Usage: curl [options...] <url>
 -d, --data <data>          HTTP POST data
 -f, --fail                 Fail silently (no output at all) on HTTP errors
 -h, --help <category>      Get help for commands
 -i, --include              Include protocol response headers in the output
 -o, --output <file>        Write to file instead of stdout
 -O, --remote-name          Write output to a file named as the remote file
 -s, --silent               Silent mode
 -T, --upload-file <file>   Transfer local FILE to destination
 -u, --user <user:password> Server user and password
 -A, --user-agent <name>    Send User-Agent <name> to server
 -v, --verbose              Make the operation more talkative
 -V, --version              Show version number and quit

This is not the full help, this menu is stripped into categories.
Use "--help category" to get an overview of all categories.
For all options use the manual or "--help all".
```

# License

Is released under MIT. See [LICENSE](./LICENSE)
