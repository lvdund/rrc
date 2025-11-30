package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCSetupRequest_IEs struct {
	Ue_Identity        InitialUE_Identity `madatory`
	EstablishmentCause EstablishmentCause `madatory`
	Spare              aper.BitString     `lb:1,ub:1,madatory`
}

func (ie *RRCSetupRequest_IEs) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Ue_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode Ue_Identity", err)
	}
	if err = ie.EstablishmentCause.Encode(w); err != nil {
		return utils.WrapError("Encode EstablishmentCause", err)
	}
	if err = w.WriteBitString(ie.Spare.Bytes, uint(ie.Spare.NumBits), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteBitString Spare", err)
	}
	return nil
}

func (ie *RRCSetupRequest_IEs) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Ue_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode Ue_Identity", err)
	}
	if err = ie.EstablishmentCause.Decode(r); err != nil {
		return utils.WrapError("Decode EstablishmentCause", err)
	}
	var tmp_bs_Spare []byte
	var n_Spare uint
	if tmp_bs_Spare, n_Spare, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadBitString Spare", err)
	}
	ie.Spare = aper.BitString{
		Bytes:   tmp_bs_Spare,
		NumBits: uint64(n_Spare),
	}
	return nil
}
