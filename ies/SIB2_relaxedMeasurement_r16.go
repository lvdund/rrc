package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_relaxedMeasurement_r16 struct {
	LowMobilityEvaluation_r16       *SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16       `optional`
	CellEdgeEvaluation_r16          *SIB2_relaxedMeasurement_r16_cellEdgeEvaluation_r16          `optional`
	CombineRelaxedMeasCondition_r16 *SIB2_relaxedMeasurement_r16_combineRelaxedMeasCondition_r16 `optional`
	HighPriorityMeasRelax_r16       *SIB2_relaxedMeasurement_r16_highPriorityMeasRelax_r16       `optional`
}

func (ie *SIB2_relaxedMeasurement_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.LowMobilityEvaluation_r16 != nil, ie.CellEdgeEvaluation_r16 != nil, ie.CombineRelaxedMeasCondition_r16 != nil, ie.HighPriorityMeasRelax_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.LowMobilityEvaluation_r16 != nil {
		if err = ie.LowMobilityEvaluation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode LowMobilityEvaluation_r16", err)
		}
	}
	if ie.CellEdgeEvaluation_r16 != nil {
		if err = ie.CellEdgeEvaluation_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CellEdgeEvaluation_r16", err)
		}
	}
	if ie.CombineRelaxedMeasCondition_r16 != nil {
		if err = ie.CombineRelaxedMeasCondition_r16.Encode(w); err != nil {
			return utils.WrapError("Encode CombineRelaxedMeasCondition_r16", err)
		}
	}
	if ie.HighPriorityMeasRelax_r16 != nil {
		if err = ie.HighPriorityMeasRelax_r16.Encode(w); err != nil {
			return utils.WrapError("Encode HighPriorityMeasRelax_r16", err)
		}
	}
	return nil
}

func (ie *SIB2_relaxedMeasurement_r16) Decode(r *aper.AperReader) error {
	var err error
	var LowMobilityEvaluation_r16Present bool
	if LowMobilityEvaluation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CellEdgeEvaluation_r16Present bool
	if CellEdgeEvaluation_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CombineRelaxedMeasCondition_r16Present bool
	if CombineRelaxedMeasCondition_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HighPriorityMeasRelax_r16Present bool
	if HighPriorityMeasRelax_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if LowMobilityEvaluation_r16Present {
		ie.LowMobilityEvaluation_r16 = new(SIB2_relaxedMeasurement_r16_lowMobilityEvaluation_r16)
		if err = ie.LowMobilityEvaluation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode LowMobilityEvaluation_r16", err)
		}
	}
	if CellEdgeEvaluation_r16Present {
		ie.CellEdgeEvaluation_r16 = new(SIB2_relaxedMeasurement_r16_cellEdgeEvaluation_r16)
		if err = ie.CellEdgeEvaluation_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CellEdgeEvaluation_r16", err)
		}
	}
	if CombineRelaxedMeasCondition_r16Present {
		ie.CombineRelaxedMeasCondition_r16 = new(SIB2_relaxedMeasurement_r16_combineRelaxedMeasCondition_r16)
		if err = ie.CombineRelaxedMeasCondition_r16.Decode(r); err != nil {
			return utils.WrapError("Decode CombineRelaxedMeasCondition_r16", err)
		}
	}
	if HighPriorityMeasRelax_r16Present {
		ie.HighPriorityMeasRelax_r16 = new(SIB2_relaxedMeasurement_r16_highPriorityMeasRelax_r16)
		if err = ie.HighPriorityMeasRelax_r16.Decode(r); err != nil {
			return utils.WrapError("Decode HighPriorityMeasRelax_r16", err)
		}
	}
	return nil
}
