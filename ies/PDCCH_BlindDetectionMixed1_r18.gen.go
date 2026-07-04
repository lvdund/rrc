// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: PDCCH-BlindDetectionMixed1-r18 (line 18464).

var pDCCHBlindDetectionMixed1R18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "pdcch-BlindDetectionCG-UE-Mixed-r18"},
	},
}

type PDCCH_BlindDetectionMixed1_r18 struct {
	Pdcch_BlindDetectionCG_UE_Mixed_r18 struct {
		Pdcch_BlindDetectionMCG_UE_Mixed_r18 int64
		Pdcch_BlindDetectionSCG_UE_Mixed_r18 int64
	}
}

func (ie *PDCCH_BlindDetectionMixed1_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(pDCCHBlindDetectionMixed1R18Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionCG-UE-Mixed-r18: sequence
	{
		{
			c := &ie.Pdcch_BlindDetectionCG_UE_Mixed_r18
			if err := e.EncodeInteger(c.Pdcch_BlindDetectionMCG_UE_Mixed_r18, per.Constrained(1, 15)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.Pdcch_BlindDetectionSCG_UE_Mixed_r18, per.Constrained(1, 15)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *PDCCH_BlindDetectionMixed1_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(pDCCHBlindDetectionMixed1R18Constraints)
	_ = seq

	// 1. pdcch-BlindDetectionCG-UE-Mixed-r18: sequence
	{
		{
			c := &ie.Pdcch_BlindDetectionCG_UE_Mixed_r18
			{
				v, err := d.DecodeInteger(per.Constrained(1, 15))
				if err != nil {
					return err
				}
				c.Pdcch_BlindDetectionMCG_UE_Mixed_r18 = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 15))
				if err != nil {
					return err
				}
				c.Pdcch_BlindDetectionSCG_UE_Mixed_r18 = v
			}
		}
	}

	return nil
}
