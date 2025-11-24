package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NR_DL_PRS_PDC_ResourceSet_r17 struct {
	PeriodicityAndOffset_r17     NR_DL_PRS_Periodicity_and_ResourceSetSlotOffset_r17 `madatory`
	NumSymbols_r17               NR_DL_PRS_PDC_ResourceSet_r17_numSymbols_r17        `madatory`
	Dl_PRS_ResourceBandwidth_r17 int64                                               `lb:1,ub:63,madatory`
	Dl_PRS_StartPRB_r17          int64                                               `lb:0,ub:2176,madatory`
	ResourceList_r17             []NR_DL_PRS_Resource_r17                            `lb:1,ub:maxNrofPRS_ResourcesPerSet_r17,madatory`
	RepFactorAndTimeGap_r17      *RepFactorAndTimeGap_r17                            `optional`
}

func (ie *NR_DL_PRS_PDC_ResourceSet_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.RepFactorAndTimeGap_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PeriodicityAndOffset_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PeriodicityAndOffset_r17", err)
	}
	if err = ie.NumSymbols_r17.Encode(w); err != nil {
		return utils.WrapError("Encode NumSymbols_r17", err)
	}
	if err = w.WriteInteger(ie.Dl_PRS_ResourceBandwidth_r17, &uper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
		return utils.WrapError("WriteInteger Dl_PRS_ResourceBandwidth_r17", err)
	}
	if err = w.WriteInteger(ie.Dl_PRS_StartPRB_r17, &uper.Constraint{Lb: 0, Ub: 2176}, false); err != nil {
		return utils.WrapError("WriteInteger Dl_PRS_StartPRB_r17", err)
	}
	tmp_ResourceList_r17 := utils.NewSequence[*NR_DL_PRS_Resource_r17]([]*NR_DL_PRS_Resource_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPRS_ResourcesPerSet_r17}, false)
	for _, i := range ie.ResourceList_r17 {
		tmp_ResourceList_r17.Value = append(tmp_ResourceList_r17.Value, &i)
	}
	if err = tmp_ResourceList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ResourceList_r17", err)
	}
	if ie.RepFactorAndTimeGap_r17 != nil {
		if err = ie.RepFactorAndTimeGap_r17.Encode(w); err != nil {
			return utils.WrapError("Encode RepFactorAndTimeGap_r17", err)
		}
	}
	return nil
}

func (ie *NR_DL_PRS_PDC_ResourceSet_r17) Decode(r *uper.UperReader) error {
	var err error
	var RepFactorAndTimeGap_r17Present bool
	if RepFactorAndTimeGap_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PeriodicityAndOffset_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PeriodicityAndOffset_r17", err)
	}
	if err = ie.NumSymbols_r17.Decode(r); err != nil {
		return utils.WrapError("Decode NumSymbols_r17", err)
	}
	var tmp_int_Dl_PRS_ResourceBandwidth_r17 int64
	if tmp_int_Dl_PRS_ResourceBandwidth_r17, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 63}, false); err != nil {
		return utils.WrapError("ReadInteger Dl_PRS_ResourceBandwidth_r17", err)
	}
	ie.Dl_PRS_ResourceBandwidth_r17 = tmp_int_Dl_PRS_ResourceBandwidth_r17
	var tmp_int_Dl_PRS_StartPRB_r17 int64
	if tmp_int_Dl_PRS_StartPRB_r17, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 2176}, false); err != nil {
		return utils.WrapError("ReadInteger Dl_PRS_StartPRB_r17", err)
	}
	ie.Dl_PRS_StartPRB_r17 = tmp_int_Dl_PRS_StartPRB_r17
	tmp_ResourceList_r17 := utils.NewSequence[*NR_DL_PRS_Resource_r17]([]*NR_DL_PRS_Resource_r17{}, uper.Constraint{Lb: 1, Ub: maxNrofPRS_ResourcesPerSet_r17}, false)
	fn_ResourceList_r17 := func() *NR_DL_PRS_Resource_r17 {
		return new(NR_DL_PRS_Resource_r17)
	}
	if err = tmp_ResourceList_r17.Decode(r, fn_ResourceList_r17); err != nil {
		return utils.WrapError("Decode ResourceList_r17", err)
	}
	ie.ResourceList_r17 = []NR_DL_PRS_Resource_r17{}
	for _, i := range tmp_ResourceList_r17.Value {
		ie.ResourceList_r17 = append(ie.ResourceList_r17, *i)
	}
	if RepFactorAndTimeGap_r17Present {
		ie.RepFactorAndTimeGap_r17 = new(RepFactorAndTimeGap_r17)
		if err = ie.RepFactorAndTimeGap_r17.Decode(r); err != nil {
			return utils.WrapError("Decode RepFactorAndTimeGap_r17", err)
		}
	}
	return nil
}
