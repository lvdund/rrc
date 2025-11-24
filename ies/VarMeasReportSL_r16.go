package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarMeasReportSL_r16 struct {
	Sl_MeasId_r16                 SL_MeasId_r16   `madatory`
	Sl_FrequencyTriggeredList_r16 []ARFCN_ValueNR `lb:1,ub:maxNrofFreqSL_r16,optional`
	Sl_NumberOfReportsSent_r16    int64           `madatory`
}

func (ie *VarMeasReportSL_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.Sl_FrequencyTriggeredList_r16) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sl_MeasId_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Sl_MeasId_r16", err)
	}
	if len(ie.Sl_FrequencyTriggeredList_r16) > 0 {
		tmp_Sl_FrequencyTriggeredList_r16 := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		for _, i := range ie.Sl_FrequencyTriggeredList_r16 {
			tmp_Sl_FrequencyTriggeredList_r16.Value = append(tmp_Sl_FrequencyTriggeredList_r16.Value, &i)
		}
		if err = tmp_Sl_FrequencyTriggeredList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_FrequencyTriggeredList_r16", err)
		}
	}
	if err = w.WriteInteger(ie.Sl_NumberOfReportsSent_r16, &uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("WriteInteger Sl_NumberOfReportsSent_r16", err)
	}
	return nil
}

func (ie *VarMeasReportSL_r16) Decode(r *uper.UperReader) error {
	var err error
	var Sl_FrequencyTriggeredList_r16Present bool
	if Sl_FrequencyTriggeredList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sl_MeasId_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Sl_MeasId_r16", err)
	}
	if Sl_FrequencyTriggeredList_r16Present {
		tmp_Sl_FrequencyTriggeredList_r16 := utils.NewSequence[*ARFCN_ValueNR]([]*ARFCN_ValueNR{}, uper.Constraint{Lb: 1, Ub: maxNrofFreqSL_r16}, false)
		fn_Sl_FrequencyTriggeredList_r16 := func() *ARFCN_ValueNR {
			return new(ARFCN_ValueNR)
		}
		if err = tmp_Sl_FrequencyTriggeredList_r16.Decode(r, fn_Sl_FrequencyTriggeredList_r16); err != nil {
			return utils.WrapError("Decode Sl_FrequencyTriggeredList_r16", err)
		}
		ie.Sl_FrequencyTriggeredList_r16 = []ARFCN_ValueNR{}
		for _, i := range tmp_Sl_FrequencyTriggeredList_r16.Value {
			ie.Sl_FrequencyTriggeredList_r16 = append(ie.Sl_FrequencyTriggeredList_r16, *i)
		}
	}
	var tmp_int_Sl_NumberOfReportsSent_r16 int64
	if tmp_int_Sl_NumberOfReportsSent_r16, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return utils.WrapError("ReadInteger Sl_NumberOfReportsSent_r16", err)
	}
	ie.Sl_NumberOfReportsSent_r16 = tmp_int_Sl_NumberOfReportsSent_r16
	return nil
}
