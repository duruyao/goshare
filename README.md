# GoShare

GoShare provides a simple Web UI and a file transfer service which using HTTP or FTP protocol.

## Install

### Download executable

Download the compiled **[releases](https://github.com/duruyao/goshare/releases)** for your platform.

## Usage

The following operation takes the Linux platform as an example.

Type `goshare -h` to show help manual.

```text

   _____       _____ _
  / ____|     / ____| |
 | |  __  ___| (___ | |__   __ _ _ __ ___
 | | |_ |/ _ \\___ \| '_ \ / _' | '__/ _ \
 | |__| | (_) |___) | | | | (_| | | |  __/
  \_____|\___/_____/|_| |_|\__,_|_|  \___|


Usage: goshare [OPTIONS]

GoShare shares file and directory by HTTP or FTP protocol

Options:
    -h, --help                  Display this help message
    --host STRING               Host address to listen (default: 'localhost:3927')
    --path STRING               Path or directory (default: '/opt/sdk')
    --scheme STRING             Scheme name (default: 'http')
    --url-prefix STRING         Custom URL prefix (default: '/')
    -v, --version               Print version information and quit

Examples:
    goshare -host example.io -path /opt/sdk
    goshare -host localhost:3927 -path /opt/sdk
    goshare --host localhost:3927 --url-prefix /sdk --path /opt/sdk
    goshare --host=localhost:3927 --url-prefix=/sdk --path=/opt/sdk

See more about GoShare at https://github.com/duruyao/goshare

```

Start goshare in the foreground.

```bash
sudo goshare --host=localhost:3927 --url-prefix=$PWD --path=$PWD
```

Start goshare in the background.

```bash
sudo goshare --host=localhost:3927 --url-prefix=$PWD --path=$PWD &
```

Share the files with your group members by using the **LAN IP** or **WAN IP** instead of `localhost`, `127.0.0.1` or `*`. 

Download a file via `wget` or `curl`.

```bash
wget <https://ip:port/path/to/file>

curl <https://ip:port/path/to/file> -o <file>
```

Download a directory via `wget`.

```bash
wget -r -np -nH -R "index.html*" <https://ip:port/path/to/dir>
```
