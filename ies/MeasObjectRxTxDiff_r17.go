package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MeasObjectRxTxDiff_r17 struct {
	Dl_Ref_r17 *MeasObjectRxTxDiff_r17_dl_Ref_r17 `optional`
}

func (ie *MeasObjectRxTxDiff_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Dl_Ref_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dl_Ref_r17 != nil {
		if err = ie.Dl_Ref_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Dl_Ref_r17", err)
		}
	}
	return nil
}

func (ie *MeasObjectRxTxDiff_r17) Decode(r *aper.AperReader) error {
	var err error
	var Dl_Ref_r17Present bool
	if Dl_Ref_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if Dl_Ref_r17Present {
		ie.Dl_Ref_r17 = new(MeasObjectRxTxDiff_r17_dl_Ref_r17)
		if err = ie.Dl_Ref_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Dl_Ref_r17", err)
		}
	}
	return nil
}
