package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type LoggedMeasurementConfiguration_v1700_IEs struct {
	SigLoggedMeasType_r17   *LoggedMeasurementConfiguration_v1700_IEs_sigLoggedMeasType_r17   `optional`
	EarlyMeasIndication_r17 *LoggedMeasurementConfiguration_v1700_IEs_earlyMeasIndication_r17 `optional`
	AreaConfiguration_v1700 *AreaConfiguration_v1700                                          `optional`
	NonCriticalExtension    interface{}                                                       `optional`
}

func (ie *LoggedMeasurementConfiguration_v1700_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SigLoggedMeasType_r17 != nil, ie.EarlyMeasIndication_r17 != nil, ie.AreaConfiguration_v1700 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SigLoggedMeasType_r17 != nil {
		if err = ie.SigLoggedMeasType_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SigLoggedMeasType_r17", err)
		}
	}
	if ie.EarlyMeasIndication_r17 != nil {
		if err = ie.EarlyMeasIndication_r17.Encode(w); err != nil {
			return utils.WrapError("Encode EarlyMeasIndication_r17", err)
		}
	}
	if ie.AreaConfiguration_v1700 != nil {
		if err = ie.AreaConfiguration_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode AreaConfiguration_v1700", err)
		}
	}
	return nil
}

func (ie *LoggedMeasurementConfiguration_v1700_IEs) Decode(r *aper.AperReader) error {
	var err error
	var SigLoggedMeasType_r17Present bool
	if SigLoggedMeasType_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var EarlyMeasIndication_r17Present bool
	if EarlyMeasIndication_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AreaConfiguration_v1700Present bool
	if AreaConfiguration_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SigLoggedMeasType_r17Present {
		ie.SigLoggedMeasType_r17 = new(LoggedMeasurementConfiguration_v1700_IEs_sigLoggedMeasType_r17)
		if err = ie.SigLoggedMeasType_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SigLoggedMeasType_r17", err)
		}
	}
	if EarlyMeasIndication_r17Present {
		ie.EarlyMeasIndication_r17 = new(LoggedMeasurementConfiguration_v1700_IEs_earlyMeasIndication_r17)
		if err = ie.EarlyMeasIndication_r17.Decode(r); err != nil {
			return utils.WrapError("Decode EarlyMeasIndication_r17", err)
		}
	}
	if AreaConfiguration_v1700Present {
		ie.AreaConfiguration_v1700 = new(AreaConfiguration_v1700)
		if err = ie.AreaConfiguration_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode AreaConfiguration_v1700", err)
		}
	}
	return nil
}
