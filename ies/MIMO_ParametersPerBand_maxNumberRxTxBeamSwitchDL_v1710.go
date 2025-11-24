package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_v1710 struct {
	Scs_480kHz_r17 *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_v1710_scs_480kHz_r17 `optional`
	Scs_960kHz_r17 *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_v1710_scs_960kHz_r17 `optional`
}

func (ie *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_v1710) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_480kHz_r17 != nil, ie.Scs_960kHz_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs_480kHz_r17 != nil {
		if err = ie.Scs_480kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_480kHz_r17", err)
		}
	}
	if ie.Scs_960kHz_r17 != nil {
		if err = ie.Scs_960kHz_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_960kHz_r17", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_v1710) Decode(r *uper.UperReader) error {
	var err error
	var Scs_480kHz_r17Present bool
	if Scs_480kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_960kHz_r17Present bool
	if Scs_960kHz_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_480kHz_r17Present {
		ie.Scs_480kHz_r17 = new(MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_v1710_scs_480kHz_r17)
		if err = ie.Scs_480kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_480kHz_r17", err)
		}
	}
	if Scs_960kHz_r17Present {
		ie.Scs_960kHz_r17 = new(MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_v1710_scs_960kHz_r17)
		if err = ie.Scs_960kHz_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_960kHz_r17", err)
		}
	}
	return nil
}
