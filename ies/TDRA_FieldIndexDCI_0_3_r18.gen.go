// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TDRA-FieldIndexDCI-0-3-r18 (line 14890).
// TDRA-FieldIndexDCI-0-3-r18 ::=         SEQUENCE (SIZE (2.. maxNrofBWPsInSetOfCells-r18)) OF INTEGER (0..maxNrofUL-Allocations-1-r18)

var tDRAFieldIndexDCI03R18SizeConstraints = per.SizeRange(2, common.MaxNrofBWPsInSetOfCells_r18)

type TDRA_FieldIndexDCI_0_3_r18 []int64

func (ie *TDRA_FieldIndexDCI_0_3_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(tDRAFieldIndexDCI03R18SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(0, common.MaxNrofUL_Allocations_1_r18)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *TDRA_FieldIndexDCI_0_3_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(tDRAFieldIndexDCI03R18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(TDRA_FieldIndexDCI_0_3_r18, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(0, common.MaxNrofUL_Allocations_1_r18))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
