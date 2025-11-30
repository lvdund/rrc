package ies

import (
	"fmt"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

const (
	IAB_IP_Address_r16_Choice_nothing uint64 = iota
	IAB_IP_Address_r16_Choice_IPv4_Address_r16
	IAB_IP_Address_r16_Choice_IPv6_Address_r16
	IAB_IP_Address_r16_Choice_IPv6_Prefix_r16
)

type IAB_IP_Address_r16 struct {
	Choice           uint64
	IPv4_Address_r16 aper.BitString `lb:32,ub:32,madatory`
	IPv6_Address_r16 aper.BitString `lb:128,ub:128,madatory`
	IPv6_Prefix_r16  aper.BitString `lb:64,ub:64,madatory`
}

func (ie *IAB_IP_Address_r16) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteChoice(ie.Choice, 3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case IAB_IP_Address_r16_Choice_IPv4_Address_r16:
		if err = w.WriteBitString(ie.IPv4_Address_r16.Bytes, uint(ie.IPv4_Address_r16.NumBits), &aper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			err = utils.WrapError("Encode IPv4_Address_r16", err)
		}
	case IAB_IP_Address_r16_Choice_IPv6_Address_r16:
		if err = w.WriteBitString(ie.IPv6_Address_r16.Bytes, uint(ie.IPv6_Address_r16.NumBits), &aper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			err = utils.WrapError("Encode IPv6_Address_r16", err)
		}
	case IAB_IP_Address_r16_Choice_IPv6_Prefix_r16:
		if err = w.WriteBitString(ie.IPv6_Prefix_r16.Bytes, uint(ie.IPv6_Prefix_r16.NumBits), &aper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			err = utils.WrapError("Encode IPv6_Prefix_r16", err)
		}
	default:
		err = fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return err
}

func (ie *IAB_IP_Address_r16) Decode(r *aper.AperReader) error {
	var err error
	if ie.Choice, err = r.ReadChoice(3, false); err != nil {
		return err
	}
	switch ie.Choice {
	case IAB_IP_Address_r16_Choice_IPv4_Address_r16:
		var tmp_bs_IPv4_Address_r16 []byte
		var n_IPv4_Address_r16 uint
		if tmp_bs_IPv4_Address_r16, n_IPv4_Address_r16, err = r.ReadBitString(&aper.Constraint{Lb: 32, Ub: 32}, false); err != nil {
			return utils.WrapError("Decode IPv4_Address_r16", err)
		}
		ie.IPv4_Address_r16 = aper.BitString{
			Bytes:   tmp_bs_IPv4_Address_r16,
			NumBits: uint64(n_IPv4_Address_r16),
		}
	case IAB_IP_Address_r16_Choice_IPv6_Address_r16:
		var tmp_bs_IPv6_Address_r16 []byte
		var n_IPv6_Address_r16 uint
		if tmp_bs_IPv6_Address_r16, n_IPv6_Address_r16, err = r.ReadBitString(&aper.Constraint{Lb: 128, Ub: 128}, false); err != nil {
			return utils.WrapError("Decode IPv6_Address_r16", err)
		}
		ie.IPv6_Address_r16 = aper.BitString{
			Bytes:   tmp_bs_IPv6_Address_r16,
			NumBits: uint64(n_IPv6_Address_r16),
		}
	case IAB_IP_Address_r16_Choice_IPv6_Prefix_r16:
		var tmp_bs_IPv6_Prefix_r16 []byte
		var n_IPv6_Prefix_r16 uint
		if tmp_bs_IPv6_Prefix_r16, n_IPv6_Prefix_r16, err = r.ReadBitString(&aper.Constraint{Lb: 64, Ub: 64}, false); err != nil {
			return utils.WrapError("Decode IPv6_Prefix_r16", err)
		}
		ie.IPv6_Prefix_r16 = aper.BitString{
			Bytes:   tmp_bs_IPv6_Prefix_r16,
			NumBits: uint64(n_IPv6_Prefix_r16),
		}
	default:
		return fmt.Errorf("invalid choice: %d", ie.Choice)
	}
	return nil
}
