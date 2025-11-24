package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ReestabNCellInfo struct {
	CellIdentity    CellIdentity   `madatory`
	Key_gNodeB_Star uper.BitString `lb:256,ub:256,madatory`
	ShortMAC_I      ShortMAC_I     `madatory`
}

func (ie *ReestabNCellInfo) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.CellIdentity.Encode(w); err != nil {
		return utils.WrapError("Encode CellIdentity", err)
	}
	if err = w.WriteBitString(ie.Key_gNodeB_Star.Bytes, uint(ie.Key_gNodeB_Star.NumBits), &uper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteBitString Key_gNodeB_Star", err)
	}
	if err = ie.ShortMAC_I.Encode(w); err != nil {
		return utils.WrapError("Encode ShortMAC_I", err)
	}
	return nil
}

func (ie *ReestabNCellInfo) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.CellIdentity.Decode(r); err != nil {
		return utils.WrapError("Decode CellIdentity", err)
	}
	var tmp_bs_Key_gNodeB_Star []byte
	var n_Key_gNodeB_Star uint
	if tmp_bs_Key_gNodeB_Star, n_Key_gNodeB_Star, err = r.ReadBitString(&uper.Constraint{Lb: 256, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadBitString Key_gNodeB_Star", err)
	}
	ie.Key_gNodeB_Star = uper.BitString{
		Bytes:   tmp_bs_Key_gNodeB_Star,
		NumBits: uint64(n_Key_gNodeB_Star),
	}
	if err = ie.ShortMAC_I.Decode(r); err != nil {
		return utils.WrapError("Decode ShortMAC_I", err)
	}
	return nil
}
