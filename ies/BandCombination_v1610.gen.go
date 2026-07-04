// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: BandCombination-v1610 (line 16698).

var bandCombinationV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "bandList-v1610", Optional: true},
		{Name: "ca-ParametersNR-v1610", Optional: true},
		{Name: "ca-ParametersNRDC-v1610", Optional: true},
		{Name: "powerClass-v1610", Optional: true},
		{Name: "powerClassNRPart-r16", Optional: true},
		{Name: "featureSetCombinationDAPS-r16", Optional: true},
		{Name: "mrdc-Parameters-v1620", Optional: true},
	},
}

var bandCombinationV1610BandListV1610Constraints = per.SizeRange(1, common.MaxSimultaneousBands)

const (
	BandCombination_v1610_PowerClass_v1610_Pc1dot5 = 0
)

var bandCombinationV1610PowerClassV1610Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_v1610_PowerClass_v1610_Pc1dot5},
}

const (
	BandCombination_v1610_PowerClassNRPart_r16_Pc1 = 0
	BandCombination_v1610_PowerClassNRPart_r16_Pc2 = 1
	BandCombination_v1610_PowerClassNRPart_r16_Pc3 = 2
	BandCombination_v1610_PowerClassNRPart_r16_Pc5 = 3
)

var bandCombinationV1610PowerClassNRPartR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{BandCombination_v1610_PowerClassNRPart_r16_Pc1, BandCombination_v1610_PowerClassNRPart_r16_Pc2, BandCombination_v1610_PowerClassNRPart_r16_Pc3, BandCombination_v1610_PowerClassNRPart_r16_Pc5},
}

type BandCombination_v1610 struct {
	BandList_v1610                []BandParameters_v1610
	Ca_ParametersNR_v1610         *CA_ParametersNR_v1610
	Ca_ParametersNRDC_v1610       *CA_ParametersNRDC_v1610
	PowerClass_v1610              *int64
	PowerClassNRPart_r16          *int64
	FeatureSetCombinationDAPS_r16 *FeatureSetCombinationId
	Mrdc_Parameters_v1620         *MRDC_Parameters_v1620
}

func (ie *BandCombination_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(bandCombinationV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.BandList_v1610 != nil, ie.Ca_ParametersNR_v1610 != nil, ie.Ca_ParametersNRDC_v1610 != nil, ie.PowerClass_v1610 != nil, ie.PowerClassNRPart_r16 != nil, ie.FeatureSetCombinationDAPS_r16 != nil, ie.Mrdc_Parameters_v1620 != nil}); err != nil {
		return err
	}

	// 2. bandList-v1610: sequence-of
	{
		if ie.BandList_v1610 != nil {
			seqOf := e.NewSequenceOfEncoder(bandCombinationV1610BandListV1610Constraints)
			if err := seqOf.EncodeLength(int64(len(ie.BandList_v1610))); err != nil {
				return err
			}
			for i := range ie.BandList_v1610 {
				if err := ie.BandList_v1610[i].Encode(e); err != nil {
					return err
				}
			}
		}
	}

	// 3. ca-ParametersNR-v1610: ref
	{
		if ie.Ca_ParametersNR_v1610 != nil {
			if err := ie.Ca_ParametersNR_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 4. ca-ParametersNRDC-v1610: ref
	{
		if ie.Ca_ParametersNRDC_v1610 != nil {
			if err := ie.Ca_ParametersNRDC_v1610.Encode(e); err != nil {
				return err
			}
		}
	}

	// 5. powerClass-v1610: enumerated
	{
		if ie.PowerClass_v1610 != nil {
			if err := e.EncodeEnumerated(*ie.PowerClass_v1610, bandCombinationV1610PowerClassV1610Constraints); err != nil {
				return err
			}
		}
	}

	// 6. powerClassNRPart-r16: enumerated
	{
		if ie.PowerClassNRPart_r16 != nil {
			if err := e.EncodeEnumerated(*ie.PowerClassNRPart_r16, bandCombinationV1610PowerClassNRPartR16Constraints); err != nil {
				return err
			}
		}
	}

	// 7. featureSetCombinationDAPS-r16: ref
	{
		if ie.FeatureSetCombinationDAPS_r16 != nil {
			if err := ie.FeatureSetCombinationDAPS_r16.Encode(e); err != nil {
				return err
			}
		}
	}

	// 8. mrdc-Parameters-v1620: ref
	{
		if ie.Mrdc_Parameters_v1620 != nil {
			if err := ie.Mrdc_Parameters_v1620.Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *BandCombination_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(bandCombinationV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. bandList-v1610: sequence-of
	{
		if seq.IsComponentPresent(0) {
			seqOf := d.NewSequenceOfDecoder(bandCombinationV1610BandListV1610Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			ie.BandList_v1610 = make([]BandParameters_v1610, n)
			for i := int64(0); i < n; i++ {
				if err := ie.BandList_v1610[i].Decode(d); err != nil {
					return err
				}
			}
		}
	}

	// 3. ca-ParametersNR-v1610: ref
	{
		if seq.IsComponentPresent(1) {
			ie.Ca_ParametersNR_v1610 = new(CA_ParametersNR_v1610)
			if err := ie.Ca_ParametersNR_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 4. ca-ParametersNRDC-v1610: ref
	{
		if seq.IsComponentPresent(2) {
			ie.Ca_ParametersNRDC_v1610 = new(CA_ParametersNRDC_v1610)
			if err := ie.Ca_ParametersNRDC_v1610.Decode(d); err != nil {
				return err
			}
		}
	}

	// 5. powerClass-v1610: enumerated
	{
		if seq.IsComponentPresent(3) {
			idx, err := d.DecodeEnumerated(bandCombinationV1610PowerClassV1610Constraints)
			if err != nil {
				return err
			}
			ie.PowerClass_v1610 = &idx
		}
	}

	// 6. powerClassNRPart-r16: enumerated
	{
		if seq.IsComponentPresent(4) {
			idx, err := d.DecodeEnumerated(bandCombinationV1610PowerClassNRPartR16Constraints)
			if err != nil {
				return err
			}
			ie.PowerClassNRPart_r16 = &idx
		}
	}

	// 7. featureSetCombinationDAPS-r16: ref
	{
		if seq.IsComponentPresent(5) {
			ie.FeatureSetCombinationDAPS_r16 = new(FeatureSetCombinationId)
			if err := ie.FeatureSetCombinationDAPS_r16.Decode(d); err != nil {
				return err
			}
		}
	}

	// 8. mrdc-Parameters-v1620: ref
	{
		if seq.IsComponentPresent(6) {
			ie.Mrdc_Parameters_v1620 = new(MRDC_Parameters_v1620)
			if err := ie.Mrdc_Parameters_v1620.Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
