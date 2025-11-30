package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MAC_Parameters struct {
	Mac_ParametersCommon   *MAC_ParametersCommon   `optional`
	Mac_ParametersXDD_Diff *MAC_ParametersXDD_Diff `optional`
}

func (ie *MAC_Parameters) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Mac_ParametersCommon != nil, ie.Mac_ParametersXDD_Diff != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mac_ParametersCommon != nil {
		if err = ie.Mac_ParametersCommon.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_ParametersCommon", err)
		}
	}
	if ie.Mac_ParametersXDD_Diff != nil {
		if err = ie.Mac_ParametersXDD_Diff.Encode(w); err != nil {
			return utils.WrapError("Encode Mac_ParametersXDD_Diff", err)
		}
	}
	return nil
}

func (ie *MAC_Parameters) Decode(r *aper.AperReader) error {
	var err error
	var Mac_ParametersCommonPresent bool
	if Mac_ParametersCommonPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Mac_ParametersXDD_DiffPresent bool
	if Mac_ParametersXDD_DiffPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Mac_ParametersCommonPresent {
		ie.Mac_ParametersCommon = new(MAC_ParametersCommon)
		if err = ie.Mac_ParametersCommon.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_ParametersCommon", err)
		}
	}
	if Mac_ParametersXDD_DiffPresent {
		ie.Mac_ParametersXDD_Diff = new(MAC_ParametersXDD_Diff)
		if err = ie.Mac_ParametersXDD_Diff.Decode(r); err != nil {
			return utils.WrapError("Decode Mac_ParametersXDD_Diff", err)
		}
	}
	return nil
}
