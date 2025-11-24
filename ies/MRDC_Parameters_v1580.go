package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_Parameters_v1580 struct {
	DynamicPowerSharingNEDC *MRDC_Parameters_v1580_dynamicPowerSharingNEDC `optional`
}

func (ie *MRDC_Parameters_v1580) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.DynamicPowerSharingNEDC != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DynamicPowerSharingNEDC != nil {
		if err = ie.DynamicPowerSharingNEDC.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicPowerSharingNEDC", err)
		}
	}
	return nil
}

func (ie *MRDC_Parameters_v1580) Decode(r *uper.UperReader) error {
	var err error
	var DynamicPowerSharingNEDCPresent bool
	if DynamicPowerSharingNEDCPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DynamicPowerSharingNEDCPresent {
		ie.DynamicPowerSharingNEDC = new(MRDC_Parameters_v1580_dynamicPowerSharingNEDC)
		if err = ie.DynamicPowerSharingNEDC.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicPowerSharingNEDC", err)
		}
	}
	return nil
}
