package skypatrolTT8850

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/golang/glog"
	"io/ioutil"
	"strings"
)

type AsciiParser struct {
	Listener ParserListener
}

func newAsciiParser(listener ParserListener) *AsciiParser {
	return &AsciiParser{listener}
}

func (parser *AsciiParser) parse(frame *Frame) {
	glog.V(5).Infof("parsing ascii position event report: %+v", frame)

	report := &Report{}
	report.MessageHeader = asciiMsg

	dr := bytes.NewReader(frame.Data[2:]) // skip message header
	allData, err := ioutil.ReadAll(dr)
	if err != nil {
		parser.NotityError(frame, err)
		return
	}

	asciiMsg := string(allData)
	fields := strings.Split(asciiMsg, ",")
	for i, val := range fields {
		glog.V(5).Infof("fields[%d] = %v", i, val)
	}

	// TODO
	sz, _ := hex.DecodeString(fields[1])
	glog.V(5).Infof("hexa mask: %+v", sz)

	report.SACK = fields[2] == "1"
	report.MessageType = fields[3]
	report.ProtocolVersion = fields[4]
	report.UniqueID = fields[5]
	report.DeviceName = fields[6]

	switch report.MessageType {
	case "GTFRI":
		parser.positionEventReport(frame, report, fields[7:])
	default:
		err := fmt.Errorf("unknown ascii message: %v", report.MessageType)
		parser.NotityError(frame, err)
	}
}

// Notifies an error to the listener
func (parser *AsciiParser) NotityError(frame *Frame, err error) {
	if parser.Listener != nil {
		go parser.Listener.ParsingError(frame, err)
	} else {
		glog.V(5).Infof("Nil listener, discarding error, frame=%+v error=%s", frame, err)
	}
}
