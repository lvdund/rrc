package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SecurityConfig struct {
	SecurityAlgorithmConfig *SecurityAlgorithmConfig `optional`
	KeyToUse                *SecurityConfig_keyToUse `optional`
}

func (ie *SecurityConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SecurityAlgorithmConfig != nil, ie.KeyToUse != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SecurityAlgorithmConfig != nil {
		if err = ie.SecurityAlgorithmConfig.Encode(w); err != nil {
			return utils.WrapError("Encode SecurityAlgorithmConfig", err)
		}
	}
	if ie.KeyToUse != nil {
		if err = ie.KeyToUse.Encode(w); err != nil {
			return utils.WrapError("Encode KeyToUse", err)
		}
	}
	return nil
}

func (ie *SecurityConfig) Decode(r *uper.UperReader) error {
	var err error
	var SecurityAlgorithmConfigPresent bool
	if SecurityAlgorithmConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var KeyToUsePresent bool
	if KeyToUsePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SecurityAlgorithmConfigPresent {
		ie.SecurityAlgorithmConfig = new(SecurityAlgorithmConfig)
		if err = ie.SecurityAlgorithmConfig.Decode(r); err != nil {
			return utils.WrapError("Decode SecurityAlgorithmConfig", err)
		}
	}
	if KeyToUsePresent {
		ie.KeyToUse = new(SecurityConfig_keyToUse)
		if err = ie.KeyToUse.Decode(r); err != nil {
			return utils.WrapError("Decode KeyToUse", err)
		}
	}
	return nil
}
