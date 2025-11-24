package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CellAccessRelatedInfo_EUTRA_EPC struct {
	Plmn_IdentityList_eutra_epc PLMN_IdentityList_EUTRA_EPC `madatory`
	TrackingAreaCode_eutra_epc  uper.BitString              `lb:16,ub:16,madatory`
	CellIdentity_eutra_epc      uper.BitString              `lb:28,ub:28,madatory`
}

func (ie *CellAccessRelatedInfo_EUTRA_EPC) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.Plmn_IdentityList_eutra_epc.Encode(w); err != nil {
		return utils.WrapError("Encode Plmn_IdentityList_eutra_epc", err)
	}
	if err = w.WriteBitString(ie.TrackingAreaCode_eutra_epc.Bytes, uint(ie.TrackingAreaCode_eutra_epc.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteBitString TrackingAreaCode_eutra_epc", err)
	}
	if err = w.WriteBitString(ie.CellIdentity_eutra_epc.Bytes, uint(ie.CellIdentity_eutra_epc.NumBits), &uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
		return utils.WrapError("WriteBitString CellIdentity_eutra_epc", err)
	}
	return nil
}

func (ie *CellAccessRelatedInfo_EUTRA_EPC) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.Plmn_IdentityList_eutra_epc.Decode(r); err != nil {
		return utils.WrapError("Decode Plmn_IdentityList_eutra_epc", err)
	}
	var tmp_bs_TrackingAreaCode_eutra_epc []byte
	var n_TrackingAreaCode_eutra_epc uint
	if tmp_bs_TrackingAreaCode_eutra_epc, n_TrackingAreaCode_eutra_epc, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadBitString TrackingAreaCode_eutra_epc", err)
	}
	ie.TrackingAreaCode_eutra_epc = uper.BitString{
		Bytes:   tmp_bs_TrackingAreaCode_eutra_epc,
		NumBits: uint64(n_TrackingAreaCode_eutra_epc),
	}
	var tmp_bs_CellIdentity_eutra_epc []byte
	var n_CellIdentity_eutra_epc uint
	if tmp_bs_CellIdentity_eutra_epc, n_CellIdentity_eutra_epc, err = r.ReadBitString(&uper.Constraint{Lb: 28, Ub: 28}, false); err != nil {
		return utils.WrapError("ReadBitString CellIdentity_eutra_epc", err)
	}
	ie.CellIdentity_eutra_epc = uper.BitString{
		Bytes:   tmp_bs_CellIdentity_eutra_epc,
		NumBits: uint64(n_CellIdentity_eutra_epc),
	}
	return nil
}
