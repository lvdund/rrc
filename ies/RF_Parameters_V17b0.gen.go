// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: RF-Parameters-v17b0 (line 23671).

var rFParametersV17b0Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "supportedBandListNR-v17b0", Optional: true},
		{Name: "supportedBandCombinationList-v17b0", Optional: true},
		{Name: "supportedBandCombinationList-UplinkTxSwitch-v17b0", Optional: true},
	},
}

var rFParametersV17b0SupportedBandListNRV17b0Constraints = per.SizeRange(1, common.MaxBands)

type RF_Parameters_V17b0 struct {
	SupportedBandListNR_V17b0                         []BandNR_V17b0
	SupportedBandCombinationList_V17b0                *BandCombinationList_V17b0
	SupportedBandCombinationList_UplinkTxSwitch_V17b0 *BandCombinationList_UplinkTxSwitch_V17b0
}

func (ie *RF_Parameters_V17b0) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(rFParametersV17b0Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.SupportedBandListNR_V17b0 != nil, ie.SupportedBandCombinationList_V17b0 != nil, ie.SupportedBandCombinationList_UplinkTxSwitch_V17b0 != nil}); err != nil {
		return err
	}

	// 2. supportedBandListNR-v17b0: sequence-of
	{
		if ie.SupportedBandListNR_V17b0 != nil {
			seqOf := e.NewSequenceOfEncoder(rFParametersV17b0SupportedBandListNRV17b0Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.SupportedBandListNR_V17b0))); err != nil {
				return err
			}
			for i := range ie.SupportedBandListNR_V17b0 {
				if err := ie.SupportedBandListNR_V17b0[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. supportedBandCombinationList-v17b0: ref
	{
		if ie.SupportedBandCombinationList_V17b0 != nil {
			if err := ie.SupportedBandCombinationList_V17b0.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. supportedBandCombinationList-UplinkTxSwitch-v17b0: ref
	{
		if ie.SupportedBandCombinationList_UplinkTxSwitch_V17b0 != nil {
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_V17b0.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *RF_Parameters_V17b0) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(rFParametersV17b0Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. supportedBandListNR-v17b0: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(rFParametersV17b0SupportedBandListNRV17b0Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.SupportedBandListNR_V17b0 = make([]BandNR_V17b0, n)
			for i := int64(0); i < n; i++ {
				if err := ie.SupportedBandListNR_V17b0[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. supportedBandCombinationList-v17b0: ref
	{
		if seq.IsComponentPresent(1) {
			ie.SupportedBandCombinationList_V17b0 = new(BandCombinationList_V17b0)
			if err := ie.SupportedBandCombinationList_V17b0.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. supportedBandCombinationList-UplinkTxSwitch-v17b0: ref
	{
		if seq.IsComponentPresent(2) {
			ie.SupportedBandCombinationList_UplinkTxSwitch_V17b0 = new(BandCombinationList_UplinkTxSwitch_V17b0)
			if err := ie.SupportedBandCombinationList_UplinkTxSwitch_V17b0.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
