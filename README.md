# binview

[![Build Status](https://travis-ci.com/Spaceface16518/binview.svg?branch=master)](https://travis-ci.com/Spaceface16518/binview)

**Ones and zeros**

**binview** is a simple command line utility and library for viewing the binary string representation of a binary file.
Sometimes viewing the "ones and zeros" of a file is more useful that just random garbage data. **binview** is here for that.

Soon, you will also be able to view the hex and octal representation of a file, as well as being able non-numerically group dump items (i.e. by UTF-8 characters)

## Usage
Run with the `-h` or `-help` flag to the the usage information.

Currently, binview yields this:

```text
-line int
    The number of bytes per line (default 4)
-path string
    A path to a binary file. If left blank, input will be taken standard input
```
_This section is in progress_
