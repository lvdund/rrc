package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_RemoteUE_Config_r17 struct {
	ThreshHighRemote_r17     *RSRP_Range               `optional`
	HystMaxRemote_r17        *Hysteresis               `optional`
	Sl_ReselectionConfig_r17 *SL_ReselectionConfig_r17 `optional`
}

func (ie *SL_RemoteUE_Config_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ThreshHighRemote_r17 != nil, ie.HystMaxRemote_r17 != nil, ie.Sl_ReselectionConfig_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.ThreshHighRemote_r17 != nil {
		if err = ie.ThreshHighRemote_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ThreshHighRemote_r17", err)
		}
	}
	if ie.HystMaxRemote_r17 != nil {
		if err = ie.HystMaxRemote_r17.Encode(w); err != nil {
			return utils.WrapError("Encode HystMaxRemote_r17", err)
		}
	}
	if ie.Sl_ReselectionConfig_r17 != nil {
		if err = ie.Sl_ReselectionConfig_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_ReselectionConfig_r17", err)
		}
	}
	return nil
}

func (ie *SL_RemoteUE_Config_r17) Decode(r *aper.AperReader) error {
	var err error
	var ThreshHighRemote_r17Present bool
	if ThreshHighRemote_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var HystMaxRemote_r17Present bool
	if HystMaxRemote_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_ReselectionConfig_r17Present bool
	if Sl_ReselectionConfig_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if ThreshHighRemote_r17Present {
		ie.ThreshHighRemote_r17 = new(RSRP_Range)
		if err = ie.ThreshHighRemote_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ThreshHighRemote_r17", err)
		}
	}
	if HystMaxRemote_r17Present {
		ie.HystMaxRemote_r17 = new(Hysteresis)
		if err = ie.HystMaxRemote_r17.Decode(r); err != nil {
			return utils.WrapError("Decode HystMaxRemote_r17", err)
		}
	}
	if Sl_ReselectionConfig_r17Present {
		ie.Sl_ReselectionConfig_r17 = new(SL_ReselectionConfig_r17)
		if err = ie.Sl_ReselectionConfig_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_ReselectionConfig_r17", err)
		}
	}
	return nil
}
