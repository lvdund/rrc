// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CondExecutionCondToReleaseList-r18 (line 6498).
// CondExecutionCondToReleaseList-r18 ::= SEQUENCE (SIZE (1.. maxNrofCondCells-r16)) OF CondReconfigId-r16

var condExecutionCondToReleaseListR18SizeConstraints = per.SizeRange(1, common.MaxNrofCondCells_r16)

type CondExecutionCondToReleaseList_r18 []CondReconfigId_r16

func (ie *CondExecutionCondToReleaseList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(condExecutionCondToReleaseListR18SizeConstraints)
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

func (ie *CondExecutionCondToReleaseList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(condExecutionCondToReleaseListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CondExecutionCondToReleaseList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
