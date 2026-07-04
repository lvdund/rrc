// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RedCapParameters-v1740 (line 23534).

var redCapParametersV1740Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ncd-SSB-ForRedCapInitialBWP-SDT-r17", Optional: true},
	},
}

const (
	RedCapParameters_v1740_Ncd_SSB_ForRedCapInitialBWP_SDT_r17_Supported = 0
)

var redCapParametersV1740NcdSSBForRedCapInitialBWPSDTR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{RedCapParameters_v1740_Ncd_SSB_ForRedCapInitialBWP_SDT_r17_Supported},
}

type RedCapParameters_v1740 struct {
	Ncd_SSB_ForRedCapInitialBWP_SDT_r17 *int64
}

func (ie *RedCapParameters_v1740) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(redCapParametersV1740Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ncd_SSB_ForRedCapInitialBWP_SDT_r17 != nil}); err != nil {
		return err
	}

	// 2. ncd-SSB-ForRedCapInitialBWP-SDT-r17: enumerated
	{
		if ie.Ncd_SSB_ForRedCapInitialBWP_SDT_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ncd_SSB_ForRedCapInitialBWP_SDT_r17, redCapParametersV1740NcdSSBForRedCapInitialBWPSDTR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RedCapParameters_v1740) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(redCapParametersV1740Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ncd-SSB-ForRedCapInitialBWP-SDT-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(redCapParametersV1740NcdSSBForRedCapInitialBWPSDTR17Constraints)
			if err != nil {
				return err
			}
			ie.Ncd_SSB_ForRedCapInitialBWP_SDT_r17 = &idx
		}
	}

	return nil
}
