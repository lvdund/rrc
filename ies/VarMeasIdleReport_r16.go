package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type VarMeasIdleReport_r16 struct {
	MeasReportIdleNR_r16    *MeasResultIdleNR_r16    `optional`
	MeasReportIdleEUTRA_r16 *MeasResultIdleEUTRA_r16 `optional`
}

func (ie *VarMeasIdleReport_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MeasReportIdleNR_r16 != nil, ie.MeasReportIdleEUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MeasReportIdleNR_r16 != nil {
		if err = ie.MeasReportIdleNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasReportIdleNR_r16", err)
		}
	}
	if ie.MeasReportIdleEUTRA_r16 != nil {
		if err = ie.MeasReportIdleEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MeasReportIdleEUTRA_r16", err)
		}
	}
	return nil
}

func (ie *VarMeasIdleReport_r16) Decode(r *uper.UperReader) error {
	var err error
	var MeasReportIdleNR_r16Present bool
	if MeasReportIdleNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MeasReportIdleEUTRA_r16Present bool
	if MeasReportIdleEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MeasReportIdleNR_r16Present {
		ie.MeasReportIdleNR_r16 = new(MeasResultIdleNR_r16)
		if err = ie.MeasReportIdleNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasReportIdleNR_r16", err)
		}
	}
	if MeasReportIdleEUTRA_r16Present {
		ie.MeasReportIdleEUTRA_r16 = new(MeasResultIdleEUTRA_r16)
		if err = ie.MeasReportIdleEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MeasReportIdleEUTRA_r16", err)
		}
	}
	return nil
}
