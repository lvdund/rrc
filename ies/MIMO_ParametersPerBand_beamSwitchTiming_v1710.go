package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MIMO_ParametersPerBand_beamSwitchTiming_v1710 struct {
	Scs_480kHz *MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz `optional`
	Scs_960kHz *MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_960kHz `optional`
}

func (ie *MIMO_ParametersPerBand_beamSwitchTiming_v1710) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_480kHz != nil, ie.Scs_960kHz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Scs_480kHz != nil {
		if err = ie.Scs_480kHz.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_480kHz", err)
		}
	}
	if ie.Scs_960kHz != nil {
		if err = ie.Scs_960kHz.Encode(w); err != nil {
			return utils.WrapError("Encode Scs_960kHz", err)
		}
	}
	return nil
}

func (ie *MIMO_ParametersPerBand_beamSwitchTiming_v1710) Decode(r *aper.AperReader) error {
	var err error
	var Scs_480kHzPresent bool
	if Scs_480kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_960kHzPresent bool
	if Scs_960kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_480kHzPresent {
		ie.Scs_480kHz = new(MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_480kHz)
		if err = ie.Scs_480kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_480kHz", err)
		}
	}
	if Scs_960kHzPresent {
		ie.Scs_960kHz = new(MIMO_ParametersPerBand_beamSwitchTiming_v1710_scs_960kHz)
		if err = ie.Scs_960kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_960kHz", err)
		}
	}
	return nil
}
