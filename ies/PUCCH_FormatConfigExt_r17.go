package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_FormatConfigExt_r17 struct {
	MaxCodeRateLP_r17 *PUCCH_MaxCodeRate `optional`
}

func (ie *PUCCH_FormatConfigExt_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxCodeRateLP_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MaxCodeRateLP_r17 != nil {
		if err = ie.MaxCodeRateLP_r17.Encode(w); err != nil {
			return utils.WrapError("Encode MaxCodeRateLP_r17", err)
		}
	}
	return nil
}

func (ie *PUCCH_FormatConfigExt_r17) Decode(r *aper.AperReader) error {
	var err error
	var MaxCodeRateLP_r17Present bool
	if MaxCodeRateLP_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MaxCodeRateLP_r17Present {
		ie.MaxCodeRateLP_r17 = new(PUCCH_MaxCodeRate)
		if err = ie.MaxCodeRateLP_r17.Decode(r); err != nil {
			return utils.WrapError("Decode MaxCodeRateLP_r17", err)
		}
	}
	return nil
}
