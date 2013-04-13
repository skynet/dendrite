package dendrite

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

var config *SourceConfig = nil
var group *TailGroup = nil

var _tg_init = func() {
	config = new(SourceConfig)
	config.Glob = "./data/solr*txt"
	config.Pattern = "(?P<line>.+)\n"
	config.OffsetDir = "tmp"
	_ = os.RemoveAll(config.OffsetDir)
	os.Mkdir(config.OffsetDir, 0777)
	output = make(chan Record, 100000)
	matches, _ := filepath.Glob(config.Glob)
	for _, m := range matches {
		os.Chtimes(m, time.Now(), time.Now())
	}

	group = NewTailGroup(*config, output)
}

func TestGroupHasTails(t *testing.T) {
	_tg_init()
	n := len(group.Tails)
	if n != 2 {
		t.Errorf("group has %d tails", n)
	}
}

func TestGroupCanPoll(t *testing.T) {
	_tg_init()
	group.Poll()
}
