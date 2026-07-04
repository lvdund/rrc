// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: UplinkTxSwitchingAssociatedBandDualUL-r18 (line 5836).

var uplinkTxSwitchingAssociatedBandDualULR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "transmitBand-r18"},
		{Name: "associatedBand-r18"},
	},
}

type UplinkTxSwitchingAssociatedBandDualUL_r18 struct {
	TransmitBand_r18   UplinkTxSwitchingBandIndex_r18
	AssociatedBand_r18 UplinkTxSwitchingBandIndex_r18
}

func (ie *UplinkTxSwitchingAssociatedBandDualUL_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uplinkTxSwitchingAssociatedBandDualULR18Constraints)
	_ = seq

	// 1. transmitBand-r18: ref
	{
		if err := ie.TransmitBand_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. associatedBand-r18: ref
	{
		if err := ie.AssociatedBand_r18.Encode(e); err != nil {
			return err
		}
	}

	return nil
}

func (ie *UplinkTxSwitchingAssociatedBandDualUL_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uplinkTxSwitchingAssociatedBandDualULR18Constraints)
	_ = seq

	// 1. transmitBand-r18: ref
	{
		if err := ie.TransmitBand_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. associatedBand-r18: ref
	{
		if err := ie.AssociatedBand_r18.Decode(d); err != nil {
			return err
		}
	}

	return nil
}
