package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasIdleCarrierEUTRA_r16_qualityThresholdEUTRA_r16 struct {
	IdleRSRP_Threshold_EUTRA_r16 *RSRP_RangeEUTRA     `optional`
	IdleRSRQ_Threshold_EUTRA_r16 *RSRQ_RangeEUTRA_r16 `optional`
}

func (ie *MeasIdleCarrierEUTRA_r16_qualityThresholdEUTRA_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.IdleRSRP_Threshold_EUTRA_r16 != nil, ie.IdleRSRQ_Threshold_EUTRA_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IdleRSRP_Threshold_EUTRA_r16 != nil {
		if err = ie.IdleRSRP_Threshold_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IdleRSRP_Threshold_EUTRA_r16", err)
		}
	}
	if ie.IdleRSRQ_Threshold_EUTRA_r16 != nil {
		if err = ie.IdleRSRQ_Threshold_EUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IdleRSRQ_Threshold_EUTRA_r16", err)
		}
	}
	return nil
}

func (ie *MeasIdleCarrierEUTRA_r16_qualityThresholdEUTRA_r16) Decode(r *aper.AperReader) error {
	var err error
	var IdleRSRP_Threshold_EUTRA_r16Present bool
	if IdleRSRP_Threshold_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IdleRSRQ_Threshold_EUTRA_r16Present bool
	if IdleRSRQ_Threshold_EUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if IdleRSRP_Threshold_EUTRA_r16Present {
		ie.IdleRSRP_Threshold_EUTRA_r16 = new(RSRP_RangeEUTRA)
		if err = ie.IdleRSRP_Threshold_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IdleRSRP_Threshold_EUTRA_r16", err)
		}
	}
	if IdleRSRQ_Threshold_EUTRA_r16Present {
		ie.IdleRSRQ_Threshold_EUTRA_r16 = new(RSRQ_RangeEUTRA_r16)
		if err = ie.IdleRSRQ_Threshold_EUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IdleRSRQ_Threshold_EUTRA_r16", err)
		}
	}
	return nil
}
