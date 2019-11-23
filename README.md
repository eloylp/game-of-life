# Conway`s game-of-life, written in Go .
[![Build Status](https://travis-ci.com/eloylp/game-of-life.svg?branch=master)](https://travis-ci.com/eloylp/game-of-life)

https://en.wikipedia.org/wiki/Conway%27s_Game_of_Life

## How to install

You must go to [the releases page](https://github.com/eloylp/game-of-life/releases)
and download the latest binary for your architecture.

You can also try this one liner install:
```bash
sudo curl -L "https://github.com/eloylp/game-of-life/releases/download/v1.0.0/gol_1.0.0_Linux_x86_64" \
-o /usr/local/bin/gol \
&& sudo chmod +x /usr/local/bin/gol
```
## How to run 

If you want a description about the arguments, just invoke help:
```bash
gol -h
```

Here is a full example of how to use this CLI:
```bash
gol -x 20 -y 20 -g 100 -i 2
```

This will bring more life to your computer !

How to run tests
```bash
go test
```

How to run benchmark tests
```bash
go test -bench=.
```
