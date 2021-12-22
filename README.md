# GoShare

GoShare provides a simple Web UI and a file transfer service which using HTTP or FTP protocol.

## Install

### Download executable

Download the compiled **[release](https://github.com/duruyao/goshare/releases)** for your platform.

### Compile from source

**Prerequisites**: [Golang (version >= 1.17)](https://golang.org/)

```bash
git clone https://github.com/duruyao/goshare.git goshare && \
  pushd goshare && \
  chmod +x build.sh && \
  ./build.sh && \
  popd
```

## Usage

The following operation takes the Linux platform as an example.

Add executable permissions to the current file.

```bash
chmod +x goshare
```

Type `./goshare -h` to show usage.

```text
USAGE:
    goshare [-h] [-v] [--url-prefix <prefix>] [-s {http, ftp}] [-a <ip:port>] [-p <path>]

OPTIONS:
    -h, --help
                    show usage
    -v, --version
                    show version
    --url-prefix <prefix>
                    url prefix
    -s {http, ftp}, --scheme {http, ftp}
                    scheme name (default: "http")
    -a <ip:port>, --address <ip:port>
                    ip address and port to listen (default: "127.0.0.1:8080")
    -p </path/to/file>,	--path </path/to/file>
                    path of file or directory to share (default: "$HOME")

EXAMPLES:
    goshare -a 10.0.13.120:8080 -p /opt/share0/releases/
    goshare --url-prefix /share/releases/ -a 10.0.13.120:8080 -p /opt/share0/releases/
    goshare --url-prefix /share/releases/ -a=10.0.13.120:8080 -p=/opt/share0/releases/
    goshare --url-prefix=/share/releases/ --address 10.0.13.120:8080 --path /opt/share0/releases/
    goshare --url-prefix=/share/releases/ --address=10.0.13.120:8080 --path=/opt/share0/releases/
```

Start GoShare.

```bash
sudo ./goshare -a 127.0.0.1:8080 -p /opt
```

Press `Ctrl_Z` to stop the service in foreground.

Type `jobs` to show all the jobs' status contains of **JOB_ID**.

Use the `fg` to start and run the service in background.

```bash
fg %JOB_ID
```

Access [http://127.0.0.1:8080/opt/](http://127.0.0.1:8080/opt/) in your browser.

![img/browser-127.0.0.1.png](img/browser-127.0.0.1.png)

Share the files with your group members by using the **LAN IP** or **WAN IP** instead of `localhost`, `127.0.0.1` or `*`. 

Download a file by `wget`.

```bash
wget [-P <Destination Directory>] [-O <Destination Path>] <https://ip:port/path/to/file>
```

Download a file by `curl`.

```bash
curl --request GET -sL --url <https://ip:port/path/to/file> --output </path/to/file>
```

Download a directory by `wget`.

```bash
wget [-P <Destination Directory>] -r -np -nH -R 'index.html*' <https://ip:port/path/to/dir/>
```
