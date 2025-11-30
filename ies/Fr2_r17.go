package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Fr2_r17 struct {
	Scs_60kHz_r17  *aper.BitString `lb:16,ub:16,optional`
	Scs_120kHz_r17 *aper.BitString `lb:16,ub:16,optional`
}

func (ie *Fr2_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_60kHz_r17 != nil, ie.Scs_120kHz_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs_60kHz_r17 != nil {
		if err = w.WriteBitString(ie.Scs_60kHz_r17.Bytes, uint(ie.Scs_60kHz_r17.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Scs_60kHz_r17", err)
		}
	}
	if ie.Scs_120kHz_r17 != nil {
		if err = w.WriteBitString(ie.Scs_120kHz_r17.Bytes, uint(ie.Scs_120kHz_r17.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Scs_120kHz_r17", err)
		}
	}
	return nil
}

func (ie *Fr2_r17) Decode(r *aper.AperReader) error {
	var err error
	var Scs_60kHz_r17Present bool
	if Scs_60kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_120kHz_r17Present bool
	if Scs_120kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_60kHz_r17Present {
		var tmp_bs_Scs_60kHz_r17 []byte
		var n_Scs_60kHz_r17 uint
		if tmp_bs_Scs_60kHz_r17, n_Scs_60kHz_r17, err = r.ReadBitString(&aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Scs_60kHz_r17", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_Scs_60kHz_r17,
			NumBits: uint64(n_Scs_60kHz_r17),
		}
		ie.Scs_60kHz_r17 = &tmp_bitstring
	}
	if Scs_120kHz_r17Present {
		var tmp_bs_Scs_120kHz_r17 []byte
		var n_Scs_120kHz_r17 uint
		if tmp_bs_Scs_120kHz_r17, n_Scs_120kHz_r17, err = r.ReadBitString(&aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Scs_120kHz_r17", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_Scs_120kHz_r17,
			NumBits: uint64(n_Scs_120kHz_r17),
		}
		ie.Scs_120kHz_r17 = &tmp_bitstring
	}
	return nil
}
