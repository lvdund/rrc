// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: CellGlobalIdList-r16 (line 26075).
// CellGlobalIdList-r16 ::=         SEQUENCE (SIZE (1..32)) OF CGI-Info-Logging-r16

var cellGlobalIdListR16SizeConstraints = per.SizeRange(1, 32)

type CellGlobalIdList_r16 []CGI_Info_Logging_r16

func (ie *CellGlobalIdList_r16) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cellGlobalIdListR16SizeConstraints)
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

func (ie *CellGlobalIdList_r16) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cellGlobalIdListR16SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CellGlobalIdList_r16, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
