package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CA_ParametersNRDC struct {
	Ca_ParametersNR_ForDC       *CA_ParametersNR         `optional`
	Ca_ParametersNR_ForDC_v1540 *CA_ParametersNR_v1540   `optional`
	Ca_ParametersNR_ForDC_v1550 *CA_ParametersNR_v1550   `optional`
	Ca_ParametersNR_ForDC_v1560 *CA_ParametersNR_v1560   `optional`
	FeatureSetCombinationDC     *FeatureSetCombinationId `optional`
}

func (ie *CA_ParametersNRDC) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Ca_ParametersNR_ForDC != nil, ie.Ca_ParametersNR_ForDC_v1540 != nil, ie.Ca_ParametersNR_ForDC_v1550 != nil, ie.Ca_ParametersNR_ForDC_v1560 != nil, ie.FeatureSetCombinationDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Ca_ParametersNR_ForDC != nil {
		if err = ie.Ca_ParametersNR_ForDC.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_ForDC", err)
		}
	}
	if ie.Ca_ParametersNR_ForDC_v1540 != nil {
		if err = ie.Ca_ParametersNR_ForDC_v1540.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_ForDC_v1540", err)
		}
	}
	if ie.Ca_ParametersNR_ForDC_v1550 != nil {
		if err = ie.Ca_ParametersNR_ForDC_v1550.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_ForDC_v1550", err)
		}
	}
	if ie.Ca_ParametersNR_ForDC_v1560 != nil {
		if err = ie.Ca_ParametersNR_ForDC_v1560.Encode(w); err != nil {
			return utils.WrapError("Encode Ca_ParametersNR_ForDC_v1560", err)
		}
	}
	if ie.FeatureSetCombinationDC != nil {
		if err = ie.FeatureSetCombinationDC.Encode(w); err != nil {
			return utils.WrapError("Encode FeatureSetCombinationDC", err)
		}
	}
	return nil
}

func (ie *CA_ParametersNRDC) Decode(r *uper.UperReader) error {
	var err error
	var Ca_ParametersNR_ForDCPresent bool
	if Ca_ParametersNR_ForDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNR_ForDC_v1540Present bool
	if Ca_ParametersNR_ForDC_v1540Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNR_ForDC_v1550Present bool
	if Ca_ParametersNR_ForDC_v1550Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Ca_ParametersNR_ForDC_v1560Present bool
	if Ca_ParametersNR_ForDC_v1560Present, err = r.ReadBool(); err != nil {
		return err
	}
	var FeatureSetCombinationDCPresent bool
	if FeatureSetCombinationDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Ca_ParametersNR_ForDCPresent {
		ie.Ca_ParametersNR_ForDC = new(CA_ParametersNR)
		if err = ie.Ca_ParametersNR_ForDC.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_ForDC", err)
		}
	}
	if Ca_ParametersNR_ForDC_v1540Present {
		ie.Ca_ParametersNR_ForDC_v1540 = new(CA_ParametersNR_v1540)
		if err = ie.Ca_ParametersNR_ForDC_v1540.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_ForDC_v1540", err)
		}
	}
	if Ca_ParametersNR_ForDC_v1550Present {
		ie.Ca_ParametersNR_ForDC_v1550 = new(CA_ParametersNR_v1550)
		if err = ie.Ca_ParametersNR_ForDC_v1550.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_ForDC_v1550", err)
		}
	}
	if Ca_ParametersNR_ForDC_v1560Present {
		ie.Ca_ParametersNR_ForDC_v1560 = new(CA_ParametersNR_v1560)
		if err = ie.Ca_ParametersNR_ForDC_v1560.Decode(r); err != nil {
			return utils.WrapError("Decode Ca_ParametersNR_ForDC_v1560", err)
		}
	}
	if FeatureSetCombinationDCPresent {
		ie.FeatureSetCombinationDC = new(FeatureSetCombinationId)
		if err = ie.FeatureSetCombinationDC.Decode(r); err != nil {
			return utils.WrapError("Decode FeatureSetCombinationDC", err)
		}
	}
	return nil
}
