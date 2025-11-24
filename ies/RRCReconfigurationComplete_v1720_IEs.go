package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationComplete_v1720_IEs struct {
	UplinkTxDirectCurrentMoreCarrierList_r17 *UplinkTxDirectCurrentMoreCarrierList_r17 `optional`
	NonCriticalExtension                     interface{}                               `optional`
}

func (ie *RRCReconfigurationComplete_v1720_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.UplinkTxDirectCurrentMoreCarrierList_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.UplinkTxDirectCurrentMoreCarrierList_r17 != nil {
		if err = ie.UplinkTxDirectCurrentMoreCarrierList_r17.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkTxDirectCurrentMoreCarrierList_r17", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationComplete_v1720_IEs) Decode(r *uper.UperReader) error {
	var err error
	var UplinkTxDirectCurrentMoreCarrierList_r17Present bool
	if UplinkTxDirectCurrentMoreCarrierList_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if UplinkTxDirectCurrentMoreCarrierList_r17Present {
		ie.UplinkTxDirectCurrentMoreCarrierList_r17 = new(UplinkTxDirectCurrentMoreCarrierList_r17)
		if err = ie.UplinkTxDirectCurrentMoreCarrierList_r17.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkTxDirectCurrentMoreCarrierList_r17", err)
		}
	}
	return nil
}
