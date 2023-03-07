package main

import (
	"container/list"
	"log"
	"strconv"
)

// Conservative
const (
	BUFFER_MAX_IOVEC       = 64
	BUFFER_SEARCH_FORWARD  = 1
	BUFFER_SEARCH_BACKWARD = 2
)

var (
	buffers       = &ge_buflist{list.New()}
	internals     = &ge_buflist{list.New()}
	errstr        = ""
	active        = &gebuf{}
	scratch       = &gebuf{}
	cursor_column = TERM_CURSOR_MIN
)

func geBufferInit(argc int, argv []string) {
	var last *gebuf
	var linenr int64
	var err error

	last = nil

	// init the buffer and internal lists
	buffers := list.New()
	internals := list.New()

	scratch = geBufferInternal("scratch")
	scratch.mode = 0644
	active = scratch
	geTermUpdateTitle()

	for i := 0; i < argc; i++ {
		if len(argv[i]) > 0 && argv[i][0] == '+' && last != nil {
			linenr, err = strconv.ParseInt(argv[i][1:], 10, 64)
			if err != nil {
				log.Fatalf("%s is a bad line number", argv[i][1:])
			}
			geBufferJumpLine(last, int(linenr), 0)
			continue
		}
		last = geBufferFile(argv[i])
		if last == nil {
			geEditorMessage("%s", geBufferStrerror())
		}
	}
	if active = buffers.Front().Value.(*gebuf); active == nil {
		active = scratch
		geTermUpdateTitle()
		geEditorShowSplash()
		geEditorSettings(active)
	}
}

func geBufferCleanup() {
	var buf *gebuf
	for buffers.head.Len() > 0 {
		buf = buffers.head.Front().Value.(*gebuf)
		geBufferFree(buf)
	}

	for internals.head.Len() > 0 {
		buf = internals.head.Front().Value.(*gebuf)
		geBufferFreeInternal(buf)
	}
}

func geBufferCloseNonactive() {
	var buf, next *gebuf
	for e := buffers.head.Front(); e != nil; e = next.list {
		buf = e.Value.(*gebuf)
		next = e.Next().Value.(*gebuf)
		if buf == active || buf.buftype == GE_BUF_TYPE_SHELLCMD {
			continue
		}
		geBufferFree(buf)
	}

	active.prev = scratch
}

func geBufferCloseShellbufs() {
	var buf, next *gebuf
	for e := buffers.head.Front(); e != nil; e = next.list {
		buf = e.Value.(*gebuf)
		next = e.Next().Value.(*gebuf)
		if buf.buftype != GE_BUF_TYPE_SHELLCMD {
			continue
		}
		geBufferFree(buf)
	}

	active.prev = scratch
}

func geBufferSetError(err string)             { errstr = err }
func geBufferSetName(buf *gebuf, name string) { buf.name = name }
func geBufferScratchActive() bool             { return active == scratch }
