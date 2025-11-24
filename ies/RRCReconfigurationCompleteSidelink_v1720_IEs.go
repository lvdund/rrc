package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationCompleteSidelink_v1720_IEs struct {
	Sl_DRX_ConfigReject_v1720 *RRCReconfigurationCompleteSidelink_v1720_IEs_sl_DRX_ConfigReject_v1720 `optional`
	NonCriticalExtension      interface{}                                                             `optional`
}

func (ie *RRCReconfigurationCompleteSidelink_v1720_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Sl_DRX_ConfigReject_v1720 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Sl_DRX_ConfigReject_v1720 != nil {
		if err = ie.Sl_DRX_ConfigReject_v1720.Encode(w); err != nil {
			return utils.WrapError("Encode Sl_DRX_ConfigReject_v1720", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationCompleteSidelink_v1720_IEs) Decode(r *uper.UperReader) error {
	var err error
	var Sl_DRX_ConfigReject_v1720Present bool
	if Sl_DRX_ConfigReject_v1720Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Sl_DRX_ConfigReject_v1720Present {
		ie.Sl_DRX_ConfigReject_v1720 = new(RRCReconfigurationCompleteSidelink_v1720_IEs_sl_DRX_ConfigReject_v1720)
		if err = ie.Sl_DRX_ConfigReject_v1720.Decode(r); err != nil {
			return utils.WrapError("Decode Sl_DRX_ConfigReject_v1720", err)
		}
	}
	return nil
}
