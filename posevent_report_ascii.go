package skypatrolTT8850

import (
	"fmt"
	"github.com/golang/glog"
	"strconv"
	"strings"
	"time"
)

const (
	timestampLayout = "20060102150405"
)

// Parses a position-event report frame. The report pointer contains the common message fields parsed.
// The remaining and unparsed message is in the fields array.
func (parser *AsciiParser) positionEventReport(frame *Frame, report *Report, fields []string) {
	per := &PositionEventReport{}
	per.Report = report

	// ReportID
	if reportId, err := strconv.Atoi(fields[0]); err != nil {
		parseErr := fmt.Errorf("Error parsing ReportId: %s", err)
		parser.NotityError(frame, parseErr)
		return
	} else {
		per.ReportID = reportId
	}

	// ReportType
	if reportType, err := strconv.Atoi(fields[1]); err != nil {
		parseErr := fmt.Errorf("Error parsing ReportType: %s", err)
		parser.NotityError(frame, parseErr)
		return
	} else {
		per.ReportType = reportType
	}

	// Number (always 1)
	if number, err := strconv.Atoi(fields[2]); err != nil {
		parseErr := fmt.Errorf("Error parsing Number: %s", err)
		parser.NotityError(frame, parseErr)
		return
	} else {
		per.Number = number
	}

	// GPS Accuracy
	if gpsAccu, err := strconv.Atoi(fields[3]); err != nil {
		parseErr := fmt.Errorf("Error parsing GPSAccuracy: %s", err)
		parser.NotityError(frame, parseErr)
		return
	} else {
		per.GPSAccuracy = gpsAccu
	}

	// Speed
	if speed, err := strconv.ParseFloat(fields[4], 64); err != nil {
		parseErr := fmt.Errorf("Error parsing Speed: %s", err)
		parser.NotityError(frame, parseErr)
		return
	} else {
		per.Speed = speed
	}

	// Azimuth
	if azimuth, err := strconv.Atoi(fields[5]); err != nil {
		parseErr := fmt.Errorf("Error parsing Azimuth: %s", err)
		parser.NotityError(frame, parseErr)
		return
	} else {
		per.Azimuth = azimuth
	}

	// Altitude
	if altitude, err := strconv.ParseFloat(fields[6], 64); err != nil {
		parseErr := fmt.Errorf("Error parsing Altitude: %s", err)
		parser.NotityError(frame, parseErr)
		return
	} else {
		per.Altitude = altitude
	}

	// Longitude
	if longitude, err := strconv.ParseFloat(fields[7], 64); err != nil {
		parseErr := fmt.Errorf("Error parsing Longitude: %s", err)
		parser.NotityError(frame, parseErr)
		return
	} else {
		per.Longitude = longitude
	}

	// Latitude
	if latitude, err := strconv.ParseFloat(fields[8], 64); err != nil {
		parseErr := fmt.Errorf("Error parsing Latitude: %s", err)
		parser.NotityError(frame, parseErr)
		return
	} else {
		per.Latitude = latitude
	}

	// GPS UTC Time
	if gpsTime, err := time.Parse(timestampLayout, fields[9]); err != nil {
		parseErr := fmt.Errorf("Error parsing GPSUTCTime: %s", err)
		parser.NotityError(frame, parseErr)
		return
	} else {
		per.GPSUTCTime = gpsTime
	}

	// MCC
	if fields[10] != "" { //FIXME: use hexa mask
		if mcc, err := strconv.Atoi(fields[10]); err != nil {
			parseErr := fmt.Errorf("Error parsing MCC: %s", err)
			parser.NotityError(frame, parseErr)
			return
		} else {
			per.MCC = mcc
		}
	}

	// MNC
	if fields[11] != "" { //FIXME: use hexa mask
		if mnc, err := strconv.Atoi(fields[11]); err != nil {
			parseErr := fmt.Errorf("Error parsing MNC: %s", err)
			parser.NotityError(frame, parseErr)
			return
		} else {
			per.MNC = mnc
		}
	}

	// LAC
	per.LAC = fields[12] //FIXME: use hexa mask

	// Cell ID
	per.CellID = fields[13]

	// Battery percentage
	if battPerc, err := strconv.Atoi(fields[14]); err != nil {
		parseErr := fmt.Errorf("Error parsing BatteryPercentage: %s", err)
		parser.NotityError(frame, parseErr)
		return
	} else {
		per.BatteryPercentage = battPerc
	}

	// Send time
	if sendTime, err := time.Parse(timestampLayout, fields[15]); err != nil {
		parseErr := fmt.Errorf("Error parsing SendTime: %s", err)
		parser.NotityError(frame, parseErr)
		return
	} else {
		per.SendTime = sendTime
	}

	// Count number
	cn := strings.Split(fields[16], "$")[0]
	per.CountNumber = cn

	parser.NotityPositionEventReport(frame, per)
}

// Notifies a position-event report to the listener
func (parser *AsciiParser) NotityPositionEventReport(frame *Frame, report *PositionEventReport) {
	if parser.Listener != nil {
		go parser.Listener.PositionEventReport(frame, report)
	} else {
		glog.V(5).Infof("Nil listener, discarding position event report, frame=%+v report=%+v", frame, report)
	}
}
