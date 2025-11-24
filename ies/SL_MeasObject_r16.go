package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasObject_r16 struct {
	FrequencyInfoSL_r16 ARFCN_ValueNR `madatory`
}

func (ie *SL_MeasObject_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.FrequencyInfoSL_r16.Encode(w); err != nil {
		return utils.WrapError("Encode FrequencyInfoSL_r16", err)
	}
	return nil
}

func (ie *SL_MeasObject_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.FrequencyInfoSL_r16.Decode(r); err != nil {
		return utils.WrapError("Decode FrequencyInfoSL_r16", err)
	}
	return nil
}
