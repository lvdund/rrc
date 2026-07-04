// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UL-DelayValueConfig-r16 (line 16244).

var uLDelayValueConfigR16Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "delay-DRBlist-r16"},
	},
}

var uLDelayValueConfigR16DelayDRBlistR16Constraints = per.SizeRange(1, common.MaxDRB)

type UL_DelayValueConfig_r16 struct {
	Delay_DRBlist_r16 []DRB_Identity
}

func (ie *UL_DelayValueConfig_r16) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLDelayValueConfigR16Constraints)
	_ = seq

	// 1. delay-DRBlist-r16: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(uLDelayValueConfigR16DelayDRBlistR16Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Delay_DRBlist_r16))); err != nil {
			return err
		}
		for i := range ie.Delay_DRBlist_r16 {
			if err := ie.Delay_DRBlist_r16[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UL_DelayValueConfig_r16) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLDelayValueConfigR16Constraints)
	_ = seq

	// 1. delay-DRBlist-r16: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(uLDelayValueConfigR16DelayDRBlistR16Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Delay_DRBlist_r16 = make([]DRB_Identity, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Delay_DRBlist_r16[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
