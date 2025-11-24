package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_relaxedMeasurement_r17 struct {
	StationaryMobilityEvaluation_r17      SIB2_relaxedMeasurement_r17_stationaryMobilityEvaluation_r17       `madatory`
	CellEdgeEvaluationWhileStationary_r17 *SIB2_relaxedMeasurement_r17_cellEdgeEvaluationWhileStationary_r17 `optional`
	CombineRelaxedMeasCondition2_r17      *SIB2_relaxedMeasurement_r17_combineRelaxedMeasCondition2_r17      `optional`
}

func (ie *SIB2_relaxedMeasurement_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CellEdgeEvaluationWhileStationary_r17 != nil, ie.CombineRelaxedMeasCondition2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.StationaryMobilityEvaluation_r17.Encode(w); err != nil {
		return utils.WrapError("Encode StationaryMobilityEvaluation_r17", err)
	}
	if ie.CellEdgeEvaluationWhileStationary_r17 != nil {
		if err = ie.CellEdgeEvaluationWhileStationary_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CellEdgeEvaluationWhileStationary_r17", err)
		}
	}
	if ie.CombineRelaxedMeasCondition2_r17 != nil {
		if err = ie.CombineRelaxedMeasCondition2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CombineRelaxedMeasCondition2_r17", err)
		}
	}
	return nil
}

func (ie *SIB2_relaxedMeasurement_r17) Decode(r *uper.UperReader) error {
	var err error
	var CellEdgeEvaluationWhileStationary_r17Present bool
	if CellEdgeEvaluationWhileStationary_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CombineRelaxedMeasCondition2_r17Present bool
	if CombineRelaxedMeasCondition2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.StationaryMobilityEvaluation_r17.Decode(r); err != nil {
		return utils.WrapError("Decode StationaryMobilityEvaluation_r17", err)
	}
	if CellEdgeEvaluationWhileStationary_r17Present {
		ie.CellEdgeEvaluationWhileStationary_r17 = new(SIB2_relaxedMeasurement_r17_cellEdgeEvaluationWhileStationary_r17)
		if err = ie.CellEdgeEvaluationWhileStationary_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CellEdgeEvaluationWhileStationary_r17", err)
		}
	}
	if CombineRelaxedMeasCondition2_r17Present {
		ie.CombineRelaxedMeasCondition2_r17 = new(SIB2_relaxedMeasurement_r17_combineRelaxedMeasCondition2_r17)
		if err = ie.CombineRelaxedMeasCondition2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CombineRelaxedMeasCondition2_r17", err)
		}
	}
	return nil
}
