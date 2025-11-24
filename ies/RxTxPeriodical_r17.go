package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RxTxPeriodical_r17 struct {
	RxTxReportInterval_r17 *RxTxReportInterval_r17             `optional`
	ReportAmount_r17       RxTxPeriodical_r17_reportAmount_r17 `madatory`
}

func (ie *RxTxPeriodical_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.RxTxReportInterval_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.RxTxReportInterval_r17 != nil {
		if err = ie.RxTxReportInterval_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RxTxReportInterval_r17", err)
		}
	}
	if err = ie.ReportAmount_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ReportAmount_r17", err)
	}
	return nil
}

func (ie *RxTxPeriodical_r17) Decode(r *uper.UperReader) error {
	var err error
	var RxTxReportInterval_r17Present bool
	if RxTxReportInterval_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if RxTxReportInterval_r17Present {
		ie.RxTxReportInterval_r17 = new(RxTxReportInterval_r17)
		if err = ie.RxTxReportInterval_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RxTxReportInterval_r17", err)
		}
	}
	if err = ie.ReportAmount_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ReportAmount_r17", err)
	}
	return nil
}
