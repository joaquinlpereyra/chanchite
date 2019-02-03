Chanchite
=========

```
       |\_,,____ 
        ( o__o \/            PIGGY IS PROUD OF YOU
        /(..)  \                YOU HAVE EARNED
       (_ )--( _)                   $424242 
       / ""--"" \ 
```

Chanchite is a ultra-small script which tells you how much $ you're making 
each time you commit your code.

Usage
========
To set _chanchite_ to run after each "git commit", a shell function is required.
Just write something like this in your .zshrc, .bashrc or whatever you use
to read the config for your shell.

Replace $SALARY with your monthly salary.

```bash
git() {
    if [[ $1 == "commit" ]]; then
        shift 1;
        command git commit "$@" && chanchite $SALARY;
    else
        command git "$@"
    fi
}
```

This will work even for your aliases, like if you have set up 'gcm' to 'git commit'.

Installation
==========
There's a binary for `linux amd64` and `darwin amd64` included in the repository. 
Just go to `/bin/$ARCH/chanchite`, download the binary and put it somewhere in your `$PATH`.

Build
=====
If you want to use `chanchite` in other architectures, you must have `golang` installed.

This should do the trick.
```bash
$ go get github.com/joaquinlpereyra/chanchite
$ go install chanchite
```

Chanchite's commitment to code quality
=============================
Chanchite was built in an hour on a late Saturday night, with absolutely no 
automated testing and on a 'it works for me' development model.

You know what to expect.
