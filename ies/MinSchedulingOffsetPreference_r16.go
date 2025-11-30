package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MinSchedulingOffsetPreference_r16 struct {
	PreferredK0_r16 *MinSchedulingOffsetPreference_r16_preferredK0_r16 `optional`
	PreferredK2_r16 *MinSchedulingOffsetPreference_r16_preferredK2_r16 `optional`
}

func (ie *MinSchedulingOffsetPreference_r16) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.PreferredK0_r16 != nil, ie.PreferredK2_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PreferredK0_r16 != nil {
		if err = ie.PreferredK0_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredK0_r16", err)
		}
	}
	if ie.PreferredK2_r16 != nil {
		if err = ie.PreferredK2_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredK2_r16", err)
		}
	}
	return nil
}

func (ie *MinSchedulingOffsetPreference_r16) Decode(r *aper.AperReader) error {
	var err error
	var PreferredK0_r16Present bool
	if PreferredK0_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PreferredK2_r16Present bool
	if PreferredK2_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if PreferredK0_r16Present {
		ie.PreferredK0_r16 = new(MinSchedulingOffsetPreference_r16_preferredK0_r16)
		if err = ie.PreferredK0_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredK0_r16", err)
		}
	}
	if PreferredK2_r16Present {
		ie.PreferredK2_r16 = new(MinSchedulingOffsetPreference_r16_preferredK2_r16)
		if err = ie.PreferredK2_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredK2_r16", err)
		}
	}
	return nil
}
