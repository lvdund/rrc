package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type OLPC_SRS_Pos_r16 struct {
	Olpc_SRS_PosBasedOnPRS_Serving_r16      *OLPC_SRS_Pos_r16_olpc_SRS_PosBasedOnPRS_Serving_r16      `optional`
	Olpc_SRS_PosBasedOnSSB_Neigh_r16        *OLPC_SRS_Pos_r16_olpc_SRS_PosBasedOnSSB_Neigh_r16        `optional`
	Olpc_SRS_PosBasedOnPRS_Neigh_r16        *OLPC_SRS_Pos_r16_olpc_SRS_PosBasedOnPRS_Neigh_r16        `optional`
	MaxNumberPathLossEstimatePerServing_r16 *OLPC_SRS_Pos_r16_maxNumberPathLossEstimatePerServing_r16 `optional`
}

func (ie *OLPC_SRS_Pos_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Olpc_SRS_PosBasedOnPRS_Serving_r16 != nil, ie.Olpc_SRS_PosBasedOnSSB_Neigh_r16 != nil, ie.Olpc_SRS_PosBasedOnPRS_Neigh_r16 != nil, ie.MaxNumberPathLossEstimatePerServing_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Olpc_SRS_PosBasedOnPRS_Serving_r16 != nil {
		if err = ie.Olpc_SRS_PosBasedOnPRS_Serving_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Olpc_SRS_PosBasedOnPRS_Serving_r16", err)
		}
	}
	if ie.Olpc_SRS_PosBasedOnSSB_Neigh_r16 != nil {
		if err = ie.Olpc_SRS_PosBasedOnSSB_Neigh_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Olpc_SRS_PosBasedOnSSB_Neigh_r16", err)
		}
	}
	if ie.Olpc_SRS_PosBasedOnPRS_Neigh_r16 != nil {
		if err = ie.Olpc_SRS_PosBasedOnPRS_Neigh_r16.Encode(w); err != nil {
			return utils.WrapError("Encode Olpc_SRS_PosBasedOnPRS_Neigh_r16", err)
		}
	}
	if ie.MaxNumberPathLossEstimatePerServing_r16 != nil {
		if err = ie.MaxNumberPathLossEstimatePerServing_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MaxNumberPathLossEstimatePerServing_r16", err)
		}
	}
	return nil
}

func (ie *OLPC_SRS_Pos_r16) Decode(r *uper.UperReader) error {
	var err error
	var Olpc_SRS_PosBasedOnPRS_Serving_r16Present bool
	if Olpc_SRS_PosBasedOnPRS_Serving_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Olpc_SRS_PosBasedOnSSB_Neigh_r16Present bool
	if Olpc_SRS_PosBasedOnSSB_Neigh_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Olpc_SRS_PosBasedOnPRS_Neigh_r16Present bool
	if Olpc_SRS_PosBasedOnPRS_Neigh_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var MaxNumberPathLossEstimatePerServing_r16Present bool
	if MaxNumberPathLossEstimatePerServing_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Olpc_SRS_PosBasedOnPRS_Serving_r16Present {
		ie.Olpc_SRS_PosBasedOnPRS_Serving_r16 = new(OLPC_SRS_Pos_r16_olpc_SRS_PosBasedOnPRS_Serving_r16)
		if err = ie.Olpc_SRS_PosBasedOnPRS_Serving_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Olpc_SRS_PosBasedOnPRS_Serving_r16", err)
		}
	}
	if Olpc_SRS_PosBasedOnSSB_Neigh_r16Present {
		ie.Olpc_SRS_PosBasedOnSSB_Neigh_r16 = new(OLPC_SRS_Pos_r16_olpc_SRS_PosBasedOnSSB_Neigh_r16)
		if err = ie.Olpc_SRS_PosBasedOnSSB_Neigh_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Olpc_SRS_PosBasedOnSSB_Neigh_r16", err)
		}
	}
	if Olpc_SRS_PosBasedOnPRS_Neigh_r16Present {
		ie.Olpc_SRS_PosBasedOnPRS_Neigh_r16 = new(OLPC_SRS_Pos_r16_olpc_SRS_PosBasedOnPRS_Neigh_r16)
		if err = ie.Olpc_SRS_PosBasedOnPRS_Neigh_r16.Decode(r); err != nil {
			return utils.WrapError("Decode Olpc_SRS_PosBasedOnPRS_Neigh_r16", err)
		}
	}
	if MaxNumberPathLossEstimatePerServing_r16Present {
		ie.MaxNumberPathLossEstimatePerServing_r16 = new(OLPC_SRS_Pos_r16_maxNumberPathLossEstimatePerServing_r16)
		if err = ie.MaxNumberPathLossEstimatePerServing_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MaxNumberPathLossEstimatePerServing_r16", err)
		}
	}
	return nil
}
