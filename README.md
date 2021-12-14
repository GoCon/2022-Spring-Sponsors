# 2022-Spring-Sponsors

## How to use

### Dump company names

Dump company names into below files.

* cmd/gentest/_companies/platinum_gold.txt 
* cmd/gentest/_companies/gold.txt 

### Generate testcodes

```sh
$ go1.17.5 generate
```

### Run lottery

```sh
$ go1.17.4 test -v -shuffle on -run "TestPlatinumGold*"
$ go1.17.4 test -v -shuffle on -run "TestGold*"
```
