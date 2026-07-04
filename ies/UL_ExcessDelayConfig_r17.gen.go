// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UL-ExcessDelayConfig-r17 (line 16251).

var uLExcessDelayConfigR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "excessDelay-DRBlist-r17"},
	},
}

var uLExcessDelayConfigR17ExcessDelayDRBlistR17Constraints = per.SizeRange(1, common.MaxDRB)

type UL_ExcessDelayConfig_r17 struct {
	ExcessDelay_DRBlist_r17 []ExcessDelay_DRB_IdentityInfo_r17
}

func (ie *UL_ExcessDelayConfig_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(uLExcessDelayConfigR17Constraints)
	_ = seq

	// 1. excessDelay-DRBlist-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(uLExcessDelayConfigR17ExcessDelayDRBlistR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.ExcessDelay_DRBlist_r17))); err != nil {
			return err
		}
		for i := range ie.ExcessDelay_DRBlist_r17 {
			if err := ie.ExcessDelay_DRBlist_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *UL_ExcessDelayConfig_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(uLExcessDelayConfigR17Constraints)
	_ = seq

	// 1. excessDelay-DRBlist-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(uLExcessDelayConfigR17ExcessDelayDRBlistR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.ExcessDelay_DRBlist_r17 = make([]ExcessDelay_DRB_IdentityInfo_r17, n)
		for i := int64(0); i < n; i++ {
			if err := ie.ExcessDelay_DRBlist_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
