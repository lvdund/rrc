package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasResultsPerCellIdleNR_r16_measIdleResultNR_r16 struct {
	Rsrp_Result_r16        *RSRP_Range                  `optional`
	Rsrq_Result_r16        *RSRQ_Range                  `optional`
	ResultsSSB_Indexes_r16 *ResultsPerSSB_IndexList_r16 `optional`
}

func (ie *MeasResultsPerCellIdleNR_r16_measIdleResultNR_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Rsrp_Result_r16 != nil, ie.Rsrq_Result_r16 != nil, ie.ResultsSSB_Indexes_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Rsrp_Result_r16 != nil {
		if err = ie.Rsrp_Result_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rsrp_Result_r16", err)
		}
	}
	if ie.Rsrq_Result_r16 != nil {
		if err = ie.Rsrq_Result_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Rsrq_Result_r16", err)
		}
	}
	if ie.ResultsSSB_Indexes_r16 != nil {
		if err = ie.ResultsSSB_Indexes_r16.Encode(w); err != nil {
			return utils.WrapError("Encode ResultsSSB_Indexes_r16", err)
		}
	}
	return nil
}

func (ie *MeasResultsPerCellIdleNR_r16_measIdleResultNR_r16) Decode(r *aper.AperReader) error {
	var err error
	var Rsrp_Result_r16Present bool
	if Rsrp_Result_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Rsrq_Result_r16Present bool
	if Rsrq_Result_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ResultsSSB_Indexes_r16Present bool
	if ResultsSSB_Indexes_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Rsrp_Result_r16Present {
		ie.Rsrp_Result_r16 = new(RSRP_Range)
		if err = ie.Rsrp_Result_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrp_Result_r16", err)
		}
	}
	if Rsrq_Result_r16Present {
		ie.Rsrq_Result_r16 = new(RSRQ_Range)
		if err = ie.Rsrq_Result_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Rsrq_Result_r16", err)
		}
	}
	if ResultsSSB_Indexes_r16Present {
		ie.ResultsSSB_Indexes_r16 = new(ResultsPerSSB_IndexList_r16)
		if err = ie.ResultsSSB_Indexes_r16.Decode(r); err != nil {
			return utils.WrapError("Decode ResultsSSB_Indexes_r16", err)
		}
	}
	return nil
}
