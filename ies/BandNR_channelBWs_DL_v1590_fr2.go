package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_channelBWs_DL_v1590_fr2 struct {
	Scs_60kHz  *uper.BitString `lb:8,ub:8,optional`
	Scs_120kHz *uper.BitString `lb:8,ub:8,optional`
}

func (ie *BandNR_channelBWs_DL_v1590_fr2) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_60kHz != nil, ie.Scs_120kHz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs_60kHz != nil {
		if err = w.WriteBitString(ie.Scs_60kHz.Bytes, uint(ie.Scs_60kHz.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Scs_60kHz", err)
		}
	}
	if ie.Scs_120kHz != nil {
		if err = w.WriteBitString(ie.Scs_120kHz.Bytes, uint(ie.Scs_120kHz.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode Scs_120kHz", err)
		}
	}
	return nil
}

func (ie *BandNR_channelBWs_DL_v1590_fr2) Decode(r *uper.UperReader) error {
	var err error
	var Scs_60kHzPresent bool
	if Scs_60kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_120kHzPresent bool
	if Scs_120kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_60kHzPresent {
		var tmp_bs_Scs_60kHz []byte
		var n_Scs_60kHz uint
		if tmp_bs_Scs_60kHz, n_Scs_60kHz, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Scs_60kHz", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Scs_60kHz,
			NumBits: uint64(n_Scs_60kHz),
		}
		ie.Scs_60kHz = &tmp_bitstring
	}
	if Scs_120kHzPresent {
		var tmp_bs_Scs_120kHz []byte
		var n_Scs_120kHz uint
		if tmp_bs_Scs_120kHz, n_Scs_120kHz, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode Scs_120kHz", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_Scs_120kHz,
			NumBits: uint64(n_Scs_120kHz),
		}
		ie.Scs_120kHz = &tmp_bitstring
	}
	return nil
}
