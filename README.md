UDPutil
=========
UDPutil makes it easy to send or receive UDP in the terminal.

### updrecv
udprecv logs messages received on the 'bind' address to stdout.

Print to stdout:
```
udprecv -bind=":65000"
```

Write results into a file for 1 second:
```
updrecv -bind=":54000" -t=1s > msgs.txt
```

### udpsnd
udpsnd sends newline-delimited messages from stdin.

Type messages yourself:
```
udpsnd -bind ":65000"
```
Pipe messages across the wire:
```
echo 'here's a message\nhere's another message' | udpsnd -bind ":50000"
```

### TODO

 - Multicast send/receive
