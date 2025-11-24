package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17 struct {
	Scs15_r17  *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs15_r17  `optional`
	Scs30_r17  *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs30_r17  `optional`
	Scs60_r17  *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs60_r17  `optional`
	Scs120_r17 *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs120_r17 `optional`
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs15_r17 != nil, ie.Scs30_r17 != nil, ie.Scs60_r17 != nil, ie.Scs120_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs15_r17 != nil {
		if err = ie.Scs15_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs15_r17", err)
		}
	}
	if ie.Scs30_r17 != nil {
		if err = ie.Scs30_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs30_r17", err)
		}
	}
	if ie.Scs60_r17 != nil {
		if err = ie.Scs60_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs60_r17", err)
		}
	}
	if ie.Scs120_r17 != nil {
		if err = ie.Scs120_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs120_r17", err)
		}
	}
	return nil
}

func (ie *PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17) Decode(r *uper.UperReader) error {
	var err error
	var Scs15_r17Present bool
	if Scs15_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs30_r17Present bool
	if Scs30_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs60_r17Present bool
	if Scs60_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs120_r17Present bool
	if Scs120_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs15_r17Present {
		ie.Scs15_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs15_r17)
		if err = ie.Scs15_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs15_r17", err)
		}
	}
	if Scs30_r17Present {
		ie.Scs30_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs30_r17)
		if err = ie.Scs30_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs30_r17", err)
		}
	}
	if Scs60_r17Present {
		ie.Scs60_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs60_r17)
		if err = ie.Scs60_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs60_r17", err)
		}
	}
	if Scs120_r17Present {
		ie.Scs120_r17 = new(PRS_ProcessingCapabilityOutsideMGinPPWperType_r17_ppw_maxNumOfDL_PRS_ResProcessedPerSlot_r17_scs120_r17)
		if err = ie.Scs120_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs120_r17", err)
		}
	}
	return nil
}
