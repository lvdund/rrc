package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17 struct {
	Scs15kHz_15kHz_r17 *CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17_scs15kHz_15kHz_r17 `optional`
	Scs15kHz_30kHz_r17 *CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17_scs15kHz_30kHz_r17 `optional`
	Scs15kHz_60kHz_r17 *CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17_scs15kHz_60kHz_r17 `optional`
	Scs30kHz_30kHz_r17 *uper.BitString                                                                          `lb:1,ub:496,optional`
	Scs30kHz_60kHz_r17 *uper.BitString                                                                          `lb:1,ub:496,optional`
	Scs60kHz_60kHz_r17 *uper.BitString                                                                          `lb:1,ub:496,optional`
}

func (ie *CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs15kHz_15kHz_r17 != nil, ie.Scs15kHz_30kHz_r17 != nil, ie.Scs15kHz_60kHz_r17 != nil, ie.Scs30kHz_30kHz_r17 != nil, ie.Scs30kHz_60kHz_r17 != nil, ie.Scs60kHz_60kHz_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs15kHz_15kHz_r17 != nil {
		if err = ie.Scs15kHz_15kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs15kHz_15kHz_r17", err)
		}
	}
	if ie.Scs15kHz_30kHz_r17 != nil {
		if err = ie.Scs15kHz_30kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs15kHz_30kHz_r17", err)
		}
	}
	if ie.Scs15kHz_60kHz_r17 != nil {
		if err = ie.Scs15kHz_60kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs15kHz_60kHz_r17", err)
		}
	}
	if ie.Scs30kHz_30kHz_r17 != nil {
		if err = w.WriteBitString(ie.Scs30kHz_30kHz_r17.Bytes, uint(ie.Scs30kHz_30kHz_r17.NumBits), &uper.Constraint{Lb: 1, Ub: 496}, false); err != nil {
			return utils.WrapError("Encode Scs30kHz_30kHz_r17", err)
		}
	}
	if ie.Scs30kHz_60kHz_r17 != nil {
		if err = w.WriteBitString(ie.Scs30kHz_60kHz_r17.Bytes, uint(ie.Scs30kHz_60kHz_r17.NumBits), &uper.Constraint{Lb: 1, Ub: 496}, false); err != nil {
			return utils.WrapError("Encode Scs30kHz_60kHz_r17", err)
		}
	}
	if ie.Scs60kHz_60kHz_r17 != nil {
		if err = w.WriteBitString(ie.Scs60kHz_60kHz_r17.Bytes, uint(ie.Scs60kHz_60kHz_r17.NumBits), &uper.Constraint{Lb: 1, Ub: 496}, false); err != nil {
			return utils.WrapError("Encode Scs60kHz_60kHz_r17", err)
		}
	}
	return nil
}

func (ie *CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17) Decode(r *uper.UperReader) error {
	var err error
	var Scs15kHz_15kHz_r17Present bool
	if Scs15kHz_15kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs15kHz_30kHz_r17Present bool
	if Scs15kHz_30kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs15kHz_60kHz_r17Present bool
	if Scs15kHz_60kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs30kHz_30kHz_r17Present bool
	if Scs30kHz_30kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs30kHz_60kHz_r17Present bool
	if Scs30kHz_60kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs60kHz_60kHz_r17Present bool
	if Scs60kHz_60kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs15kHz_15kHz_r17Present {
		ie.Scs15kHz_15kHz_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17_scs15kHz_15kHz_r17)
		if err = ie.Scs15kHz_15kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs15kHz_15kHz_r17", err)
		}
	}
	if Scs15kHz_30kHz_r17Present {
		ie.Scs15kHz_30kHz_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17_scs15kHz_30kHz_r17)
		if err = ie.Scs15kHz_30kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs15kHz_30kHz_r17", err)
		}
	}
	if Scs15kHz_60kHz_r17Present {
		ie.Scs15kHz_60kHz_r17 = new(CrossCarrierSchedulingSCell_SpCell_r17_supportedSCS_Combinations_r17_scs15kHz_60kHz_r17)
		if err = ie.Scs15kHz_60kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs15kHz_60kHz_r17", err)
		}
	}
	if Scs30kHz_30kHz_r17Present {
		var tmp_bs_Scs30kHz_30kHz_r17 []byte
		var n_Scs30kHz_30kHz_r17 uint
		if tmp_bs_Scs30kHz_30kHz_r17, n_Scs30kHz_30kHz_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 496}, false); err != nil {
			return utils.WrapError("Decode Scs30kHz_30kHz_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Scs30kHz_30kHz_r17,
			NumBits: uint64(n_Scs30kHz_30kHz_r17),
		}
		ie.Scs30kHz_30kHz_r17 = &tmp_bitstring
	}
	if Scs30kHz_60kHz_r17Present {
		var tmp_bs_Scs30kHz_60kHz_r17 []byte
		var n_Scs30kHz_60kHz_r17 uint
		if tmp_bs_Scs30kHz_60kHz_r17, n_Scs30kHz_60kHz_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 496}, false); err != nil {
			return utils.WrapError("Decode Scs30kHz_60kHz_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Scs30kHz_60kHz_r17,
			NumBits: uint64(n_Scs30kHz_60kHz_r17),
		}
		ie.Scs30kHz_60kHz_r17 = &tmp_bitstring
	}
	if Scs60kHz_60kHz_r17Present {
		var tmp_bs_Scs60kHz_60kHz_r17 []byte
		var n_Scs60kHz_60kHz_r17 uint
		if tmp_bs_Scs60kHz_60kHz_r17, n_Scs60kHz_60kHz_r17, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 496}, false); err != nil {
			return utils.WrapError("Decode Scs60kHz_60kHz_r17", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Scs60kHz_60kHz_r17,
			NumBits: uint64(n_Scs60kHz_60kHz_r17),
		}
		ie.Scs60kHz_60kHz_r17 = &tmp_bitstring
	}
	return nil
}
