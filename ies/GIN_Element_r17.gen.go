// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: GIN-Element-r17 (line 4497).

var gINElementR17Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "plmn-Identity-r17"},
		{Name: "nid-List-r17"},
	},
}

var gINElementR17NidListR17Constraints = per.SizeRange(1, common.MaxGIN_r17)

type GIN_Element_r17 struct {
	Plmn_Identity_r17 PLMN_Identity
	Nid_List_r17      []NID_r16
}

func (ie *GIN_Element_r17) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(gINElementR17Constraints)
	_ = seq

	// 1. plmn-Identity-r17: ref
	{
		if err := ie.Plmn_Identity_r17.Encode(e); err != nil {
			return err
		}
	}

	// 2. nid-List-r17: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(gINElementR17NidListR17Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Nid_List_r17))); err != nil {
			return err
		}
		for i := range ie.Nid_List_r17 {
			if err := ie.Nid_List_r17[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *GIN_Element_r17) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(gINElementR17Constraints)
	_ = seq

	// 1. plmn-Identity-r17: ref
	{
		if err := ie.Plmn_Identity_r17.Decode(d); err != nil {
			return err
		}
	}

	// 2. nid-List-r17: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(gINElementR17NidListR17Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Nid_List_r17 = make([]NID_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Nid_List_r17[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
