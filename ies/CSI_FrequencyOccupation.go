package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type CSI_FrequencyOccupation struct {
	StartingRB int64 `lb:0,ub:maxNrofPhysicalResourceBlocks_1,madatory`
	NrofRBs    int64 `lb:24,ub:maxNrofPhysicalResourceBlocksPlus1,madatory`
}

func (ie *CSI_FrequencyOccupation) Encode(w *aper.AperWriter) error {
	var err error
	if err = w.WriteInteger(ie.StartingRB, &aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("WriteInteger StartingRB", err)
	}
	if err = w.WriteInteger(ie.NrofRBs, &aper.Constraint{Lb: 24, Ub: maxNrofPhysicalResourceBlocksPlus1}, false); err != nil {
		return utils.WrapError("WriteInteger NrofRBs", err)
	}
	return nil
}

func (ie *CSI_FrequencyOccupation) Decode(r *aper.AperReader) error {
	var err error
	var tmp_int_StartingRB int64
	if tmp_int_StartingRB, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: maxNrofPhysicalResourceBlocks_1}, false); err != nil {
		return utils.WrapError("ReadInteger StartingRB", err)
	}
	ie.StartingRB = tmp_int_StartingRB
	var tmp_int_NrofRBs int64
	if tmp_int_NrofRBs, err = r.ReadInteger(&aper.Constraint{Lb: 24, Ub: maxNrofPhysicalResourceBlocksPlus1}, false); err != nil {
		return utils.WrapError("ReadInteger NrofRBs", err)
	}
	ie.NrofRBs = tmp_int_NrofRBs
	return nil
}
