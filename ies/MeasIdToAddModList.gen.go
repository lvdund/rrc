// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasIdToAddModList (line 9326).
// MeasIdToAddModList ::=              SEQUENCE (SIZE (1..maxNrofMeasId)) OF MeasIdToAddMod

var measIdToAddModListSizeConstraints = per.SizeRange(1, common.MaxNrofMeasId)

type MeasIdToAddModList []MeasIdToAddMod

func (ie *MeasIdToAddModList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measIdToAddModListSizeConstraints)
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

func (ie *MeasIdToAddModList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measIdToAddModListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasIdToAddModList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
