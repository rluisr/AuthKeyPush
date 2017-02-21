## AuthKeyPush
Deploy public key to bastion server of ssh.

## Overview
1. Access `http://localhost:9000`
2. Put name and public key.
3. Submit

## Install
### Install revel
```
$ go get github.com/revel/revel
$ go get github.com/revel/cmd/revel

```
### Deploy
```
$ cd $GOPATH/src
$ revel new git@github.com:rluisr/AuthKeyPush
```

### First setting
**conf/ssh.json**  
example:
```
[
  {
    "host": "13.112.xx.xx",
    "user": "ec2-user",
    "port": "22",
    "key": "/path/to/secret_key"
  },
  {
    "host": "192.168.xx.xx",
    "user": "foo",
    "port": "49155",
    "password": "foobar"
  }
]
```

### Second setting
Put secret key `/path/to/secret_key`.

## Run
