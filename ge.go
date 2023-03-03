package main

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
	Cmd  string
	List tailqEntry[gehist]
}

type geHistList tailqHead[gehist]

type geline struct {
	flag    uint32
	data    interface{}
	maxsz   uint
	length  uint
	columns uint
}

const GELINE_ALLOCATED = 1 << 1
