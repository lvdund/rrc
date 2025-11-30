package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC_v1640 struct {
	Ca_ParametersNR_ForDC_v1640 *CA_ParametersNR_v1640 `optional`
}

func (ie *CA_ParametersNRDC_v1640) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNR_ForDC_v1640 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ca_ParametersNR_ForDC_v1640 != nil {
		if err = ie.Ca_ParametersNR_ForDC_v1640.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_ForDC_v1640", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC_v1640) Decode(r *aper.AperReader) error {
	var err error
	var Ca_ParametersNR_ForDC_v1640Present bool
	if Ca_ParametersNR_ForDC_v1640Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Ca_ParametersNR_ForDC_v1640Present {
		ie.Ca_ParametersNR_ForDC_v1640 = new(CA_ParametersNR_v1640)
		if err = ie.Ca_ParametersNR_ForDC_v1640.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_ForDC_v1640", err)
		}
	}
	return nil
}
