package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdleCarrierNR_r16_qualityThreshold_r16 struct {
	IdleRSRP_Threshold_NR_r16 *RSRP_Range `optional`
	IdleRSRQ_Threshold_NR_r16 *RSRQ_Range `optional`
}

func (ie *MeasIdleCarrierNR_r16_qualityThreshold_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.IdleRSRP_Threshold_NR_r16 != nil, ie.IdleRSRQ_Threshold_NR_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IdleRSRP_Threshold_NR_r16 != nil {
		if err = ie.IdleRSRP_Threshold_NR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IdleRSRP_Threshold_NR_r16", err)
		}
	}
	if ie.IdleRSRQ_Threshold_NR_r16 != nil {
		if err = ie.IdleRSRQ_Threshold_NR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IdleRSRQ_Threshold_NR_r16", err)
		}
	}
	return nil
}

func (ie *MeasIdleCarrierNR_r16_qualityThreshold_r16) Decode(r *uper.UperReader) error {
	var err error
	var IdleRSRP_Threshold_NR_r16Present bool
	if IdleRSRP_Threshold_NR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IdleRSRQ_Threshold_NR_r16Present bool
	if IdleRSRQ_Threshold_NR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if IdleRSRP_Threshold_NR_r16Present {
		ie.IdleRSRP_Threshold_NR_r16 = new(RSRP_Range)
		if err = ie.IdleRSRP_Threshold_NR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IdleRSRP_Threshold_NR_r16", err)
		}
	}
	if IdleRSRQ_Threshold_NR_r16Present {
		ie.IdleRSRQ_Threshold_NR_r16 = new(RSRQ_Range)
		if err = ie.IdleRSRQ_Threshold_NR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IdleRSRQ_Threshold_NR_r16", err)
		}
	}
	return nil
}
