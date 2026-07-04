// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SI-SchedulingInfo-v1800 (line 15075).

var sISchedulingInfoV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "si-RequestConfigMSG1-Repetition-r18", Optional: true},
		{Name: "si-RequestConfigRedCap-MSG1-Repetition-r18", Optional: true},
		{Name: "si-RequestConfigSUL-MSG1-Repetition-r18", Optional: true},
	},
}

type SI_SchedulingInfo_v1800 struct {
	Si_RequestConfigMSG1_Repetition_r18        *SI_RequestConfigRepetition_r18
	Si_RequestConfigRedCap_MSG1_Repetition_r18 *SI_RequestConfigRepetition_r18
	Si_RequestConfigSUL_MSG1_Repetition_r18    *SI_RequestConfigRepetition_r18
}

func (ie *SI_SchedulingInfo_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sISchedulingInfoV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Si_RequestConfigMSG1_Repetition_r18 != nil, ie.Si_RequestConfigRedCap_MSG1_Repetition_r18 != nil, ie.Si_RequestConfigSUL_MSG1_Repetition_r18 != nil}); err != nil {
		return err
	}

	// 2. si-RequestConfigMSG1-Repetition-r18: ref
	{
		if ie.Si_RequestConfigMSG1_Repetition_r18 != nil {
			if err := ie.Si_RequestConfigMSG1_Repetition_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. si-RequestConfigRedCap-MSG1-Repetition-r18: ref
	{
		if ie.Si_RequestConfigRedCap_MSG1_Repetition_r18 != nil {
			if err := ie.Si_RequestConfigRedCap_MSG1_Repetition_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. si-RequestConfigSUL-MSG1-Repetition-r18: ref
	{
		if ie.Si_RequestConfigSUL_MSG1_Repetition_r18 != nil {
			if err := ie.Si_RequestConfigSUL_MSG1_Repetition_r18.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SI_SchedulingInfo_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sISchedulingInfoV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. si-RequestConfigMSG1-Repetition-r18: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Si_RequestConfigMSG1_Repetition_r18 = new(SI_RequestConfigRepetition_r18)
			if err := ie.Si_RequestConfigMSG1_Repetition_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. si-RequestConfigRedCap-MSG1-Repetition-r18: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Si_RequestConfigRedCap_MSG1_Repetition_r18 = new(SI_RequestConfigRepetition_r18)
			if err := ie.Si_RequestConfigRedCap_MSG1_Repetition_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. si-RequestConfigSUL-MSG1-Repetition-r18: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Si_RequestConfigSUL_MSG1_Repetition_r18 = new(SI_RequestConfigRepetition_r18)
			if err := ie.Si_RequestConfigSUL_MSG1_Repetition_r18.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
