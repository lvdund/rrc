// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasObjectToRemoveList (line 9125).
// MeasObjectToRemoveList ::=              SEQUENCE (SIZE (1..maxNrofObjectId)) OF MeasObjectId

var measObjectToRemoveListSizeConstraints = per.SizeRange(1, common.MaxNrofObjectId)

type MeasObjectToRemoveList []MeasObjectId

func (ie *MeasObjectToRemoveList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measObjectToRemoveListSizeConstraints)
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

func (ie *MeasObjectToRemoveList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measObjectToRemoveListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasObjectToRemoveList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
