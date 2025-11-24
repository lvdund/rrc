package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type SRS_ResourceConfigCLI_r16 struct {
	Srs_Resource_r16     SRS_Resource      `madatory`
	Srs_SCS_r16          SubcarrierSpacing `madatory`
	RefServCellIndex_r16 *ServCellIndex    `optional`
	RefBWP_r16           BWP_Id            `madatory`
}

func (ie *SRS_ResourceConfigCLI_r16) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.RefServCellIndex_r16 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Srs_Resource_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_Resource_r16", err)
	}
	if err = ie.Srs_SCS_r16.Encode(w); err != nil {
		return utils.WrapError("Encode Srs_SCS_r16", err)
	}
	if ie.RefServCellIndex_r16 != nil {
		if err = ie.RefServCellIndex_r16.Encode(w); err != nil {
			return utils.WrapError("Encode RefServCellIndex_r16", err)
		}
	}
	if err = ie.RefBWP_r16.Encode(w); err != nil {
		return utils.WrapError("Encode RefBWP_r16", err)
	}
	return nil
}

func (ie *SRS_ResourceConfigCLI_r16) Decode(r *uper.UperReader) error {
	var err error
	var RefServCellIndex_r16Present bool
	if RefServCellIndex_r16Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Srs_Resource_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Srs_Resource_r16", err)
	}
	if err = ie.Srs_SCS_r16.Decode(r); err != nil {
		return utils.WrapError("Decode Srs_SCS_r16", err)
	}
	if RefServCellIndex_r16Present {
		ie.RefServCellIndex_r16 = new(ServCellIndex)
		if err = ie.RefServCellIndex_r16.Decode(r); err != nil {
			return utils.WrapError("Decode RefServCellIndex_r16", err)
		}
	}
	if err = ie.RefBWP_r16.Decode(r); err != nil {
		return utils.WrapError("Decode RefBWP_r16", err)
	}
	return nil
}
