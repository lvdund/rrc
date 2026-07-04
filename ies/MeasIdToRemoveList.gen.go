// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasIdToRemoveList (line 9127).
// MeasIdToRemoveList ::=                  SEQUENCE (SIZE (1..maxNrofMeasId)) OF MeasId

var measIdToRemoveListSizeConstraints = per.SizeRange(1, common.MaxNrofMeasId)

type MeasIdToRemoveList []MeasId

func (ie *MeasIdToRemoveList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measIdToRemoveListSizeConstraints)
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

func (ie *MeasIdToRemoveList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measIdToRemoveListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasIdToRemoveList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
