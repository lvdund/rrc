// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: WayPoint-r18 (line 2984).

var wayPointR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "wayPointLocation-r18"},
		{Name: "timeStamp-r18", Optional: true},
	},
}

var wayPointR18WayPointLocationR18Constraints = per.SizeConstraints{}

type WayPoint_r18 struct {
	WayPointLocation_r18 []byte
	TimeStamp_r18        *AbsoluteTimeInfo_r16
}

func (ie *WayPoint_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(wayPointR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.TimeStamp_r18 != nil}); err != nil {
		return err
	}

	// 2. wayPointLocation-r18: octet-string
	{
		if err := e.EncodeOctetString(ie.WayPointLocation_r18, wayPointR18WayPointLocationR18Constraints); err != nil {
			return err
		}
	}

	// 3. timeStamp-r18: ref
	{
		if ie.TimeStamp_r18 != nil {
			if err := ie.TimeStamp_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *WayPoint_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(wayPointR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. wayPointLocation-r18: octet-string
	{
		v0, err := d.DecodeOctetString(wayPointR18WayPointLocationR18Constraints)
		if err != nil {
			return err
		}
		ie.WayPointLocation_r18 = v0
	}

	// 3. timeStamp-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.TimeStamp_r18 = new(AbsoluteTimeInfo_r16)
			if err := ie.TimeStamp_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
