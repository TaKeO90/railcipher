## Rail fence cipher in golang



### Quick Start

```sh
$ go build .
```

* encryption Example

```sh
$ ./railcipher  ./railcipher -key 5 -word "hello world" -encrypt
```
Excpected Output : `h_dewlolrol`

* decryption Example

```sh
$ ./railcipher -key 5 -word h_dewlolrol
```
Excpected Output : `hello_world`
