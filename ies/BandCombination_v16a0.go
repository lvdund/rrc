package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v16a0 struct {
	Ca_ParametersNR_v16a0   *CA_ParametersNR_v16a0   `optional`
	Ca_ParametersNRDC_v16a0 *CA_ParametersNRDC_v16a0 `optional`
}

func (ie *BandCombination_v16a0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNR_v16a0 != nil, ie.Ca_ParametersNRDC_v16a0 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ca_ParametersNR_v16a0 != nil {
		if err = ie.Ca_ParametersNR_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_v16a0", err)
		}
	}
	if ie.Ca_ParametersNRDC_v16a0 != nil {
		if err = ie.Ca_ParametersNRDC_v16a0.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNRDC_v16a0", err)
		}
	}
	return nil
}

func (ie *BandCombination_v16a0) Decode(r *uper.UperReader) error {
	var err error
	var Ca_ParametersNR_v16a0Present bool
	if Ca_ParametersNR_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNRDC_v16a0Present bool
	if Ca_ParametersNRDC_v16a0Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ca_ParametersNR_v16a0Present {
		ie.Ca_ParametersNR_v16a0 = new(CA_ParametersNR_v16a0)
		if err = ie.Ca_ParametersNR_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_v16a0", err)
		}
	}
	if Ca_ParametersNRDC_v16a0Present {
		ie.Ca_ParametersNRDC_v16a0 = new(CA_ParametersNRDC_v16a0)
		if err = ie.Ca_ParametersNRDC_v16a0.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNRDC_v16a0", err)
		}
	}
	return nil
}
