// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MBS-FSAI-List-r17 (line 4598).
// MBS-FSAI-List-r17 ::= SEQUENCE (SIZE (1..maxFSAI-MBS-r17)) OF MBS-FSAI-r17

var mBSFSAIListR17SizeConstraints = per.SizeRange(1, common.MaxFSAI_MBS_r17)

type MBS_FSAI_List_r17 []MBS_FSAI_r17

func (ie *MBS_FSAI_List_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mBSFSAIListR17SizeConstraints)
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

func (ie *MBS_FSAI_List_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mBSFSAIListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MBS_FSAI_List_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
