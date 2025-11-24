package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultCBR_NR_r16 struct {
	Sl_poolReportIdentity_r16 SL_ResourcePoolID_r16 `madatory`
	Sl_CBR_ResultsNR_r16      SL_CBR_r16            `madatory`
}

func (ie *MeasResultCBR_NR_r16) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Sl_poolReportIdentity_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_poolReportIdentity_r16", err)
	}
	if err = ie.Sl_CBR_ResultsNR_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_CBR_ResultsNR_r16", err)
	}
	return nil
}

func (ie *MeasResultCBR_NR_r16) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Sl_poolReportIdentity_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_poolReportIdentity_r16", err)
	}
	if err = ie.Sl_CBR_ResultsNR_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_CBR_ResultsNR_r16", err)
	}
	return nil
}
