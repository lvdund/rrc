package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1650 struct {
	Ca_ParametersNRDC_v1650 *CA_ParametersNRDC_v1650 `optional`
}

func (ie *BandCombination_v1650) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNRDC_v1650 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ca_ParametersNRDC_v1650 != nil {
		if err = ie.Ca_ParametersNRDC_v1650.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNRDC_v1650", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1650) Decode(r *aper.AperReader) error {
	var err error
	var Ca_ParametersNRDC_v1650Present bool
	if Ca_ParametersNRDC_v1650Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ca_ParametersNRDC_v1650Present {
		ie.Ca_ParametersNRDC_v1650 = new(CA_ParametersNRDC_v1650)
		if err = ie.Ca_ParametersNRDC_v1650.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNRDC_v1650", err)
		}
	}
	return nil
}
