package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationComplete_v1640_IEs struct {
	UplinkTxDirectCurrentTwoCarrierList_r16 *UplinkTxDirectCurrentTwoCarrierList_r16 `optional`
	NonCriticalExtension                    *RRCReconfigurationComplete_v1700_IEs    `optional`
}

func (ie *RRCReconfigurationComplete_v1640_IEs) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.UplinkTxDirectCurrentTwoCarrierList_r16 != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.UplinkTxDirectCurrentTwoCarrierList_r16 != nil {
		if err = ie.UplinkTxDirectCurrentTwoCarrierList_r16.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkTxDirectCurrentTwoCarrierList_r16", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationComplete_v1640_IEs) Decode(r *aper.AperReader) error {
	var err error
	var UplinkTxDirectCurrentTwoCarrierList_r16Present bool
	if UplinkTxDirectCurrentTwoCarrierList_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if UplinkTxDirectCurrentTwoCarrierList_r16Present {
		ie.UplinkTxDirectCurrentTwoCarrierList_r16 = new(UplinkTxDirectCurrentTwoCarrierList_r16)
		if err = ie.UplinkTxDirectCurrentTwoCarrierList_r16.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkTxDirectCurrentTwoCarrierList_r16", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCReconfigurationComplete_v1700_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
