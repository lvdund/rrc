// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ULDedicatedMessageSegment-r16-IEs (line 3622).

var uLDedicatedMessageSegmentR16IEsConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "segmentNumber-r16"},
		{Name: "rrc-MessageSegmentContainer-r16"},
		{Name: "rrc-MessageSegmentType-r16"},
		{Name: "lateNonCriticalExtension", Optional: true},
		{Name: "nonCriticalExtension", Optional: true},
	},
}

var uLDedicatedMessageSegmentR16IEsSegmentNumberR16Constraints = per.Constrained(0, 15)

var uLDedicatedMessageSegmentR16IEsRrcMessageSegmentContainerR16Constraints = per.SizeConstraints{}

const (
	ULDedicatedMessageSegment_r16_IEs_Rrc_MessageSegmentType_r16_NotLastSegment = 0
	ULDedicatedMessageSegment_r16_IEs_Rrc_MessageSegmentType_r16_LastSegment    = 1
)

var uLDedicatedMessageSegmentR16IEsRrcMessageSegmentTypeR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ULDedicatedMessageSegment_r16_IEs_Rrc_MessageSegmentType_r16_NotLastSegment, ULDedicatedMessageSegment_r16_IEs_Rrc_MessageSegmentType_r16_LastSegment},
}

var uLDedicatedMessageSegmentR16IEsLateNonCriticalExtensionConstraints = per.SizeConstraints{}

type ULDedicatedMessageSegment_r16_IEs struct {
	SegmentNumber_r16               int64
	Rrc_MessageSegmentContainer_r16 []byte
	Rrc_MessageSegmentType_r16      int64
	LateNonCriticalExtension        []byte
	NonCriticalExtension            *struct{}
}

func (ie *ULDedicatedMessageSegment_r16_IEs) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLDedicatedMessageSegmentR16IEsConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.LateNonCriticalExtension != nil, ie.NonCriticalExtension != nil}); err != nil {
		return err
	}

	// 2. segmentNumber-r16: integer
	{
		if err := e.EncodeInteger(ie.SegmentNumber_r16, uLDedicatedMessageSegmentR16IEsSegmentNumberR16Constraints); err != nil {
			return err
		}
	}

	// 3. rrc-MessageSegmentContainer-r16: octet-string
	{
		if err := e.EncodeOctetString(ie.Rrc_MessageSegmentContainer_r16, uLDedicatedMessageSegmentR16IEsRrcMessageSegmentContainerR16Constraints); err != nil {
			return err
		}
	}

	// 4. rrc-MessageSegmentType-r16: enumerated
	{
		if err := e.EncodeEnumerated(ie.Rrc_MessageSegmentType_r16, uLDedicatedMessageSegmentR16IEsRrcMessageSegmentTypeR16Constraints); err != nil {
			return err
		}
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if ie.LateNonCriticalExtension != nil {
			if err := e.EncodeOctetString(ie.LateNonCriticalExtension, uLDedicatedMessageSegmentR16IEsLateNonCriticalExtensionConstraints); err != nil {
				return err
			}
		}
	}

	// 6. nonCriticalExtension: sequence
	{
		if ie.NonCriticalExtension != nil {
			// empty SEQUENCE {}: no content to encode.
		}
	}

	return nil
}

func (ie *ULDedicatedMessageSegment_r16_IEs) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLDedicatedMessageSegmentR16IEsConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. segmentNumber-r16: integer
	{
		v0, err := d.DecodeInteger(uLDedicatedMessageSegmentR16IEsSegmentNumberR16Constraints)
		if err != nil {
			return err
		}
		ie.SegmentNumber_r16 = v0
	}

	// 3. rrc-MessageSegmentContainer-r16: octet-string
	{
		v1, err := d.DecodeOctetString(uLDedicatedMessageSegmentR16IEsRrcMessageSegmentContainerR16Constraints)
		if err != nil {
			return err
		}
		ie.Rrc_MessageSegmentContainer_r16 = v1
	}

	// 4. rrc-MessageSegmentType-r16: enumerated
	{
		v2, err := d.DecodeEnumerated(uLDedicatedMessageSegmentR16IEsRrcMessageSegmentTypeR16Constraints)
		if err != nil {
			return err
		}
		ie.Rrc_MessageSegmentType_r16 = v2
	}

	// 5. lateNonCriticalExtension: octet-string
	{
		if seq.IsComponentPresent(3) {
			v, err := d.DecodeOctetString(uLDedicatedMessageSegmentR16IEsLateNonCriticalExtensionConstraints)
			if err != nil {
				return err
			}
			ie.LateNonCriticalExtension = v
		}
	}

	// 6. nonCriticalExtension: sequence
	{
		if seq.IsComponentPresent(4) {
			ie.NonCriticalExtension = &struct{}{}
		}
	}

	return nil
}
