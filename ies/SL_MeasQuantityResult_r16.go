package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SL_MeasQuantityResult_r16 struct {
	Sl_RSRP_r16 *RSRP_Range `optional`
}

func (ie *SL_MeasQuantityResult_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_RSRP_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_RSRP_r16 != nil {
		if err = ie.Sl_RSRP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_RSRP_r16", err)
		}
	}
	return nil
}

func (ie *SL_MeasQuantityResult_r16) Decode(r *aper.AperReader) error {
	var err error
	var Sl_RSRP_r16Present bool
	if Sl_RSRP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_RSRP_r16Present {
		ie.Sl_RSRP_r16 = new(RSRP_Range)
		if err = ie.Sl_RSRP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_RSRP_r16", err)
		}
	}
	return nil
}
