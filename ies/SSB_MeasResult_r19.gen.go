// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SSB-MeasResult-r19 (line 3520).

var sSBMeasResultR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "ssb-Id-r19"},
		{Name: "l1-RSRP-r19"},
	},
}

type SSB_MeasResult_r19 struct {
	Ssb_Id_r19  SSB_Index
	L1_RSRP_r19 RSRP_Range
}

func (ie *SSB_MeasResult_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sSBMeasResultR19Constraints)
	_ = seq

	// 1. ssb-Id-r19: ref
	{
		if err := ie.Ssb_Id_r19.Encode(e); err != nil {
			return err
		}
	}

	// 2. l1-RSRP-r19: ref
	{
		if err := ie.L1_RSRP_r19.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SSB_MeasResult_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sSBMeasResultR19Constraints)
	_ = seq

	// 1. ssb-Id-r19: ref
	{
		if err := ie.Ssb_Id_r19.Decode(d); err != nil {
			return err
		}
	}

	// 2. l1-RSRP-r19: ref
	{
		if err := ie.L1_RSRP_r19.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
