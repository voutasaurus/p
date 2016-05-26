p [![Build Status](https://travis-ci.org/voutasaurus/p.svg?branch=master)](https://travis-ci.org/voutasaurus/p)
=======

p is a simple random password generator.

Usage
=====

Once p is compiled and added to PATH, you can run commands like the following examples:

p
- Generates a random string of length 40 from lowercase, uppercase, and digits (no symbols or whitespace)

p -l 10
- Specifies a length of 10 (character sets are default)

p -c a*
- Specifies character sets lowercase, uppercase, and symbols. The a* can be any combination of letter and a symbol (ensure any symbols are escaped properly in your terminal).

p -l 30 -c 1*
- Generates a random string of length 30 from digits and symbols. Again the 1% could have just as easily been 2*.

p -l 1 -c a
- Generates a random string of length 1 from lowercase and uppercase.

p -h
- Prints the help output.


Build Instructions
==================

First install and set up Go:
Install the latest version of Go (https://golang.org/dl/)

Set $GOPATH and create three directories in $GOPATH called bin, src and pkg. Add $GOPATH/bin to $PATH

Go get the package:
```
   go get github.com/voutasaurus/p
```
