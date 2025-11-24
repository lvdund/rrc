package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type BandCombination_v1640 struct {
	Ca_ParametersNR_v1640   *CA_ParametersNR_v1640   `optional`
	Ca_ParametersNRDC_v1640 *CA_ParametersNRDC_v1640 `optional`
}

func (ie *BandCombination_v1640) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNR_v1640 != nil, ie.Ca_ParametersNRDC_v1640 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ca_ParametersNR_v1640 != nil {
		if err = ie.Ca_ParametersNR_v1640.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_v1640", err)
		}
	}
	if ie.Ca_ParametersNRDC_v1640 != nil {
		if err = ie.Ca_ParametersNRDC_v1640.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNRDC_v1640", err)
		}
	}
	return nil
}

func (ie *BandCombination_v1640) Decode(r *uper.UperReader) error {
	var err error
	var Ca_ParametersNR_v1640Present bool
	if Ca_ParametersNR_v1640Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNRDC_v1640Present bool
	if Ca_ParametersNRDC_v1640Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ca_ParametersNR_v1640Present {
		ie.Ca_ParametersNR_v1640 = new(CA_ParametersNR_v1640)
		if err = ie.Ca_ParametersNR_v1640.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_v1640", err)
		}
	}
	if Ca_ParametersNRDC_v1640Present {
		ie.Ca_ParametersNRDC_v1640 = new(CA_ParametersNRDC_v1640)
		if err = ie.Ca_ParametersNRDC_v1640.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNRDC_v1640", err)
		}
	}
	return nil
}
