package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v15g0 struct {
	Ca_ParametersNR_v15g0   *CA_ParametersNR_v15g0   `optional`
	Ca_ParametersNRDC_v15g0 *CA_ParametersNRDC_v15g0 `optional`
	Mrdc_Parameters_v15g0   *MRDC_Parameters_v15g0   `optional`
}

func (ie *BandCombination_v15g0) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNR_v15g0 != nil, ie.Ca_ParametersNRDC_v15g0 != nil, ie.Mrdc_Parameters_v15g0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ca_ParametersNR_v15g0 != nil {
		if err = ie.Ca_ParametersNR_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_v15g0", err)
		}
	}
	if ie.Ca_ParametersNRDC_v15g0 != nil {
		if err = ie.Ca_ParametersNRDC_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNRDC_v15g0", err)
		}
	}
	if ie.Mrdc_Parameters_v15g0 != nil {
		if err = ie.Mrdc_Parameters_v15g0.Encode(w); err != nil {
			return utils.WrapError("Encode Mrdc_Parameters_v15g0", err)
		}
	}
	return nil
}

func (ie *BandCombination_v15g0) Decode(r *aper.AperReader) error {
	var err error
	var Ca_ParametersNR_v15g0Present bool
	if Ca_ParametersNR_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNRDC_v15g0Present bool
	if Ca_ParametersNRDC_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Mrdc_Parameters_v15g0Present bool
	if Mrdc_Parameters_v15g0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ca_ParametersNR_v15g0Present {
		ie.Ca_ParametersNR_v15g0 = new(CA_ParametersNR_v15g0)
		if err = ie.Ca_ParametersNR_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_v15g0", err)
		}
	}
	if Ca_ParametersNRDC_v15g0Present {
		ie.Ca_ParametersNRDC_v15g0 = new(CA_ParametersNRDC_v15g0)
		if err = ie.Ca_ParametersNRDC_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNRDC_v15g0", err)
		}
	}
	if Mrdc_Parameters_v15g0Present {
		ie.Mrdc_Parameters_v15g0 = new(MRDC_Parameters_v15g0)
		if err = ie.Mrdc_Parameters_v15g0.Decode(r); err != nil {
			return utils.WrapError("Decode Mrdc_Parameters_v15g0", err)
		}
	}
	return nil
}
