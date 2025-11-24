package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CG_ConfigInfo_v1700_IEs struct {
	CandidateCellListCPC_r17                  *CandidateCellListCPC_r17                                          `optional`
	TwoPHRModeMCG_r17                         *CG_ConfigInfo_v1700_IEs_twoPHRModeMCG_r17                         `optional`
	LowMobilityEvaluationConnectedInPCell_r17 *CG_ConfigInfo_v1700_IEs_lowMobilityEvaluationConnectedInPCell_r17 `optional`
	NonCriticalExtension                      *CG_ConfigInfo_v1730_IEs                                           `optional`
}

func (ie *CG_ConfigInfo_v1700_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CandidateCellListCPC_r17 != nil, ie.TwoPHRModeMCG_r17 != nil, ie.LowMobilityEvaluationConnectedInPCell_r17 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CandidateCellListCPC_r17 != nil {
		if err = ie.CandidateCellListCPC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CandidateCellListCPC_r17", err)
		}
	}
	if ie.TwoPHRModeMCG_r17 != nil {
		if err = ie.TwoPHRModeMCG_r17.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPHRModeMCG_r17", err)
		}
	}
	if ie.LowMobilityEvaluationConnectedInPCell_r17 != nil {
		if err = ie.LowMobilityEvaluationConnectedInPCell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode LowMobilityEvaluationConnectedInPCell_r17", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *CG_ConfigInfo_v1700_IEs) Decode(r *uper.UperReader) error {
	var err error
	var CandidateCellListCPC_r17Present bool
	if CandidateCellListCPC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPHRModeMCG_r17Present bool
	if TwoPHRModeMCG_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LowMobilityEvaluationConnectedInPCell_r17Present bool
	if LowMobilityEvaluationConnectedInPCell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CandidateCellListCPC_r17Present {
		ie.CandidateCellListCPC_r17 = new(CandidateCellListCPC_r17)
		if err = ie.CandidateCellListCPC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CandidateCellListCPC_r17", err)
		}
	}
	if TwoPHRModeMCG_r17Present {
		ie.TwoPHRModeMCG_r17 = new(CG_ConfigInfo_v1700_IEs_twoPHRModeMCG_r17)
		if err = ie.TwoPHRModeMCG_r17.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPHRModeMCG_r17", err)
		}
	}
	if LowMobilityEvaluationConnectedInPCell_r17Present {
		ie.LowMobilityEvaluationConnectedInPCell_r17 = new(CG_ConfigInfo_v1700_IEs_lowMobilityEvaluationConnectedInPCell_r17)
		if err = ie.LowMobilityEvaluationConnectedInPCell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode LowMobilityEvaluationConnectedInPCell_r17", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(CG_ConfigInfo_v1730_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
