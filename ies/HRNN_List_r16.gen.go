// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: HRNN-List-r16 (line 4254).
// HRNN-List-r16 ::=           SEQUENCE (SIZE (1..maxNPN-r16)) OF HRNN-r16

var hRNNListR16SizeConstraints = per.SizeRange(1, common.MaxNPN_r16)

type HRNN_List_r16 []HRNN_r16

func (ie *HRNN_List_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(hRNNListR16SizeConstraints)
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

func (ie *HRNN_List_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(hRNNListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(HRNN_List_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
