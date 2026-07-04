// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: FeatureSetUplinkPerCC-v1540 (line 20555).

var featureSetUplinkPerCCV1540Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "mimo-NonCB-PUSCH", Optional: true},
	},
}

type FeatureSetUplinkPerCC_v1540 struct {
	Mimo_NonCB_PUSCH *struct {
		MaxNumberSRS_ResourcePerSet         int64
		MaxNumberSimultaneousSRS_ResourceTx int64
	}
}

func (ie *FeatureSetUplinkPerCC_v1540) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(featureSetUplinkPerCCV1540Constraints)

	// 1. Root OPTIONAL bitmap, one bit per optional/default field.
	if err := seq.EncodePreamble([]bool{ie.Mimo_NonCB_PUSCH != nil}); err != nil {
		return err
	}

	// 2. mimo-NonCB-PUSCH: sequence
	{
		if ie.Mimo_NonCB_PUSCH != nil {
			c := ie.Mimo_NonCB_PUSCH
			if err := e.EncodeInteger(c.MaxNumberSRS_ResourcePerSet, per.Constrained(1, 4)); err != nil {
				return err
			}
			if err := e.EncodeInteger(c.MaxNumberSimultaneousSRS_ResourceTx, per.Constrained(1, 4)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *FeatureSetUplinkPerCC_v1540) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(featureSetUplinkPerCCV1540Constraints)

	// 1. Decode root OPTIONAL bitmap.
	if err := seq.DecodePreamble(); err != nil {
		return err
	}

	// 2. mimo-NonCB-PUSCH: sequence
	{
		if seq.IsComponentPresent(0) {
			ie.Mimo_NonCB_PUSCH = &struct {
				MaxNumberSRS_ResourcePerSet         int64
				MaxNumberSimultaneousSRS_ResourceTx int64
			}{}
			c := ie.Mimo_NonCB_PUSCH
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberSRS_ResourcePerSet = v
			}
			{
				v, err := d.DecodeInteger(per.Constrained(1, 4))
				if err != nil {
					return err
				}
				c.MaxNumberSimultaneousSRS_ResourceTx = v
			}
		}
	}

	return nil
}
