// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: CAG-Config-r18 (line 26087).

var cAGConfigR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity-r18"},
		{Name: "cag-IdentityList-r18"},
	},
}

var cAGConfigR18CagIdentityListR18Constraints = per.SizeRange(1, common.MaxNPN_r16)

type CAG_Config_r18 struct {
	Plmn_Identity_r18    PLMN_Identity
	Cag_IdentityList_r18 []per.BitString
}

func (ie *CAG_Config_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(cAGConfigR18Constraints)
	_ = seq

	// 1. plmn-Identity-r18: ref
	{
		if err := ie.Plmn_Identity_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. cag-IdentityList-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(cAGConfigR18CagIdentityListR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Cag_IdentityList_r18))); err != nil {
			return err
		}
		for i := range ie.Cag_IdentityList_r18 {
			if err := e.EncodeBitString(ie.Cag_IdentityList_r18[i], per.FixedSize(32)); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *CAG_Config_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(cAGConfigR18Constraints)
	_ = seq

	// 1. plmn-Identity-r18: ref
	{
		if err := ie.Plmn_Identity_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. cag-IdentityList-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(cAGConfigR18CagIdentityListR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Cag_IdentityList_r18 = make([]per.BitString, n)
		for i := int64(0); i < n; i++ {
			v, err := d.DecodeBitString(per.FixedSize(32))
			if err != nil {
				return err
			}
			ie.Cag_IdentityList_r18[i] = v
		}
	}

	return nil
}
