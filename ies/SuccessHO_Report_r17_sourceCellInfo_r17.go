package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SuccessHO_Report_r17_sourceCellInfo_r17 struct {
	SourcePCellId_r17    CGI_Info_Logging_r16                                          `madatory`
	SourceCellMeas_r17   *MeasResultSuccessHONR_r17                                    `optional`
	Rlf_InSourceDAPS_r17 *SuccessHO_Report_r17_sourceCellInfo_r17_rlf_InSourceDAPS_r17 `optional`
}

func (ie *SuccessHO_Report_r17_sourceCellInfo_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SourceCellMeas_r17 != nil, ie.Rlf_InSourceDAPS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SourcePCellId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode SourcePCellId_r17", err)
	}
	if ie.SourceCellMeas_r17 != nil {
		if err = ie.SourceCellMeas_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SourceCellMeas_r17", err)
		}
	}
	if ie.Rlf_InSourceDAPS_r17 != nil {
		if err = ie.Rlf_InSourceDAPS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Rlf_InSourceDAPS_r17", err)
		}
	}
	return nil
}

func (ie *SuccessHO_Report_r17_sourceCellInfo_r17) Decode(r *uper.UperReader) error {
	var err error
	var SourceCellMeas_r17Present bool
	if SourceCellMeas_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rlf_InSourceDAPS_r17Present bool
	if Rlf_InSourceDAPS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SourcePCellId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode SourcePCellId_r17", err)
	}
	if SourceCellMeas_r17Present {
		ie.SourceCellMeas_r17 = new(MeasResultSuccessHONR_r17)
		if err = ie.SourceCellMeas_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SourceCellMeas_r17", err)
		}
	}
	if Rlf_InSourceDAPS_r17Present {
		ie.Rlf_InSourceDAPS_r17 = new(SuccessHO_Report_r17_sourceCellInfo_r17_rlf_InSourceDAPS_r17)
		if err = ie.Rlf_InSourceDAPS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Rlf_InSourceDAPS_r17", err)
		}
	}
	return nil
}
