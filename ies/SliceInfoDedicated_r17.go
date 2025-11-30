package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SliceInfoDedicated_r17 struct {
	Nsag_IdentityInfo_r17               NSAG_IdentityInfo_r17       `madatory`
	Nsag_CellReselectionPriority_r17    *CellReselectionPriority    `optional`
	Nsag_CellReselectionSubPriority_r17 *CellReselectionSubPriority `optional`
}

func (ie *SliceInfoDedicated_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Nsag_CellReselectionPriority_r17 != nil, ie.Nsag_CellReselectionSubPriority_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Nsag_IdentityInfo_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Nsag_IdentityInfo_r17", err)
	}
	if ie.Nsag_CellReselectionPriority_r17 != nil {
		if err = ie.Nsag_CellReselectionPriority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Nsag_CellReselectionPriority_r17", err)
		}
	}
	if ie.Nsag_CellReselectionSubPriority_r17 != nil {
		if err = ie.Nsag_CellReselectionSubPriority_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Nsag_CellReselectionSubPriority_r17", err)
		}
	}
	return nil
}

func (ie *SliceInfoDedicated_r17) Decode(r *aper.AperReader) error {
	var err error
	var Nsag_CellReselectionPriority_r17Present bool
	if Nsag_CellReselectionPriority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Nsag_CellReselectionSubPriority_r17Present bool
	if Nsag_CellReselectionSubPriority_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Nsag_IdentityInfo_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Nsag_IdentityInfo_r17", err)
	}
	if Nsag_CellReselectionPriority_r17Present {
		ie.Nsag_CellReselectionPriority_r17 = new(CellReselectionPriority)
		if err = ie.Nsag_CellReselectionPriority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Nsag_CellReselectionPriority_r17", err)
		}
	}
	if Nsag_CellReselectionSubPriority_r17Present {
		ie.Nsag_CellReselectionSubPriority_r17 = new(CellReselectionSubPriority)
		if err = ie.Nsag_CellReselectionSubPriority_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Nsag_CellReselectionSubPriority_r17", err)
		}
	}
	return nil
}
