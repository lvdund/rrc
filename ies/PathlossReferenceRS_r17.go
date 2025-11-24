package ies

import (
	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type PathlossReferenceRS_r17 struct {
	PathlossReferenceRS_Id_r17 PathlossReferenceRS_Id_r17                  `madatory`
	ReferenceSignal_r17        PathlossReferenceRS_r17_referenceSignal_r17 `madatory`
	AdditionalPCI_r17          *AdditionalPCIIndex_r17                     `optional`
}

func (ie *PathlossReferenceRS_r17) Encode(w *uper.UperWriter) error {
	var err error
	preambleBits := []bool{ie.AdditionalPCI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.PathlossReferenceRS_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode PathlossReferenceRS_Id_r17", err)
	}
	if err = ie.ReferenceSignal_r17.Encode(w); err != nil {
		return utils.WrapError("Encode ReferenceSignal_r17", err)
	}
	if ie.AdditionalPCI_r17 != nil {
		if err = ie.AdditionalPCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalPCI_r17", err)
		}
	}
	return nil
}

func (ie *PathlossReferenceRS_r17) Decode(r *uper.UperReader) error {
	var err error
	var AdditionalPCI_r17Present bool
	if AdditionalPCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.PathlossReferenceRS_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode PathlossReferenceRS_Id_r17", err)
	}
	if err = ie.ReferenceSignal_r17.Decode(r); err != nil {
		return utils.WrapError("Decode ReferenceSignal_r17", err)
	}
	if AdditionalPCI_r17Present {
		ie.AdditionalPCI_r17 = new(AdditionalPCIIndex_r17)
		if err = ie.AdditionalPCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalPCI_r17", err)
		}
	}
	return nil
}
