// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UplinkConfigCommon (line 16299).

var uplinkConfigCommonConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyInfoUL", Optional: true},
		{Name: "initialUplinkBWP", Optional: true},
		{Name: "dummy"},
	},
}

type UplinkConfigCommon struct {
	FrequencyInfoUL  *FrequencyInfoUL
	InitialUplinkBWP *BWP_UplinkCommon
	Dummy            TimeAlignmentTimer
}

func (ie *UplinkConfigCommon) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkConfigCommonConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.FrequencyInfoUL != nil, ie.InitialUplinkBWP != nil}); err != nil {
		return err
	}

	// 2. frequencyInfoUL: ref
	{
		if ie.FrequencyInfoUL != nil {
			if err := ie.FrequencyInfoUL.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. initialUplinkBWP: ref
	{
		if ie.InitialUplinkBWP != nil {
			if err := ie.InitialUplinkBWP.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. dummy: ref
	{
		if err := ie.Dummy.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UplinkConfigCommon) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkConfigCommonConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. frequencyInfoUL: ref
	{
		if seq.IsComponentPresent(0) {
			ie.FrequencyInfoUL = new(FrequencyInfoUL)
			if err := ie.FrequencyInfoUL.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. initialUplinkBWP: ref
	{
		if seq.IsComponentPresent(1) {
			ie.InitialUplinkBWP = new(BWP_UplinkCommon)
			if err := ie.InitialUplinkBWP.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. dummy: ref
	{
		if err := ie.Dummy.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
