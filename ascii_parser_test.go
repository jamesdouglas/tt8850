package skypatrolTT8850

import (
	"flag"
	. "gopkg.in/check.v1"
)

type AsciiParserSuite struct{}

func (s *AsciiParserSuite) SetUpSuite(c *C) {
	// activate glog
	flag.Set("alsologtostderr", "true")
	flag.Set("log_dir", ".")
	flag.Set("v", "10")
}

var _ = Suite(&AsciiParserSuite{})
