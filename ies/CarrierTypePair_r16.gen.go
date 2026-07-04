// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CarrierTypePair-r16 (line 18331).

var carrierTypePairR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "carrierForCSI-Measurement-r16"},
		{Name: "carrierForCSI-Reporting-r16"},
	},
}

type CarrierTypePair_r16 struct {
	CarrierForCSI_Measurement_r16 PUCCH_Grp_CarrierTypes_r16
	CarrierForCSI_Reporting_r16   PUCCH_Grp_CarrierTypes_r16
}

func (ie *CarrierTypePair_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(carrierTypePairR16Constraints)
	_ = seq

	// 1. carrierForCSI-Measurement-r16: ref
	{
		if err := ie.CarrierForCSI_Measurement_r16.Encode(e); err != nil {
			return err
		}
	}

	// 2. carrierForCSI-Reporting-r16: ref
	{
		if err := ie.CarrierForCSI_Reporting_r16.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *CarrierTypePair_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(carrierTypePairR16Constraints)
	_ = seq

	// 1. carrierForCSI-Measurement-r16: ref
	{
		if err := ie.CarrierForCSI_Measurement_r16.Decode(d); err != nil {
			return err
		}
	}

	// 2. carrierForCSI-Reporting-r16: ref
	{
		if err := ie.CarrierForCSI_Reporting_r16.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
