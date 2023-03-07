package main

import (
	"container/list"
	"time"

	"golang.org/x/sys/unix"
)

const (
	GE_GREP_CMD = "!rg -uuu --line-number "
	GE_FIND_CMD = "!find . -type f -name "
	errno_s     = "no error" // this is not supported in Go, so hardcode it is!

	GE_MAX_POLL      = 128
	GE_MAX_FILE_SIZE = (1024 * 1024 * 1024)

	GE_BUFFER_SEARCH_NORMAL   = 0
	GE_BUFFER_SEARCH_PREVIOUS = 1
	GE_BUFFER_SEARCH_NEXT     = 2

	GE_EDITOR_MODE_NORMAL     = 0
	GE_EDITOR_MODE_INSERT     = 1
	GE_EDITOR_MODE_COMMAND    = 2
	GE_EDITOR_MODE_BUFLIST    = 3
	GE_EDITOR_MODE_SEARCH     = 4
	GE_EDITOR_MODE_SELECT     = 5
	GE_EDITOR_MODE_NORMAL_CMD = 6
	GE_EDITOR_MODE_MAX        = 7

	TERM_COLOR_BLACK   = 0
	TERM_COLOR_RED     = 1
	TERM_COLOR_GREEN   = 2
	TERM_COLOR_YELLOW  = 3
	TERM_COLOR_BLUE    = 4
	TERM_COLOR_MAGENTA = 5
	TERM_COLOR_CYAN    = 6
	TERM_COLOR_WHITE   = 7

	TERM_COLOR_BRIGHT = 40
	TERM_COLOR_FG     = 30
	TERM_COLOR_BG     = 40

	TERM_CURSOR_MIN = 1
	TERM_ESCAPE     = "\x1b["

	TERM_SEQUENCE_CLEAR_CURSOR_DOWN = TERM_ESCAPE + "J"
	TERM_SEQUENCE_CLEAR_CURSOR_UP   = TERM_ESCAPE + "1J"
	TERM_SEQUENCE_CLEAR_ONLY        = TERM_ESCAPE + "2J"
	TERM_SEQUENCE_CLEAR             = TERM_ESCAPE + "2J" + TERM_ESCAPE + "1;1H"
	TERM_SEQUENCE_CURSOR_UP         = TERM_ESCAPE + "0A"
	TERM_SEQUENCE_CURSOR_DOWN       = TERM_ESCAPE + "0B"
	TERM_SEQUENCE_CURSOR_RIGHT      = TERM_ESCAPE + "0C"
	TERM_SEQUENCE_CURSOR_LEFT       = TERM_ESCAPE + "0D"
	TERM_SEQUENCE_CURSOR_SAVE       = "\0337"
	TERM_SEQUENCE_CURSOR_RESTORE    = "\0338"
	TERM_SEQUENCE_LINE_ERASE        = TERM_ESCAPE + "K"

	TERM_SEQUENCE_ATTR_OFF       = TERM_ESCAPE + "m"
	TERM_SEQUENCE_ATTR_BOLD      = TERM_ESCAPE + "1m"
	TERM_SEQUENCE_ATTR_REVERSE   = TERM_ESCAPE + "7m"
	TERM_SEQUENCE_FMT_SET_COLOR  = TERM_ESCAPE + "%dm"
	TERM_SEQUENCE_FMT_SET_CURSOR = TERM_ESCAPE + "%zu;%zuH"

	TERM_SEQUENCE_ALTERNATE_ON  = TERM_ESCAPE + "?1049h"
	TERM_SEQUENCE_ALTERNATE_OFF = TERM_ESCAPE + "?1049l"

	GE_FILE_TYPE_PLAIN = iota
	GE_FILE_TYPE_C
	GE_FILE_TYPE_PYTHON
	GE_FILE_TYPE_DIFF
	GE_FILE_TYPE_JS
	GE_FILE_TYPE_SHELL
	GE_FILE_TYPE_SWIFT
	GE_FILE_TYPE_YAML
	GE_FILE_TYPE_JSON
	GE_FILE_TYPE_DIRLIST
	GE_FILE_TYPE_HTML
	GE_FILE_TYPE_CSS
	GE_FILE_TYPE_GO
	GE_FILE_TYPE_LATEX
	GE_FILE_TYPE_LUA

	GE_TAB_WIDTH_DEFAULT  = 4
	GE_TAB_EXPAND_DEFAULT = 0

	/*
	 * Gamified statistics, because I can.
	 */

	GE_XP_GROWTH    = 15
	GE_XP_INITIAL   = 100
	GE_XP_PER_AWARD = 100
)

type gegame struct {
	Xp    uint32
	Opens uint32
}

type geconf struct {
	TabWidth  int // default: 4
	TabExpand int // default: no
	TabShow   int // default: yes
}

/*
 * A history entry for a cmd executed via cmdbug or select-execute
 */

type gehist struct {
	cmd string
}

// Define the history list as a head of the tail queue
type ge_histlist struct {
	head *list.List
}

// represents a single line in a file
const GE_LINE_ALLOCATED = 1 << 1

type geline struct {
	flags   uint32      // flags
	data    interface{} // line data
	maxsz   uint        // size of allocation in case line is allocated
	length  uint        // length of the line in bytes
	columns uint        // length of the line in columns
}

/*
 * A marker and its associated line in a cebuf.
 */
const (
	GE_MARK_MIN      = '0'
	GE_MARK_MAX      = 'z'
	GE_MARK_OFFSET   = GE_MARK_MIN
	GE_MARK_PREVIOUS = '\''
	GE_MARK_SELEXEC  = '.'
)

type gemark struct {
	set  bool
	line int
	col  int
	off  int
}

/*
 * A running process that is attached to a buffer.
 */
const GE_PROC_AUTO_SCROLL = (1 << 1)

type geproc struct {
	pid   int          // Process id
	ofd   int          // File descriptor to read from
	pfd   *unix.PollFd // Set from ge_buffer_proc_gather() until ge_buffer_proc_dispatch()
	first int          // XXX merge into flags?
	flags int          // Aux flags
	idx   uint64       // Line number index when command started
	cnt   uint64       // Number of bytes read in total
	cmd   string       // The command that was run
	buf   *gebuf       // Pointer back to the owning buffer
}

const (
	GE_BUFFER_DIRTY = 0x0001
	GE_BUFFER_RO    = 0x0004
)

const (
	GE_BUF_TYPE_DEFAULT  = 0
	GE_BUF_TYPE_DIRLIST  = 1
	GE_BUF_TYPE_SHELLCMD = 2
)

type gebuf struct {
	internal   bool
	buftype    uint16
	flags      uint32
	type_      uint32
	data       []byte
	maxsz      int
	length     int
	prev       *gebuf
	path       string
	mode       uint32
	mtime      time.Time
	name       string
	cursorLine int
	line       int
	column     int
	width      int
	height     int
	origLine   int
	origColumn int
	top        int
	loff       int
	lcnt       int
	lines      []*geline
	markers    [GE_MARK_MAX]gemark
	prevmark   gemark
	selend     gemark
	selmark    gemark
	selstart   gemark
	selexec    gemark
	proc       *geproc
	cb         func(*gebuf, uint8)
	intdata    interface{}
	list       *list.Element
}

type ge_buflist struct {
	head *list.List
}
