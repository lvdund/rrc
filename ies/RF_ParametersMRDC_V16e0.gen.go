// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RF-ParametersMRDC-v16e0 (line 24808).

var rFParametersMRDCV16e0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandCombinationList-UplinkTxSwitch-v16e0", Optional: true},
	},
}

type RF_ParametersMRDC_V16e0 struct {
	SupportedBandCombinationList_UplinkTxSwitch_V16e0 *BandCombinationList_UplinkTxSwitch_V16e0
}

func (ie *RF_ParametersMRDC_V16e0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rFParametersMRDCV16e0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandCombinationList_UplinkTxSwitch_V16e0 != nil}); err != nil {
		return err
	}

	// 2. supportedBandCombinationList-UplinkTxSwitch-v16e0: ref
	{
		if ie.SupportedBandCombinationList_UplinkTxSwitch_V16e0 != nil {
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_V16e0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RF_ParametersMRDC_V16e0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rFParametersMRDCV16e0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandCombinationList-UplinkTxSwitch-v16e0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_V16e0 = new(BandCombinationList_UplinkTxSwitch_V16e0)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_V16e0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
