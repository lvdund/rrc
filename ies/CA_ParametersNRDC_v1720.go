package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC_v1720 struct {
	Ca_ParametersNR_ForDC_v1700 *CA_ParametersNR_v1700 `optional`
	Ca_ParametersNR_ForDC_v1720 *CA_ParametersNR_v1720 `optional`
}

func (ie *CA_ParametersNRDC_v1720) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNR_ForDC_v1700 != nil, ie.Ca_ParametersNR_ForDC_v1720 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ca_ParametersNR_ForDC_v1700 != nil {
		if err = ie.Ca_ParametersNR_ForDC_v1700.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_ForDC_v1700", err)
		}
	}
	if ie.Ca_ParametersNR_ForDC_v1720 != nil {
		if err = ie.Ca_ParametersNR_ForDC_v1720.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_ForDC_v1720", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC_v1720) Decode(r *uper.UperReader) error {
	var err error
	var Ca_ParametersNR_ForDC_v1700Present bool
	if Ca_ParametersNR_ForDC_v1700Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNR_ForDC_v1720Present bool
	if Ca_ParametersNR_ForDC_v1720Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ca_ParametersNR_ForDC_v1700Present {
		ie.Ca_ParametersNR_ForDC_v1700 = new(CA_ParametersNR_v1700)
		if err = ie.Ca_ParametersNR_ForDC_v1700.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_ForDC_v1700", err)
		}
	}
	if Ca_ParametersNR_ForDC_v1720Present {
		ie.Ca_ParametersNR_ForDC_v1720 = new(CA_ParametersNR_v1720)
		if err = ie.Ca_ParametersNR_ForDC_v1720.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_ForDC_v1720", err)
		}
	}
	return nil
}
