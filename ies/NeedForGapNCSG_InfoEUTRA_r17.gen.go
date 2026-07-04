// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NeedForGapNCSG-InfoEUTRA-r17 (line 10489).

var needForGapNCSGInfoEUTRAR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "needForNCSG-EUTRA-r17"},
	},
}

var needForGapNCSGInfoEUTRAR17NeedForNCSGEUTRAR17Constraints = per.SizeRange(1, common.MaxBandsEUTRA)

type NeedForGapNCSG_InfoEUTRA_r17 struct {
	NeedForNCSG_EUTRA_r17 []NeedForNCSG_EUTRA_r17
}

func (ie *NeedForGapNCSG_InfoEUTRA_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(needForGapNCSGInfoEUTRAR17Constraints)
	_ = seq

	// 1. needForNCSG-EUTRA-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(needForGapNCSGInfoEUTRAR17NeedForNCSGEUTRAR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.NeedForNCSG_EUTRA_r17))); err != nil {
			return err
		}
		for i := range ie.NeedForNCSG_EUTRA_r17 {
			if err := ie.NeedForNCSG_EUTRA_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *NeedForGapNCSG_InfoEUTRA_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(needForGapNCSGInfoEUTRAR17Constraints)
	_ = seq

	// 1. needForNCSG-EUTRA-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(needForGapNCSGInfoEUTRAR17NeedForNCSGEUTRAR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.NeedForNCSG_EUTRA_r17 = make([]NeedForNCSG_EUTRA_r17, n)
		for i := int64(0); i < n; i++ {
			if err := ie.NeedForNCSG_EUTRA_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
