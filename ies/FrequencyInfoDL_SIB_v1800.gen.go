// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FrequencyInfoDL-SIB-v1800 (line 8422).

var frequencyInfoDLSIBV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyBandListAerial-r18"},
	},
}

type FrequencyInfoDL_SIB_v1800 struct {
	FrequencyBandListAerial_r18 MultiFrequencyBandListNR_Aerial_SIB_r18
}

func (ie *FrequencyInfoDL_SIB_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(frequencyInfoDLSIBV1800Constraints)
	_ = seq

	// 1. frequencyBandListAerial-r18: ref
	{
		if err := ie.FrequencyBandListAerial_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *FrequencyInfoDL_SIB_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(frequencyInfoDLSIBV1800Constraints)
	_ = seq

	// 1. frequencyBandListAerial-r18: ref
	{
		if err := ie.FrequencyBandListAerial_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
