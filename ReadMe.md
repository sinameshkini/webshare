# WebShare

## Overview

**WebShare** is a simple and efficient tool for sharing files over HTTP. With WebShare, you can quickly set up a web service to share a directory's contents for reading and writing data. The tool supports basic authentication for secure access, making it ideal for both local network and internet usage.

## Features

- **Easy to Use**: Simple command-line interface to start sharing your files.
- **Cross-Platform**: Works on all major platforms.
- **Basic Authentication**: Protect your shared directories with basic auth.

## Installation

To install WebShare, download the executable from the [GitHub Releases](https://github.com/sinameshkini/webshare/releases) page corresponding to your platform.

## Usage

### Basic Command

To run WebShare, use the following command:

```shell
webshare -p <port> -r <path/to/read-only-directory> -w <path/to/read-write-directory>
```

### Command-Line Options

- `-p, --port <port>`: Specify the port to run the web service on (default: 8080).
- `-r, --readonly <path/to/read-only-directory>`: Set the directory that can be accessed for reading.
- `-w, --readwrite <path/to/read-write-directory>`: Set the directory that can be accessed for both reading and writing.
- `-U, --username <username>`: Set the username for basic authentication.
- `-P, --password <password>`: Set the password for basic authentication.

### Example

Running WebShare on port 8080 with a read-only directory and a read-write directory:

```shell
webshare -p 8080 -r /path/to/read-only -w /path/to/read-write -U myusername -P mypassword
```

### Accessing the Web Service

Once the service is running, you can access the shared directories from any device on the network by navigating to:

```
http://<your-ip>:<port>
```

For example, if you are running the service on `localhost` and port `4242`, access it at:

```
http://localhost:4242
```

### Preview
```shell
$ webshare -h
 __        __         _       ____    _
 \ \      / /   ___  | |__   / ___|  | |__     __ _   _ __    ___
  \ \ /\ / /   / _ \ | '_ \  \___ \  | '_ \   / _` | | '__|  / _ \
   \ V  V /   |  __/ | |_) |  ___) | | | | | | (_| | | |    |  __/
    \_/\_/     \___| |_.__/  |____/  |_| |_|  \__,_| |_|     \___|

version: 0.0.2 
Available on https://github.com/sinameshkini/webshare

        WebShare is a user-friendly application designed for quick and secure file sharing over HTTP/HTTPS. 
        It allows you to easily set up a web service that provides access to specified directories for both 
        reading and writing data. With cross-platform compatibility and support for basic authentication, WebShare 
        ensures that your files can be accessed securely from any device on your local network or over the internet.
        Whether you need to share files for collaboration or distribute data efficiently,
        WebShare offers a straightforward solution with minimal setup.

Usage:
  webshare [flags]

Flags:
  -h, --help               help for webshare
  -P, --password string    Set the password for basic authentication.
  -p, --port int           Specifies the port number on which the web service will run. (default 4242)
  -r, --readonly string    Set the directory that can be accessed for reading. (default ".")
  -w, --readwrite string   Set the directory that can be accessed for both reading and writing.
  -U, --username string    Set the username for basic authentication.
```

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve the project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

Thanks to all the contributors who have helped make WebShare better.
