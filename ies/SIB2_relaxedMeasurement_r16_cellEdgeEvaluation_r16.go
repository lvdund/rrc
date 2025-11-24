package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_relaxedMeasurement_r16_cellEdgeEvaluation_r16 struct {
	S_SearchThresholdP_r16 ReselectionThreshold   `madatory`
	S_SearchThresholdQ_r16 *ReselectionThresholdQ `optional`
}

func (ie *SIB2_relaxedMeasurement_r16_cellEdgeEvaluation_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.S_SearchThresholdQ_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.S_SearchThresholdP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode S_SearchThresholdP_r16", err)
	}
	if ie.S_SearchThresholdQ_r16 != nil {
		if err = ie.S_SearchThresholdQ_r16.Encode(w); err != nil {
			return utils.WrapError("Encode S_SearchThresholdQ_r16", err)
		}
	}
	return nil
}

func (ie *SIB2_relaxedMeasurement_r16_cellEdgeEvaluation_r16) Decode(r *uper.UperReader) error {
	var err error
	var S_SearchThresholdQ_r16Present bool
	if S_SearchThresholdQ_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.S_SearchThresholdP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode S_SearchThresholdP_r16", err)
	}
	if S_SearchThresholdQ_r16Present {
		ie.S_SearchThresholdQ_r16 = new(ReselectionThresholdQ)
		if err = ie.S_SearchThresholdQ_r16.Decode(r); err != nil {
			return utils.WrapError("Decode S_SearchThresholdQ_r16", err)
		}
	}
	return nil
}
