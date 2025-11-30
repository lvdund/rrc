package ies

import (
	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type MPE_Resource_r17 struct {
	Mpe_ResourceId_r17      MPE_ResourceId_r17                       `madatory`
	Cell_r17                *ServCellIndex                           `optional`
	AdditionalPCI_r17       *AdditionalPCIIndex_r17                  `optional`
	Mpe_ReferenceSignal_r17 MPE_Resource_r17_mpe_ReferenceSignal_r17 `madatory`
}

func (ie *MPE_Resource_r17) Encode(w *aper.AperWriter) error {
	var err error
	preambleBits := []bool{ie.Cell_r17 != nil, ie.AdditionalPCI_r17 != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Mpe_ResourceId_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mpe_ResourceId_r17", err)
	}
	if ie.Cell_r17 != nil {
		if err = ie.Cell_r17.Encode(w); err != nil {
			return utils.WrapError("Encode Cell_r17", err)
		}
	}
	if ie.AdditionalPCI_r17 != nil {
		if err = ie.AdditionalPCI_r17.Encode(w); err != nil {
			return utils.WrapError("Encode AdditionalPCI_r17", err)
		}
	}
	if err = ie.Mpe_ReferenceSignal_r17.Encode(w); err != nil {
		return utils.WrapError("Encode Mpe_ReferenceSignal_r17", err)
	}
	return nil
}

func (ie *MPE_Resource_r17) Decode(r *aper.AperReader) error {
	var err error
	var Cell_r17Present bool
	if Cell_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	var AdditionalPCI_r17Present bool
	if AdditionalPCI_r17Present, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Mpe_ResourceId_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mpe_ResourceId_r17", err)
	}
	if Cell_r17Present {
		ie.Cell_r17 = new(ServCellIndex)
		if err = ie.Cell_r17.Decode(r); err != nil {
			return utils.WrapError("Decode Cell_r17", err)
		}
	}
	if AdditionalPCI_r17Present {
		ie.AdditionalPCI_r17 = new(AdditionalPCIIndex_r17)
		if err = ie.AdditionalPCI_r17.Decode(r); err != nil {
			return utils.WrapError("Decode AdditionalPCI_r17", err)
		}
	}
	if err = ie.Mpe_ReferenceSignal_r17.Decode(r); err != nil {
		return utils.WrapError("Decode Mpe_ReferenceSignal_r17", err)
	}
	return nil
}
