package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_cellReselectionInfoCommon struct {
	NrofSS_BlocksToAverage          *int64                                                    `lb:2,ub:maxNrofSS_BlocksToAverage,optional`
	AbsThreshSS_BlocksConsolidation *ThresholdNR                                              `optional`
	RangeToBestCell                 *RangeToBestCell                                          `optional`
	Q_Hyst                          SIB2_cellReselectionInfoCommon_q_Hyst                     `madatory`
	SpeedStateReselectionPars       *SIB2_cellReselectionInfoCommon_speedStateReselectionPars `optional`
}

func (ie *SIB2_cellReselectionInfoCommon) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.NrofSS_BlocksToAverage != nil, ie.AbsThreshSS_BlocksConsolidation != nil, ie.RangeToBestCell != nil, ie.SpeedStateReselectionPars != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.NrofSS_BlocksToAverage != nil {
		if err = w.WriteInteger(*ie.NrofSS_BlocksToAverage, &aper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Encode NrofSS_BlocksToAverage", err)
		}
	}
	if ie.AbsThreshSS_BlocksConsolidation != nil {
		if err = ie.AbsThreshSS_BlocksConsolidation.Encode(w); err != nil {
			return utils.WrapError("Encode AbsThreshSS_BlocksConsolidation", err)
		}
	}
	if ie.RangeToBestCell != nil {
		if err = ie.RangeToBestCell.Encode(w); err != nil {
			return utils.WrapError("Encode RangeToBestCell", err)
		}
	}
	if err = ie.Q_Hyst.Encode(w); err != nil {
		return utils.WrapError("Encode Q_Hyst", err)
	}
	if ie.SpeedStateReselectionPars != nil {
		if err = ie.SpeedStateReselectionPars.Encode(w); err != nil {
			return utils.WrapError("Encode SpeedStateReselectionPars", err)
		}
	}
	return nil
}

func (ie *SIB2_cellReselectionInfoCommon) Decode(r *aper.AperReader) error {
	var err error
	var NrofSS_BlocksToAveragePresent bool
	if NrofSS_BlocksToAveragePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AbsThreshSS_BlocksConsolidationPresent bool
	if AbsThreshSS_BlocksConsolidationPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var RangeToBestCellPresent bool
	if RangeToBestCellPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SpeedStateReselectionParsPresent bool
	if SpeedStateReselectionParsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if NrofSS_BlocksToAveragePresent {
		var tmp_int_NrofSS_BlocksToAverage int64
		if tmp_int_NrofSS_BlocksToAverage, err = r.ReadInteger(&aper.Constraint{Lb: 2, Ub: maxNrofSS_BlocksToAverage}, false); err != nil {
			return utils.WrapError("Decode NrofSS_BlocksToAverage", err)
		}
		ie.NrofSS_BlocksToAverage = &tmp_int_NrofSS_BlocksToAverage
	}
	if AbsThreshSS_BlocksConsolidationPresent {
		ie.AbsThreshSS_BlocksConsolidation = new(ThresholdNR)
		if err = ie.AbsThreshSS_BlocksConsolidation.Decode(r); err != nil {
			return utils.WrapError("Decode AbsThreshSS_BlocksConsolidation", err)
		}
	}
	if RangeToBestCellPresent {
		ie.RangeToBestCell = new(RangeToBestCell)
		if err = ie.RangeToBestCell.Decode(r); err != nil {
			return utils.WrapError("Decode RangeToBestCell", err)
		}
	}
	if err = ie.Q_Hyst.Decode(r); err != nil {
		return utils.WrapError("Decode Q_Hyst", err)
	}
	if SpeedStateReselectionParsPresent {
		ie.SpeedStateReselectionPars = new(SIB2_cellReselectionInfoCommon_speedStateReselectionPars)
		if err = ie.SpeedStateReselectionPars.Decode(r); err != nil {
			return utils.WrapError("Decode SpeedStateReselectionPars", err)
		}
	}
	return nil
}
