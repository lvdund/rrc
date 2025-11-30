package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_PreconfigGeneral_r16 struct {
	Sl_TDD_Configuration_r16 *TDD_UL_DL_ConfigCommon `optional`
	ReservedBits_r16         *aper.BitString         `lb:2,ub:2,optional`
}

func (ie *SL_PreconfigGeneral_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_TDD_Configuration_r16 != nil, ie.ReservedBits_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_TDD_Configuration_r16 != nil {
		if err = ie.Sl_TDD_Configuration_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_TDD_Configuration_r16", err)
		}
	}
	if ie.ReservedBits_r16 != nil {
		if err = w.WriteBitString(ie.ReservedBits_r16.Bytes, uint(ie.ReservedBits_r16.NumBits), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Encode ReservedBits_r16", err)
		}
	}
	return nil
}

func (ie *SL_PreconfigGeneral_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_TDD_Configuration_r16Present bool
	if Sl_TDD_Configuration_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ReservedBits_r16Present bool
	if ReservedBits_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_TDD_Configuration_r16Present {
		ie.Sl_TDD_Configuration_r16 = new(TDD_UL_DL_ConfigCommon)
		if err = ie.Sl_TDD_Configuration_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_TDD_Configuration_r16", err)
		}
	}
	if ReservedBits_r16Present {
		var tmp_bs_ReservedBits_r16 []byte
		var n_ReservedBits_r16 uint
		if tmp_bs_ReservedBits_r16, n_ReservedBits_r16, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
			return utils.WrapError("Decode ReservedBits_r16", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_ReservedBits_r16,
			NumBits: uint64(n_ReservedBits_r16),
		}
		ie.ReservedBits_r16 = &tmp_bitstring
	}
	return nil
}
