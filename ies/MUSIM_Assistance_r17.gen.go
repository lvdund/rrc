// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MUSIM-Assistance-r17 (line 2608).

var mUSIMAssistanceR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "musim-PreferredRRC-State-r17", Optional: true},
		{Name: "musim-GapPreferenceList-r17", Optional: true},
	},
}

const (
	MUSIM_Assistance_r17_Musim_PreferredRRC_State_r17_Idle           = 0
	MUSIM_Assistance_r17_Musim_PreferredRRC_State_r17_Inactive       = 1
	MUSIM_Assistance_r17_Musim_PreferredRRC_State_r17_OutOfConnected = 2
)

var mUSIMAssistanceR17MusimPreferredRRCStateR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{MUSIM_Assistance_r17_Musim_PreferredRRC_State_r17_Idle, MUSIM_Assistance_r17_Musim_PreferredRRC_State_r17_Inactive, MUSIM_Assistance_r17_Musim_PreferredRRC_State_r17_OutOfConnected},
}

type MUSIM_Assistance_r17 struct {
	Musim_PreferredRRC_State_r17 *int64
	Musim_GapPreferenceList_r17  *MUSIM_GapPreferenceList_r17
}

func (ie *MUSIM_Assistance_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(mUSIMAssistanceR17Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Musim_PreferredRRC_State_r17 != nil, ie.Musim_GapPreferenceList_r17 != nil}); err != nil {
		return err
	}

	// 2. musim-PreferredRRC-State-r17: enumerated
	{
		if ie.Musim_PreferredRRC_State_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Musim_PreferredRRC_State_r17, mUSIMAssistanceR17MusimPreferredRRCStateR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. musim-GapPreferenceList-r17: ref
	{
		if ie.Musim_GapPreferenceList_r17 != nil {
			if err := ie.Musim_GapPreferenceList_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *MUSIM_Assistance_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(mUSIMAssistanceR17Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. musim-PreferredRRC-State-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(mUSIMAssistanceR17MusimPreferredRRCStateR17Constraints)
			if err != nil {
				return err
			}
			ie.Musim_PreferredRRC_State_r17 = &idx
		}
	}

	// 3. musim-GapPreferenceList-r17: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Musim_GapPreferenceList_r17 = new(MUSIM_GapPreferenceList_r17)
			if err := ie.Musim_GapPreferenceList_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
