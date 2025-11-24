package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16 struct {
	TwoPorts_r16                 *uper.BitString                                                                    `lb:2,ub:2,optional`
	FourPortsNonCoherent_r16     *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsNonCoherent_r16     `optional`
	FourPortsPartialCoherent_r16 *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16 `optional`
}

func (ie *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.TwoPorts_r16 != nil, ie.FourPortsNonCoherent_r16 != nil, ie.FourPortsPartialCoherent_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.TwoPorts_r16 != nil {
		if err = w.WriteBitString(ie.TwoPorts_r16.Bytes, uint(ie.TwoPorts_r16.NumBits), &uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode TwoPorts_r16", err)
		}
	}
	if ie.FourPortsNonCoherent_r16 != nil {
		if err = ie.FourPortsNonCoherent_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FourPortsNonCoherent_r16", err)
		}
	}
	if ie.FourPortsPartialCoherent_r16 != nil {
		if err = ie.FourPortsPartialCoherent_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FourPortsPartialCoherent_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16) Decode(r *uper.UperReader) error {
	var err error
	var TwoPorts_r16Present bool
	if TwoPorts_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var FourPortsNonCoherent_r16Present bool
	if FourPortsNonCoherent_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var FourPortsPartialCoherent_r16Present bool
	if FourPortsPartialCoherent_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if TwoPorts_r16Present {
		var tmp_bs_TwoPorts_r16 []byte
		var n_TwoPorts_r16 uint
		if tmp_bs_TwoPorts_r16, n_TwoPorts_r16, err = r.ReadBitString(&uper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode TwoPorts_r16", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_TwoPorts_r16,
			NumBits: uint64(n_TwoPorts_r16),
		}
		ie.TwoPorts_r16 = &tmp_bitstring
	}
	if FourPortsNonCoherent_r16Present {
		ie.FourPortsNonCoherent_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsNonCoherent_r16)
		if err = ie.FourPortsNonCoherent_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FourPortsNonCoherent_r16", err)
		}
	}
	if FourPortsPartialCoherent_r16Present {
		ie.FourPortsPartialCoherent_r16 = new(FeatureSetUplink_v1610_ul_FullPwrMode2_TPMIGroup_r16_fourPortsPartialCoherent_r16)
		if err = ie.FourPortsPartialCoherent_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FourPortsPartialCoherent_r16", err)
		}
	}
	return nil
}
