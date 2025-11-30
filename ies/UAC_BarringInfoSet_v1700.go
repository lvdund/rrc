package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UAC_BarringInfoSet_v1700 struct {
	Uac_BarringFactorForAI3_r17 *UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17 `optional`
}

func (ie *UAC_BarringInfoSet_v1700) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Uac_BarringFactorForAI3_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Uac_BarringFactorForAI3_r17 != nil {
		if err = ie.Uac_BarringFactorForAI3_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Uac_BarringFactorForAI3_r17", err)
		}
	}
	return nil
}

func (ie *UAC_BarringInfoSet_v1700) Decode(r *aper.AperReader) error {
	var err error
	var Uac_BarringFactorForAI3_r17Present bool
	if Uac_BarringFactorForAI3_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Uac_BarringFactorForAI3_r17Present {
		ie.Uac_BarringFactorForAI3_r17 = new(UAC_BarringInfoSet_v1700_uac_BarringFactorForAI3_r17)
		if err = ie.Uac_BarringFactorForAI3_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Uac_BarringFactorForAI3_r17", err)
		}
	}
	return nil
}
