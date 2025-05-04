## ge (go ed(1) modern clone)

### Derivative Work: [(from thimrc)](https://github.com/thimrc/ed)

```markdown
Go clone of [ed(1)](https://man.openbsd.org/ed.1), the famous
line-oriented text editor that is originally from the 1970s. Simply put,
_the UNIX text editor_.

## Differences
This version of ed aims to be a bug for bug implementation of the
original. The only thing that differs is that this version uses RE2
instead of BRE (basic regular expresions). The reason for this is that
the Go programming languages standard library uses that in the
[regexp](https://pkg.go.dev/regexp) package.
```

### Additions

This port extends thimc/ed's port by adding the following:

- [ ] finish all of thimc's TODOs.
- [x] modify POSIX `z` command (browse backward)
- [x] implement `b` command (browse forward) from Plan9 ed.  
- [ ] readline support via [liner](https://github.com/peterh/liner)

### Todo

	grep -rn 'TODO' .


### Installation


	go build
	./ed file


### Future


1. Add treesitter support for structural regexp (a spin of sam's regexp
model)  
1. Vi keybindings (possibly pivot off liner and use an `.inputrc`
   supportable readline library for Go.


### Non-future


1. syntax highlighting support
1. feature creep
