package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB2_cellReselectionServingFreqInfo struct {
	S_NonIntraSearchP          *ReselectionThreshold       `optional`
	S_NonIntraSearchQ          *ReselectionThresholdQ      `optional`
	ThreshServingLowP          ReselectionThreshold        `madatory`
	ThreshServingLowQ          *ReselectionThresholdQ      `optional`
	CellReselectionPriority    CellReselectionPriority     `madatory`
	CellReselectionSubPriority *CellReselectionSubPriority `optional`
}

func (ie *SIB2_cellReselectionServingFreqInfo) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.S_NonIntraSearchP != nil, ie.S_NonIntraSearchQ != nil, ie.ThreshServingLowQ != nil, ie.CellReselectionSubPriority != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.S_NonIntraSearchP != nil {
		if err = ie.S_NonIntraSearchP.Encode(w); err != nil {
			return utils.WrapError("Encode S_NonIntraSearchP", err)
		}
	}
	if ie.S_NonIntraSearchQ != nil {
		if err = ie.S_NonIntraSearchQ.Encode(w); err != nil {
			return utils.WrapError("Encode S_NonIntraSearchQ", err)
		}
	}
	if err = ie.ThreshServingLowP.Encode(w); err != nil {
		return utils.WrapError("Encode ThreshServingLowP", err)
	}
	if ie.ThreshServingLowQ != nil {
		if err = ie.ThreshServingLowQ.Encode(w); err != nil {
			return utils.WrapError("Encode ThreshServingLowQ", err)
		}
	}
	if err = ie.CellReselectionPriority.Encode(w); err != nil {
		return utils.WrapError("Encode CellReselectionPriority", err)
	}
	if ie.CellReselectionSubPriority != nil {
		if err = ie.CellReselectionSubPriority.Encode(w); err != nil {
			return utils.WrapError("Encode CellReselectionSubPriority", err)
		}
	}
	return nil
}

func (ie *SIB2_cellReselectionServingFreqInfo) Decode(r *aper.AperReader) error {
	var err error
	var S_NonIntraSearchPPresent bool
	if S_NonIntraSearchPPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var S_NonIntraSearchQPresent bool
	if S_NonIntraSearchQPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var ThreshServingLowQPresent bool
	if ThreshServingLowQPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var CellReselectionSubPriorityPresent bool
	if CellReselectionSubPriorityPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if S_NonIntraSearchPPresent {
		ie.S_NonIntraSearchP = new(ReselectionThreshold)
		if err = ie.S_NonIntraSearchP.Decode(r); err != nil {
			return utils.WrapError("Decode S_NonIntraSearchP", err)
		}
	}
	if S_NonIntraSearchQPresent {
		ie.S_NonIntraSearchQ = new(ReselectionThresholdQ)
		if err = ie.S_NonIntraSearchQ.Decode(r); err != nil {
			return utils.WrapError("Decode S_NonIntraSearchQ", err)
		}
	}
	if err = ie.ThreshServingLowP.Decode(r); err != nil {
		return utils.WrapError("Decode ThreshServingLowP", err)
	}
	if ThreshServingLowQPresent {
		ie.ThreshServingLowQ = new(ReselectionThresholdQ)
		if err = ie.ThreshServingLowQ.Decode(r); err != nil {
			return utils.WrapError("Decode ThreshServingLowQ", err)
		}
	}
	if err = ie.CellReselectionPriority.Decode(r); err != nil {
		return utils.WrapError("Decode CellReselectionPriority", err)
	}
	if CellReselectionSubPriorityPresent {
		ie.CellReselectionSubPriority = new(CellReselectionSubPriority)
		if err = ie.CellReselectionSubPriority.Decode(r); err != nil {
			return utils.WrapError("Decode CellReselectionSubPriority", err)
		}
	}
	return nil
}
