// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SharedSpectrumChAccessParamsPerBand-v1650 (line 24964).

var sharedSpectrumChAccessParamsPerBandV1650Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "extendedSearchSpaceSwitchWithDCI-r16", Optional: true},
	},
}

const (
	SharedSpectrumChAccessParamsPerBand_v1650_ExtendedSearchSpaceSwitchWithDCI_r16_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandV1650ExtendedSearchSpaceSwitchWithDCIR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_v1650_ExtendedSearchSpaceSwitchWithDCI_r16_Supported},
}

type SharedSpectrumChAccessParamsPerBand_v1650 struct {
	ExtendedSearchSpaceSwitchWithDCI_r16 *int64
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1650) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sharedSpectrumChAccessParamsPerBandV1650Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.ExtendedSearchSpaceSwitchWithDCI_r16 != nil}); err != nil {
		return err
	}

	// 2. extendedSearchSpaceSwitchWithDCI-r16: enumerated
	{
		if ie.ExtendedSearchSpaceSwitchWithDCI_r16 != nil {
			if err := e.EncodeEnumerated(*ie.ExtendedSearchSpaceSwitchWithDCI_r16, sharedSpectrumChAccessParamsPerBandV1650ExtendedSearchSpaceSwitchWithDCIR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1650) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sharedSpectrumChAccessParamsPerBandV1650Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. extendedSearchSpaceSwitchWithDCI-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandV1650ExtendedSearchSpaceSwitchWithDCIR16Constraints)
			if err != nil {
				return err
			}
			ie.ExtendedSearchSpaceSwitchWithDCI_r16 = &idx
		}
	}

	return nil
}
