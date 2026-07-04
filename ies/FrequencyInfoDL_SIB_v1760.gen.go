// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FrequencyInfoDL-SIB-v1760 (line 8418).

var frequencyInfoDLSIBV1760Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyBandList-v1760"},
	},
}

type FrequencyInfoDL_SIB_v1760 struct {
	FrequencyBandList_v1760 MultiFrequencyBandListNR_SIB_v1760
}

func (ie *FrequencyInfoDL_SIB_v1760) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(frequencyInfoDLSIBV1760Constraints)
	_ = seq

	// 1. frequencyBandList-v1760: ref
	{
		if err := ie.FrequencyBandList_v1760.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *FrequencyInfoDL_SIB_v1760) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(frequencyInfoDLSIBV1760Constraints)
	_ = seq

	// 1. frequencyBandList-v1760: ref
	{
		if err := ie.FrequencyBandList_v1760.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
