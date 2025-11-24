package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PUCCH_ResourceSet struct {
	Pucch_ResourceSetId PUCCH_ResourceSetId `madatory`
	ResourceList        []PUCCH_ResourceId  `lb:1,ub:maxNrofPUCCH_ResourcesPerSet,madatory`
	MaxPayloadSize      *int64              `lb:4,ub:256,optional`
}

func (ie *PUCCH_ResourceSet) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MaxPayloadSize != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Pucch_ResourceSetId.Encode(w); err != nil {
		return utils.WrapError("Encode Pucch_ResourceSetId", err)
	}
	tmp_ResourceList := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourcesPerSet}, false)
	for _, i := range ie.ResourceList {
		tmp_ResourceList.Value = append(tmp_ResourceList.Value, &i)
	}
	if err = tmp_ResourceList.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceList", err)
	}
	if ie.MaxPayloadSize != nil {
		if err = w.WriteInteger(*ie.MaxPayloadSize, &uper.Constraint{Lb: 4, Ub: 256}, false); err != nil {
			return utils.WrapError("Encode MaxPayloadSize", err)
		}
	}
	return nil
}

func (ie *PUCCH_ResourceSet) Decode(r *uper.UperReader) error {
	var err error
	var MaxPayloadSizePresent bool
	if MaxPayloadSizePresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Pucch_ResourceSetId.Decode(r); err != nil {
		return utils.WrapError("Decode Pucch_ResourceSetId", err)
	}
	tmp_ResourceList := utils.NewSequence[*PUCCH_ResourceId]([]*PUCCH_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofPUCCH_ResourcesPerSet}, false)
	fn_ResourceList := func() *PUCCH_ResourceId {
		return new(PUCCH_ResourceId)
	}
	if err = tmp_ResourceList.Decode(r, fn_ResourceList); err != nil {
		return utils.WrapError("Decode ResourceList", err)
	}
	ie.ResourceList = []PUCCH_ResourceId{}
	for _, i := range tmp_ResourceList.Value {
		ie.ResourceList = append(ie.ResourceList, *i)
	}
	if MaxPayloadSizePresent {
		var tmp_int_MaxPayloadSize int64
		if tmp_int_MaxPayloadSize, err = r.ReadInteger(&uper.Constraint{Lb: 4, Ub: 256}, false); err != nil {
			return utils.WrapError("Decode MaxPayloadSize", err)
		}
		ie.MaxPayloadSize = &tmp_int_MaxPayloadSize
	}
	return nil
}
