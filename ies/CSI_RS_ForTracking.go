package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CSI_RS_ForTracking struct {
	MaxBurstLength                   int64 `lb:1,ub:2,madatory`
	MaxSimultaneousResourceSetsPerCC int64 `lb:1,ub:8,madatory`
	MaxConfiguredResourceSetsPerCC   int64 `lb:1,ub:64,madatory`
	MaxConfiguredResourceSetsAllCC   int64 `lb:1,ub:256,madatory`
}

func (ie *CSI_RS_ForTracking) Encode(w *uper.UperWriter) error {
	var err error
	if err = w.WriteInteger(ie.MaxBurstLength, &uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("WriteInteger MaxBurstLength", err)
	}
	if err = w.WriteInteger(ie.MaxSimultaneousResourceSetsPerCC, &uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("WriteInteger MaxSimultaneousResourceSetsPerCC", err)
	}
	if err = w.WriteInteger(ie.MaxConfiguredResourceSetsPerCC, &uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("WriteInteger MaxConfiguredResourceSetsPerCC", err)
	}
	if err = w.WriteInteger(ie.MaxConfiguredResourceSetsAllCC, &uper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return utils.WrapError("WriteInteger MaxConfiguredResourceSetsAllCC", err)
	}
	return nil
}

func (ie *CSI_RS_ForTracking) Decode(r *uper.UperReader) error {
	var err error
	var tmp_int_MaxBurstLength int64
	if tmp_int_MaxBurstLength, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 2}, false); err != nil {
		return utils.WrapError("ReadInteger MaxBurstLength", err)
	}
	ie.MaxBurstLength = tmp_int_MaxBurstLength
	var tmp_int_MaxSimultaneousResourceSetsPerCC int64
	if tmp_int_MaxSimultaneousResourceSetsPerCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 8}, false); err != nil {
		return utils.WrapError("ReadInteger MaxSimultaneousResourceSetsPerCC", err)
	}
	ie.MaxSimultaneousResourceSetsPerCC = tmp_int_MaxSimultaneousResourceSetsPerCC
	var tmp_int_MaxConfiguredResourceSetsPerCC int64
	if tmp_int_MaxConfiguredResourceSetsPerCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 64}, false); err != nil {
		return utils.WrapError("ReadInteger MaxConfiguredResourceSetsPerCC", err)
	}
	ie.MaxConfiguredResourceSetsPerCC = tmp_int_MaxConfiguredResourceSetsPerCC
	var tmp_int_MaxConfiguredResourceSetsAllCC int64
	if tmp_int_MaxConfiguredResourceSetsAllCC, err = r.ReadInteger(&uper.Constraint{Lb: 1, Ub: 256}, false); err != nil {
		return utils.WrapError("ReadInteger MaxConfiguredResourceSetsAllCC", err)
	}
	ie.MaxConfiguredResourceSetsAllCC = tmp_int_MaxConfiguredResourceSetsAllCC
	return nil
}
