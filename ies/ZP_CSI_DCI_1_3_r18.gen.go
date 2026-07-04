// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: ZP-CSI-DCI-1-3-r18 (line 14878).
// ZP-CSI-DCI-1-3-r18 ::=                 SEQUENCE (SIZE (1.. maxNrofCellsInSet-r18)) OF BIT STRING (SIZE (1..2))

var zPCSIDCI13R18SizeConstraints = per.SizeRange(1, common.MaxNrofCellsInSet_r18)

type ZP_CSI_DCI_1_3_r18 []per.BitString

func (ie *ZP_CSI_DCI_1_3_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(zPCSIDCI13R18SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeBitString((*ie)[i], per.SizeRange(1, 2)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *ZP_CSI_DCI_1_3_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(zPCSIDCI13R18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(ZP_CSI_DCI_1_3_r18, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeBitString(per.SizeRange(1, 2))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
