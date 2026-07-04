// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: BandParameters-v1610 (line 17057).

var bandParametersV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-TxSwitch-v1610", Optional: true},
	},
}

const (
	BandParameters_v1610_Srs_TxSwitch_v1610_SupportedSRS_TxPortSwitch_v1610_T1r1_T1r2                = 0
	BandParameters_v1610_Srs_TxSwitch_v1610_SupportedSRS_TxPortSwitch_v1610_T1r1_T1r2_T1r4           = 1
	BandParameters_v1610_Srs_TxSwitch_v1610_SupportedSRS_TxPortSwitch_v1610_T1r1_T1r2_T2r2_T2r4      = 2
	BandParameters_v1610_Srs_TxSwitch_v1610_SupportedSRS_TxPortSwitch_v1610_T1r1_T1r2_T2r2_T1r4_T2r4 = 3
	BandParameters_v1610_Srs_TxSwitch_v1610_SupportedSRS_TxPortSwitch_v1610_T1r1_T2r2                = 4
	BandParameters_v1610_Srs_TxSwitch_v1610_SupportedSRS_TxPortSwitch_v1610_T1r1_T2r2_T4r4           = 5
)

var bandParametersV1610SrsTxSwitchV1610SupportedSRSTxPortSwitchV1610Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandParameters_v1610_Srs_TxSwitch_v1610_SupportedSRS_TxPortSwitch_v1610_T1r1_T1r2, BandParameters_v1610_Srs_TxSwitch_v1610_SupportedSRS_TxPortSwitch_v1610_T1r1_T1r2_T1r4, BandParameters_v1610_Srs_TxSwitch_v1610_SupportedSRS_TxPortSwitch_v1610_T1r1_T1r2_T2r2_T2r4, BandParameters_v1610_Srs_TxSwitch_v1610_SupportedSRS_TxPortSwitch_v1610_T1r1_T1r2_T2r2_T1r4_T2r4, BandParameters_v1610_Srs_TxSwitch_v1610_SupportedSRS_TxPortSwitch_v1610_T1r1_T2r2, BandParameters_v1610_Srs_TxSwitch_v1610_SupportedSRS_TxPortSwitch_v1610_T1r1_T2r2_T4r4},
}

type BandParameters_v1610 struct {
	Srs_TxSwitch_v1610 *struct{ SupportedSRS_TxPortSwitch_v1610 int64 }
}

func (ie *BandParameters_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandParametersV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Srs_TxSwitch_v1610 != nil}); err != nil {
		return err
	}

	// 2. srs-TxSwitch-v1610: sequence
	{
		if ie.Srs_TxSwitch_v1610 != nil {
			c := ie.Srs_TxSwitch_v1610
			if err := e.EncodeEnumerated(c.SupportedSRS_TxPortSwitch_v1610, bandParametersV1610SrsTxSwitchV1610SupportedSRSTxPortSwitchV1610Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandParameters_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandParametersV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. srs-TxSwitch-v1610: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Srs_TxSwitch_v1610 = &struct{ SupportedSRS_TxPortSwitch_v1610 int64 }{}
			c := ie.Srs_TxSwitch_v1610
			{
				v, err := d.DecodeEnumerated(bandParametersV1610SrsTxSwitchV1610SupportedSRSTxPortSwitchV1610Constraints)
				if err != nil {
					return err
				}
				c.SupportedSRS_TxPortSwitch_v1610 = v
			}
		}
	}

	return nil
}
