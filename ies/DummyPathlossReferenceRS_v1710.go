package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type DummyPathlossReferenceRS_v1710 struct {
	Pusch_PathlossReferenceRS_Id_r17 PUSCH_PathlossReferenceRS_Id_r17 `madatory`
	AdditionalPCI_r17                *AdditionalPCIIndex_r17          `optional`
}

func (ie *DummyPathlossReferenceRS_v1710) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.AdditionalPCI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Pusch_PathlossReferenceRS_Id_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Pusch_PathlossReferenceRS_Id_r17", err)
	}
	if ie.AdditionalPCI_r17 != nil {
		if err = ie.AdditionalPCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalPCI_r17", err)
		}
	}
	return nil
}

func (ie *DummyPathlossReferenceRS_v1710) Decode(r *aper.AperReader) error {
	var err error
	var AdditionalPCI_r17Present bool
	if AdditionalPCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Pusch_PathlossReferenceRS_Id_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Pusch_PathlossReferenceRS_Id_r17", err)
	}
	if AdditionalPCI_r17Present {
		ie.AdditionalPCI_r17 = new(AdditionalPCIIndex_r17)
		if err = ie.AdditionalPCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalPCI_r17", err)
		}
	}
	return nil
}
