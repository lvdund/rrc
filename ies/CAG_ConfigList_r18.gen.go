// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CAG-ConfigList-r18 (line 26085).
// CAG-ConfigList-r18 ::=           SEQUENCE (SIZE (1..maxNPN-r16)) OF CAG-Config-r18

var cAGConfigListR18SizeConstraints = per.SizeRange(1, common.MaxNPN_r16)

type CAG_ConfigList_r18 []CAG_Config_r18

func (ie *CAG_ConfigList_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(cAGConfigListR18SizeConstraints)
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

func (ie *CAG_ConfigList_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(cAGConfigListR18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(CAG_ConfigList_r18, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
