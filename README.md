# GoFS

GoFS provides a simple Web UI for files storing server.

## Install

Download **[release](https://github.com/duruyao/gofs/releases)** for your platform.

## Usage

The following demo is for users of Linux OS.

```bash
$ chmod +x GoFS-Linux-amd64
$ ./GoFS-Linux-amd64 -v
2021.11.16
$ ./GoFS-Linux-amd64 -h
Usage of ./GoFS-Linux-amd64:
  -a string
        listening address in "<ip>:<port>" format (default "127.0.0.1:8080")
  -f string
        handling local file path in "/.../<path>" format (default "/home/duruyao")
$ sudo ./GoFS-Linux-amd64 -a 127.0.0.1:8080 -p /opt
GoFS is listening on 127.0.0.1:8080 and handling /opt/ ...
...
```

Access [http://127.0.0.1:8080](http://127.0.0.1:8080) in your browser.

![img/browser-127.0.0.1.png](img/browser-127.0.0.1.png)

Using the LAN IP or WAN IP instead of `localhost` or `127.0.0.1` to share the files with your group members.

Download file by browser or `curl`.

```bash
$ curl --url 'http://ip:port/path/to/file' --output './path/to/file'
```