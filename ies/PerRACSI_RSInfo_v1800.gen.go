// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PerRACSI-RSInfo-v1800 (line 3191).

var perRACSIRSInfoV1800Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "allPreamblesBlocked", Optional: true},
		{Name: "lbt-Detected-r18", Optional: true},
	},
}

const (
	PerRACSI_RSInfo_v1800_AllPreamblesBlocked_True = 0
)

var perRACSIRSInfoV1800AllPreamblesBlockedConstraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PerRACSI_RSInfo_v1800_AllPreamblesBlocked_True},
}

const (
	PerRACSI_RSInfo_v1800_Lbt_Detected_r18_True = 0
)

var perRACSIRSInfoV1800LbtDetectedR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PerRACSI_RSInfo_v1800_Lbt_Detected_r18_True},
}

type PerRACSI_RSInfo_v1800 struct {
	AllPreamblesBlocked *int64
	Lbt_Detected_r18    *int64
}

func (ie *PerRACSI_RSInfo_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(perRACSIRSInfoV1800Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.AllPreamblesBlocked != nil, ie.Lbt_Detected_r18 != nil}); err != nil {
		return err
	}

	// 3. allPreamblesBlocked: enumerated
	{
		if ie.AllPreamblesBlocked != nil {
			if err := e.EncodeEnumerated(*ie.AllPreamblesBlocked, perRACSIRSInfoV1800AllPreamblesBlockedConstraints); err != nil {
				return err
			}
		}
	}

	// 4. lbt-Detected-r18: enumerated
	{
		if ie.Lbt_Detected_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Lbt_Detected_r18, perRACSIRSInfoV1800LbtDetectedR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PerRACSI_RSInfo_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(perRACSIRSInfoV1800Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. allPreamblesBlocked: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(perRACSIRSInfoV1800AllPreamblesBlockedConstraints)
			if err != nil {
				return err
			}
			ie.AllPreamblesBlocked = &idx
		}
	}

	// 4. lbt-Detected-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(perRACSIRSInfoV1800LbtDetectedR18Constraints)
			if err != nil {
				return err
			}
			ie.Lbt_Detected_r18 = &idx
		}
	}

	return nil
}
