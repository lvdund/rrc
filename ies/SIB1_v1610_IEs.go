package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB1_v1610_IEs struct {
	IdleModeMeasurementsEUTRA_r16 *SIB1_v1610_IEs_idleModeMeasurementsEUTRA_r16 `optional`
	IdleModeMeasurementsNR_r16    *SIB1_v1610_IEs_idleModeMeasurementsNR_r16    `optional`
	PosSI_SchedulingInfo_r16      *PosSI_SchedulingInfo_r16                     `optional`
	NonCriticalExtension          *SIB1_v1630_IEs                               `optional`
}

func (ie *SIB1_v1610_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.IdleModeMeasurementsEUTRA_r16 != nil, ie.IdleModeMeasurementsNR_r16 != nil, ie.PosSI_SchedulingInfo_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.IdleModeMeasurementsEUTRA_r16 != nil {
		if err = ie.IdleModeMeasurementsEUTRA_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IdleModeMeasurementsEUTRA_r16", err)
		}
	}
	if ie.IdleModeMeasurementsNR_r16 != nil {
		if err = ie.IdleModeMeasurementsNR_r16.Encode(w); err != nil {
			return utils.WrapError("Encode IdleModeMeasurementsNR_r16", err)
		}
	}
	if ie.PosSI_SchedulingInfo_r16 != nil {
		if err = ie.PosSI_SchedulingInfo_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PosSI_SchedulingInfo_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB1_v1610_IEs) Decode(r *aper.AperReader) error {
	var err error
	var IdleModeMeasurementsEUTRA_r16Present bool
	if IdleModeMeasurementsEUTRA_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var IdleModeMeasurementsNR_r16Present bool
	if IdleModeMeasurementsNR_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PosSI_SchedulingInfo_r16Present bool
	if PosSI_SchedulingInfo_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if IdleModeMeasurementsEUTRA_r16Present {
		ie.IdleModeMeasurementsEUTRA_r16 = new(SIB1_v1610_IEs_idleModeMeasurementsEUTRA_r16)
		if err = ie.IdleModeMeasurementsEUTRA_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IdleModeMeasurementsEUTRA_r16", err)
		}
	}
	if IdleModeMeasurementsNR_r16Present {
		ie.IdleModeMeasurementsNR_r16 = new(SIB1_v1610_IEs_idleModeMeasurementsNR_r16)
		if err = ie.IdleModeMeasurementsNR_r16.Decode(r); err != nil {
			return utils.WrapError("Decode IdleModeMeasurementsNR_r16", err)
		}
	}
	if PosSI_SchedulingInfo_r16Present {
		ie.PosSI_SchedulingInfo_r16 = new(PosSI_SchedulingInfo_r16)
		if err = ie.PosSI_SchedulingInfo_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PosSI_SchedulingInfo_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(SIB1_v1630_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
