package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultCellSFTD_NR struct {
	PhysCellId                PhysCellId  `madatory`
	Sfn_OffsetResult          int64       `lb:0,ub:1023,madatory`
	FrameBoundaryOffsetResult int64       `lb:-30720,ub:30719,madatory`
	Rsrp_Result               *RSRP_Range `optional`
}

func (ie *MeasResultCellSFTD_NR) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Rsrp_Result != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PhysCellId.Encode(w); err != nil {
		return utils.WrapError("Encode PhysCellId", err)
	}
	if err = w.WriteInteger(ie.Sfn_OffsetResult, &aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("WriteInteger Sfn_OffsetResult", err)
	}
	if err = w.WriteInteger(ie.FrameBoundaryOffsetResult, &aper.Constraint{Lb: -30720, Ub: 30719}, false); err != nil {
		return utils.WrapError("WriteInteger FrameBoundaryOffsetResult", err)
	}
	if ie.Rsrp_Result != nil {
		if err = ie.Rsrp_Result.Encode(w); err != nil {
			return utils.WrapError("Encode Rsrp_Result", err)
		}
	}
	return nil
}

func (ie *MeasResultCellSFTD_NR) Decode(r *aper.AperReader) error {
	var err error
	var Rsrp_ResultPresent bool
	if Rsrp_ResultPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PhysCellId.Decode(r); err != nil {
		return utils.WrapError("Decode PhysCellId", err)
	}
	var tmp_int_Sfn_OffsetResult int64
	if tmp_int_Sfn_OffsetResult, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 1023}, false); err != nil {
		return utils.WrapError("ReadInteger Sfn_OffsetResult", err)
	}
	ie.Sfn_OffsetResult = tmp_int_Sfn_OffsetResult
	var tmp_int_FrameBoundaryOffsetResult int64
	if tmp_int_FrameBoundaryOffsetResult, err = r.ReadInteger(&aper.Constraint{Lb: -30720, Ub: 30719}, false); err != nil {
		return utils.WrapError("ReadInteger FrameBoundaryOffsetResult", err)
	}
	ie.FrameBoundaryOffsetResult = tmp_int_FrameBoundaryOffsetResult
	if Rsrp_ResultPresent {
		ie.Rsrp_Result = new(RSRP_Range)
		if err = ie.Rsrp_Result.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrp_Result", err)
		}
	}
	return nil
}
