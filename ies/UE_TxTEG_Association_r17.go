package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_TxTEG_Association_r17 struct {
	Ue_TxTEG_ID_r17                     int64                   `lb:0,ub:maxNrOfTxTEG_ID_1_r17,madatory`
	Nr_TimeStamp_r17                    NR_TimeStamp_r17        `madatory`
	AssociatedSRS_PosResourceIdList_r17 []SRS_PosResourceId_r16 `lb:1,ub:maxNrofSRS_PosResources_r16,madatory`
	ServCellId_r17                      *ServCellIndex          `optional`
}

func (ie *UE_TxTEG_Association_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.ServCellId_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteInteger(ie.Ue_TxTEG_ID_r17, &aper.Constraint{Lb: 0, Ub: maxNrOfTxTEG_ID_1_r17}, false); err != nil {
		return utils.WrapError("WriteInteger Ue_TxTEG_ID_r17", err)
	}
	if err = ie.Nr_TimeStamp_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Nr_TimeStamp_r17", err)
	}
	tmp_AssociatedSRS_PosResourceIdList_r17 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
	for _, i := range ie.AssociatedSRS_PosResourceIdList_r17 {
		tmp_AssociatedSRS_PosResourceIdList_r17.Value = append(tmp_AssociatedSRS_PosResourceIdList_r17.Value, &i)
	}
	if err = tmp_AssociatedSRS_PosResourceIdList_r17.Encode(w); err != nil {
		return utils.WrapError("Encode AssociatedSRS_PosResourceIdList_r17", err)
	}
	if ie.ServCellId_r17 != nil {
		if err = ie.ServCellId_r17.Encode(w); err != nil {
			return utils.WrapError("Encode ServCellId_r17", err)
		}
	}
	return nil
}

func (ie *UE_TxTEG_Association_r17) Decode(r *aper.AperReader) error {
	var err error
	var ServCellId_r17Present bool
	if ServCellId_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_int_Ue_TxTEG_ID_r17 int64
	if tmp_int_Ue_TxTEG_ID_r17, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrOfTxTEG_ID_1_r17}, false); err != nil {
		return utils.WrapError("ReadInteger Ue_TxTEG_ID_r17", err)
	}
	ie.Ue_TxTEG_ID_r17 = tmp_int_Ue_TxTEG_ID_r17
	if err = ie.Nr_TimeStamp_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Nr_TimeStamp_r17", err)
	}
	tmp_AssociatedSRS_PosResourceIdList_r17 := utils.NewSequence[*SRS_PosResourceId_r16]([]*SRS_PosResourceId_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofSRS_PosResources_r16}, false)
	fn_AssociatedSRS_PosResourceIdList_r17 := func() *SRS_PosResourceId_r16 {
		return new(SRS_PosResourceId_r16)
	}
	if err = tmp_AssociatedSRS_PosResourceIdList_r17.Decode(r, fn_AssociatedSRS_PosResourceIdList_r17); err != nil {
		return utils.WrapError("Decode AssociatedSRS_PosResourceIdList_r17", err)
	}
	ie.AssociatedSRS_PosResourceIdList_r17 = []SRS_PosResourceId_r16{}
	for _, i := range tmp_AssociatedSRS_PosResourceIdList_r17.Value {
		ie.AssociatedSRS_PosResourceIdList_r17 = append(ie.AssociatedSRS_PosResourceIdList_r17, *i)
	}
	if ServCellId_r17Present {
		ie.ServCellId_r17 = new(ServCellIndex)
		if err = ie.ServCellId_r17.Decode(r); err != nil {
			return utils.WrapError("Decode ServCellId_r17", err)
		}
	}
	return nil
}
