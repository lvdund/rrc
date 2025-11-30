package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_timeDurationForQCL struct {
	Scs_60kHz  *FeatureSetDownlink_timeDurationForQCL_scs_60kHz  `optional`
	Scs_120kHz *FeatureSetDownlink_timeDurationForQCL_scs_120kHz `optional`
}

func (ie *FeatureSetDownlink_timeDurationForQCL) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_60kHz != nil, ie.Scs_120kHz != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
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
	return nil
}

func (ie *FeatureSetDownlink_timeDurationForQCL) Decode(r *aper.AperReader) error {
	var err error
	var Scs_60kHzPresent bool
	if Scs_60kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Scs_120kHzPresent bool
	if Scs_120kHzPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Scs_60kHzPresent {
		ie.Scs_60kHz = new(FeatureSetDownlink_timeDurationForQCL_scs_60kHz)
		if err = ie.Scs_60kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_60kHz", err)
		}
	}
	if Scs_120kHzPresent {
		ie.Scs_120kHz = new(FeatureSetDownlink_timeDurationForQCL_scs_120kHz)
		if err = ie.Scs_120kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_120kHz", err)
		}
	}
	return nil
}
