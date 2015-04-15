package skypatrolTT8850

import (
	"github.com/golang/glog"
	. "gopkg.in/check.v1"
	"sync"
	"testing"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	TestingT(t)
}

// Listener implementation for testing purposes
type TestListener struct {
	Wg                  sync.WaitGroup       // Call wait on this before any assertion
	ExpectPositionEvent bool                 // true if a position-event report is expected
	ExpectError         bool                 // true if an parsing error is expected
	Frame               *Frame               // testing frame
	Report              *PositionEventReport // notified report
	Error               error                // notified error
	C                   *C                   // testing context
}

// Creates a new listener, wich expects an position-event report
func NewExpectPositionEventReportListener(c *C) *TestListener {
	listener := &TestListener{}
	listener.C = c
	listener.Wg.Add(1)
	listener.ExpectPositionEvent = true
	return listener
}

// Creates a new listenes, wich expects a parsing error
func NewExpectErrorListener(c *C) *TestListener {
	listener := &TestListener{}
	listener.C = c
	listener.Wg.Add(1)
	listener.ExpectError = true
	return listener
}

func (s *TestListener) PositionEventReport(frame *Frame, report *PositionEventReport) {
	if !s.ExpectPositionEvent {
		s.C.Errorf("No position-event report expected for frame %+v", frame)
	}
	s.Frame = frame
	s.Report = report
	glog.Infof("PositionEventReport:\n\tframe: %+v\n\treport: %+v\n\tpos-event report: %+v", frame, report.Report, report)
	s.Wg.Done()
}

func (s *TestListener) ParsingError(frame *Frame, err error) {
	if !s.ExpectError {
		s.C.Errorf("No error expected for frame %+v", frame)
	}
	s.Frame = frame
	s.Error = err
	glog.Infof("ParsingError:\n\tframe: %+v\n\terror: %+v", frame, err)
	s.Wg.Done()
}
