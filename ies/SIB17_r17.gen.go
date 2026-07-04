// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB17-r17 (line 4407).

var sIB17R17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "segmentNumber-r17"},
		{Name: "segmentType-r17"},
		{Name: "segmentContainer-r17"},
	},
}

var sIB17R17SegmentNumberR17Constraints = per.Constrained(0, 63)

const (
	SIB17_r17_SegmentType_r17_NotLastSegment = 0
	SIB17_r17_SegmentType_r17_LastSegment    = 1
)

var sIB17R17SegmentTypeR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB17_r17_SegmentType_r17_NotLastSegment, SIB17_r17_SegmentType_r17_LastSegment},
}

var sIB17R17SegmentContainerR17Constraints = per.SizeConstraints{}

type SIB17_r17 struct {
	SegmentNumber_r17    int64
	SegmentType_r17      int64
	SegmentContainer_r17 []byte
}

func (ie *SIB17_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB17R17Constraints)
	_ = seq

	// 1. segmentNumber-r17: integer
	{
		if err := e.EncodeInteger(ie.SegmentNumber_r17, sIB17R17SegmentNumberR17Constraints); err != nil {
			return err
		}
	}

	// 2. segmentType-r17: enumerated
	{
		if err := e.EncodeEnumerated(ie.SegmentType_r17, sIB17R17SegmentTypeR17Constraints); err != nil {
			return err
		}
	}

	// 3. segmentContainer-r17: octet-string
	{
		if err := e.EncodeOctetString(ie.SegmentContainer_r17, sIB17R17SegmentContainerR17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SIB17_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB17R17Constraints)
	_ = seq

	// 1. segmentNumber-r17: integer
	{
		v0, err := d.DecodeInteger(sIB17R17SegmentNumberR17Constraints)
		if err != nil {
			return err
		}
		ie.SegmentNumber_r17 = v0
	}

	// 2. segmentType-r17: enumerated
	{
		v1, err := d.DecodeEnumerated(sIB17R17SegmentTypeR17Constraints)
		if err != nil {
			return err
		}
		ie.SegmentType_r17 = v1
	}

	// 3. segmentContainer-r17: octet-string
	{
		v2, err := d.DecodeOctetString(sIB17R17SegmentContainerR17Constraints)
		if err != nil {
			return err
		}
		ie.SegmentContainer_r17 = v2
	}

	return nil
}
