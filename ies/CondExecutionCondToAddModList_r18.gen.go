// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CondExecutionCondToAddModList-r18 (line 6489).
// CondExecutionCondToAddModList-r18 ::= SEQUENCE (SIZE (1.. maxNrofCondCells-r16)) OF CondExecutionCondToAddMod-r18

var condExecutionCondToAddModListR18SizeConstraints = per.SizeRange(1, common.MaxNrofCondCells_r16)

type CondExecutionCondToAddModList_r18 []CondExecutionCondToAddMod_r18

func (ie *CondExecutionCondToAddModList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(condExecutionCondToAddModListR18SizeConstraints)
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

func (ie *CondExecutionCondToAddModList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(condExecutionCondToAddModListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CondExecutionCondToAddModList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
