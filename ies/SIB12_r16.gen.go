// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SIB12-r16 (line 4272).

var sIB12R16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "segmentNumber-r16"},
		{Name: "segmentType-r16"},
		{Name: "segmentContainer-r16"},
	},
}

var sIB12R16SegmentNumberR16Constraints = per.Constrained(0, 63)

const (
	SIB12_r16_SegmentType_r16_NotLastSegment = 0
	SIB12_r16_SegmentType_r16_LastSegment    = 1
)

var sIB12R16SegmentTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SIB12_r16_SegmentType_r16_NotLastSegment, SIB12_r16_SegmentType_r16_LastSegment},
}

var sIB12R16SegmentContainerR16Constraints = per.SizeConstraints{}

type SIB12_r16 struct {
	SegmentNumber_r16    int64
	SegmentType_r16      int64
	SegmentContainer_r16 []byte
}

func (ie *SIB12_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sIB12R16Constraints)
	_ = seq

	// 1. segmentNumber-r16: integer
	{
		if err := e.EncodeInteger(ie.SegmentNumber_r16, sIB12R16SegmentNumberR16Constraints); err != nil {
			return err
		}
	}

	// 2. segmentType-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.SegmentType_r16, sIB12R16SegmentTypeR16Constraints); err != nil {
			return err
		}
	}

	// 3. segmentContainer-r16: octet-string
	{
		if err := e.EncodeOctetString(ie.SegmentContainer_r16, sIB12R16SegmentContainerR16Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SIB12_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sIB12R16Constraints)
	_ = seq

	// 1. segmentNumber-r16: integer
	{
		v0, err := d.DecodeInteger(sIB12R16SegmentNumberR16Constraints)
		if err != nil {
			return err
		}
		ie.SegmentNumber_r16 = v0
	}

	// 2. segmentType-r16: enumerated
	{
		v1, err := d.DecodeEnumerated(sIB12R16SegmentTypeR16Constraints)
		if err != nil {
			return err
		}
		ie.SegmentType_r16 = v1
	}

	// 3. segmentContainer-r16: octet-string
	{
		v2, err := d.DecodeOctetString(sIB12R16SegmentContainerR16Constraints)
		if err != nil {
			return err
		}
		ie.SegmentContainer_r16 = v2
	}

	return nil
}
