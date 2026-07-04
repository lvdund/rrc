// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SRS-OffsetCombo-r18 (line 14884).
// SRS-OffsetCombo-r18 ::=                SEQUENCE (SIZE (1.. maxNrofCellsInSet-r18)) OF INTEGER (0..3)

var sRSOffsetComboR18SizeConstraints = per.SizeRange(1, common.MaxNrofCellsInSet_r18)

type SRS_OffsetCombo_r18 []int64

func (ie *SRS_OffsetCombo_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sRSOffsetComboR18SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeInteger((*ie)[i], per.Constrained(0, 3)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *SRS_OffsetCombo_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sRSOffsetComboR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SRS_OffsetCombo_r18, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeInteger(per.Constrained(0, 3))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
