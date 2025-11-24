package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DMRS_UplinkConfig struct {
	Dmrs_Type                  *DMRS_UplinkConfig_dmrs_Type                  `optional`
	Dmrs_AdditionalPosition    *DMRS_UplinkConfig_dmrs_AdditionalPosition    `optional`
	PhaseTrackingRS            *PTRS_UplinkConfig                            `optional,setuprelease`
	MaxLength                  *DMRS_UplinkConfig_maxLength                  `optional`
	TransformPrecodingDisabled *DMRS_UplinkConfig_transformPrecodingDisabled `lb:0,ub:65535,optional`
	TransformPrecodingEnabled  *DMRS_UplinkConfig_transformPrecodingEnabled  `lb:0,ub:1007,optional,ext`
}

func (ie *DMRS_UplinkConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Dmrs_Type != nil, ie.Dmrs_AdditionalPosition != nil, ie.PhaseTrackingRS != nil, ie.MaxLength != nil, ie.TransformPrecodingDisabled != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dmrs_Type != nil {
		if err = ie.Dmrs_Type.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_Type", err)
		}
	}
	if ie.Dmrs_AdditionalPosition != nil {
		if err = ie.Dmrs_AdditionalPosition.Encode(w); err != nil {
			return utils.WrapError("Encode Dmrs_AdditionalPosition", err)
		}
	}
	if ie.PhaseTrackingRS != nil {
		tmp_PhaseTrackingRS := utils.SetupRelease[*PTRS_UplinkConfig]{
			Setup: ie.PhaseTrackingRS,
		}
		if err = tmp_PhaseTrackingRS.Encode(w); err != nil {
			return utils.WrapError("Encode PhaseTrackingRS", err)
		}
	}
	if ie.MaxLength != nil {
		if err = ie.MaxLength.Encode(w); err != nil {
			return utils.WrapError("Encode MaxLength", err)
		}
	}
	if ie.TransformPrecodingDisabled != nil {
		if err = ie.TransformPrecodingDisabled.Encode(w); err != nil {
			return utils.WrapError("Encode TransformPrecodingDisabled", err)
		}
	}
	return nil
}

func (ie *DMRS_UplinkConfig) Decode(r *uper.UperReader) error {
	var err error
	var Dmrs_TypePresent bool
	if Dmrs_TypePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Dmrs_AdditionalPositionPresent bool
	if Dmrs_AdditionalPositionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var PhaseTrackingRSPresent bool
	if PhaseTrackingRSPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxLengthPresent bool
	if MaxLengthPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TransformPrecodingDisabledPresent bool
	if TransformPrecodingDisabledPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Dmrs_TypePresent {
		ie.Dmrs_Type = new(DMRS_UplinkConfig_dmrs_Type)
		if err = ie.Dmrs_Type.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_Type", err)
		}
	}
	if Dmrs_AdditionalPositionPresent {
		ie.Dmrs_AdditionalPosition = new(DMRS_UplinkConfig_dmrs_AdditionalPosition)
		if err = ie.Dmrs_AdditionalPosition.Decode(r); err != nil {
			return utils.WrapError("Decode Dmrs_AdditionalPosition", err)
		}
	}
	if PhaseTrackingRSPresent {
		tmp_PhaseTrackingRS := utils.SetupRelease[*PTRS_UplinkConfig]{}
		if err = tmp_PhaseTrackingRS.Decode(r); err != nil {
			return utils.WrapError("Decode PhaseTrackingRS", err)
		}
		ie.PhaseTrackingRS = tmp_PhaseTrackingRS.Setup
	}
	if MaxLengthPresent {
		ie.MaxLength = new(DMRS_UplinkConfig_maxLength)
		if err = ie.MaxLength.Decode(r); err != nil {
			return utils.WrapError("Decode MaxLength", err)
		}
	}
	if TransformPrecodingDisabledPresent {
		ie.TransformPrecodingDisabled = new(DMRS_UplinkConfig_transformPrecodingDisabled)
		if err = ie.TransformPrecodingDisabled.Decode(r); err != nil {
			return utils.WrapError("Decode TransformPrecodingDisabled", err)
		}
	}
	return nil
}
