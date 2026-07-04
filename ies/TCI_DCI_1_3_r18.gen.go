// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: TCI-DCI-1-3-r18 (line 14880).
// TCI-DCI-1-3-r18 ::=                    SEQUENCE (SIZE (2.. maxNrofCellsInSet-r18)) OF BIT STRING (SIZE (3))

var tCIDCI13R18SizeConstraints = per.SizeRange(2, common.MaxNrofCellsInSet_r18)

type TCI_DCI_1_3_r18 []per.BitString

func (ie *TCI_DCI_1_3_r18) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(tCIDCI13R18SizeConstraints)
	if err := seqOf.EncodeLength(int64(len(*ie))); err != nil {
		return err
	}
	for i := range *ie {
		if err := e.EncodeBitString((*ie)[i], per.FixedSize(3)); err != nil {
			return err
		}
	}
	return nil
}

func (ie *TCI_DCI_1_3_r18) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(tCIDCI13R18SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(TCI_DCI_1_3_r18, n)
	for i := int64(0); i < n; i++ {
		v, err := d.DecodeBitString(per.FixedSize(3))
		if err != nil {
			return err
		}
		(*ie)[i] = v
	}
	return nil
}
