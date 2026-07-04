// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CSI-RS-MeasResult-r19 (line 3515).

var cSIRSMeasResultR19Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "resourceId-r19"},
		{Name: "l1-RSRP-r19"},
	},
}

type CSI_RS_MeasResult_r19 struct {
	ResourceId_r19 NZP_CSI_RS_ResourceId
	L1_RSRP_r19    RSRP_Range
}

func (ie *CSI_RS_MeasResult_r19) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cSIRSMeasResultR19Constraints)
	_ = seq

	// 1. resourceId-r19: ref
	{
		if err := ie.ResourceId_r19.Encode(e); err != nil {
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

func (ie *CSI_RS_MeasResult_r19) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cSIRSMeasResultR19Constraints)
	_ = seq

	// 1. resourceId-r19: ref
	{
		if err := ie.ResourceId_r19.Decode(d); err != nil {
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
