package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NeedForGapsInfoNR_r16 struct {
	IntraFreq_needForGap_r16 NeedForGapsIntraFreqList_r16 `madatory`
	InterFreq_needForGap_r16 NeedForGapsBandListNR_r16    `madatory`
}

func (ie *NeedForGapsInfoNR_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.IntraFreq_needForGap_r16.Encode(w); err != nil {
		return utils.WrapError("Encode IntraFreq_needForGap_r16", err)
	}
	if err = ie.InterFreq_needForGap_r16.Encode(w); err != nil {
		return utils.WrapError("Encode InterFreq_needForGap_r16", err)
	}
	return nil
}

func (ie *NeedForGapsInfoNR_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.IntraFreq_needForGap_r16.Decode(r); err != nil {
		return utils.WrapError("Decode IntraFreq_needForGap_r16", err)
	}
	if err = ie.InterFreq_needForGap_r16.Decode(r); err != nil {
		return utils.WrapError("Decode InterFreq_needForGap_r16", err)
	}
	return nil
}
