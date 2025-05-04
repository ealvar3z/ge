# ge (go ed(1) clone)

## Original Go [Port (from thimrc)](https://github.com/thimrc/ed)

```md
Go clone of [ed(1)](https://man.openbsd.org/ed.1), the famous
line-oriented text editor that is originally from the 1970s. Simply
put, _the UNIX text editor_.

## Differences
This version of ed aims to be a bug for bug implementation of the
original. The only thing that differs is that this version uses RE2
instead of BRE (basic regular expresions). The reason for this is that
the Go programming languages standard library uses that in the
[regexp](https://pkg.go.dev/regexp) package.
```

## Additions

This port extends [thimc/ed](https://github.com/thimc/ed)'s port by adding
the following:

1. readline support via [liner](https://github.com/peterh/liner)
1. adds `m` command.  
1. adds `b` command from Plan9 ed.  

## Todo

	godoc -notes 'TODO'


## Installation


	go build
	./ed file


## Future


1. Add treesitter support for structural regexp (a spin of sam's regexp
model)  
1. Vi keybindings (possibly pivot off liner and use an `.inputrc`
   supportable readline library for Go.


## Non-future


1. syntax highlighting support
1. feature creep
