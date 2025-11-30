package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC_v1630 struct {
	Ca_ParametersNR_ForDC_v1610 *CA_ParametersNR_v1610 `optional`
	Ca_ParametersNR_ForDC_v1630 *CA_ParametersNR_v1630 `optional`
}

func (ie *CA_ParametersNRDC_v1630) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNR_ForDC_v1610 != nil, ie.Ca_ParametersNR_ForDC_v1630 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ca_ParametersNR_ForDC_v1610 != nil {
		if err = ie.Ca_ParametersNR_ForDC_v1610.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_ForDC_v1610", err)
		}
	}
	if ie.Ca_ParametersNR_ForDC_v1630 != nil {
		if err = ie.Ca_ParametersNR_ForDC_v1630.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_ForDC_v1630", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC_v1630) Decode(r *aper.AperReader) error {
	var err error
	var Ca_ParametersNR_ForDC_v1610Present bool
	if Ca_ParametersNR_ForDC_v1610Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNR_ForDC_v1630Present bool
	if Ca_ParametersNR_ForDC_v1630Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ca_ParametersNR_ForDC_v1610Present {
		ie.Ca_ParametersNR_ForDC_v1610 = new(CA_ParametersNR_v1610)
		if err = ie.Ca_ParametersNR_ForDC_v1610.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_ForDC_v1610", err)
		}
	}
	if Ca_ParametersNR_ForDC_v1630Present {
		ie.Ca_ParametersNR_ForDC_v1630 = new(CA_ParametersNR_v1630)
		if err = ie.Ca_ParametersNR_ForDC_v1630.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_ForDC_v1630", err)
		}
	}
	return nil
}
