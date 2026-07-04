// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: DownlinkConfigCommonSIB-v1760 (line 7849).

var downlinkConfigCommonSIBV1760Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "frequencyInfoDL-v1760"},
	},
}

type DownlinkConfigCommonSIB_v1760 struct {
	FrequencyInfoDL_v1760 FrequencyInfoDL_SIB_v1760
}

func (ie *DownlinkConfigCommonSIB_v1760) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(downlinkConfigCommonSIBV1760Constraints)
	_ = seq

	// 1. frequencyInfoDL-v1760: ref
	{
		if err := ie.FrequencyInfoDL_v1760.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *DownlinkConfigCommonSIB_v1760) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(downlinkConfigCommonSIBV1760Constraints)
	_ = seq

	// 1. frequencyInfoDL-v1760: ref
	{
		if err := ie.FrequencyInfoDL_v1760.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
