// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SharedSpectrumChAccessParamsPerBand-v1710 (line 24969).

var sharedSpectrumChAccessParamsPerBandV1710Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ul-Semi-StaticChAccessDependentConfig-r17", Optional: true},
		{Name: "ul-Semi-StaticChAccessIndependentConfig-r17", Optional: true},
	},
}

const (
	SharedSpectrumChAccessParamsPerBand_v1710_Ul_Semi_StaticChAccessDependentConfig_r17_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandV1710UlSemiStaticChAccessDependentConfigR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_v1710_Ul_Semi_StaticChAccessDependentConfig_r17_Supported},
}

const (
	SharedSpectrumChAccessParamsPerBand_v1710_Ul_Semi_StaticChAccessIndependentConfig_r17_Supported = 0
)

var sharedSpectrumChAccessParamsPerBandV1710UlSemiStaticChAccessIndependentConfigR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{SharedSpectrumChAccessParamsPerBand_v1710_Ul_Semi_StaticChAccessIndependentConfig_r17_Supported},
}

type SharedSpectrumChAccessParamsPerBand_v1710 struct {
	Ul_Semi_StaticChAccessDependentConfig_r17   *int64
	Ul_Semi_StaticChAccessIndependentConfig_r17 *int64
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1710) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sharedSpectrumChAccessParamsPerBandV1710Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Ul_Semi_StaticChAccessDependentConfig_r17 != nil, ie.Ul_Semi_StaticChAccessIndependentConfig_r17 != nil}); err != nil {
		return err
	}

	// 2. ul-Semi-StaticChAccessDependentConfig-r17: enumerated
	{
		if ie.Ul_Semi_StaticChAccessDependentConfig_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_Semi_StaticChAccessDependentConfig_r17, sharedSpectrumChAccessParamsPerBandV1710UlSemiStaticChAccessDependentConfigR17Constraints); err != nil {
				return err
			}
		}
	}

	// 3. ul-Semi-StaticChAccessIndependentConfig-r17: enumerated
	{
		if ie.Ul_Semi_StaticChAccessIndependentConfig_r17 != nil {
			if err := e.EncodeEnumerated(*ie.Ul_Semi_StaticChAccessIndependentConfig_r17, sharedSpectrumChAccessParamsPerBandV1710UlSemiStaticChAccessIndependentConfigR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SharedSpectrumChAccessParamsPerBand_v1710) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sharedSpectrumChAccessParamsPerBandV1710Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. ul-Semi-StaticChAccessDependentConfig-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandV1710UlSemiStaticChAccessDependentConfigR17Constraints)
			if err != nil {
				return err
			}
			ie.Ul_Semi_StaticChAccessDependentConfig_r17 = &idx
		}
	}

	// 3. ul-Semi-StaticChAccessIndependentConfig-r17: enumerated
	{
		if seq.IsComponentPresent(1) {
			idx, err := d.DecodeEnumerated(sharedSpectrumChAccessParamsPerBandV1710UlSemiStaticChAccessIndependentConfigR17Constraints)
			if err != nil {
				return err
			}
			ie.Ul_Semi_StaticChAccessIndependentConfig_r17 = &idx
		}
	}

	return nil
}
