package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type RRCReconfigurationComplete_v1530_IEs struct {
	UplinkTxDirectCurrentList *UplinkTxDirectCurrentList            `optional`
	NonCriticalExtension      *RRCReconfigurationComplete_v1560_IEs `optional`
}

func (ie *RRCReconfigurationComplete_v1530_IEs) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.UplinkTxDirectCurrentList != nil, ie.NonCriticalExtension != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.UplinkTxDirectCurrentList != nil {
		if err = ie.UplinkTxDirectCurrentList.Encode(w); err != nil {
			return utils.WrapError("Encode UplinkTxDirectCurrentList", err)
		}
	}
	if ie.NonCriticalExtension != nil {
		if err = ie.NonCriticalExtension.Encode(w); err != nil {
			return utils.WrapError("Encode NonCriticalExtension", err)
		}
	}
	return nil
}

func (ie *RRCReconfigurationComplete_v1530_IEs) Decode(r *uper.UperReader) error {
	var err error
	var UplinkTxDirectCurrentListPresent bool
	if UplinkTxDirectCurrentListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var NonCriticalExtensionPresent bool
	if NonCriticalExtensionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if UplinkTxDirectCurrentListPresent {
		ie.UplinkTxDirectCurrentList = new(UplinkTxDirectCurrentList)
		if err = ie.UplinkTxDirectCurrentList.Decode(r); err != nil {
			return utils.WrapError("Decode UplinkTxDirectCurrentList", err)
		}
	}
	if NonCriticalExtensionPresent {
		ie.NonCriticalExtension = new(RRCReconfigurationComplete_v1560_IEs)
		if err = ie.NonCriticalExtension.Decode(r); err != nil {
			return utils.WrapError("Decode NonCriticalExtension", err)
		}
	}
	return nil
}
