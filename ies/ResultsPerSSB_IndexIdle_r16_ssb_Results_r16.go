package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ResultsPerSSB_IndexIdle_r16_ssb_Results_r16 struct {
	Ssb_RSRP_Result_r16 *RSRP_Range `optional`
	Ssb_RSRQ_Result_r16 *RSRQ_Range `optional`
}

func (ie *ResultsPerSSB_IndexIdle_r16_ssb_Results_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ssb_RSRP_Result_r16 != nil, ie.Ssb_RSRQ_Result_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ssb_RSRP_Result_r16 != nil {
		if err = ie.Ssb_RSRP_Result_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_RSRP_Result_r16", err)
		}
	}
	if ie.Ssb_RSRQ_Result_r16 != nil {
		if err = ie.Ssb_RSRQ_Result_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_RSRQ_Result_r16", err)
		}
	}
	return nil
}

func (ie *ResultsPerSSB_IndexIdle_r16_ssb_Results_r16) Decode(r *aper.AperReader) error {
	var err error
	var Ssb_RSRP_Result_r16Present bool
	if Ssb_RSRP_Result_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ssb_RSRQ_Result_r16Present bool
	if Ssb_RSRQ_Result_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ssb_RSRP_Result_r16Present {
		ie.Ssb_RSRP_Result_r16 = new(RSRP_Range)
		if err = ie.Ssb_RSRP_Result_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_RSRP_Result_r16", err)
		}
	}
	if Ssb_RSRQ_Result_r16Present {
		ie.Ssb_RSRQ_Result_r16 = new(RSRQ_Range)
		if err = ie.Ssb_RSRQ_Result_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_RSRQ_Result_r16", err)
		}
	}
	return nil
}
