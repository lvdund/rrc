package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SL_InterUE_CoordinationConfig_r17 struct {
	Sl_InterUE_CoordinationScheme1_r17 *SL_InterUE_CoordinationScheme1_r17 `optional`
	Sl_InterUE_CoordinationScheme2_r17 *SL_InterUE_CoordinationScheme2_r17 `optional`
}

func (ie *SL_InterUE_CoordinationConfig_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_InterUE_CoordinationScheme1_r17 != nil, ie.Sl_InterUE_CoordinationScheme2_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_InterUE_CoordinationScheme1_r17 != nil {
		if err = ie.Sl_InterUE_CoordinationScheme1_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_InterUE_CoordinationScheme1_r17", err)
		}
	}
	if ie.Sl_InterUE_CoordinationScheme2_r17 != nil {
		if err = ie.Sl_InterUE_CoordinationScheme2_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_InterUE_CoordinationScheme2_r17", err)
		}
	}
	return nil
}

func (ie *SL_InterUE_CoordinationConfig_r17) Decode(r *uper.UperReader) error {
	var err error
	var Sl_InterUE_CoordinationScheme1_r17Present bool
	if Sl_InterUE_CoordinationScheme1_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Sl_InterUE_CoordinationScheme2_r17Present bool
	if Sl_InterUE_CoordinationScheme2_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_InterUE_CoordinationScheme1_r17Present {
		ie.Sl_InterUE_CoordinationScheme1_r17 = new(SL_InterUE_CoordinationScheme1_r17)
		if err = ie.Sl_InterUE_CoordinationScheme1_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_InterUE_CoordinationScheme1_r17", err)
		}
	}
	if Sl_InterUE_CoordinationScheme2_r17Present {
		ie.Sl_InterUE_CoordinationScheme2_r17 = new(SL_InterUE_CoordinationScheme2_r17)
		if err = ie.Sl_InterUE_CoordinationScheme2_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_InterUE_CoordinationScheme2_r17", err)
		}
	}
	return nil
}
