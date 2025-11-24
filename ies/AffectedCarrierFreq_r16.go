package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type AffectedCarrierFreq_r16 struct {
	CarrierFreq_r16           ARFCN_ValueNR                                     `madatory`
	InterferenceDirection_r16 AffectedCarrierFreq_r16_interferenceDirection_r16 `madatory`
}

func (ie *AffectedCarrierFreq_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.CarrierFreq_r16.Encode(w); err != nil {
		return utils.WrapError("Encode CarrierFreq_r16", err)
	}
	if err = ie.InterferenceDirection_r16.Encode(w); err != nil {
		return utils.WrapError("Encode InterferenceDirection_r16", err)
	}
	return nil
}

func (ie *AffectedCarrierFreq_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.CarrierFreq_r16.Decode(r); err != nil {
		return utils.WrapError("Decode CarrierFreq_r16", err)
	}
	if err = ie.InterferenceDirection_r16.Decode(r); err != nil {
		return utils.WrapError("Decode InterferenceDirection_r16", err)
	}
	return nil
}
