package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type IAB_IP_AddressConfiguration_r16 struct {
	Iab_IP_AddressIndex_r16      IAB_IP_AddressIndex_r16 `madatory`
	Iab_IP_Address_r16           *IAB_IP_Address_r16     `optional`
	Iab_IP_Usage_r16             *IAB_IP_Usage_r16       `optional`
	Iab_donor_DU_BAP_Address_r16 *aper.BitString         `lb:10,ub:10,optional`
}

func (ie *IAB_IP_AddressConfiguration_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Iab_IP_Address_r16 != nil, ie.Iab_IP_Usage_r16 != nil, ie.Iab_donor_DU_BAP_Address_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Iab_IP_AddressIndex_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Iab_IP_AddressIndex_r16", err)
	}
	if ie.Iab_IP_Address_r16 != nil {
		if err = ie.Iab_IP_Address_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Iab_IP_Address_r16", err)
		}
	}
	if ie.Iab_IP_Usage_r16 != nil {
		if err = ie.Iab_IP_Usage_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Iab_IP_Usage_r16", err)
		}
	}
	if ie.Iab_donor_DU_BAP_Address_r16 != nil {
		if err = w.WriteBitString(ie.Iab_donor_DU_BAP_Address_r16.Bytes, uint(ie.Iab_donor_DU_BAP_Address_r16.NumBits), &aper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Encode Iab_donor_DU_BAP_Address_r16", err)
		}
	}
	return nil
}

func (ie *IAB_IP_AddressConfiguration_r16) Decode(r *aper.AperReader) error {
	var err error
	var Iab_IP_Address_r16Present bool
	if Iab_IP_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Iab_IP_Usage_r16Present bool
	if Iab_IP_Usage_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Iab_donor_DU_BAP_Address_r16Present bool
	if Iab_donor_DU_BAP_Address_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Iab_IP_AddressIndex_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Iab_IP_AddressIndex_r16", err)
	}
	if Iab_IP_Address_r16Present {
		ie.Iab_IP_Address_r16 = new(IAB_IP_Address_r16)
		if err = ie.Iab_IP_Address_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Iab_IP_Address_r16", err)
		}
	}
	if Iab_IP_Usage_r16Present {
		ie.Iab_IP_Usage_r16 = new(IAB_IP_Usage_r16)
		if err = ie.Iab_IP_Usage_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Iab_IP_Usage_r16", err)
		}
	}
	if Iab_donor_DU_BAP_Address_r16Present {
		var tmp_bs_Iab_donor_DU_BAP_Address_r16 []byte
		var n_Iab_donor_DU_BAP_Address_r16 uint
		if tmp_bs_Iab_donor_DU_BAP_Address_r16, n_Iab_donor_DU_BAP_Address_r16, err = r.ReadBitString(&aper.Constraint{Lb: 10, Ub: 10}, false); err != nil {
			return utils.WrapError("Decode Iab_donor_DU_BAP_Address_r16", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_Iab_donor_DU_BAP_Address_r16,
			NumBits: uint64(n_Iab_donor_DU_BAP_Address_r16),
		}
		ie.Iab_donor_DU_BAP_Address_r16 = &tmp_bitstring
	}
	return nil
}
