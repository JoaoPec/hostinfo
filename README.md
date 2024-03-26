# hostinfo

Hostinfo is a simple command-line application written in Go that retrieves the IP address and name server of web addresses.
Installation

Make sure you have Go installed on your system. You can download and install it from golang.org.
Clone the repository:

```bash

git clone https://github.com/JoaoPec/hostinfo.git
```
Navigate to the project directory:

```bash

cd hostinfo

```

Build the project:

```go

    go build
```

## Usage

The application provides a command-line interface with the following command:
(By default, it searches for "google.com".)
```bash

./hostinfo ip --host <web_address>
```
Replace <web_address> with the desired web address for which you want to retrieve the IP address and name server. By default, it searches for "google.com".

Example:

```bash

./hostinfo ip --host=example.com
```

This will display the IP address(es) and name server(s) associated with the provided web address.

To see the server name run the following command:

```bash

./hostinfo server --host <web_address>
```
Replace <web_address> with the desired web address for which you want to retrieve the server name.
