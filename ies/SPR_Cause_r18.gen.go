// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SPR-Cause-r18 (line 3478).

var sPRCauseR18Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "t304-cause-r18", Optional: true},
		{Name: "t310-cause-r18", Optional: true},
		{Name: "t312-cause-r18", Optional: true},
	},
}

const (
	SPR_Cause_r18_T304_Cause_r18_True = 0
)

var sPRCauseR18T304CauseR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SPR_Cause_r18_T304_Cause_r18_True},
}

const (
	SPR_Cause_r18_T310_Cause_r18_True = 0
)

var sPRCauseR18T310CauseR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SPR_Cause_r18_T310_Cause_r18_True},
}

const (
	SPR_Cause_r18_T312_Cause_r18_True = 0
)

var sPRCauseR18T312CauseR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SPR_Cause_r18_T312_Cause_r18_True},
}

type SPR_Cause_r18 struct {
	T304_Cause_r18 *int64
	T310_Cause_r18 *int64
	T312_Cause_r18 *int64
}

func (ie *SPR_Cause_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sPRCauseR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.T304_Cause_r18 != nil, ie.T310_Cause_r18 != nil, ie.T312_Cause_r18 != nil}); err != nil {
		return err
	}

	// 3. t304-cause-r18: enumerated
	{
		if ie.T304_Cause_r18 != nil {
			if err := e.EncodeEnumerated(*ie.T304_Cause_r18, sPRCauseR18T304CauseR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. t310-cause-r18: enumerated
	{
		if ie.T310_Cause_r18 != nil {
			if err := e.EncodeEnumerated(*ie.T310_Cause_r18, sPRCauseR18T310CauseR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. t312-cause-r18: enumerated
	{
		if ie.T312_Cause_r18 != nil {
			if err := e.EncodeEnumerated(*ie.T312_Cause_r18, sPRCauseR18T312CauseR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SPR_Cause_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sPRCauseR18Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. t304-cause-r18: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sPRCauseR18T304CauseR18Constraints)
			if err != nil {
				return err
			}
			ie.T304_Cause_r18 = &idx
		}
	}

	// 4. t310-cause-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sPRCauseR18T310CauseR18Constraints)
			if err != nil {
				return err
			}
			ie.T310_Cause_r18 = &idx
		}
	}

	// 5. t312-cause-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(sPRCauseR18T312CauseR18Constraints)
			if err != nil {
				return err
			}
			ie.T312_Cause_r18 = &idx
		}
	}

	return nil
}
