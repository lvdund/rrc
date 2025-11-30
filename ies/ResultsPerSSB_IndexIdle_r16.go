package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type ResultsPerSSB_IndexIdle_r16 struct {
	Ssb_Index_r16   SSB_Index                                    `madatory`
	Ssb_Results_r16 *ResultsPerSSB_IndexIdle_r16_ssb_Results_r16 `optional`
}

func (ie *ResultsPerSSB_IndexIdle_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ssb_Results_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Ssb_Index_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Ssb_Index_r16", err)
	}
	if ie.Ssb_Results_r16 != nil {
		if err = ie.Ssb_Results_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Ssb_Results_r16", err)
		}
	}
	return nil
}

func (ie *ResultsPerSSB_IndexIdle_r16) Decode(r *aper.AperReader) error {
	var err error
	var Ssb_Results_r16Present bool
	if Ssb_Results_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Ssb_Index_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Ssb_Index_r16", err)
	}
	if Ssb_Results_r16Present {
		ie.Ssb_Results_r16 = new(ResultsPerSSB_IndexIdle_r16_ssb_Results_r16)
		if err = ie.Ssb_Results_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Ssb_Results_r16", err)
		}
	}
	return nil
}
