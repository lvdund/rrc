// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: ERedCapParameters-r18 (line 19413).

var eRedCapParametersR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportOfERedCap-r18"},
		{Name: "eRedCapNotReducedBB-BW-r18", Optional: true},
		{Name: "eRedCapIgnoreCapabilityFiltering-r18", Optional: true},
	},
}

const (
	ERedCapParameters_r18_SupportOfERedCap_r18_Supported = 0
)

var eRedCapParametersR18SupportOfERedCapR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ERedCapParameters_r18_SupportOfERedCap_r18_Supported},
}

const (
	ERedCapParameters_r18_ERedCapNotReducedBB_BW_r18_Supported = 0
)

var eRedCapParametersR18ERedCapNotReducedBBBWR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ERedCapParameters_r18_ERedCapNotReducedBB_BW_r18_Supported},
}

const (
	ERedCapParameters_r18_ERedCapIgnoreCapabilityFiltering_r18_Supported = 0
)

var eRedCapParametersR18ERedCapIgnoreCapabilityFilteringR18Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{ERedCapParameters_r18_ERedCapIgnoreCapabilityFiltering_r18_Supported},
}

type ERedCapParameters_r18 struct {
	SupportOfERedCap_r18                 int64
	ERedCapNotReducedBB_BW_r18           *int64
	ERedCapIgnoreCapabilityFiltering_r18 *int64
}

func (ie *ERedCapParameters_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(eRedCapParametersR18Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ERedCapNotReducedBB_BW_r18 != nil, ie.ERedCapIgnoreCapabilityFiltering_r18 != nil}); err != nil {
		return err
	}

	// 2. supportOfERedCap-r18: enumerated
	{
		if err := e.EncodeEnumerated(ie.SupportOfERedCap_r18, eRedCapParametersR18SupportOfERedCapR18Constraints); err != nil {
			return err
		}
	}

	// 3. eRedCapNotReducedBB-BW-r18: enumerated
	{
		if ie.ERedCapNotReducedBB_BW_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ERedCapNotReducedBB_BW_r18, eRedCapParametersR18ERedCapNotReducedBBBWR18Constraints); err != nil {
				return err
			}
		}
	}

	// 4. eRedCapIgnoreCapabilityFiltering-r18: enumerated
	{
		if ie.ERedCapIgnoreCapabilityFiltering_r18 != nil {
			if err := e.EncodeEnumerated(*ie.ERedCapIgnoreCapabilityFiltering_r18, eRedCapParametersR18ERedCapIgnoreCapabilityFilteringR18Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *ERedCapParameters_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(eRedCapParametersR18Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportOfERedCap-r18: enumerated
	{
		v0, err := d.DecodeEnumerated(eRedCapParametersR18SupportOfERedCapR18Constraints)
		if err != nil {
			return err
		}
		ie.SupportOfERedCap_r18 = v0
	}

	// 3. eRedCapNotReducedBB-BW-r18: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(eRedCapParametersR18ERedCapNotReducedBBBWR18Constraints)
			if err != nil {
				return err
			}
			ie.ERedCapNotReducedBB_BW_r18 = &idx
		}
	}

	// 4. eRedCapIgnoreCapabilityFiltering-r18: enumerated
	{
		if seq.IsComponentPresent(2) {
			idx, err := d.DecodeEnumerated(eRedCapParametersR18ERedCapIgnoreCapabilityFilteringR18Constraints)
			if err != nil {
				return err
			}
			ie.ERedCapIgnoreCapabilityFiltering_r18 = &idx
		}
	}

	return nil
}
