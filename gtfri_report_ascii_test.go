package skypatrolTT8850

import (
	"encoding/hex"
	. "gopkg.in/check.v1"
)

func (s *AsciiParserSuite) TestParse(c *C) {
	// test frame
	hexaMsg := "00042C303035372C302C47544652492C3032303130332C3836373834343030313739343333312C2C302C302C312C312C312E312C302C3636302E362C2D37302E3538323034372C2D33332E3431343735342C32303135303131343139343033322C2C2C2C2C3130302C32303135303131343139343033332C3236394524"
	data, err := hex.DecodeString(hexaMsg)
	c.Assert(err, IsNil)

	// test listener
	listener := NewExpectPositionEventReportListener(c)
	parser := newAsciiParser(listener)

	// test
	frame := NewFrame(data)
	parser.parse(frame)

	// wait listener
	listener.Wg.Wait()

	// check report
	r := listener.Report
	c.Assert(r, NotNil)
	c.Assert(r.MessageHeader, Equals, int16(4))
	c.Assert(r.Mask, Equals, int16(0)) //TODO
	c.Assert(r.SACK, Equals, false)
	c.Assert(r.MessageType, Equals, "GTFRI")
	c.Assert(r.ProtocolVersion, Equals, "020103")
	c.Assert(r.UniqueID, Equals, "867844001794331")
	c.Assert(r.DeviceName, Equals, "")
	c.Assert(r.ReportID, Equals, 0)
	c.Assert(r.ReportType, Equals, 0)
	c.Assert(r.Number, Equals, 1)
	c.Assert(r.GPSAccuracy, Equals, 1)
	c.Assert(r.Speed, Equals, float64(1.1))
	c.Assert(r.Azimuth, Equals, 0)
	c.Assert(r.Altitude, Equals, float64(660.6))
	c.Assert(r.Longitude, Equals, float64(-70.582047))
	c.Assert(r.Latitude, Equals, float64(-33.414754))
	c.Assert(r.GPSUTCTime.Format(timestampLayout), Equals, "20150114194032")
	c.Assert(r.MCC, Equals, 0)
	c.Assert(r.MNC, Equals, 0)
	c.Assert(r.LAC, Equals, "")
	c.Assert(r.CellID, Equals, "")
	c.Assert(r.BatteryPercentage, Equals, 100)
	c.Assert(r.SendTime.Format(timestampLayout), Equals, "20150114194033")
	c.Assert(r.CountNumber, Equals, "269E")
}

func (s *AsciiParserSuite) TestParseGTFRI(c *C) {
	// test frame
	hexaMsg := "00042c303035462c302c47544652492c3032303130302c3133353739303234363831313232302c2c302c302c312c312c342e332c39322c37302e302c3132312e3335343333352c33312e3232323037332c32303039303231343031333235342c303436302c303030302c313864382c363134312c39302c32303039303231343039333235342c3131463024"
	data, err := hex.DecodeString(hexaMsg)
	c.Assert(err, IsNil)

	// test listener
	listener := NewExpectPositionEventReportListener(c)
	parser := newAsciiParser(listener)

	// test
	frame := NewFrame(data)
	parser.parse(frame)

	// wait listener
	listener.Wg.Wait()

	// check report
	r := listener.Report
	c.Assert(r, NotNil)
	c.Assert(r.MessageHeader, Equals, int16(4))
	c.Assert(r.Mask, Equals, int16(0)) //TODO
	c.Assert(r.SACK, Equals, false)
	c.Assert(r.MessageType, Equals, "GTFRI")
	c.Assert(r.ProtocolVersion, Equals, "020100")
	c.Assert(r.UniqueID, Equals, "135790246811220")
	c.Assert(r.DeviceName, Equals, "")
	c.Assert(r.ReportID, Equals, 0)
	c.Assert(r.ReportType, Equals, 0)
	c.Assert(r.Number, Equals, 1)
	c.Assert(r.GPSAccuracy, Equals, 1)
	c.Assert(r.Speed, Equals, float64(4.3))
	c.Assert(r.Azimuth, Equals, 92)
	c.Assert(r.Altitude, Equals, float64(70.0))
	c.Assert(r.Longitude, Equals, float64(121.354335))
	c.Assert(r.Latitude, Equals, float64(31.222073))
	c.Assert(r.GPSUTCTime.Format(timestampLayout), Equals, "20090214013254")
	c.Assert(r.MCC, Equals, 460)
	c.Assert(r.MNC, Equals, 0)
	c.Assert(r.LAC, Equals, "18d8")
	c.Assert(r.CellID, Equals, "6141")
	c.Assert(r.BatteryPercentage, Equals, 90)
	c.Assert(r.SendTime.Format(timestampLayout), Equals, "20090214093254")
	c.Assert(r.CountNumber, Equals, "11F0")
}
