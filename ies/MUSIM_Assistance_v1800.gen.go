// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-Assistance-v1800 (line 2616).

var mUSIMAssistanceV1800Constraints = per.SequenceConstraints{
	Extensible: true,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-GapPriorityPreferenceList-r18", Optional: true},
		{Name: "musim-GapKeepPreference-r18", Optional: true},
		{Name: "musim-CapRestriction-r18", Optional: true},
		{Name: "musim-NeedForGapsInfoNR-r18", Optional: true},
	},
}

const (
	MUSIM_Assistance_v1800_Musim_GapKeepPreference_r18_True = 0
)

var mUSIMAssistanceV1800MusimGapKeepPreferenceR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MUSIM_Assistance_v1800_Musim_GapKeepPreference_r18_True},
}

type MUSIM_Assistance_v1800 struct {
	Musim_GapPriorityPreferenceList_r18 *MUSIM_GapPriorityPreferenceList_r18
	Musim_GapKeepPreference_r18         *int64
	Musim_CapRestriction_r18            *MUSIM_CapRestriction_r18
	Musim_NeedForGapsInfoNR_r18         *NeedForGapsInfoNR_r16
}

func (ie *MUSIM_Assistance_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMAssistanceV1800Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.EncodeExtensionBit(false); err != nil {
		return err
	}

	// 2. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Musim_GapPriorityPreferenceList_r18 != nil, ie.Musim_GapKeepPreference_r18 != nil, ie.Musim_CapRestriction_r18 != nil, ie.Musim_NeedForGapsInfoNR_r18 != nil}); err != nil {
		return err
	}

	// 3. musim-GapPriorityPreferenceList-r18: ref
	{
		if ie.Musim_GapPriorityPreferenceList_r18 != nil {
			if err := ie.Musim_GapPriorityPreferenceList_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. musim-GapKeepPreference-r18: enumerated
	{
		if ie.Musim_GapKeepPreference_r18 != nil {
			if err := e.EncodeEnumerated(*ie.Musim_GapKeepPreference_r18, mUSIMAssistanceV1800MusimGapKeepPreferenceR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. musim-CapRestriction-r18: ref
	{
		if ie.Musim_CapRestriction_r18 != nil {
			if err := ie.Musim_CapRestriction_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. musim-NeedForGapsInfoNR-r18: ref
	{
		if ie.Musim_NeedForGapsInfoNR_r18 != nil {
			if err := ie.Musim_NeedForGapsInfoNR_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MUSIM_Assistance_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMAssistanceV1800Constraints)

	// 1. SEQUENCE extension bit.
	if err := seq.DecodeExtensionBit(); err != nil {
		return err
	}

	// 2. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 3. musim-GapPriorityPreferenceList-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Musim_GapPriorityPreferenceList_r18 = new(MUSIM_GapPriorityPreferenceList_r18)
			if err := ie.Musim_GapPriorityPreferenceList_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. musim-GapKeepPreference-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(mUSIMAssistanceV1800MusimGapKeepPreferenceR18Constraints)
			if err != nil {
				return err
			}
			ie.Musim_GapKeepPreference_r18 = &idx
		}
	}

	// 5. musim-CapRestriction-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Musim_CapRestriction_r18 = new(MUSIM_CapRestriction_r18)
			if err := ie.Musim_CapRestriction_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. musim-NeedForGapsInfoNR-r18: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Musim_NeedForGapsInfoNR_r18 = new(NeedForGapsInfoNR_r16)
			if err := ie.Musim_NeedForGapsInfoNR_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
