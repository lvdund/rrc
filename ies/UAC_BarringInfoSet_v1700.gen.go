// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UAC-BarringInfoSet-v1700 (line 16189).

var uACBarringInfoSetV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "uac-BarringFactorForAI3-r17", Optional: true},
	},
}

const (
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P00 = 0
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P05 = 1
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P10 = 2
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P15 = 3
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P20 = 4
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P25 = 5
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P30 = 6
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P40 = 7
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P50 = 8
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P60 = 9
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P70 = 10
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P75 = 11
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P80 = 12
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P85 = 13
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P90 = 14
	UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P95 = 15
)

var uACBarringInfoSetV1700UacBarringFactorForAI3R17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P00, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P05, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P10, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P15, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P20, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P25, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P30, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P40, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P50, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P60, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P70, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P75, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P80, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P85, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P90, UAC_BarringInfoSet_v1700_Uac_BarringFactorForAI3_r17_P95},
}

type UAC_BarringInfoSet_v1700 struct {
	Uac_BarringFactorForAI3_r17 *int64
}

func (ie *UAC_BarringInfoSet_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uACBarringInfoSetV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Uac_BarringFactorForAI3_r17 != nil}); err != nil {
		return err
	}

	// 2. uac-BarringFactorForAI3-r17: enumerated
	{
		if ie.Uac_BarringFactorForAI3_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Uac_BarringFactorForAI3_r17, uACBarringInfoSetV1700UacBarringFactorForAI3R17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UAC_BarringInfoSet_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uACBarringInfoSetV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. uac-BarringFactorForAI3-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(uACBarringInfoSetV1700UacBarringFactorForAI3R17Constraints)
			if err != nil {
				return err
			}
			ie.Uac_BarringFactorForAI3_r17 = &idx
		}
	}

	return nil
}
