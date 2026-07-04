// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: MeasResultSRS-RSRP-r16 (line 9860).

var measResultSRSRSRPR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "srs-ResourceId-r16"},
		{Name: "srs-RSRP-Result-r16"},
	},
}

type MeasResultSRS_RSRP_r16 struct {
	Srs_ResourceId_r16  SRS_ResourceId
	Srs_RSRP_Result_r16 SRS_RSRP_Range_r16
}

func (ie *MeasResultSRS_RSRP_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(measResultSRSRSRPR16Constraints)
	_ = seq

	// 1. srs-ResourceId-r16: ref
	{
		if err := ie.Srs_ResourceId_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. srs-RSRP-Result-r16: ref
	{
		if err := ie.Srs_RSRP_Result_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *MeasResultSRS_RSRP_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(measResultSRSRSRPR16Constraints)
	_ = seq

	// 1. srs-ResourceId-r16: ref
	{
		if err := ie.Srs_ResourceId_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. srs-RSRP-Result-r16: ref
	{
		if err := ie.Srs_RSRP_Result_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
