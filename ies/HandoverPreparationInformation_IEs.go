package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type HandoverPreparationInformation_IEs struct {
	Ue_CapabilityRAT_List UE_CapabilityRAT_ContainerList `madatory`
	SourceConfig          *AS_Config                     `optional`
	Rrm_Config            *RRM_Config                    `optional`
	As_Context            *AS_Context                    `optional`
	NonCriticalExtension  interface{}                    `optional`
}

func (ie *HandoverPreparationInformation_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SourceConfig != nil, ie.Rrm_Config != nil, ie.As_Context != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SourceConfig != nil {
		if err = ie.SourceConfig.Encode(w); err != nil {
			return utils.WrapError("Encode SourceConfig", err)
		}
	}
	if ie.Rrm_Config != nil {
		if err = ie.Rrm_Config.Encode(w); err != nil {
			return utils.WrapError("Encode Rrm_Config", err)
		}
	}
	if ie.As_Context != nil {
		if err = ie.As_Context.Encode(w); err != nil {
			return utils.WrapError("Encode As_Context", err)
		}
	}
	return nil
}

func (ie *HandoverPreparationInformation_IEs) Decode(r *uper.UperReader) error {
	var err error
	var SourceConfigPresent bool
	if SourceConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Rrm_ConfigPresent bool
	if Rrm_ConfigPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var As_ContextPresent bool
	if As_ContextPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SourceConfigPresent {
		ie.SourceConfig = new(AS_Config)
		if err = ie.SourceConfig.Decode(r); err != nil {
			return utils.WrapError("Decode SourceConfig", err)
		}
	}
	if Rrm_ConfigPresent {
		ie.Rrm_Config = new(RRM_Config)
		if err = ie.Rrm_Config.Decode(r); err != nil {
			return utils.WrapError("Decode Rrm_Config", err)
		}
	}
	if As_ContextPresent {
		ie.As_Context = new(AS_Context)
		if err = ie.As_Context.Decode(r); err != nil {
			return utils.WrapError("Decode As_Context", err)
		}
	}
	return nil
}
