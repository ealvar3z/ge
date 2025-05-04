package main

// undoType determines how the history entry should behandled,
// undoTypeDelete removes lines and undoTypeAdd adds lines.
type undoType int

const (
	undoTypeAdd undoType = iota
	undoTypeDelete
)

type undoAction struct {
	cursor
	typ   undoType
	lines []string
}

type undo struct {
	action  []undoAction // history for a command in progress
	history [][]undoAction
	global  []undoAction
}

func (u *undo) clear() { u.action = []undoAction{} }

func (u *undo) reset() {
	u.clear()
	u.history = [][]undoAction{}
}
func (u *undo) pop(ed *Editor) error {
	if len(u.history) < 1 {
		return ErrNothingToUndo
	}
	action := u.history[len(u.history)-1]
	for i := len(action) - 1; i >= 0; i-- {
		a := action[i]
		// fix the slice bounds panic error
		idx1 := a.first - 1
		if idx1 < 0 {
			idx1 = 0
		}
		if idx1 > len(ed.file.lines) {
			idx1 = len(ed.file.lines)
		}
		idx2 := a.second
		if idx2 < 0 {
			idx2 = 0
		}
		if idx2 > len(ed.file.lines) {
			idx2 = len(ed.file.lines)
		}
		before := ed.file.lines[:idx1]
		after := ed.file.lines[idx1:]
		switch a.typ {
		case undoTypeDelete:
			after = ed.file.lines[idx2:]
			ed.file.lines = append(before, after...)
		case undoTypeAdd:
			ed.file.lines = append(before, append(a.lines, after...)...)
		}
		ed.dot = a.dot
		ed.file.dirty = true
	}
	if count := len(u.history) - 1; count > 0 {
		u.history = u.history[:count]
	} else {
		u.history = [][]undoAction{}
	}
	return nil
}

func (u *undo) append(typ undoType, cur cursor, lines []string) {
	u.action = append(u.action, undoAction{
		typ:    typ,
		cursor: cur,
		lines:  lines,
	})
}

func (u *undo) store(g bool) {
	if g {
		u.global = append(u.global, u.action...)
	} else {
		u.history = append(u.history, u.action)
	}
	u.clear()
}

func (u *undo) storeGlobal() {
	u.history = append(u.history, u.global)
	u.clear()
}
