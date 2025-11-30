package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MinSchedulingOffsetPreferenceExt_r17 struct {
	PreferredK0_r17 *MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17 `optional`
	PreferredK2_r17 *MinSchedulingOffsetPreferenceExt_r17_preferredK2_r17 `optional`
}

func (ie *MinSchedulingOffsetPreferenceExt_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PreferredK0_r17 != nil, ie.PreferredK2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PreferredK0_r17 != nil {
		if err = ie.PreferredK0_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredK0_r17", err)
		}
	}
	if ie.PreferredK2_r17 != nil {
		if err = ie.PreferredK2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredK2_r17", err)
		}
	}
	return nil
}

func (ie *MinSchedulingOffsetPreferenceExt_r17) Decode(r *aper.AperReader) error {
	var err error
	var PreferredK0_r17Present bool
	if PreferredK0_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PreferredK2_r17Present bool
	if PreferredK2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if PreferredK0_r17Present {
		ie.PreferredK0_r17 = new(MinSchedulingOffsetPreferenceExt_r17_preferredK0_r17)
		if err = ie.PreferredK0_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredK0_r17", err)
		}
	}
	if PreferredK2_r17Present {
		ie.PreferredK2_r17 = new(MinSchedulingOffsetPreferenceExt_r17_preferredK2_r17)
		if err = ie.PreferredK2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredK2_r17", err)
		}
	}
	return nil
}
