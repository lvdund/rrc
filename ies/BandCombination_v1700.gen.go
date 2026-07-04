// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-v1700 (line 16745).

var bandCombinationV1700Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ca-ParametersNR-v1700", Optional: true},
		{Name: "ca-ParametersNRDC-v1700", Optional: true},
		{Name: "mrdc-Parameters-v1700", Optional: true},
		{Name: "bandList-v1710", Optional: true},
		{Name: "supportedBandCombListPerBC-SL-RelayDiscovery-r17", Optional: true},
		{Name: "supportedBandCombListPerBC-SL-NonRelayDiscovery-r17", Optional: true},
	},
}

var bandCombinationV1700BandListV1710Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

var bandCombinationV1700SupportedBandCombListPerBCSLRelayDiscoveryR17Constraints = per.SizeRange(1, common.MaxBandComb)

var bandCombinationV1700SupportedBandCombListPerBCSLNonRelayDiscoveryR17Constraints = per.SizeRange(1, common.MaxBandComb)

type BandCombination_v1700 struct {
	Ca_ParametersNR_v1700                               *CA_ParametersNR_v1700
	Ca_ParametersNRDC_v1700                             *CA_ParametersNRDC_v1700
	Mrdc_Parameters_v1700                               *MRDC_Parameters_v1700
	BandList_v1710                                      []BandParameters_v1710
	SupportedBandCombListPerBC_SL_RelayDiscovery_r17    *per.BitString
	SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17 *per.BitString
}

func (ie *BandCombination_v1700) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1700Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersNR_v1700 != nil, ie.Ca_ParametersNRDC_v1700 != nil, ie.Mrdc_Parameters_v1700 != nil, ie.BandList_v1710 != nil, ie.SupportedBandCombListPerBC_SL_RelayDiscovery_r17 != nil, ie.SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17 != nil}); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1700: ref
	{
		if ie.Ca_ParametersNR_v1700 != nil {
			if err := ie.Ca_ParametersNR_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1700: ref
	{
		if ie.Ca_ParametersNRDC_v1700 != nil {
			if err := ie.Ca_ParametersNRDC_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. mrdc-Parameters-v1700: ref
	{
		if ie.Mrdc_Parameters_v1700 != nil {
			if err := ie.Mrdc_Parameters_v1700.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. bandList-v1710: sequence-of
	{
		if ie.BandList_v1710 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationV1700BandListV1710Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.BandList_v1710))); err != nil {
				return err
			}
			for i := range ie.BandList_v1710 {
				if err := ie.BandList_v1710[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 6. supportedBandCombListPerBC-SL-RelayDiscovery-r17: bit-string
	{
		if ie.SupportedBandCombListPerBC_SL_RelayDiscovery_r17 != nil {
			if err := e.EncodeBitString(*ie.SupportedBandCombListPerBC_SL_RelayDiscovery_r17, bandCombinationV1700SupportedBandCombListPerBCSLRelayDiscoveryR17Constraints); err != nil {
				return err
			}
		}
	}

	// 7. supportedBandCombListPerBC-SL-NonRelayDiscovery-r17: bit-string
	{
		if ie.SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17 != nil {
			if err := e.EncodeBitString(*ie.SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17, bandCombinationV1700SupportedBandCombListPerBCSLNonRelayDiscoveryR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1700) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1700Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ca-ParametersNR-v1700: ref
	{
		if seq.IsComponentPresent(0) {
			ie.Ca_ParametersNR_v1700 = new(CA_ParametersNR_v1700)
			if err := ie.Ca_ParametersNR_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. ca-ParametersNRDC-v1700: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNRDC_v1700 = new(CA_ParametersNRDC_v1700)
			if err := ie.Ca_ParametersNRDC_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. mrdc-Parameters-v1700: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Mrdc_Parameters_v1700 = new(MRDC_Parameters_v1700)
			if err := ie.Mrdc_Parameters_v1700.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. bandList-v1710: sequence-of
	{
		if seq.IsComponentPresent(3) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationV1700BandListV1710Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.BandList_v1710 = make([]BandParameters_v1710, n)
			for i := int64(0); i < n; i++ {
				if err := ie.BandList_v1710[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 6. supportedBandCombListPerBC-SL-RelayDiscovery-r17: bit-string
	{
		if seq.IsComponentPresent(4) {
			v, err := d.DecodeBitString(bandCombinationV1700SupportedBandCombListPerBCSLRelayDiscoveryR17Constraints)
			if err != nil {
				return err
			}
			ie.SupportedBandCombListPerBC_SL_RelayDiscovery_r17 = &v
		}
	}

	// 7. supportedBandCombListPerBC-SL-NonRelayDiscovery-r17: bit-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeBitString(bandCombinationV1700SupportedBandCombListPerBCSLNonRelayDiscoveryR17Constraints)
			if err != nil {
				return err
			}
			ie.SupportedBandCombListPerBC_SL_NonRelayDiscovery_r17 = &v
		}
	}

	return nil
}
