package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PRS_ProcessingCapabilityOutsideMGinPPWperType_r17 struct {
	PrsProcessingType_r17                      PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_prsProcessingType_r17                       `madatory`
	Ppw_dl_PRS_BufferType_r17                  PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_dl_PRS_BufferType_r17                   `madatory`
	Ppw_durationOfPRS_Processing_r17           *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17           `optional`
	Ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17 *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17 `optional`
	Ppw_maxNumOfDL_Bandwidth_r17               *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_Bandwidth_r17               `optional,ext`
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ppw_durationOfPRS_Processing_r17 != nil, ie.Ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PrsProcessingType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PrsProcessingType_r17", err)
	}
	if err = ie.Ppw_dl_PRS_BufferType_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Ppw_dl_PRS_BufferType_r17", err)
	}
	if ie.Ppw_durationOfPRS_Processing_r17 != nil {
		if err = ie.Ppw_durationOfPRS_Processing_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ppw_durationOfPRS_Processing_r17", err)
		}
	}
	if ie.Ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17 != nil {
		if err = ie.Ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17", err)
		}
	}
	return nil
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17) Decode(r *aper.AperReader) error {
	var err error
	var Ppw_durationOfPRS_Processing_r17Present bool
	if Ppw_durationOfPRS_Processing_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17Present bool
	if Ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PrsProcessingType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PrsProcessingType_r17", err)
	}
	if err = ie.Ppw_dl_PRS_BufferType_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Ppw_dl_PRS_BufferType_r17", err)
	}
	if Ppw_durationOfPRS_Processing_r17Present {
		ie.Ppw_durationOfPRS_Processing_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_durationOfPRS_Processing_r17)
		if err = ie.Ppw_durationOfPRS_Processing_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ppw_durationOfPRS_Processing_r17", err)
		}
	}
	if Ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17Present {
		ie.Ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17)
		if err = ie.Ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17", err)
		}
	}
	return nil
}
