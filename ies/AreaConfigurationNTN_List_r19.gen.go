// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: AreaConfigurationNTN-List-r19 (line 599).
// AreaConfigurationNTN-List-r19 ::= SEQUENCE (SIZE (1..maxNrofAreaNTN-r19)) OF AreaConfigurationNTN-r19

var areaConfigurationNTNListR19SizeConstraints = per.SizeRange(1, common.MaxNrofAreaNTN_r19)

type AreaConfigurationNTN_List_r19 []AreaConfigurationNTN_r19

func (ie *AreaConfigurationNTN_List_r19) Encode(e *per.Encoder) error {
	seqOf := e.NewSequenceOfEncoder(areaConfigurationNTNListR19SizeConstraints)
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

func (ie *AreaConfigurationNTN_List_r19) Decode(d *per.Decoder) error {
	seqOf := d.NewSequenceOfDecoder(areaConfigurationNTNListR19SizeConstraints)
	n, err := seqOf.DecodeLength()
	if err != nil {
		return err
	}
	*ie = make(AreaConfigurationNTN_List_r19, n)
	for i := int64(0); i < n; i++ {
		if err := (*ie)[i].Decode(d); err != nil {
			return err
		}
	}
	return nil
}
