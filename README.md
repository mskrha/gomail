[![Go Report Card](https://goreportcard.com/badge/github.com/mskrha/gomail)](https://goreportcard.com/report/github.com/mskrha/gomail)

## gomail

### Description
Very simple replacement of mail command.

### Build
```shell
git clone https://github.com/mskrha/gomail.git
cd gomail
make
```

### Usage
```shell
echo 'Test test test.' | ./gomail -f from@local.domain -r to@local.domain -s 'Test 123'
```
