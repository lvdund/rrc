// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-v1800 (line 16796).

var bandCombinationV1800Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-v1800", Optional: true},
		{Name: "ca-ParametersNRDC-v1800", Optional: true},
		{Name: "supportedBandCombListPerBC-SL-U2U-RelayDiscovery-r18", Optional: true},
		{Name: "bandList-v1810", Optional: true},
	},
}

var bandCombinationV1800SupportedBandCombListPerBCSLU2URelayDiscoveryR18Constraints = per.SizeRange(1, common.MaxBandComb)

var bandCombinationV1800BandListV1810Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

type BandCombination_v1800 struct {
	Ca_ParametersNR_v1800                                *CA_ParametersNR_v1800
	Ca_ParametersNRDC_v1800                              *CA_ParametersNRDC_v1800
	SupportedBandCombListPerBC_SL_U2U_RelayDiscovery_r18 *per.BitString
	BandList_v1810                                       []BandParameters_v1810
}

func (ie *BandCombination_v1800) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1800Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_v1800 != nil, ie.Ca_ParametersNRDC_v1800 != nil, ie.SupportedBandCombListPerBC_SL_U2U_RelayDiscovery_r18 != nil, ie.BandList_v1810 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1800: ref
	{
		if ie.Ca_ParametersNR_v1800 != nil {
			if err := ie.Ca_ParametersNR_v1800.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1800: ref
	{
		if ie.Ca_ParametersNRDC_v1800 != nil {
			if err := ie.Ca_ParametersNRDC_v1800.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. supportedBandCombListPerBC-SL-U2U-RelayDiscovery-r18: bit-string
	{
		if ie.SupportedBandCombListPerBC_SL_U2U_RelayDiscovery_r18 != nil {
			if err := e.EncodeBitString(*ie.SupportedBandCombListPerBC_SL_U2U_RelayDiscovery_r18, bandCombinationV1800SupportedBandCombListPerBCSLU2URelayDiscoveryR18Constraints); err != nil {
				return err
			}
		}
	}

	// 5. bandList-v1810: sequence-of
	{
		if ie.BandList_v1810 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationV1800BandListV1810Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.BandList_v1810))); err != nil {
				return err
			}
			for i := range ie.BandList_v1810 {
				if err := ie.BandList_v1810[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1800) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1800Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1800: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_v1800 = new(CA_ParametersNR_v1800)
			if err := ie.Ca_ParametersNR_v1800.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1800: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNRDC_v1800 = new(CA_ParametersNRDC_v1800)
			if err := ie.Ca_ParametersNRDC_v1800.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. supportedBandCombListPerBC-SL-U2U-RelayDiscovery-r18: bit-string
	{
		if seq.IsComponentPresent(2) {
			v, err := d.DecodeBitString(bandCombinationV1800SupportedBandCombListPerBCSLU2URelayDiscoveryR18Constraints)
			if err != nil {
				return err
			}
			ie.SupportedBandCombListPerBC_SL_U2U_RelayDiscovery_r18 = &v
		}
	}

	// 5. bandList-v1810: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationV1800BandListV1810Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.BandList_v1810 = make([]BandParameters_v1810, n)
			for i := int64(0); i < n; i++ {
				if err := ie.BandList_v1810[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
