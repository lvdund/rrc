// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: RF-Parameters-v16a0 (line 23657).

var rFParametersV16a0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandCombinationList-v16a0", Optional: true},
		{Name: "supportedBandCombinationList-UplinkTxSwitch-v16a0", Optional: true},
	},
}

type RF_Parameters_V16a0 struct {
	SupportedBandCombinationList_V16a0                *BandCombinationList_V16a0
	SupportedBandCombinationList_UplinkTxSwitch_V16a0 *BandCombinationList_UplinkTxSwitch_V16a0
}

func (ie *RF_Parameters_V16a0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rFParametersV16a0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandCombinationList_V16a0 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_V16a0 != nil}); err != nil {
		return err
	}

	// 2. supportedBandCombinationList-v16a0: ref
	{
		if ie.SupportedBandCombinationList_V16a0 != nil {
			if err := ie.SupportedBandCombinationList_V16a0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. supportedBandCombinationList-UplinkTxSwitch-v16a0: ref
	{
		if ie.SupportedBandCombinationList_UplinkTxSwitch_V16a0 != nil {
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_V16a0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RF_Parameters_V16a0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rFParametersV16a0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandCombinationList-v16a0: ref
	{
		if seq.IsComponentPresent(0) {
			ie.SupportedBandCombinationList_V16a0 = new(BandCombinationList_V16a0)
			if err := ie.SupportedBandCombinationList_V16a0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. supportedBandCombinationList-UplinkTxSwitch-v16a0: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_V16a0 = new(BandCombinationList_UplinkTxSwitch_V16a0)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_V16a0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
