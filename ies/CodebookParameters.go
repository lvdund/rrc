package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type CodebookParameters struct {
	Type1               *CodebookParameters_type1               `lb:1,ub:maxNrofCSI_RS_Resources,optional`
	Type2               *CodebookParameters_type2               `lb:1,ub:maxNrofCSI_RS_Resources,optional`
	Type2_PortSelection *CodebookParameters_type2_PortSelection `lb:1,ub:maxNrofCSI_RS_Resources,optional`
}

func (ie *CodebookParameters) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.Type1 != nil, ie.Type2 != nil, ie.Type2_PortSelection != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Type1 != nil {
		if err = ie.Type1.Encode(w); err != nil {
			return utils.WrapError("Encode Type1", err)
		}
	}
	if ie.Type2 != nil {
		if err = ie.Type2.Encode(w); err != nil {
			return utils.WrapError("Encode Type2", err)
		}
	}
	if ie.Type2_PortSelection != nil {
		if err = ie.Type2_PortSelection.Encode(w); err != nil {
			return utils.WrapError("Encode Type2_PortSelection", err)
		}
	}
	return nil
}

func (ie *CodebookParameters) Decode(r *uper.UperReader) error {
	var err error
	var Type1Present bool
	if Type1Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type2Present bool
	if Type2Present, err = r.ReadBool(); err != nil {
		return err
	}
	var Type2_PortSelectionPresent bool
	if Type2_PortSelectionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Type1Present {
		ie.Type1 = new(CodebookParameters_type1)
		if err = ie.Type1.Decode(r); err != nil {
			return utils.WrapError("Decode Type1", err)
		}
	}
	if Type2Present {
		ie.Type2 = new(CodebookParameters_type2)
		if err = ie.Type2.Decode(r); err != nil {
			return utils.WrapError("Decode Type2", err)
		}
	}
	if Type2_PortSelectionPresent {
		ie.Type2_PortSelection = new(CodebookParameters_type2_PortSelection)
		if err = ie.Type2_PortSelection.Decode(r); err != nil {
			return utils.WrapError("Decode Type2_PortSelection", err)
		}
	}
	return nil
}
