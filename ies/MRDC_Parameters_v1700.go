package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1700 struct {
	CondPSCellAdditionENDC_r17               *MRDC_Parameters_v1700_condPSCellAdditionENDC_r17               `optional`
	Scg_ActivationDeactivationENDC_r17       *MRDC_Parameters_v1700_scg_ActivationDeactivationENDC_r17       `optional`
	Scg_ActivationDeactivationResumeENDC_r17 *MRDC_Parameters_v1700_scg_ActivationDeactivationResumeENDC_r17 `optional`
}

func (ie *MRDC_Parameters_v1700) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.CondPSCellAdditionENDC_r17 != nil, ie.Scg_ActivationDeactivationENDC_r17 != nil, ie.Scg_ActivationDeactivationResumeENDC_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.CondPSCellAdditionENDC_r17 != nil {
		if err = ie.CondPSCellAdditionENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CondPSCellAdditionENDC_r17", err)
		}
	}
	if ie.Scg_ActivationDeactivationENDC_r17 != nil {
		if err = ie.Scg_ActivationDeactivationENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scg_ActivationDeactivationENDC_r17", err)
		}
	}
	if ie.Scg_ActivationDeactivationResumeENDC_r17 != nil {
		if err = ie.Scg_ActivationDeactivationResumeENDC_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scg_ActivationDeactivationResumeENDC_r17", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1700) Decode(r *uper.UperReader) error {
	var err error
	var CondPSCellAdditionENDC_r17Present bool
	if CondPSCellAdditionENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scg_ActivationDeactivationENDC_r17Present bool
	if Scg_ActivationDeactivationENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scg_ActivationDeactivationResumeENDC_r17Present bool
	if Scg_ActivationDeactivationResumeENDC_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if CondPSCellAdditionENDC_r17Present {
		ie.CondPSCellAdditionENDC_r17 = new(MRDC_Parameters_v1700_condPSCellAdditionENDC_r17)
		if err = ie.CondPSCellAdditionENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode CondPSCellAdditionENDC_r17", err)
		}
	}
	if Scg_ActivationDeactivationENDC_r17Present {
		ie.Scg_ActivationDeactivationENDC_r17 = new(MRDC_Parameters_v1700_scg_ActivationDeactivationENDC_r17)
		if err = ie.Scg_ActivationDeactivationENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scg_ActivationDeactivationENDC_r17", err)
		}
	}
	if Scg_ActivationDeactivationResumeENDC_r17Present {
		ie.Scg_ActivationDeactivationResumeENDC_r17 = new(MRDC_Parameters_v1700_scg_ActivationDeactivationResumeENDC_r17)
		if err = ie.Scg_ActivationDeactivationResumeENDC_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scg_ActivationDeactivationResumeENDC_r17", err)
		}
	}
	return nil
}
