// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UplinkConfigCommon-v1700 (line 16305).

var uplinkConfigCommonV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "initialUplinkBWP-RedCap-r17", Optional: true},
	},
}

type UplinkConfigCommon_v1700 struct {
	InitialUplinkBWP_RedCap_r17 *BWP_UplinkCommon
}

func (ie *UplinkConfigCommon_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkConfigCommonV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.InitialUplinkBWP_RedCap_r17 != nil}); err != nil {
		return err
	}

	// 2. initialUplinkBWP-RedCap-r17: ref
	{
		if ie.InitialUplinkBWP_RedCap_r17 != nil {
			if err := ie.InitialUplinkBWP_RedCap_r17.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UplinkConfigCommon_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkConfigCommonV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. initialUplinkBWP-RedCap-r17: ref
	{
		if seq.IsComponentPresent(0) {
			ie.InitialUplinkBWP_RedCap_r17 = new(BWP_UplinkCommon)
			if err := ie.InitialUplinkBWP_RedCap_r17.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
