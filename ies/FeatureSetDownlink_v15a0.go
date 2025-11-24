package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlink_v15a0 struct {
	SupportedSRS_Resources *SRS_Resources `optional`
}

func (ie *FeatureSetDownlink_v15a0) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.SupportedSRS_Resources != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.SupportedSRS_Resources != nil {
		if err = ie.SupportedSRS_Resources.Encode(w); err != nil {
			return utils.WrapError("Encode SupportedSRS_Resources", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlink_v15a0) Decode(r *uper.UperReader) error {
	var err error
	var SupportedSRS_ResourcesPresent bool
	if SupportedSRS_ResourcesPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if SupportedSRS_ResourcesPresent {
		ie.SupportedSRS_Resources = new(SRS_Resources)
		if err = ie.SupportedSRS_Resources.Decode(r); err != nil {
			return utils.WrapError("Decode SupportedSRS_Resources", err)
		}
	}
	return nil
}
