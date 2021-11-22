# GoFS

GoFS provides a simple Web UI and a transfering files service which using HTTP protocol.

## Install

### Install binary

Download **[release](https://github.com/duruyao/gofs/releases)** for your platform.

### Compile from source

**Prerequisites**: [Golang](https://golang.org/)

```bash
git clone https://github.com/duruyao/gofs.git && \
  pushd gofs && \
  chmod +x build.sh && \
  ./build.sh && \
  popd
```

## Usage

The following demo is for users of Linux OS.

Add executable permissions to the current file.

```bash
chmod +x GoFS-Linux-amd64
```

Show version.

```bash
./GoFS-Linux-amd64 -v
```

Show help guide.

```bash
./GoFS-Linux-amd64 -h
```

Start up GoFS.

```
sudo ./GoFS-Linux-amd64 -a 127.0.0.1:8080 -p /opt
```

Access [http://127.0.0.1:8080/opt/](http://127.0.0.1:8080/opt/) in your browser.

![img/browser-127.0.0.1.png](img/browser-127.0.0.1.png)

Use the **LAN IP** or **WAN IP** instead of `localhost` or `127.0.0.1` to share the files with your group members.

Download a file by `wget`.

```bash
wget [-P 'Destination Directory'] [-O 'Destination Path'] 'http://ip:port/path/to/file'
```

Download a directory by `wget`.

```bash
wget [-P 'Destination Directory'] -r -np -nH -R 'index.html*' 'http://ip:port/path/to/dir/'
```
