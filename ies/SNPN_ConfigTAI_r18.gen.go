// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: SNPN-ConfigTAI-r18 (line 26107).

var sNPNConfigTAIR18Constraints = per.SequenceConstraints{
	Extensible: false,
	RootComponents: []per.ComponentInfo{
		{Name: "tai-Identity-r18"},
		{Name: "nid-IdentityList-r18"},
	},
}

var sNPNConfigTAIR18NidIdentityListR18Constraints = per.SizeRange(1, common.MaxNPN_r16)

type SNPN_ConfigTAI_r18 struct {
	Tai_Identity_r18     TrackingAreaIdentity_r16
	Nid_IdentityList_r18 []NID_r16
}

func (ie *SNPN_ConfigTAI_r18) Encode(e *per.Encoder) error {
	seq := e.NewSequenceEncoder(sNPNConfigTAIR18Constraints)
	_ = seq

	// 1. tai-Identity-r18: ref
	{
		if err := ie.Tai_Identity_r18.Encode(e); err != nil {
			return err
		}
	}

	// 2. nid-IdentityList-r18: sequence-of
	{
		seqOf := e.NewSequenceOfEncoder(sNPNConfigTAIR18NidIdentityListR18Constraints)
		if err := seqOf.EncodeLength(int64(len(ie.Nid_IdentityList_r18))); err != nil {
			return err
		}
		for i := range ie.Nid_IdentityList_r18 {
			if err := ie.Nid_IdentityList_r18[i].Encode(e); err != nil {
				return err
			}
		}
	}

	return nil
}

func (ie *SNPN_ConfigTAI_r18) Decode(d *per.Decoder) error {
	seq := d.NewSequenceDecoder(sNPNConfigTAIR18Constraints)
	_ = seq

	// 1. tai-Identity-r18: ref
	{
		if err := ie.Tai_Identity_r18.Decode(d); err != nil {
			return err
		}
	}

	// 2. nid-IdentityList-r18: sequence-of
	{
		seqOf := d.NewSequenceOfDecoder(sNPNConfigTAIR18NidIdentityListR18Constraints)
		n, err := seqOf.DecodeLength()
		if err != nil {
			return err
		}
		ie.Nid_IdentityList_r18 = make([]NID_r16, n)
		for i := int64(0); i < n; i++ {
			if err := ie.Nid_IdentityList_r18[i].Decode(d); err != nil {
				return err
			}
		}
	}

	return nil
}
