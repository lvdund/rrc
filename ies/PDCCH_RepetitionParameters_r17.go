package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PDCCH_RepetitionParameters_r17 struct {
	SupportedMode_r17   PDCCH_RepetitionParameters_r17_supportedMode_r17    `madatory`
	LimitX_PerCC_r17    *PDCCH_RepetitionParameters_r17_limitX_PerCC_r17    `optional`
	LimitX_AcrossCC_r17 *PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17 `optional`
}

func (ie *PDCCH_RepetitionParameters_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.LimitX_PerCC_r17 != nil, ie.LimitX_AcrossCC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.SupportedMode_r17.Encode(w); err != nil {
		return utils.WrapError("Encode SupportedMode_r17", err)
	}
	if ie.LimitX_PerCC_r17 != nil {
		if err = ie.LimitX_PerCC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode LimitX_PerCC_r17", err)
		}
	}
	if ie.LimitX_AcrossCC_r17 != nil {
		if err = ie.LimitX_AcrossCC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode LimitX_AcrossCC_r17", err)
		}
	}
	return nil
}

func (ie *PDCCH_RepetitionParameters_r17) Decode(r *uper.UperReader) error {
	var err error
	var LimitX_PerCC_r17Present bool
	if LimitX_PerCC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LimitX_AcrossCC_r17Present bool
	if LimitX_AcrossCC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.SupportedMode_r17.Decode(r); err != nil {
		return utils.WrapError("Decode SupportedMode_r17", err)
	}
	if LimitX_PerCC_r17Present {
		ie.LimitX_PerCC_r17 = new(PDCCH_RepetitionParameters_r17_limitX_PerCC_r17)
		if err = ie.LimitX_PerCC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode LimitX_PerCC_r17", err)
		}
	}
	if LimitX_AcrossCC_r17Present {
		ie.LimitX_AcrossCC_r17 = new(PDCCH_RepetitionParameters_r17_limitX_AcrossCC_r17)
		if err = ie.LimitX_AcrossCC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode LimitX_AcrossCC_r17", err)
		}
	}
	return nil
}
