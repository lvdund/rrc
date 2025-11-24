package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MinSchedulingOffsetPreference_r16_preferredK0_r16 struct {
	PreferredK0_SCS_15kHz_r16  *MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_15kHz_r16  `optional`
	PreferredK0_SCS_30kHz_r16  *MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_30kHz_r16  `optional`
	PreferredK0_SCS_60kHz_r16  *MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_60kHz_r16  `optional`
	PreferredK0_SCS_120kHz_r16 *MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_120kHz_r16 `optional`
}

func (ie *MinSchedulingOffsetPreference_r16_preferredK0_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.PreferredK0_SCS_15kHz_r16 != nil, ie.PreferredK0_SCS_30kHz_r16 != nil, ie.PreferredK0_SCS_60kHz_r16 != nil, ie.PreferredK0_SCS_120kHz_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PreferredK0_SCS_15kHz_r16 != nil {
		if err = ie.PreferredK0_SCS_15kHz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredK0_SCS_15kHz_r16", err)
		}
	}
	if ie.PreferredK0_SCS_30kHz_r16 != nil {
		if err = ie.PreferredK0_SCS_30kHz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredK0_SCS_30kHz_r16", err)
		}
	}
	if ie.PreferredK0_SCS_60kHz_r16 != nil {
		if err = ie.PreferredK0_SCS_60kHz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredK0_SCS_60kHz_r16", err)
		}
	}
	if ie.PreferredK0_SCS_120kHz_r16 != nil {
		if err = ie.PreferredK0_SCS_120kHz_r16.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredK0_SCS_120kHz_r16", err)
		}
	}
	return nil
}

func (ie *MinSchedulingOffsetPreference_r16_preferredK0_r16) Decode(r *uper.UperReader) error {
	var err error
	var PreferredK0_SCS_15kHz_r16Present bool
	if PreferredK0_SCS_15kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PreferredK0_SCS_30kHz_r16Present bool
	if PreferredK0_SCS_30kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PreferredK0_SCS_60kHz_r16Present bool
	if PreferredK0_SCS_60kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PreferredK0_SCS_120kHz_r16Present bool
	if PreferredK0_SCS_120kHz_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if PreferredK0_SCS_15kHz_r16Present {
		ie.PreferredK0_SCS_15kHz_r16 = new(MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_15kHz_r16)
		if err = ie.PreferredK0_SCS_15kHz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredK0_SCS_15kHz_r16", err)
		}
	}
	if PreferredK0_SCS_30kHz_r16Present {
		ie.PreferredK0_SCS_30kHz_r16 = new(MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_30kHz_r16)
		if err = ie.PreferredK0_SCS_30kHz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredK0_SCS_30kHz_r16", err)
		}
	}
	if PreferredK0_SCS_60kHz_r16Present {
		ie.PreferredK0_SCS_60kHz_r16 = new(MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_60kHz_r16)
		if err = ie.PreferredK0_SCS_60kHz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredK0_SCS_60kHz_r16", err)
		}
	}
	if PreferredK0_SCS_120kHz_r16Present {
		ie.PreferredK0_SCS_120kHz_r16 = new(MinSchedulingOffsetPreference_r16_preferredK0_r16_preferredK0_SCS_120kHz_r16)
		if err = ie.PreferredK0_SCS_120kHz_r16.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredK0_SCS_120kHz_r16", err)
		}
	}
	return nil
}
