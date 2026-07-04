// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MeasObjectToAddModList (line 9645).
// MeasObjectToAddModList ::=                  SEQUENCE (SIZE (1..maxNrofObjectId)) OF MeasObjectToAddMod

var measObjectToAddModListSizeConstraints = per.SizeRange(1, common.MaxNrofObjectId)

type MeasObjectToAddModList []MeasObjectToAddMod

func (ie *MeasObjectToAddModList) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(measObjectToAddModListSizeConstraints)
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

func (ie *MeasObjectToAddModList) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(measObjectToAddModListSizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MeasObjectToAddModList, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
