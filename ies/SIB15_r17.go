package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SIB15_r17 struct {
	CommonPLMNsWithDisasterCondition_r17 []PLMN_Identity              `lb:1,ub:maxPLMN,optional`
	ApplicableDisasterInfoList_r17       []ApplicableDisasterInfo_r17 `lb:1,ub:maxPLMN,optional`
	LateNonCriticalExtension             *[]byte                      `optional`
}

func (ie *SIB15_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{len(ie.CommonPLMNsWithDisasterCondition_r17) > 0, len(ie.ApplicableDisasterInfoList_r17) > 0, ie.LateNonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if len(ie.CommonPLMNsWithDisasterCondition_r17) > 0 {
		tmp_CommonPLMNsWithDisasterCondition_r17 := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, aper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		for _, i := range ie.CommonPLMNsWithDisasterCondition_r17 {
			tmp_CommonPLMNsWithDisasterCondition_r17.Value = append(tmp_CommonPLMNsWithDisasterCondition_r17.Value, &i)
		}
		if err = tmp_CommonPLMNsWithDisasterCondition_r17.Encode(w); err != nil {
			return utils.WrapError("Encode CommonPLMNsWithDisasterCondition_r17", err)
		}
	}
	if len(ie.ApplicableDisasterInfoList_r17) > 0 {
		tmp_ApplicableDisasterInfoList_r17 := utils.NewSequence[*ApplicableDisasterInfo_r17]([]*ApplicableDisasterInfo_r17{}, aper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		for _, i := range ie.ApplicableDisasterInfoList_r17 {
			tmp_ApplicableDisasterInfoList_r17.Value = append(tmp_ApplicableDisasterInfoList_r17.Value, &i)
		}
		if err = tmp_ApplicableDisasterInfoList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ApplicableDisasterInfoList_r17", err)
		}
	}
	if ie.LateNonCriticalExtension != nil {
		if err = w.WriteOctetString(*ie.LateNonCriticalExtension, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Encode LateNonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *SIB15_r17) Decode(r *aper.AperReader) error {
	var err error
	var CommonPLMNsWithDisasterCondition_r17Present bool
	if CommonPLMNsWithDisasterCondition_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var ApplicableDisasterInfoList_r17Present bool
	if ApplicableDisasterInfoList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var LateNonCriticalExtensionPresent bool
	if LateNonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if CommonPLMNsWithDisasterCondition_r17Present {
		tmp_CommonPLMNsWithDisasterCondition_r17 := utils.NewSequence[*PLMN_Identity]([]*PLMN_Identity{}, aper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		fn_CommonPLMNsWithDisasterCondition_r17 := func() *PLMN_Identity {
			return new(PLMN_Identity)
		}
		if err = tmp_CommonPLMNsWithDisasterCondition_r17.Decode(r, fn_CommonPLMNsWithDisasterCondition_r17); err != nil {
			return utils.WrapError("Decode CommonPLMNsWithDisasterCondition_r17", err)
		}
		ie.CommonPLMNsWithDisasterCondition_r17 = []PLMN_Identity{}
		for _, i := range tmp_CommonPLMNsWithDisasterCondition_r17.Value {
			ie.CommonPLMNsWithDisasterCondition_r17 = append(ie.CommonPLMNsWithDisasterCondition_r17, *i)
		}
	}
	if ApplicableDisasterInfoList_r17Present {
		tmp_ApplicableDisasterInfoList_r17 := utils.NewSequence[*ApplicableDisasterInfo_r17]([]*ApplicableDisasterInfo_r17{}, aper.Constraint{Lb: 1, Ub: maxPLMN}, false)
		fn_ApplicableDisasterInfoList_r17 := func() *ApplicableDisasterInfo_r17 {
			return new(ApplicableDisasterInfo_r17)
		}
		if err = tmp_ApplicableDisasterInfoList_r17.Decode(r, fn_ApplicableDisasterInfoList_r17); err != nil {
			return utils.WrapError("Decode ApplicableDisasterInfoList_r17", err)
		}
		ie.ApplicableDisasterInfoList_r17 = []ApplicableDisasterInfo_r17{}
		for _, i := range tmp_ApplicableDisasterInfoList_r17.Value {
			ie.ApplicableDisasterInfoList_r17 = append(ie.ApplicableDisasterInfoList_r17, *i)
		}
	}
	if LateNonCriticalExtensionPresent {
		var tmp_os_LateNonCriticalExtension []byte
		if tmp_os_LateNonCriticalExtension, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return utils.WrapError("Decode LateNonCriticalExtension", err)
		}
		ie.LateNonCriticalExtension = &tmp_os_LateNonCriticalExtension
	}
	return nil
}
