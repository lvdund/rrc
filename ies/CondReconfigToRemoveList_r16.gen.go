// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CondReconfigToRemoveList-r16 (line 6515).
// CondReconfigToRemoveList-r16 ::=     SEQUENCE (SIZE (1.. maxNrofCondCells-r16)) OF CondReconfigId-r16

var condReconfigToRemoveListR16SizeConstraints = per.SizeRange(1, common.MaxNrofCondCells_r16)

type CondReconfigToRemoveList_r16 []CondReconfigId_r16

func (ie *CondReconfigToRemoveList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(condReconfigToRemoveListR16SizeConstraints)
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

func (ie *CondReconfigToRemoveList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(condReconfigToRemoveListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CondReconfigToRemoveList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
