package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type ServingCellConfigCommonSIB_ssb_PositionsInBurst struct {
	InOneGroup    uper.BitString  `lb:8,ub:8,madatory`
	GroupPresence *uper.BitString `lb:8,ub:8,optional`
}

func (ie *ServingCellConfigCommonSIB_ssb_PositionsInBurst) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.GroupPresence != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = w.WriteBitString(ie.InOneGroup.Bytes, uint(ie.InOneGroup.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteBitString InOneGroup", err)
	}
	if ie.GroupPresence != nil {
		if err = w.WriteBitString(ie.GroupPresence.Bytes, uint(ie.GroupPresence.NumBits), &uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Encode GroupPresence", err)
		}
	}
	return nil
}

func (ie *ServingCellConfigCommonSIB_ssb_PositionsInBurst) Decode(r *uper.UperReader) error {
	var err error
	var GroupPresencePresent bool
	if GroupPresencePresent, err = r.ReadBool(); err != nil {
		return err
	}
	var tmp_bs_InOneGroup []byte
	var n_InOneGroup uint
	if tmp_bs_InOneGroup, n_InOneGroup, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadBitString InOneGroup", err)
	}
	ie.InOneGroup = uper.BitString{
		Bytes:   tmp_bs_InOneGroup,
		NumBits: uint64(n_InOneGroup),
	}
	if GroupPresencePresent {
		var tmp_bs_GroupPresence []byte
		var n_GroupPresence uint
		if tmp_bs_GroupPresence, n_GroupPresence, err = r.ReadBitString(&uper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
			return utils.WrapError("Decode GroupPresence", err)
		}
		tmp_bitstring := uper.BitString{
			Bytes:   tmp_bs_GroupPresence,
			NumBits: uint64(n_GroupPresence),
		}
		ie.GroupPresence = &tmp_bitstring
	}
	return nil
}
