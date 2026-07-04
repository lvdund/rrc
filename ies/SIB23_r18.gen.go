// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB23-r18 (line 4632).

var sIB23R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "segmentNumber-r18"},
		{Name: "segmentType-r18"},
		{Name: "segmentContainer-r18"},
	},
}

var sIB23R18SegmentNumberR18Constraints = per.Constrained(0, 63)

const (
	SIB23_r18_SegmentType_r18_NotLastSegment = 0
	SIB23_r18_SegmentType_r18_LastSegment    = 1
)

var sIB23R18SegmentTypeR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB23_r18_SegmentType_r18_NotLastSegment, SIB23_r18_SegmentType_r18_LastSegment},
}

var sIB23R18SegmentContainerR18Constraints = per.SizeConstraints{}

type SIB23_r18 struct {
	SegmentNumber_r18    int64
	SegmentType_r18      int64
	SegmentContainer_r18 []byte
}

func (ie *SIB23_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB23R18Constraints)
	_ = seq

	// 1. segmentNumber-r18: integer
	{
		if err := e.EncodeInteger(ie.SegmentNumber_r18, sIB23R18SegmentNumberR18Constraints); err != nil {
			return err
		}
	}

	// 2. segmentType-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.SegmentType_r18, sIB23R18SegmentTypeR18Constraints); err != nil {
			return err
		}
	}

	// 3. segmentContainer-r18: octet-string
	{
		if err := e.EncodeOctetString(ie.SegmentContainer_r18, sIB23R18SegmentContainerR18Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SIB23_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB23R18Constraints)
	_ = seq

	// 1. segmentNumber-r18: integer
	{
		v0, err := d.DecodeInteger(sIB23R18SegmentNumberR18Constraints)
		if err != nil {
			return err
		}
		ie.SegmentNumber_r18 = v0
	}

	// 2. segmentType-r18: enumerated
	{
		v1, err := d.DecodeEnumerated(sIB23R18SegmentTypeR18Constraints)
		if err != nil {
			return err
		}
		ie.SegmentType_r18 = v1
	}

	// 3. segmentContainer-r18: octet-string
	{
		v2, err := d.DecodeOctetString(sIB23R18SegmentContainerR18Constraints)
		if err != nil {
			return err
		}
		ie.SegmentContainer_r18 = v2
	}

	return nil
}
