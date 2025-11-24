package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PTRS_UplinkConfig struct {
	TransformPrecoderDisabled *PTRS_UplinkConfig_transformPrecoderDisabled `lb:2,ub:2,optional`
	TransformPrecoderEnabled  *PTRS_UplinkConfig_transformPrecoderEnabled  `lb:5,ub:5,optional`
}

func (ie *PTRS_UplinkConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.TransformPrecoderDisabled != nil, ie.TransformPrecoderEnabled != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.TransformPrecoderDisabled != nil {
		if err = ie.TransformPrecoderDisabled.Encode(w); err != nil {
			return utils.WrapError("Encode TransformPrecoderDisabled", err)
		}
	}
	if ie.TransformPrecoderEnabled != nil {
		if err = ie.TransformPrecoderEnabled.Encode(w); err != nil {
			return utils.WrapError("Encode TransformPrecoderEnabled", err)
		}
	}
	return nil
}

func (ie *PTRS_UplinkConfig) Decode(r *uper.UperReader) error {
	var err error
	var TransformPrecoderDisabledPresent bool
	if TransformPrecoderDisabledPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TransformPrecoderEnabledPresent bool
	if TransformPrecoderEnabledPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if TransformPrecoderDisabledPresent {
		ie.TransformPrecoderDisabled = new(PTRS_UplinkConfig_transformPrecoderDisabled)
		if err = ie.TransformPrecoderDisabled.Decode(r); err != nil {
			return utils.WrapError("Decode TransformPrecoderDisabled", err)
		}
	}
	if TransformPrecoderEnabledPresent {
		ie.TransformPrecoderEnabled = new(PTRS_UplinkConfig_transformPrecoderEnabled)
		if err = ie.TransformPrecoderEnabled.Decode(r); err != nil {
			return utils.WrapError("Decode TransformPrecoderEnabled", err)
		}
	}
	return nil
}
