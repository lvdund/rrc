package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SearchSpaceSwitchConfig_r17 struct {
	SearchSpaceSwitchTimer_r17 *SCS_SpecificDuration_r17 `optional`
	SearchSpaceSwitchDelay_r17 *int64                    `lb:10,ub:52,optional`
}

func (ie *SearchSpaceSwitchConfig_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.SearchSpaceSwitchTimer_r17 != nil, ie.SearchSpaceSwitchDelay_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SearchSpaceSwitchTimer_r17 != nil {
		if err = ie.SearchSpaceSwitchTimer_r17.Encode(w); err != nil {
			return utils.WrapError("Encode SearchSpaceSwitchTimer_r17", err)
		}
	}
	if ie.SearchSpaceSwitchDelay_r17 != nil {
		if err = w.WriteInteger(*ie.SearchSpaceSwitchDelay_r17, &aper.Constraint{Lb: 10, Ub: 52}, false); err != nil {
			return utils.WrapError("Encode SearchSpaceSwitchDelay_r17", err)
		}
	}
	return nil
}

func (ie *SearchSpaceSwitchConfig_r17) Decode(r *aper.AperReader) error {
	var err error
	var SearchSpaceSwitchTimer_r17Present bool
	if SearchSpaceSwitchTimer_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SearchSpaceSwitchDelay_r17Present bool
	if SearchSpaceSwitchDelay_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if SearchSpaceSwitchTimer_r17Present {
		ie.SearchSpaceSwitchTimer_r17 = new(SCS_SpecificDuration_r17)
		if err = ie.SearchSpaceSwitchTimer_r17.Decode(r); err != nil {
			return utils.WrapError("Decode SearchSpaceSwitchTimer_r17", err)
		}
	}
	if SearchSpaceSwitchDelay_r17Present {
		var tmp_int_SearchSpaceSwitchDelay_r17 int64
		if tmp_int_SearchSpaceSwitchDelay_r17, err = r.ReadInteger(&aper.Constraint{Lb: 10, Ub: 52}, false); err != nil {
			return utils.WrapError("Decode SearchSpaceSwitchDelay_r17", err)
		}
		ie.SearchSpaceSwitchDelay_r17 = &tmp_int_SearchSpaceSwitchDelay_r17
	}
	return nil
}
