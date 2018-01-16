![ls-go](./img/ls-go.png)

A more colorful, user-friendly implementation of `ls` written in [Go](https://golang.org/).

You want to be able to glean a lot of information as quickly as possible from `ls`.
Colors can help, your mind parse the information.
You can configure `ls` to color the output a little bit.
Configuring `ls` is a hassle though, and the colors are limited.

Instead, you can use `ls-go`.
It is highly colored by default.
It has much fewer flags so you can get the behavior you want more easily.
The colors are beautiful and semantic.
A terminal with xterm-256 colors is **required*.*


## Install

```sh
$ go get github.com/acarl005/ls-go
```

## Usage

![demo-1](./img/demo-1.png)

Of course, you can use an alias to save some typing and get your favorite options.

![demo-2](./img/demo-2.png)

### Prior Art

This is inspired by [monsterkodi/color-ls](https://github.com/monsterkodi/color-ls), ported to Go, with various modifications.

