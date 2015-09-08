# Piet-on-Go
A Piet interpreter written in Go

## Usage
```sh
$ export GOPATH=$PWD
$ make
$ bin/pietongo.o -i [filename]
```
If you want to use debug mode,
```sh
$ bin/pietongo.o --debug -i [filename]
```

## NOTE
- Additional colors will be regarded as black.
- This program is unstable. Pull requests are welcome.