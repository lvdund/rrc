// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasResultServMOList (line 9751).
// MeasResultServMOList ::=                SEQUENCE (SIZE (1..maxNrofServingCells)) OF MeasResultServMO

var measResultServMOListSizeConstraints = per.SizeRange(1, common.MaxNrofServingCells)

type MeasResultServMOList []MeasResultServMO

func (ie *MeasResultServMOList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measResultServMOListSizeConstraints)
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

func (ie *MeasResultServMOList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measResultServMOListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasResultServMOList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
