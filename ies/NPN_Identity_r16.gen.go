// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
	"github.com/lvdund/rrc/common"
)

// Source: NPN-Identity-r16 (line 10560).

var nPNIdentityR16Constraints = per.ChoiceConstraints{
	Extensible: false,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "pni-npn-r16"},
		{Name: "snpn-r16"},
	},
}

const (
	NPN_Identity_r16_Pni_Npn_r16 = 0
	NPN_Identity_r16_Snpn_r16    = 1
)

var nPNIdentityR16PniNpnR16CagIdentityListR16Constraints = per.SizeRange(1, common.MaxNPN_r16)

var nPNIdentityR16SnpnR16NidListR16Constraints = per.SizeRange(1, common.MaxNPN_r16)

type NPN_Identity_r16 struct {
	Choice      int
	Pni_Npn_r16 *struct {
		Plmn_Identity_r16    PLMN_Identity
		Cag_IdentityList_r16 []CAG_IdentityInfo_r16
	}
	Snpn_r16 *struct {
		Plmn_Identity_r16 PLMN_Identity
		Nid_List_r16      []NID_r16
	}
}

func (ie *NPN_Identity_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(nPNIdentityR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case NPN_Identity_r16_Pni_Npn_r16:
		c := (*ie.Pni_Npn_r16)
		if err := c.Plmn_Identity_r16.Encode(e); err != nil {
			return err
		}
		{
			seqOf := e.NewSequenceOfEncoder(nPNIdentityR16PniNpnR16CagIdentityListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Cag_IdentityList_r16))); err != nil {
				return err
			}
			for i := range c.Cag_IdentityList_r16 {
				if err := c.Cag_IdentityList_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	case NPN_Identity_r16_Snpn_r16:
		c := (*ie.Snpn_r16)
		if err := c.Plmn_Identity_r16.Encode(e); err != nil {
			return err
		}
		{
			seqOf := e.NewSequenceOfEncoder(nPNIdentityR16SnpnR16NidListR16Constraints)
			if err := seqOf.EncodeLength(int64(len(c.Nid_List_r16))); err != nil {
				return err
			}
			for i := range c.Nid_List_r16 {
				if err := c.Nid_List_r16[i].Encode(e); err != nil {
					return err
				}
			}
		}
	default:
		// Choice index out of root range.
		return &per.ConstraintViolationError{
			TypeName:   "NPN-Identity-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *NPN_Identity_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(nPNIdentityR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "NPN-Identity-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case NPN_Identity_r16_Pni_Npn_r16:
		ie.Pni_Npn_r16 = &struct {
			Plmn_Identity_r16    PLMN_Identity
			Cag_IdentityList_r16 []CAG_IdentityInfo_r16
		}{}
		c := (*ie.Pni_Npn_r16)
		{
			if err := c.Plmn_Identity_r16.Decode(d); err != nil {
				return err
			}
		}
		{
			seqOf := d.NewSequenceOfDecoder(nPNIdentityR16PniNpnR16CagIdentityListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Cag_IdentityList_r16 = make([]CAG_IdentityInfo_r16, n)
			for i := int64(0); i < n; i++ {
				if err := c.Cag_IdentityList_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	case NPN_Identity_r16_Snpn_r16:
		ie.Snpn_r16 = &struct {
			Plmn_Identity_r16 PLMN_Identity
			Nid_List_r16      []NID_r16
		}{}
		c := (*ie.Snpn_r16)
		{
			if err := c.Plmn_Identity_r16.Decode(d); err != nil {
				return err
			}
		}
		{
			seqOf := d.NewSequenceOfDecoder(nPNIdentityR16SnpnR16NidListR16Constraints)
			n, err := seqOf.DecodeLength()
			if err != nil {
				return err
			}
			c.Nid_List_r16 = make([]NID_r16, n)
			for i := int64(0); i < n; i++ {
				if err := c.Nid_List_r16[i].Decode(d); err != nil {
					return err
				}
			}
		}
	default:
		return &per.DecodeError{TypeName: "NPN-Identity-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
