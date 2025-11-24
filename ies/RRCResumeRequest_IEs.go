package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCResumeRequest_IEs struct {
	ResumeIdentity ShortI_RNTI_Value `madatory`
	ResumeMAC_I    uper.BitString    `lb:16,ub:16,madatory`
	ResumeCause    ResumeCause       `madatory`
	Spare          uper.BitString    `lb:1,ub:1,madatory`
}

func (ie *RRCResumeRequest_IEs) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.ResumeIdentity.Encode(w); err != nil {
		return utils.WrapError("Encode ResumeIdentity", err)
	}
	if err = w.WriteBitString(ie.ResumeMAC_I.Bytes, uint(ie.ResumeMAC_I.NumBits), &uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("WriteBitString ResumeMAC_I", err)
	}
	if err = ie.ResumeCause.Encode(w); err != nil {
		return utils.WrapError("Encode ResumeCause", err)
	}
	if err = w.WriteBitString(ie.Spare.Bytes, uint(ie.Spare.NumBits), &uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("WriteBitString Spare", err)
	}
	return nil
}

func (ie *RRCResumeRequest_IEs) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.ResumeIdentity.Decode(r); err != nil {
		return utils.WrapError("Decode ResumeIdentity", err)
	}
	var tmp_bs_ResumeMAC_I []byte
	var n_ResumeMAC_I uint
	if tmp_bs_ResumeMAC_I, n_ResumeMAC_I, err = r.ReadBitString(&uper.Constraint{Lb: 16, Ub: 16}, false); err != nil {
		return utils.WrapError("ReadBitString ResumeMAC_I", err)
	}
	ie.ResumeMAC_I = uper.BitString{
		Bytes:   tmp_bs_ResumeMAC_I,
		NumBits: uint64(n_ResumeMAC_I),
	}
	if err = ie.ResumeCause.Decode(r); err != nil {
		return utils.WrapError("Decode ResumeCause", err)
	}
	var tmp_bs_Spare []byte
	var n_Spare uint
	if tmp_bs_Spare, n_Spare, err = r.ReadBitString(&uper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return utils.WrapError("ReadBitString Spare", err)
	}
	ie.Spare = uper.BitString{
		Bytes:   tmp_bs_Spare,
		NumBits: uint64(n_Spare),
	}
	return nil
}
