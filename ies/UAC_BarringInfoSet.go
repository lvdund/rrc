package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringInfoSet struct {
	Uac_BarringFactor            UAC_BarringInfoSet_uac_BarringFactor `madatory`
	Uac_BarringTime              UAC_BarringInfoSet_uac_BarringTime   `madatory`
	Uac_BarringForAccessIdentity aper.BitString                       `lb:7,ub:7,madatory`
}

func (ie *UAC_BarringInfoSet) Encode(w *aper.AperWriter) error {
	var err error
	if err = ie.Uac_BarringFactor.Encode(w); err != nil {
		return utils.WrapError("Encode Uac_BarringFactor", err)
	}
	if err = ie.Uac_BarringTime.Encode(w); err != nil {
		return utils.WrapError("Encode Uac_BarringTime", err)
	}
	if err = w.WriteBitString(ie.Uac_BarringForAccessIdentity.Bytes, uint(ie.Uac_BarringForAccessIdentity.NumBits), &aper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
		return utils.WrapError("WriteBitString Uac_BarringForAccessIdentity", err)
	}
	return nil
}

func (ie *UAC_BarringInfoSet) Decode(r *aper.AperReader) error {
	var err error
	if err = ie.Uac_BarringFactor.Decode(r); err != nil {
		return utils.WrapError("Decode Uac_BarringFactor", err)
	}
	if err = ie.Uac_BarringTime.Decode(r); err != nil {
		return utils.WrapError("Decode Uac_BarringTime", err)
	}
	var tmp_bs_Uac_BarringForAccessIdentity []byte
	var n_Uac_BarringForAccessIdentity uint
	if tmp_bs_Uac_BarringForAccessIdentity, n_Uac_BarringForAccessIdentity, err = r.ReadBitString(&aper.Constraint{Lb: 7, Ub: 7}, false); err != nil {
		return utils.WrapError("ReadBitString Uac_BarringForAccessIdentity", err)
	}
	ie.Uac_BarringForAccessIdentity = aper.BitString{
		Bytes:   tmp_bs_Uac_BarringForAccessIdentity,
		NumBits: uint64(n_Uac_BarringForAccessIdentity),
	}
	return nil
}
