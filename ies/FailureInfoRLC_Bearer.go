package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FailureInfoRLC_Bearer struct {
	CellGroupId            CellGroupId                       `madatory`
	LogicalChannelIdentity LogicalChannelIdentity            `madatory`
	FailureType            FailureInfoRLC_Bearer_failureType `madatory`
}

func (ie *FailureInfoRLC_Bearer) Encode(w *uper.UperWriter) error {
	var err error
	if err = ie.CellGroupId.Encode(w); err != nil {
		return utils.WrapError("Encode CellGroupId", err)
	}
	if err = ie.LogicalChannelIdentity.Encode(w); err != nil {
		return utils.WrapError("Encode LogicalChannelIdentity", err)
	}
	if err = ie.FailureType.Encode(w); err != nil {
		return utils.WrapError("Encode FailureType", err)
	}
	return nil
}

func (ie *FailureInfoRLC_Bearer) Decode(r *uper.UperReader) error {
	var err error
	if err = ie.CellGroupId.Decode(r); err != nil {
		return utils.WrapError("Decode CellGroupId", err)
	}
	if err = ie.LogicalChannelIdentity.Decode(r); err != nil {
		return utils.WrapError("Decode LogicalChannelIdentity", err)
	}
	if err = ie.FailureType.Decode(r); err != nil {
		return utils.WrapError("Decode FailureType", err)
	}
	return nil
}
