// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCP-ParametersMRDC-v1610 (line 22768).

var pDCPParametersMRDCV1610Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "scg-DRB-NR-IAB-r16", Optional: true},
	},
}

const (
	PDCP_ParametersMRDC_v1610_Scg_DRB_NR_IAB_r16_Supported = 0
)

var pDCPParametersMRDCV1610ScgDRBNRIABR16Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{PDCP_ParametersMRDC_v1610_Scg_DRB_NR_IAB_r16_Supported},
}

type PDCP_ParametersMRDC_v1610 struct {
	Scg_DRB_NR_IAB_r16 *int64
}

func (ie *PDCP_ParametersMRDC_v1610) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCPParametersMRDCV1610Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Scg_DRB_NR_IAB_r16 != nil}); err != nil {
		return err
	}

	// 2. scg-DRB-NR-IAB-r16: enumerated
	{
		if ie.Scg_DRB_NR_IAB_r16 != nil {
			if err := e.EncodeEnumerated(*ie.Scg_DRB_NR_IAB_r16, pDCPParametersMRDCV1610ScgDRBNRIABR16Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDCP_ParametersMRDC_v1610) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCPParametersMRDCV1610Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. scg-DRB-NR-IAB-r16: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(pDCPParametersMRDCV1610ScgDRBNRIABR16Constraints)
			if err != nil {
				return err
			}
			ie.Scg_DRB_NR_IAB_r16 = &idx
		}
	}

	return nil
}
