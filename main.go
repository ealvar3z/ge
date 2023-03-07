package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
)

type fileType struct {
	ext   string
	_type uint32
}

var fileTypes = []fileType{
	{".c", GE_FILE_TYPE_C},
	{".cpp", GE_FILE_TYPE_C},
	{".h", GE_FILE_TYPE_C},
	{".py", GE_FILE_TYPE_PYTHON},
	{".diff", GE_FILE_TYPE_DIFF},
	{".patch", GE_FILE_TYPE_DIFF},
	{".js", GE_FILE_TYPE_JS},
	{".sh", GE_FILE_TYPE_SHELL},
	{".swift", GE_FILE_TYPE_SWIFT},
	{".yml", GE_FILE_TYPE_YAML},
	{".yaml", GE_FILE_TYPE_YAML},
	{".json", GE_FILE_TYPE_JSON},
	{".html", GE_FILE_TYPE_HTML},
	{".css", GE_FILE_TYPE_CSS},
	{".go", GE_FILE_TYPE_GO},
	{".tex", GE_FILE_TYPE_LATEX},
	{".latex", GE_FILE_TYPE_LATEX},
	{".lua", GE_FILE_TYPE_LUA},
}

var (
	fp       *os.File
	debug    bool
	lameMode bool
)

// my config
var config = struct {
	tabShow   int
	tabWidth  int
	tabExpand int
}{
	tabShow:   1,
	tabWidth:  GE_TAB_WIDTH_DEFAULT,
	tabExpand: GE_TAB_EXPAND_DEFAULT,
}

func die(format string, args ...interface{}) {
	geTermRestore()
	fmt.Fprintf(os.Stderr, format+"\n", args...)
	os.Exit(1)
}

func main() {
	flag.BoolVar(&debug, "d", false, "enable debug mode")
	flag.BoolVar(&lameMode, "l", false, "enable lame mode")
	versionFlag := flag.Bool("v", false, "print version")

	flag.Parse()

	if *versionFlag {
		fmt.Println("go editor 0.1")
		os.Exit(0)
	}

	if debug {
		var err error
		fp, err = os.OpenFile("ge.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			die("failed to open debug log: %v", err)
		}
		logrus.SetOutput(fp)
		logrus.SetLevel(logrus.DebugLevel)
	}

	args := flag.Args()
	if len(args) < 1 {
		die("usage: ge <file> [file...]")
	}
	path, err := filepath.Abs(args[0])
	if err != nil {
		die("failed to get absolute path of %s: %v", args[0], err)
	}
	args[0] = path

	// setup
	geTermSetup()

	// init
	geEditorInit()
	geGameInit()
	geHistInit()
	geBufferInit(args)

	geEditorLoop()

	// cleanup
	geBufferCleanup()
	geTermRestore()

	if debug {
		if err := fp.Close(); err != nil {
			die("failed to close debug log: %v", err)
		}
	}
}

func geFileTypeDetect(buf *gebuf) {
	buf.type_ = GE_FILE_TYPE_PLAIN

	ext := filepath.Ext(buf.path)
	for _, ft := range fileTypes {
		if ext == ft.ext {
			buf.type_ = ft._type
			break
		}
	}
	logrus.Debugf("'%s' is type '%d'", buf.path, buf.type_)
}

func geLameMode() bool { return lame }
