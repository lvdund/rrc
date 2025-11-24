package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCSystemInfoRequest_IEs struct {
	Requested_SI_List uper.BitString `lb:maxSI_Message,ub:maxSI_Message,madatory`
	Spare             uper.BitString `lb:12,ub:12,madatory`
}

func (ie *RRCSystemInfoRequest_IEs) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteBitString(ie.Requested_SI_List.Bytes, uint(ie.Requested_SI_List.NumBits), &uper.Constraint{Lb: maxSI_Message, Ub: maxSI_Message}, false); err != nil {
		return utils.WrapError("WriteBitString Requested_SI_List", err)
	}
	if err = w.WriteBitString(ie.Spare.Bytes, uint(ie.Spare.NumBits), &uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
		return utils.WrapError("WriteBitString Spare", err)
	}
	return nil
}

func (ie *RRCSystemInfoRequest_IEs) Decode(r *uper.UperReader) error {
	var err error
	var tmp_bs_Requested_SI_List []byte
	var n_Requested_SI_List uint
	if tmp_bs_Requested_SI_List, n_Requested_SI_List, err = r.ReadBitString(&uper.Constraint{Lb: maxSI_Message, Ub: maxSI_Message}, false); err != nil {
		return utils.WrapError("ReadBitString Requested_SI_List", err)
	}
	ie.Requested_SI_List = uper.BitString{
		Bytes:   tmp_bs_Requested_SI_List,
		NumBits: uint64(n_Requested_SI_List),
	}
	var tmp_bs_Spare []byte
	var n_Spare uint
	if tmp_bs_Spare, n_Spare, err = r.ReadBitString(&uper.Constraint{Lb: 12, Ub: 12}, false); err != nil {
		return utils.WrapError("ReadBitString Spare", err)
	}
	ie.Spare = uper.BitString{
		Bytes:   tmp_bs_Spare,
		NumBits: uint64(n_Spare),
	}
	return nil
}
