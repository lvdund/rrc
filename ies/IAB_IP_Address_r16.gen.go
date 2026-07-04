// Code generated from docs/NR-RRC-Definitions.asn by codegen.
// DO NOT EDIT.

package ies

import (
	"github.com/lvdund/asn1go/per"
)

// Source: IAB-IP-Address-r16 (line 26217).

var iABIPAddressR16Constraints = per.ChoiceConstraints{
	Extensible: true,
	RootAlternatives: []per.AlternativeInfo{
		{Name: "iPv4-Address-r16"},
		{Name: "iPv6-Address-r16"},
		{Name: "iPv6-Prefix-r16"},
	},
}

const (
	IAB_IP_Address_r16_IPv4_Address_r16 = 0
	IAB_IP_Address_r16_IPv6_Address_r16 = 1
	IAB_IP_Address_r16_IPv6_Prefix_r16  = 2
)

type IAB_IP_Address_r16 struct {
	Choice           int
	IPv4_Address_r16 *per.BitString
	IPv6_Address_r16 *per.BitString
	IPv6_Prefix_r16  *per.BitString
}

func (ie *IAB_IP_Address_r16) Encode(e *per.Encoder) error {
	choiceEnc := e.NewChoiceEncoder(iABIPAddressR16Constraints)
	isExt := false
	if err := choiceEnc.EncodeChoice(int64(ie.Choice), isExt, nil); err != nil {
		return err
	}
	switch ie.Choice {
	case IAB_IP_Address_r16_IPv4_Address_r16:
		if err := e.EncodeBitString((*ie.IPv4_Address_r16), per.FixedSize(32)); err != nil {
			return err
		}
	case IAB_IP_Address_r16_IPv6_Address_r16:
		if err := e.EncodeBitString((*ie.IPv6_Address_r16), per.FixedSize(128)); err != nil {
			return err
		}
	case IAB_IP_Address_r16_IPv6_Prefix_r16:
		if err := e.EncodeBitString((*ie.IPv6_Prefix_r16), per.FixedSize(64)); err != nil {
			return err
		}
	default:
		// Extension alternative: not modeled; bail out.
		return &per.ConstraintViolationError{
			TypeName:   "IAB-IP-Address-r16",
			Value:      int64(ie.Choice),
			Constraint: "choice index out of root range",
		}
	}
	return nil
}

func (ie *IAB_IP_Address_r16) Decode(d *per.Decoder) error {
	choiceDec := d.NewChoiceDecoder(iABIPAddressR16Constraints)
	index, isExt, _, err := choiceDec.DecodeChoice()
	if err != nil {
		return err
	}
	ie.Choice = int(index)
	if isExt {
		return &per.DecodeError{TypeName: "IAB-IP-Address-r16", Reason: "extension alternative not supported", Position: 0}
	}
	switch index {
	case IAB_IP_Address_r16_IPv4_Address_r16:
		ie.IPv4_Address_r16 = new(per.BitString)
		v, err := d.DecodeBitString(per.FixedSize(32))
		if err != nil {
			return err
		}
		(*ie.IPv4_Address_r16) = v
	case IAB_IP_Address_r16_IPv6_Address_r16:
		ie.IPv6_Address_r16 = new(per.BitString)
		v, err := d.DecodeBitString(per.FixedSize(128))
		if err != nil {
			return err
		}
		(*ie.IPv6_Address_r16) = v
	case IAB_IP_Address_r16_IPv6_Prefix_r16:
		ie.IPv6_Prefix_r16 = new(per.BitString)
		v, err := d.DecodeBitString(per.FixedSize(64))
		if err != nil {
			return err
		}
		(*ie.IPv6_Prefix_r16) = v
	default:
		return &per.DecodeError{TypeName: "IAB-IP-Address-r16", Reason: "choice index out of root range", Position: 0}
	}
	return nil
}
