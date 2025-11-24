package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReestablishmentRequest_IEs struct {
	Ue_Identity          ReestabUE_Identity   `madatory`
	ReestablishmentCause ReestablishmentCause `madatory`
	Spare                uper.BitString       `lb:1,ub:1,madatory`
}

func (ie *RRCReestablishmentRequest_IEs) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Ue_Identity.Encode(w); err != nil {
		return utils.WrapError("Encode Ue_Identity", err)
	}
	if err = ie.ReestablishmentCause.Encode(w); err != nil {
		return utils.WrapError("Encode ReestablishmentCause", err)
	}
	if err = w.WriteBitString(ie.Spare.Bytes, uint(ie.Spare.NumBits), &uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteBitString Spare", err)
	}
	return nil
}

func (ie *RRCReestablishmentRequest_IEs) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Ue_Identity.Decode(r); err != nil {
		return utils.WrapError("Decode Ue_Identity", err)
	}
	if err = ie.ReestablishmentCause.Decode(r); err != nil {
		return utils.WrapError("Decode ReestablishmentCause", err)
	}
	var tmp_bs_Spare []byte
	var n_Spare uint
	if tmp_bs_Spare, n_Spare, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadBitString Spare", err)
	}
	ie.Spare = uper.BitString{
		Bytes:   tmp_bs_Spare,
		NumBits: uint64(n_Spare),
	}
	return nil
}
