UDPutil
=========
UDPutil makes it easy to send or receive UDP via the command line.

### updrecv
udprecv logs messages received on the 'bind' address to stdout.

Usage:
```bash
udprecv -bind=":65000"
```

### udpsnd
udpsnd sends newline-delimited messages from stdin.

Usage:
```bash
udpsnd -bind ":65000"
```
