package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Uplink_powerControl_r17 struct {
	Ul_powercontrolId_r17  Uplink_powerControlId_r17 `madatory`
	P0AlphaSetforPUSCH_r17 *P0AlphaSet_r17           `optional`
	P0AlphaSetforPUCCH_r17 *P0AlphaSet_r17           `optional`
	P0AlphaSetforSRS_r17   *P0AlphaSet_r17           `optional`
}

func (ie *Uplink_powerControl_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.P0AlphaSetforPUSCH_r17 != nil, ie.P0AlphaSetforPUCCH_r17 != nil, ie.P0AlphaSetforSRS_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Ul_powercontrolId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Ul_powercontrolId_r17", err)
	}
	if ie.P0AlphaSetforPUSCH_r17 != nil {
		if err = ie.P0AlphaSetforPUSCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode P0AlphaSetforPUSCH_r17", err)
		}
	}
	if ie.P0AlphaSetforPUCCH_r17 != nil {
		if err = ie.P0AlphaSetforPUCCH_r17.Encode(w); err != nil {
			return utils.WrapError("Encode P0AlphaSetforPUCCH_r17", err)
		}
	}
	if ie.P0AlphaSetforSRS_r17 != nil {
		if err = ie.P0AlphaSetforSRS_r17.Encode(w); err != nil {
			return utils.WrapError("Encode P0AlphaSetforSRS_r17", err)
		}
	}
	return nil
}

func (ie *Uplink_powerControl_r17) Decode(r *aper.AperReader) error {
	var err error
	var P0AlphaSetforPUSCH_r17Present bool
	if P0AlphaSetforPUSCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var P0AlphaSetforPUCCH_r17Present bool
	if P0AlphaSetforPUCCH_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var P0AlphaSetforSRS_r17Present bool
	if P0AlphaSetforSRS_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Ul_powercontrolId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Ul_powercontrolId_r17", err)
	}
	if P0AlphaSetforPUSCH_r17Present {
		ie.P0AlphaSetforPUSCH_r17 = new(P0AlphaSet_r17)
		if err = ie.P0AlphaSetforPUSCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode P0AlphaSetforPUSCH_r17", err)
		}
	}
	if P0AlphaSetforPUCCH_r17Present {
		ie.P0AlphaSetforPUCCH_r17 = new(P0AlphaSet_r17)
		if err = ie.P0AlphaSetforPUCCH_r17.Decode(r); err != nil {
			return utils.WrapError("Decode P0AlphaSetforPUCCH_r17", err)
		}
	}
	if P0AlphaSetforSRS_r17Present {
		ie.P0AlphaSetforSRS_r17 = new(P0AlphaSet_r17)
		if err = ie.P0AlphaSetforSRS_r17.Decode(r); err != nil {
			return utils.WrapError("Decode P0AlphaSetforSRS_r17", err)
		}
	}
	return nil
}
