package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResult2EUTRA_r16 struct {
	CarrierFreq_r16    ARFCN_ValueEUTRA    `madatory`
	MeasResultList_r16 MeasResultListEUTRA `madatory`
}

func (ie *MeasResult2EUTRA_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.CarrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq_r16", err)
	}
	if err = ie.MeasResultList_r16.Encode(w); err != nil {
		return utils.WrapError("Encode MeasResultList_r16", err)
	}
	return nil
}

func (ie *MeasResult2EUTRA_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.CarrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq_r16", err)
	}
	if err = ie.MeasResultList_r16.Decode(r); err != nil {
		return utils.WrapError("Decode MeasResultList_r16", err)
	}
	return nil
}
