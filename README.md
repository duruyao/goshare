# GoFS

GoFS provides a simple Web UI and a transfering files service which using HTTP protocol.

## Install

### Download executable

Download the compiled **[release](https://github.com/duruyao/gofs/releases)** for your platform.

### Compile from source

**Prerequisites**: [Golang (version >= 1.17)](https://golang.org/)

```bash
git clone https://github.com/duruyao/gofs.git gofs && \
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

Type `./GoFS-Linux-amd64 -h` to show usage.

```text
USAGE:
    {{.AppPath}} [-h] [-v] [--url-prefix <prefix>] [-s {http, https, ftp}] [-a <address>] [-p <path>]

OPTIONS:
    -h, --help
                    show usage
    -v, --version
                    show version
    --url-prefix <prefix>
                    url prefix
    -s {http, https, ftp}, --scheme {http, https, ftp}
                    scheme name (default: "{{.DefaultScheme}}")
    -a <ip:port>, --address <ip:port>
                    listening address (default: "{{.DefaultAddr}}")
    -p </path/to/file>,	--path </path/to/file>
                    handing path or directory (default: "{{.DefaultPath}}")

EXAMPLES:
    {{.AppPath}} -a 10.0.13.120:8080 -p /opt/share0/releases/
    {{.AppPath}} --url-prefix /share/releases/ -a 10.0.13.120:8080 -p /opt/share0/releases/
    {{.AppPath}} --url-prefix /share/releases/ -a=10.0.13.120:8080 -p=/opt/share0/releases/
    {{.AppPath}} --url-prefix=/share/releases/ --address 10.0.13.120:8080 --path /opt/share0/releases/
    {{.AppPath}} --url-prefix=/share/releases/ --address=10.0.13.120:8080 --path=/opt/share0/releases/
```

Start GoFS.

```bash
sudo ./GoFS-Linux-amd64 -a 127.0.0.1:8080 -p /opt
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
