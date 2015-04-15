package skypatrolTT8850

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/golang/glog"
	"time"
)

const (
	asciiMsg = 4
	hexaMsg  = 5
)

// A frame to parse
type Frame struct {
	Timestamp time.Time
	Data      []byte
}

// Creates a frame from a byte slice
func NewFrame(data []byte) *Frame {
	frame := &Frame{}
	frame.Timestamp = time.Now()
	frame.Data = data
	return frame
}

// Common report fields
type Report struct {
	MessageHeader   int16
	Mask            int16
	SACK            bool
	MessageType     string
	ProtocolVersion string
	UniqueID        string //IMEI
	DeviceName      string
}

// Position and event report field (check section 3.3.1 of the spec)
type PositionEventReport struct {
	*Report
	ReportID          int
	ReportType        int
	Number            int
	GPSAccuracy       int
	Speed             float64
	Azimuth           int
	Altitude          float64
	Longitude         float64
	Latitude          float64
	GPSUTCTime        time.Time
	MCC               int
	MNC               int
	LAC               string //hexa
	CellID            string //hexa
	BatteryPercentage int
	SendTime          time.Time
	CountNumber       string //hexa
}

// Listener to get notified of parsing result
type ParserListener interface {
	PositionEventReport(frame *Frame, report *PositionEventReport)

	ParsingError(frame *Frame, err error)
}

// Parser, to parse the parseable
type Parser struct {
	Listener ParserListener // get results through this

	asciiParser *AsciiParser //ascii msg parser

	ByteOrder binary.ByteOrder // seems to be big endian
}

// Returns a new parser
func NewParser(listener ParserListener) *Parser {
	asciiParser := newAsciiParser(listener)
	return &Parser{listener, asciiParser, binary.BigEndian}
}

func (parser *Parser) Parse(data []byte) {
	glog.V(5).Infof("parsing frame: %+v", data)

	frame := NewFrame(data)

	// get message header (two bytes)
	var mh uint16
	dr := bytes.NewReader(data)
	err := binary.Read(dr, parser.ByteOrder, &mh)
	if err != nil {
		parser.NotityError(frame, err)
		return
	}

	// parse according to message header
	if mh == asciiMsg {
		parser.asciiParser.parse(frame)
	} else if mh == hexaMsg {
		err := fmt.Errorf("Hexa frames not beign processed, discarding frame %+v", frame)
		parser.NotityError(frame, err)
	} else {
		err := fmt.Errorf("unknown message header %d, discarding frame %+v", mh, frame)
		parser.NotityError(frame, err)
	}
}

// Notifies an error to the listener
func (parser *Parser) NotityError(frame *Frame, err error) {
	if parser.Listener != nil {
		go parser.Listener.ParsingError(frame, err)
	} else {
		glog.V(5).Infof("Nil listener, discarding error, frame=%+v error=%s", frame, err)
	}
}
