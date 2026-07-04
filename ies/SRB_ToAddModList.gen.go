// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: SRB-ToAddModList (line 13131).
// SRB-ToAddModList ::=                    SEQUENCE (SIZE (1..2)) OF SRB-ToAddMod

var sRBToAddModListSizeConstraints = per.SizeRange(1, 2)

type SRB_ToAddModList []SRB_ToAddMod

func (ie *SRB_ToAddModList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(sRBToAddModListSizeConstraints)
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

func (ie *SRB_ToAddModList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(sRBToAddModListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(SRB_ToAddModList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
