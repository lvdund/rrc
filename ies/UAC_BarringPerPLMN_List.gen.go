// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: UAC-BarringPerPLMN-List (line 16207).
// UAC-BarringPerPLMN-List ::=         SEQUENCE (SIZE (1.. maxPLMN)) OF UAC-BarringPerPLMN

var uACBarringPerPLMNListSizeConstraints = per.SizeRange(1, common.MaxPLMN)

type UAC_BarringPerPLMN_List []UAC_BarringPerPLMN

func (ie *UAC_BarringPerPLMN_List) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(uACBarringPerPLMNListSizeConstraints)
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

func (ie *UAC_BarringPerPLMN_List) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(uACBarringPerPLMNListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(UAC_BarringPerPLMN_List, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
