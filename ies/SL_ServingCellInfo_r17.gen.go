// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SL-ServingCellInfo-r17 (line 28249).

var sLServingCellInfoR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "sl-PhysCellId-r17"},
		{Name: "sl-CarrierFreqNR-r17"},
	},
}

type SL_ServingCellInfo_r17 struct {
	Sl_PhysCellId_r17    PhysCellId
	Sl_CarrierFreqNR_r17 ARFCN_ValueNR
}

func (ie *SL_ServingCellInfo_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sLServingCellInfoR17Constraints)
	_ = seq

	// 1. sl-PhysCellId-r17: ref
	{
		if err := ie.Sl_PhysCellId_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. sl-CarrierFreqNR-r17: ref
	{
		if err := ie.Sl_CarrierFreqNR_r17.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *SL_ServingCellInfo_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sLServingCellInfoR17Constraints)
	_ = seq

	// 1. sl-PhysCellId-r17: ref
	{
		if err := ie.Sl_PhysCellId_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. sl-CarrierFreqNR-r17: ref
	{
		if err := ie.Sl_CarrierFreqNR_r17.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
