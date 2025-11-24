package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_relaxedMeasurement_r17_cellEdgeEvaluationWhileStationary_r17 struct {
	S_SearchThresholdP2_r17 ReselectionThreshold   `madatory`
	S_SearchThresholdQ2_r17 *ReselectionThresholdQ `optional`
}

func (ie *SIB2_relaxedMeasurement_r17_cellEdgeEvaluationWhileStationary_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.S_SearchThresholdQ2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.S_SearchThresholdP2_r17.Encode(w); err != nil {
		return utils.WrapError("Encode S_SearchThresholdP2_r17", err)
	}
	if ie.S_SearchThresholdQ2_r17 != nil {
		if err = ie.S_SearchThresholdQ2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode S_SearchThresholdQ2_r17", err)
		}
	}
	return nil
}

func (ie *SIB2_relaxedMeasurement_r17_cellEdgeEvaluationWhileStationary_r17) Decode(r *uper.UperReader) error {
	var err error
	var S_SearchThresholdQ2_r17Present bool
	if S_SearchThresholdQ2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.S_SearchThresholdP2_r17.Decode(r); err != nil {
		return utils.WrapError("Decode S_SearchThresholdP2_r17", err)
	}
	if S_SearchThresholdQ2_r17Present {
		ie.S_SearchThresholdQ2_r17 = new(ReselectionThresholdQ)
		if err = ie.S_SearchThresholdQ2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode S_SearchThresholdQ2_r17", err)
		}
	}
	return nil
}
