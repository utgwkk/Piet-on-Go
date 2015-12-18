[![Build Status](https://travis-ci.org/utgw/Piet-on-Go.svg?branch=master)](https://travis-ci.org/utgw/Piet-on-Go)

# Piet-on-Go
A Piet interpreter written in Go

## Usage
```sh
$ make
$ ./pietongo -i [filename]
```
If you want to use debug mode,
```sh
$ ./pietongo --debug -i [filename]
```
or you want to specify the pixel of codel,
```sh
$ ./pietongo --codel [pixel size] -i [filename]
```
## NOTE
- Additional colors will be regarded as black.
- This program is unstable. Pull requests are welcome.