package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Fr1_r17 struct {
	Scs_15kHz_r17 *uper.BitString `lb:16,ub:16,optional`
	Scs_30kHz_r17 *uper.BitString `lb:16,ub:16,optional`
	Scs_60kHz_r17 *uper.BitString `lb:16,ub:16,optional`
}

func (ie *Fr1_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_15kHz_r17 != nil, ie.Scs_30kHz_r17 != nil, ie.Scs_60kHz_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs_15kHz_r17 != nil {
		if err = w.WriteBitString(ie.Scs_15kHz_r17.Bytes, uint(ie.Scs_15kHz_r17.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Scs_15kHz_r17", err)
		}
	}
	if ie.Scs_30kHz_r17 != nil {
		if err = w.WriteBitString(ie.Scs_30kHz_r17.Bytes, uint(ie.Scs_30kHz_r17.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Scs_30kHz_r17", err)
		}
	}
	if ie.Scs_60kHz_r17 != nil {
		if err = w.WriteBitString(ie.Scs_60kHz_r17.Bytes, uint(ie.Scs_60kHz_r17.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Scs_60kHz_r17", err)
		}
	}
	return nil
}

func (ie *Fr1_r17) Decode(r *uper.UperReader) error {
	var err error
	var Scs_15kHz_r17Present bool
	if Scs_15kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_30kHz_r17Present bool
	if Scs_30kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_60kHz_r17Present bool
	if Scs_60kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_15kHz_r17Present {
		var tmp_bs_Scs_15kHz_r17 []byte
		var n_Scs_15kHz_r17 uint
		if tmp_bs_Scs_15kHz_r17, n_Scs_15kHz_r17, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Scs_15kHz_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Scs_15kHz_r17,
			NumBits: uint64(n_Scs_15kHz_r17),
		}
		ie.Scs_15kHz_r17 = &tmp_bitstring
	}
	if Scs_30kHz_r17Present {
		var tmp_bs_Scs_30kHz_r17 []byte
		var n_Scs_30kHz_r17 uint
		if tmp_bs_Scs_30kHz_r17, n_Scs_30kHz_r17, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Scs_30kHz_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Scs_30kHz_r17,
			NumBits: uint64(n_Scs_30kHz_r17),
		}
		ie.Scs_30kHz_r17 = &tmp_bitstring
	}
	if Scs_60kHz_r17Present {
		var tmp_bs_Scs_60kHz_r17 []byte
		var n_Scs_60kHz_r17 uint
		if tmp_bs_Scs_60kHz_r17, n_Scs_60kHz_r17, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Scs_60kHz_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Scs_60kHz_r17,
			NumBits: uint64(n_Scs_60kHz_r17),
		}
		ie.Scs_60kHz_r17 = &tmp_bitstring
	}
	return nil
}
