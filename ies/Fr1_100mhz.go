package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Fr1_100mhz struct {
	Scs_15kHz *Fr1_100mhz_scs_15kHz `optional`
	Scs_30kHz *Fr1_100mhz_scs_30kHz `optional`
	Scs_60kHz *Fr1_100mhz_scs_60kHz `optional`
}

func (ie *Fr1_100mhz) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Scs_15kHz != nil, ie.Scs_30kHz != nil, ie.Scs_60kHz != nil}
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
	return nil
}

func (ie *Fr1_100mhz) Decode(r *aper.AperReader) error {
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
	if Scs_15kHzPresent {
		ie.Scs_15kHz = new(Fr1_100mhz_scs_15kHz)
		if err = ie.Scs_15kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_15kHz", err)
		}
	}
	if Scs_30kHzPresent {
		ie.Scs_30kHz = new(Fr1_100mhz_scs_30kHz)
		if err = ie.Scs_30kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_30kHz", err)
		}
	}
	if Scs_60kHzPresent {
		ie.Scs_60kHz = new(Fr1_100mhz_scs_60kHz)
		if err = ie.Scs_60kHz.Decode(r); err != nil {
			return utils.WrapError("Decode Scs_60kHz", err)
		}
	}
	return nil
}
