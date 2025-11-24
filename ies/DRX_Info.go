package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type DRX_Info struct {
	Drx_LongCycleStartOffset DRX_Info_drx_LongCycleStartOffset `lb:0,ub:9,madatory`
	ShortDRX                 *DRX_Info_shortDRX                `lb:1,ub:16,optional`
}

func (ie *DRX_Info) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.ShortDRX != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Drx_LongCycleStartOffset.Encode(w); err != nil {
		return utils.WrapError("Encode Drx_LongCycleStartOffset", err)
	}
	if ie.ShortDRX != nil {
		if err = ie.ShortDRX.Encode(w); err != nil {
			return utils.WrapError("Encode ShortDRX", err)
		}
	}
	return nil
}

func (ie *DRX_Info) Decode(r *uper.UperReader) error {
	var err error
	var ShortDRXPresent bool
	if ShortDRXPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Drx_LongCycleStartOffset.Decode(r); err != nil {
		return utils.WrapError("Decode Drx_LongCycleStartOffset", err)
	}
	if ShortDRXPresent {
		ie.ShortDRX = new(DRX_Info_shortDRX)
		if err = ie.ShortDRX.Decode(r); err != nil {
			return utils.WrapError("Decode ShortDRX", err)
		}
	}
	return nil
}
