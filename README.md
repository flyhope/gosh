# gosh
terminal ui ssh client

## Be dependent on program
1. zssh
2. sshpass (if use auto input password)

### Mac
```bash
brew install zssh sshpass
```

### Ubuntu
```bash
sudo apt install zssh sshpass 
```

### CentOS / Redhat
```bash
yum install zssh sshpass
```

### Windows

1. Use WSL install ubuntu.
2. Look for `ubuntu` install.

## Usage
1. copy `config.simple.xml` to `config.xml`
2. run gosh

## Support
* [x] custom command path
* [x] custom username/password/port/PrivateKey

## Todo
* [ ] set config with ui
