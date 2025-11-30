package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_channelBWs_DL_v1590_fr1 struct {
	Scs_15kHz *aper.BitString `lb:16,ub:16,optional`
	Scs_30kHz *aper.BitString `lb:16,ub:16,optional`
	Scs_60kHz *aper.BitString `lb:16,ub:16,optional`
}

func (ie *BandNR_channelBWs_DL_v1590_fr1) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_15kHz != nil, ie.Scs_30kHz != nil, ie.Scs_60kHz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs_15kHz != nil {
		if err = w.WriteBitString(ie.Scs_15kHz.Bytes, uint(ie.Scs_15kHz.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Scs_15kHz", err)
		}
	}
	if ie.Scs_30kHz != nil {
		if err = w.WriteBitString(ie.Scs_30kHz.Bytes, uint(ie.Scs_30kHz.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Scs_30kHz", err)
		}
	}
	if ie.Scs_60kHz != nil {
		if err = w.WriteBitString(ie.Scs_60kHz.Bytes, uint(ie.Scs_60kHz.NumBits), &aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Encode Scs_60kHz", err)
		}
	}
	return nil
}

func (ie *BandNR_channelBWs_DL_v1590_fr1) Decode(r *aper.AperReader) error {
	var err error
	var Scs_15kHzPresent bool
	if Scs_15kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_30kHzPresent bool
	if Scs_30kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_60kHzPresent bool
	if Scs_60kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_15kHzPresent {
		var tmp_bs_Scs_15kHz []byte
		var n_Scs_15kHz uint
		if tmp_bs_Scs_15kHz, n_Scs_15kHz, err = r.ReadBitString(&aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Scs_15kHz", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_Scs_15kHz,
			NumBits: uint64(n_Scs_15kHz),
		}
		ie.Scs_15kHz = &tmp_bitstring
	}
	if Scs_30kHzPresent {
		var tmp_bs_Scs_30kHz []byte
		var n_Scs_30kHz uint
		if tmp_bs_Scs_30kHz, n_Scs_30kHz, err = r.ReadBitString(&aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Scs_30kHz", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_Scs_30kHz,
			NumBits: uint64(n_Scs_30kHz),
		}
		ie.Scs_30kHz = &tmp_bitstring
	}
	if Scs_60kHzPresent {
		var tmp_bs_Scs_60kHz []byte
		var n_Scs_60kHz uint
		if tmp_bs_Scs_60kHz, n_Scs_60kHz, err = r.ReadBitString(&aper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
			return utils.WrapError("Decode Scs_60kHz", err)
		}
		tmp_bitstring := aper.BitString{
			Bytes:   tmp_bs_Scs_60kHz,
			NumBits: uint64(n_Scs_60kHz),
		}
		ie.Scs_60kHz = &tmp_bitstring
	}
	return nil
}
