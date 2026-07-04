// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: NZP-CSI-RS-Pairing-r17 (line 10885).

var nZPCSIRSPairingR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "nzp-CSI-RS-ResourceId1-r17"},
		{Name: "nzp-CSI-RS-ResourceId2-r17"},
	},
}

var nZPCSIRSPairingR17NzpCSIRSResourceId1R17Constraints = per.Constrained(1, 7)

var nZPCSIRSPairingR17NzpCSIRSResourceId2R17Constraints = per.Constrained(1, 7)

type NZP_CSI_RS_Pairing_r17 struct {
	Nzp_CSI_RS_ResourceId1_r17 int64
	Nzp_CSI_RS_ResourceId2_r17 int64
}

func (ie *NZP_CSI_RS_Pairing_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(nZPCSIRSPairingR17Constraints)
	_ = seq

	// 1. nzp-CSI-RS-ResourceId1-r17: integer
	{
		if err := e.EncodeInteger(ie.Nzp_CSI_RS_ResourceId1_r17, nZPCSIRSPairingR17NzpCSIRSResourceId1R17Constraints); err != nil {
			return err
		}
	}

	// 2. nzp-CSI-RS-ResourceId2-r17: integer
	{
		if err := e.EncodeInteger(ie.Nzp_CSI_RS_ResourceId2_r17, nZPCSIRSPairingR17NzpCSIRSResourceId2R17Constraints); err != nil {
			return err
		}
	}

	return nil
}

func (ie *NZP_CSI_RS_Pairing_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(nZPCSIRSPairingR17Constraints)
	_ = seq

	// 1. nzp-CSI-RS-ResourceId1-r17: integer
	{
		v0, err := d.DecodeInteger(nZPCSIRSPairingR17NzpCSIRSResourceId1R17Constraints)
		if err != nil {
			return err
		}
		ie.Nzp_CSI_RS_ResourceId1_r17 = v0
	}

	// 2. nzp-CSI-RS-ResourceId2-r17: integer
	{
		v1, err := d.DecodeInteger(nZPCSIRSPairingR17NzpCSIRSResourceId2R17Constraints)
		if err != nil {
			return err
		}
		ie.Nzp_CSI_RS_ResourceId2_r17 = v1
	}

	return nil
}
