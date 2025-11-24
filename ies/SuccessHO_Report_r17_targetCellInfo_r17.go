package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SuccessHO_Report_r17_targetCellInfo_r17 struct {
	TargetPCellId_r17  CGI_Info_Logging_r16       `madatory`
	TargetCellMeas_r17 *MeasResultSuccessHONR_r17 `optional`
}

func (ie *SuccessHO_Report_r17_targetCellInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.TargetCellMeas_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.TargetPCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode TargetPCellId_r17", err)
	}
	if ie.TargetCellMeas_r17 != nil {
		if err = ie.TargetCellMeas_r17.Encode(w); err != nil {
			return utils.WrapError("Encode TargetCellMeas_r17", err)
		}
	}
	return nil
}

func (ie *SuccessHO_Report_r17_targetCellInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var TargetCellMeas_r17Present bool
	if TargetCellMeas_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.TargetPCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode TargetPCellId_r17", err)
	}
	if TargetCellMeas_r17Present {
		ie.TargetCellMeas_r17 = new(MeasResultSuccessHONR_r17)
		if err = ie.TargetCellMeas_r17.Decode(r); err != nil {
			return utils.WrapError("Decode TargetCellMeas_r17", err)
		}
	}
	return nil
}
