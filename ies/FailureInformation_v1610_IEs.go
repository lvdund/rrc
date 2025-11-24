package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FailureInformation_v1610_IEs struct {
	FailureInfoDAPS_r16  *FailureInfoDAPS_r16 `optional`
	NonCriticalExtension interface{}          `optional`
}

func (ie *FailureInformation_v1610_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.FailureInfoDAPS_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.FailureInfoDAPS_r16 != nil {
		if err = ie.FailureInfoDAPS_r16.Encode(w); err != nil {
			return utils.WrapError("Encode FailureInfoDAPS_r16", err)
		}
	}
	return nil
}

func (ie *FailureInformation_v1610_IEs) Decode(r *uper.UperReader) error {
	var err error
	var FailureInfoDAPS_r16Present bool
	if FailureInfoDAPS_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if FailureInfoDAPS_r16Present {
		ie.FailureInfoDAPS_r16 = new(FailureInfoDAPS_r16)
		if err = ie.FailureInfoDAPS_r16.Decode(r); err != nil {
			return utils.WrapError("Decode FailureInfoDAPS_r16", err)
		}
	}
	return nil
}
