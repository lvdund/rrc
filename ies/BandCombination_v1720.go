package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1720 struct {
	Ca_ParametersNR_v1720   *CA_ParametersNR_v1720   `optional`
	Ca_ParametersNRDC_v1720 *CA_ParametersNRDC_v1720 `optional`
}

func (ie *BandCombination_v1720) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNR_v1720 != nil, ie.Ca_ParametersNRDC_v1720 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ca_ParametersNR_v1720 != nil {
		if err = ie.Ca_ParametersNR_v1720.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_v1720", err)
		}
	}
	if ie.Ca_ParametersNRDC_v1720 != nil {
		if err = ie.Ca_ParametersNRDC_v1720.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNRDC_v1720", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1720) Decode(r *aper.AperReader) error {
	var err error
	var Ca_ParametersNR_v1720Present bool
	if Ca_ParametersNR_v1720Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNRDC_v1720Present bool
	if Ca_ParametersNRDC_v1720Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ca_ParametersNR_v1720Present {
		ie.Ca_ParametersNR_v1720 = new(CA_ParametersNR_v1720)
		if err = ie.Ca_ParametersNR_v1720.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_v1720", err)
		}
	}
	if Ca_ParametersNRDC_v1720Present {
		ie.Ca_ParametersNRDC_v1720 = new(CA_ParametersNRDC_v1720)
		if err = ie.Ca_ParametersNRDC_v1720.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNRDC_v1720", err)
		}
	}
	return nil
}
