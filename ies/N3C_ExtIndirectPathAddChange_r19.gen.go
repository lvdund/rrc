// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: N3C-ExtIndirectPathAddChange-r19 (line 1131).
// N3C-ExtIndirectPathAddChange-r19 ::=            SEQUENCE (SIZE(1..maxNrofN3C-RelayUE-r19)) OF N3C-RelayUE-Info-r18

var n3CExtIndirectPathAddChangeR19SizeConstraints = per.SizeRange(1, common.MaxNrofN3C_RelayUE_r19)

type N3C_ExtIndirectPathAddChange_r19 []N3C_RelayUE_Info_r18

func (ie *N3C_ExtIndirectPathAddChange_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(n3CExtIndirectPathAddChangeR19SizeConstraints)
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

func (ie *N3C_ExtIndirectPathAddChange_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(n3CExtIndirectPathAddChangeR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(N3C_ExtIndirectPathAddChange_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
