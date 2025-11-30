package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_BlindDetectionCA_Mixed_r17 struct {
	Pdcch_BlindDetectionCA1_r17 *int64 `lb:1,ub:15,optional`
	Pdcch_BlindDetectionCA2_r17 *int64 `lb:1,ub:15,optional`
}

func (ie *PDCCH_BlindDetectionCA_Mixed_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Pdcch_BlindDetectionCA1_r17 != nil, ie.Pdcch_BlindDetectionCA2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Pdcch_BlindDetectionCA1_r17 != nil {
		if err = w.WriteInteger(*ie.Pdcch_BlindDetectionCA1_r17, &aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionCA1_r17", err)
		}
	}
	if ie.Pdcch_BlindDetectionCA2_r17 != nil {
		if err = w.WriteInteger(*ie.Pdcch_BlindDetectionCA2_r17, &aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Encode Pdcch_BlindDetectionCA2_r17", err)
		}
	}
	return nil
}

func (ie *PDCCH_BlindDetectionCA_Mixed_r17) Decode(r *aper.AperReader) error {
	var err error
	var Pdcch_BlindDetectionCA1_r17Present bool
	if Pdcch_BlindDetectionCA1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdcch_BlindDetectionCA2_r17Present bool
	if Pdcch_BlindDetectionCA2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Pdcch_BlindDetectionCA1_r17Present {
		var tmp_int_Pdcch_BlindDetectionCA1_r17 int64
		if tmp_int_Pdcch_BlindDetectionCA1_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionCA1_r17", err)
		}
		ie.Pdcch_BlindDetectionCA1_r17 = &tmp_int_Pdcch_BlindDetectionCA1_r17
	}
	if Pdcch_BlindDetectionCA2_r17Present {
		var tmp_int_Pdcch_BlindDetectionCA2_r17 int64
		if tmp_int_Pdcch_BlindDetectionCA2_r17, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 15}, false); err != nil {
			return utils.WrapError("Decode Pdcch_BlindDetectionCA2_r17", err)
		}
		ie.Pdcch_BlindDetectionCA2_r17 = &tmp_int_Pdcch_BlindDetectionCA2_r17
	}
	return nil
}
