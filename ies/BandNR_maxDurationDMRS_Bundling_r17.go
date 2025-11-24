package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandNR_maxDurationDMRS_Bundling_r17 struct {
	Fdd_r17 *BandNR_maxDurationDMRS_Bundling_r17_fdd_r17 `optional`
	Tdd_r17 *BandNR_maxDurationDMRS_Bundling_r17_tdd_r17 `optional`
}

func (ie *BandNR_maxDurationDMRS_Bundling_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Fdd_r17 != nil, ie.Tdd_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Fdd_r17 != nil {
		if err = ie.Fdd_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Fdd_r17", err)
		}
	}
	if ie.Tdd_r17 != nil {
		if err = ie.Tdd_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Tdd_r17", err)
		}
	}
	return nil
}

func (ie *BandNR_maxDurationDMRS_Bundling_r17) Decode(r *uper.UperReader) error {
	var err error
	var Fdd_r17Present bool
	if Fdd_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Tdd_r17Present bool
	if Tdd_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Fdd_r17Present {
		ie.Fdd_r17 = new(BandNR_maxDurationDMRS_Bundling_r17_fdd_r17)
		if err = ie.Fdd_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Fdd_r17", err)
		}
	}
	if Tdd_r17Present {
		ie.Tdd_r17 = new(BandNR_maxDurationDMRS_Bundling_r17_tdd_r17)
		if err = ie.Tdd_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Tdd_r17", err)
		}
	}
	return nil
}
