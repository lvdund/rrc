// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: FlightPathInfoReportConfig-r18 (line 2935).

var flightPathInfoReportConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "maxWayPointNumber-r18"},
		{Name: "includeTimeStamp-r18", Optional: true},
	},
}

var flightPathInfoReportConfigR18MaxWayPointNumberR18Constraints = per.Constrained(1, common.MaxWayPoint_r18)

const (
	FlightPathInfoReportConfig_r18_IncludeTimeStamp_r18_True = 0
)

var flightPathInfoReportConfigR18IncludeTimeStampR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{FlightPathInfoReportConfig_r18_IncludeTimeStamp_r18_True},
}

type FlightPathInfoReportConfig_r18 struct {
	MaxWayPointNumber_r18 int64
	IncludeTimeStamp_r18  *int64
}

func (ie *FlightPathInfoReportConfig_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(flightPathInfoReportConfigR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.IncludeTimeStamp_r18 != nil}); err != nil {
		return err
	}

	// 2. maxWayPointNumber-r18: integer
	{
		if err := e.EncodeInteger(ie.MaxWayPointNumber_r18, flightPathInfoReportConfigR18MaxWayPointNumberR18Constraints); err != nil {
			return err
		}
	}

	// 3. includeTimeStamp-r18: enumerated
	{
		if ie.IncludeTimeStamp_r18 != nil {
			if err := e.EncodeEnumerated(*ie.IncludeTimeStamp_r18, flightPathInfoReportConfigR18IncludeTimeStampR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FlightPathInfoReportConfig_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(flightPathInfoReportConfigR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. maxWayPointNumber-r18: integer
	{
		v0, err := d.DecodeInteger(flightPathInfoReportConfigR18MaxWayPointNumberR18Constraints)
		if err != nil {
			return err
		}
		ie.MaxWayPointNumber_r18 = v0
	}

	// 3. includeTimeStamp-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(flightPathInfoReportConfigR18IncludeTimeStampR18Constraints)
			if err != nil {
				return err
			}
			ie.IncludeTimeStamp_r18 = &idx
		}
	}

	return nil
}
