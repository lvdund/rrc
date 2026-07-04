// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CA-ParametersNR-v1760 (line 17604).

var cAParametersNRV1760Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "prioSCellPRACH-OverSP-PeriodicSRS-Support-r17", Optional: true},
	},
}

const (
	CA_ParametersNR_v1760_PrioSCellPRACH_OverSP_PeriodicSRS_Support_r17_Supported = 0
)

var cAParametersNRV1760PrioSCellPRACHOverSPPeriodicSRSSupportR17Constraints = per.EnumeratedConstraints{
	Extensible: false,
	RootValues: []int64{CA_ParametersNR_v1760_PrioSCellPRACH_OverSP_PeriodicSRS_Support_r17_Supported},
}

type CA_ParametersNR_v1760 struct {
	PrioSCellPRACH_OverSP_PeriodicSRS_Support_r17 *int64
}

func (ie *CA_ParametersNR_v1760) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAParametersNRV1760Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.PrioSCellPRACH_OverSP_PeriodicSRS_Support_r17 != nil}); err != nil {
		return err
	}

	// 2. prioSCellPRACH-OverSP-PeriodicSRS-Support-r17: enumerated
	{
		if ie.PrioSCellPRACH_OverSP_PeriodicSRS_Support_r17 != nil {
			if err := e.EncodeEnumerated(*ie.PrioSCellPRACH_OverSP_PeriodicSRS_Support_r17, cAParametersNRV1760PrioSCellPRACHOverSPPeriodicSRSSupportR17Constraints); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CA_ParametersNR_v1760) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAParametersNRV1760Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. prioSCellPRACH-OverSP-PeriodicSRS-Support-r17: enumerated
	{
		if seq.IsComponentPresent(0) {
			idx, err := d.DecodeEnumerated(cAParametersNRV1760PrioSCellPRACHOverSPPeriodicSRSSupportR17Constraints)
			if err != nil {
				return err
			}
			ie.PrioSCellPRACH_OverSP_PeriodicSRS_Support_r17 = &idx
		}
	}

	return nil
}
