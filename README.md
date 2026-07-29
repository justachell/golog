# golog
### Simple and usefull utility for web-site status checking.

## Installation
```bash
# clone repository and move in golog directory
cd golog

# build program
go build
```

## Usage
Run utility with web-site URL, like this
```bash
./golog github.com
```
or this
```bash
./golog https://github.com
```

### Output example
```text
[INFO] status 200 OK
[INFO] status 200 OK
[ERR]  status 404 Not Found
[ERR]  timeout
```
