// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: MRB-ToAddModList-r17 (line 13178).
// MRB-ToAddModList-r17 ::=                SEQUENCE (SIZE (1..maxMRB-r17)) OF MRB-ToAddMod-r17

var mRBToAddModListR17SizeConstraints = per.SizeRange(1, common.MaxMRB_r17)

type MRB_ToAddModList_r17 []MRB_ToAddMod_r17

func (ie *MRB_ToAddModList_r17) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(mRBToAddModListR17SizeConstraints)
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

func (ie *MRB_ToAddModList_r17) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(mRBToAddModListR17SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(MRB_ToAddModList_r17, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
