# WebShare

## Overview

**WebShare** is a simple and efficient tool for sharing files over HTTP/HTTPS. With WebShare, you can quickly set up a web service to share a directory's contents for reading and writing data. The tool supports basic authentication for secure access, making it ideal for both local network and internet usage.

## Features

- **Easy to Use**: Simple command-line interface to start sharing your files.
- **Cross-Platform**: Works on all major platforms.
- **HTTP/HTTPS Support**: Share files securely over the web.
- **Basic Authentication**: Protect your shared directories with basic auth.

## Installation

To install WebShare, download the executable from the [GitHub Releases](#) page corresponding to your platform.

## Usage

### Basic Command

To run WebShare, use the following command:

```shell
webshare -p <port> -r <path/to/read-only-directory> -rw <path/to/read-write-directory>
```

### Command-Line Options

- `-p, --port <port>`: Specify the port to run the web service on (default: 8080).
- `-r, --readonly <path/to/read-only-directory>`: Set the directory that can be accessed for reading.
- `-rw, --readwrite <path/to/read-write-directory>`: Set the directory that can be accessed for both reading and writing.
- `-u, --username <username>`: Set the username for basic authentication.
- `-P, --password <password>`: Set the password for basic authentication.
- `--https`: Enable HTTPS (requires certificate and key files).
- `--cert <path/to/cert>`: Specify the path to the SSL certificate file.
- `--key <path/to/key>`: Specify the path to the SSL key file.

### Example

Running WebShare on port 8080 with a read-only directory and a read-write directory:

```shell
webshare -p 8080 -r /path/to/read-only -rw /path/to/read-write -u myusername -P mypassword
```

### Accessing the Web Service

Once the service is running, you can access the shared directories from any device on the network by navigating to:

```
http://<your-ip>:<port>
```

For example, if you are running the service on `localhost` and port `8080`, access it at:

```
http://localhost:8080
```

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve the project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgements

Thanks to all the contributors who have helped make WebShare better.
