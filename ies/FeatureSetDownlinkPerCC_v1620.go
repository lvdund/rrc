package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type FeatureSetDownlinkPerCC_v1620 struct {
	MultiDCI_MultiTRP_r16  *MultiDCI_MultiTRP_r16                                `optional`
	SupportFDM_SchemeB_r16 *FeatureSetDownlinkPerCC_v1620_supportFDM_SchemeB_r16 `optional`
}

func (ie *FeatureSetDownlinkPerCC_v1620) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.MultiDCI_MultiTRP_r16 != nil, ie.SupportFDM_SchemeB_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.MultiDCI_MultiTRP_r16 != nil {
		if err = ie.MultiDCI_MultiTRP_r16.Encode(w); err != nil {
			return utils.WrapError("Encode MultiDCI_MultiTRP_r16", err)
		}
	}
	if ie.SupportFDM_SchemeB_r16 != nil {
		if err = ie.SupportFDM_SchemeB_r16.Encode(w); err != nil {
			return utils.WrapError("Encode SupportFDM_SchemeB_r16", err)
		}
	}
	return nil
}

func (ie *FeatureSetDownlinkPerCC_v1620) Decode(r *uper.UperReader) error {
	var err error
	var MultiDCI_MultiTRP_r16Present bool
	if MultiDCI_MultiTRP_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	var SupportFDM_SchemeB_r16Present bool
	if SupportFDM_SchemeB_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if MultiDCI_MultiTRP_r16Present {
		ie.MultiDCI_MultiTRP_r16 = new(MultiDCI_MultiTRP_r16)
		if err = ie.MultiDCI_MultiTRP_r16.Decode(r); err != nil {
			return utils.WrapError("Decode MultiDCI_MultiTRP_r16", err)
		}
	}
	if SupportFDM_SchemeB_r16Present {
		ie.SupportFDM_SchemeB_r16 = new(FeatureSetDownlinkPerCC_v1620_supportFDM_SchemeB_r16)
		if err = ie.SupportFDM_SchemeB_r16.Decode(r); err != nil {
			return utils.WrapError("Decode SupportFDM_SchemeB_r16", err)
		}
	}
	return nil
}
