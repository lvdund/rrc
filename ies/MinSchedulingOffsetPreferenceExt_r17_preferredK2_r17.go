package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MinSchedulingOffsetPreferenceExt_r17_preferredK2_r17 struct {
	PreferredK2_SCS_480kHz_r17 *MinSchedulingOffsetPreferenceExt_r17_preferredK2_r17_preferredK2_SCS_480kHz_r17 `optional`
	PreferredK2_SCS_960kHz_r17 *MinSchedulingOffsetPreferenceExt_r17_preferredK2_r17_preferredK2_SCS_960kHz_r17 `optional`
}

func (ie *MinSchedulingOffsetPreferenceExt_r17_preferredK2_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.PreferredK2_SCS_480kHz_r17 != nil, ie.PreferredK2_SCS_960kHz_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.PreferredK2_SCS_480kHz_r17 != nil {
		if err = ie.PreferredK2_SCS_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredK2_SCS_480kHz_r17", err)
		}
	}
	if ie.PreferredK2_SCS_960kHz_r17 != nil {
		if err = ie.PreferredK2_SCS_960kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode PreferredK2_SCS_960kHz_r17", err)
		}
	}
	return nil
}

func (ie *MinSchedulingOffsetPreferenceExt_r17_preferredK2_r17) Decode(r *uper.UperReader) error {
	var err error
	var PreferredK2_SCS_480kHz_r17Present bool
	if PreferredK2_SCS_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var PreferredK2_SCS_960kHz_r17Present bool
	if PreferredK2_SCS_960kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if PreferredK2_SCS_480kHz_r17Present {
		ie.PreferredK2_SCS_480kHz_r17 = new(MinSchedulingOffsetPreferenceExt_r17_preferredK2_r17_preferredK2_SCS_480kHz_r17)
		if err = ie.PreferredK2_SCS_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredK2_SCS_480kHz_r17", err)
		}
	}
	if PreferredK2_SCS_960kHz_r17Present {
		ie.PreferredK2_SCS_960kHz_r17 = new(MinSchedulingOffsetPreferenceExt_r17_preferredK2_r17_preferredK2_SCS_960kHz_r17)
		if err = ie.PreferredK2_SCS_960kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode PreferredK2_SCS_960kHz_r17", err)
		}
	}
	return nil
}
