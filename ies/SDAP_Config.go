package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SDAP_Config struct {
	Pdu_Session              PDU_SessionID             `madatory`
	Sdap_HeaderDL            SDAP_Config_sdap_HeaderDL `madatory`
	Sdap_HeaderUL            SDAP_Config_sdap_HeaderUL `madatory`
	DefaultDRB               bool                      `madatory`
	MappedQoS_FlowsToAdd     []QFI                     `lb:1,ub:maxNrofQFIs,optional`
	MappedQoS_FlowsToRelease []QFI                     `lb:1,ub:maxNrofQFIs,optional`
}

func (ie *SDAP_Config) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{len(ie.MappedQoS_FlowsToAdd) > 0, len(ie.MappedQoS_FlowsToRelease) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Pdu_Session.Encode(w); err != nil {
		return utils.WrapError("Encode Pdu_Session", err)
	}
	if err = ie.Sdap_HeaderDL.Encode(w); err != nil {
		return utils.WrapError("Encode Sdap_HeaderDL", err)
	}
	if err = ie.Sdap_HeaderUL.Encode(w); err != nil {
		return utils.WrapError("Encode Sdap_HeaderUL", err)
	}
	if err = w.WriteBoolean(ie.DefaultDRB); err != nil {
		return utils.WrapError("WriteBoolean DefaultDRB", err)
	}
	if len(ie.MappedQoS_FlowsToAdd) > 0 {
		tmp_MappedQoS_FlowsToAdd := utils.NewSequence[*QFI]([]*QFI{}, uper.Constraint{Lb: 1, Ub: maxNrofQFIs}, false)
		for _, i := range ie.MappedQoS_FlowsToAdd {
			tmp_MappedQoS_FlowsToAdd.Value = append(tmp_MappedQoS_FlowsToAdd.Value, &i)
		}
		if err = tmp_MappedQoS_FlowsToAdd.Encode(w); err != nil {
			return utils.WrapError("Encode MappedQoS_FlowsToAdd", err)
		}
	}
	if len(ie.MappedQoS_FlowsToRelease) > 0 {
		tmp_MappedQoS_FlowsToRelease := utils.NewSequence[*QFI]([]*QFI{}, uper.Constraint{Lb: 1, Ub: maxNrofQFIs}, false)
		for _, i := range ie.MappedQoS_FlowsToRelease {
			tmp_MappedQoS_FlowsToRelease.Value = append(tmp_MappedQoS_FlowsToRelease.Value, &i)
		}
		if err = tmp_MappedQoS_FlowsToRelease.Encode(w); err != nil {
			return utils.WrapError("Encode MappedQoS_FlowsToRelease", err)
		}
	}
	return nil
}

func (ie *SDAP_Config) Decode(r *uper.UperReader) error {
	var err error
	var MappedQoS_FlowsToAddPresent bool
	if MappedQoS_FlowsToAddPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var MappedQoS_FlowsToReleasePresent bool
	if MappedQoS_FlowsToReleasePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Pdu_Session.Decode(r); err != nil {
		return utils.WrapError("Decode Pdu_Session", err)
	}
	if err = ie.Sdap_HeaderDL.Decode(r); err != nil {
		return utils.WrapError("Decode Sdap_HeaderDL", err)
	}
	if err = ie.Sdap_HeaderUL.Decode(r); err != nil {
		return utils.WrapError("Decode Sdap_HeaderUL", err)
	}
	var tmp_bool_DefaultDRB bool
	if tmp_bool_DefaultDRB, err = r.ReadBoolean(); err != nil {
		return utils.WrapError("ReadBoolean DefaultDRB", err)
	}
	ie.DefaultDRB = tmp_bool_DefaultDRB
	if MappedQoS_FlowsToAddPresent {
		tmp_MappedQoS_FlowsToAdd := utils.NewSequence[*QFI]([]*QFI{}, uper.Constraint{Lb: 1, Ub: maxNrofQFIs}, false)
		fn_MappedQoS_FlowsToAdd := func() *QFI {
			return new(QFI)
		}
		if err = tmp_MappedQoS_FlowsToAdd.Decode(r, fn_MappedQoS_FlowsToAdd); err != nil {
			return utils.WrapError("Decode MappedQoS_FlowsToAdd", err)
		}
		ie.MappedQoS_FlowsToAdd = []QFI{}
		for _, i := range tmp_MappedQoS_FlowsToAdd.Value {
			ie.MappedQoS_FlowsToAdd = append(ie.MappedQoS_FlowsToAdd, *i)
		}
	}
	if MappedQoS_FlowsToReleasePresent {
		tmp_MappedQoS_FlowsToRelease := utils.NewSequence[*QFI]([]*QFI{}, uper.Constraint{Lb: 1, Ub: maxNrofQFIs}, false)
		fn_MappedQoS_FlowsToRelease := func() *QFI {
			return new(QFI)
		}
		if err = tmp_MappedQoS_FlowsToRelease.Decode(r, fn_MappedQoS_FlowsToRelease); err != nil {
			return utils.WrapError("Decode MappedQoS_FlowsToRelease", err)
		}
		ie.MappedQoS_FlowsToRelease = []QFI{}
		for _, i := range tmp_MappedQoS_FlowsToRelease.Value {
			ie.MappedQoS_FlowsToRelease = append(ie.MappedQoS_FlowsToRelease, *i)
		}
	}
	return nil
}
