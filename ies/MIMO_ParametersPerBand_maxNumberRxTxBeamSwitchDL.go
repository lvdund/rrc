package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL struct {
	Scs_15kHz  *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_15kHz  `optional`
	Scs_30kHz  *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_30kHz  `optional`
	Scs_60kHz  *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_60kHz  `optional`
	Scs_120kHz *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_120kHz `optional`
	Scs_240kHz *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_240kHz `optional`
}

func (ie *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_15kHz != nil, ie.Scs_30kHz != nil, ie.Scs_60kHz != nil, ie.Scs_120kHz != nil, ie.Scs_240kHz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs_15kHz != nil {
		if err = ie.Scs_15kHz.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_15kHz", err)
		}
	}
	if ie.Scs_30kHz != nil {
		if err = ie.Scs_30kHz.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_30kHz", err)
		}
	}
	if ie.Scs_60kHz != nil {
		if err = ie.Scs_60kHz.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_60kHz", err)
		}
	}
	if ie.Scs_120kHz != nil {
		if err = ie.Scs_120kHz.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_120kHz", err)
		}
	}
	if ie.Scs_240kHz != nil {
		if err = ie.Scs_240kHz.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_240kHz", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL) Decode(r *uper.UperReader) error {
	var err error
	var Scs_15kHzPresent bool
	if Scs_15kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_30kHzPresent bool
	if Scs_30kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_60kHzPresent bool
	if Scs_60kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_120kHzPresent bool
	if Scs_120kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_240kHzPresent bool
	if Scs_240kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_15kHzPresent {
		ie.Scs_15kHz = new(MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_15kHz)
		if err = ie.Scs_15kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_15kHz", err)
		}
	}
	if Scs_30kHzPresent {
		ie.Scs_30kHz = new(MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_30kHz)
		if err = ie.Scs_30kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_30kHz", err)
		}
	}
	if Scs_60kHzPresent {
		ie.Scs_60kHz = new(MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_60kHz)
		if err = ie.Scs_60kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_60kHz", err)
		}
	}
	if Scs_120kHzPresent {
		ie.Scs_120kHz = new(MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_120kHz)
		if err = ie.Scs_120kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_120kHz", err)
		}
	}
	if Scs_240kHzPresent {
		ie.Scs_240kHz = new(MIMO_ParametersPerBand_maxNumberRxTxBeamSwitchDL_scs_240kHz)
		if err = ie.Scs_240kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_240kHz", err)
		}
	}
	return nil
}
