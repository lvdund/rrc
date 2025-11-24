package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MeasAndMobParametersMRDC_Common_v1700 struct {
	CondPSCellChangeParameters_r17 *MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17 `optional`
	CondHandoverWithSCG_ENDC_r17   *MeasAndMobParametersMRDC_Common_v1700_condHandoverWithSCG_ENDC_r17   `optional`
	CondHandoverWithSCG_NEDC_r17   *MeasAndMobParametersMRDC_Common_v1700_condHandoverWithSCG_NEDC_r17   `optional`
}

func (ie *MeasAndMobParametersMRDC_Common_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CondPSCellChangeParameters_r17 != nil, ie.CondHandoverWithSCG_ENDC_r17 != nil, ie.CondHandoverWithSCG_NEDC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CondPSCellChangeParameters_r17 != nil {
		if err = ie.CondPSCellChangeParameters_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CondPSCellChangeParameters_r17", err)
		}
	}
	if ie.CondHandoverWithSCG_ENDC_r17 != nil {
		if err = ie.CondHandoverWithSCG_ENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CondHandoverWithSCG_ENDC_r17", err)
		}
	}
	if ie.CondHandoverWithSCG_NEDC_r17 != nil {
		if err = ie.CondHandoverWithSCG_NEDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CondHandoverWithSCG_NEDC_r17", err)
		}
	}
	return nil
}

func (ie *MeasAndMobParametersMRDC_Common_v1700) Decode(r *uper.UperReader) error {
	var err error
	var CondPSCellChangeParameters_r17Present bool
	if CondPSCellChangeParameters_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CondHandoverWithSCG_ENDC_r17Present bool
	if CondHandoverWithSCG_ENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var CondHandoverWithSCG_NEDC_r17Present bool
	if CondHandoverWithSCG_NEDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if CondPSCellChangeParameters_r17Present {
		ie.CondPSCellChangeParameters_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condPSCellChangeParameters_r17)
		if err = ie.CondPSCellChangeParameters_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CondPSCellChangeParameters_r17", err)
		}
	}
	if CondHandoverWithSCG_ENDC_r17Present {
		ie.CondHandoverWithSCG_ENDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condHandoverWithSCG_ENDC_r17)
		if err = ie.CondHandoverWithSCG_ENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CondHandoverWithSCG_ENDC_r17", err)
		}
	}
	if CondHandoverWithSCG_NEDC_r17Present {
		ie.CondHandoverWithSCG_NEDC_r17 = new(MeasAndMobParametersMRDC_Common_v1700_condHandoverWithSCG_NEDC_r17)
		if err = ie.CondHandoverWithSCG_NEDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CondHandoverWithSCG_NEDC_r17", err)
		}
	}
	return nil
}
