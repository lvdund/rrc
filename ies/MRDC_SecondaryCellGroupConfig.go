package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type MRDC_SecondaryCellGroupConfig struct {
	Mrdc_ReleaseAndAdd      *MRDC_SecondaryCellGroupConfig_mrdc_ReleaseAndAdd     `optional`
	Mrdc_SecondaryCellGroup MRDC_SecondaryCellGroupConfig_mrdc_SecondaryCellGroup `madatory`
}

func (ie *MRDC_SecondaryCellGroupConfig) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Mrdc_ReleaseAndAdd != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mrdc_ReleaseAndAdd != nil {
		if err = ie.Mrdc_ReleaseAndAdd.Encode(w); err != nil {
			return utils.WrapError("Encode Mrdc_ReleaseAndAdd", err)
		}
	}
	if err = ie.Mrdc_SecondaryCellGroup.Encode(w); err != nil {
		return utils.WrapError("Encode Mrdc_SecondaryCellGroup", err)
	}
	return nil
}

func (ie *MRDC_SecondaryCellGroupConfig) Decode(r *uper.UperReader) error {
	var err error
	var Mrdc_ReleaseAndAddPresent bool
	if Mrdc_ReleaseAndAddPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Mrdc_ReleaseAndAddPresent {
		ie.Mrdc_ReleaseAndAdd = new(MRDC_SecondaryCellGroupConfig_mrdc_ReleaseAndAdd)
		if err = ie.Mrdc_ReleaseAndAdd.Decode(r); err != nil {
			return utils.WrapError("Decode Mrdc_ReleaseAndAdd", err)
		}
	}
	if err = ie.Mrdc_SecondaryCellGroup.Decode(r); err != nil {
		return utils.WrapError("Decode Mrdc_SecondaryCellGroup", err)
	}
	return nil
}
