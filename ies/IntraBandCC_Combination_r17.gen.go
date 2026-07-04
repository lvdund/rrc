// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: IntraBandCC-Combination-r17 (line 5790).
// IntraBandCC-Combination-r17::=      SEQUENCE (SIZE(1.. maxNrofServingCells)) OF CC-State-r17

var intraBandCCCombinationR17SizeConstraints = per.SizeRange(1, common.MaxNrofServingCells)

type IntraBandCC_Combination_r17 []CC_State_r17

func (ie *IntraBandCC_Combination_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(intraBandCCCombinationR17SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := (*ie)[i].Encode(e); err != nil {
			return err
		}
	}
	return nil
}

func (ie *IntraBandCC_Combination_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(intraBandCCCombinationR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(IntraBandCC_Combination_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
