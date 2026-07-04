// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination (line 16650).

var bandCombinationConstraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandList"},
		{Name: "featureSetCombination"},
		{Name: "ca-ParametersEUTRA", Optional: true},
		{Name: "ca-ParametersNR", Optional: true},
		{Name: "mrdc-Parameters", Optional: true},
		{Name: "supportedBandwidthCombinationSet", Optional: true},
		{Name: "powerClass-v1530", Optional: true},
	},
}

var bandCombinationBandListConstraints = per.SizeRange(1, common.MaxSimultaneousBands)

var bandCombinationSupportedBandwidthCombinationSetConstraints = per.SizeRange(1, 32)

const (
	BandCombination_PowerClass_v1530_Pc2 = 0
)

var bandCombinationPowerClassV1530Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_PowerClass_v1530_Pc2},
}

type BandCombination struct {
	BandList                         []BandParameters
	FeatureSetCombination            FeatureSetCombinationId
	Ca_ParametersEUTRA               *CA_ParametersEUTRA
	Ca_ParametersNR                  *CA_ParametersNR
	Mrdc_Parameters                  *MRDC_Parameters
	SupportedBandwidthCombinationSet *per.BitString
	PowerClass_v1530                 *int64
}

func (ie *BandCombination) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationConstraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ca_ParametersEUTRA != nil, ie.Ca_ParametersNR != nil, ie.Mrdc_Parameters != nil, ie.SupportedBandwidthCombinationSet != nil, ie.PowerClass_v1530 != nil}); err != nil {
		return err
	}

	// 2. bandList: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(bandCombinationBandListConstraints)
		if err := seqOf.EncodeLength(int64(len(ie.BandList))); err != nil {
			return err
		}
		for i := range ie.BandList {
			if err := ie.BandList[i].Encode(e); err != nil {
				return err
			}
		}
	}

	// 3. featureSetCombination: ref
	{
		if err := ie.FeatureSetCombination.Encode(e); err != nil {
			return err
		}
	}

	// 4. ca-ParametersEUTRA: ref
	{
		if ie.Ca_ParametersEUTRA != nil {
			if err := ie.Ca_ParametersEUTRA.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. ca-ParametersNR: ref
	{
		if ie.Ca_ParametersNR != nil {
			if err := ie.Ca_ParametersNR.Encode(e); err != nil {
				return err
			}
		}
	}

	// 6. mrdc-Parameters: ref
	{
		if ie.Mrdc_Parameters != nil {
			if err := ie.Mrdc_Parameters.Encode(e); err != nil {
				return err
			}
		}
	}

	// 7. supportedBandwidthCombinationSet: bit-string
	{
		if ie.SupportedBandwidthCombinationSet != nil {
			if err := e.EncodeBitString(*ie.SupportedBandwidthCombinationSet, bandCombinationSupportedBandwidthCombinationSetConstraints); err != nil {
				return err
			}
		}
	}

	// 8. powerClass-v1530: enumerated
	{
		if ie.PowerClass_v1530 != nil {
			if err := e.EncodeEnumerated(*ie.PowerClass_v1530, bandCombinationPowerClassV1530Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationConstraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandList: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(bandCombinationBandListConstraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.BandList = make([]BandParameters, n)
		for i := int64(0); i < n; i++ {
			if err := ie.BandList[i].Decode(d); err != nil {
				return err
			}
		}
	}

	// 3. featureSetCombination: ref
	{
		if err := ie.FeatureSetCombination.Decode(d); err != nil {
			return err
		}
	}

	// 4. ca-ParametersEUTRA: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ca_ParametersEUTRA = new(CA_ParametersEUTRA)
			if err := ie.Ca_ParametersEUTRA.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. ca-ParametersNR: ref
	{
		if seq.IsComponentPresent(3) {
			ie.Ca_ParametersNR = new(CA_ParametersNR)
			if err := ie.Ca_ParametersNR.Decode(d); err != nil {
				return err
			}
		}
	}

	// 6. mrdc-Parameters: ref
	{
		if seq.IsComponentPresent(4) {
			ie.Mrdc_Parameters = new(MRDC_Parameters)
			if err := ie.Mrdc_Parameters.Decode(d); err != nil {
				return err
			}
		}
	}

	// 7. supportedBandwidthCombinationSet: bit-string
	{
		if seq.IsComponentPresent(5) {
			v, err := d.DecodeBitString(bandCombinationSupportedBandwidthCombinationSetConstraints)
			if err != nil {
				return err
			}
			ie.SupportedBandwidthCombinationSet = &v
		}
	}

	// 8. powerClass-v1530: enumerated
	{
		if seq.IsComponentPresent(6) {
			idx, err := d.DecodeEnumerated(bandCombinationPowerClassV1530Constraints)
			if err != nil {
				return err
			}
			ie.PowerClass_v1530 = &idx
		}
	}

	return nil
}
