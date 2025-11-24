package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersFR2 struct {
	Dummy                                               *Phy_ParametersFR2_dummy                                               `optional`
	Pdsch_RE_MappingFR2_PerSymbol                       *Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSymbol                       `optional`
	PCell_FR2                                           *Phy_ParametersFR2_pCell_FR2                                           `optional,ext-1`
	Pdsch_RE_MappingFR2_PerSlot                         *Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot                         `optional,ext-1`
	DefaultSpatialRelationPathlossRS_r16                *Phy_ParametersFR2_defaultSpatialRelationPathlossRS_r16                `optional,ext-2`
	SpatialRelationUpdateAP_SRS_r16                     *Phy_ParametersFR2_spatialRelationUpdateAP_SRS_r16                     `optional,ext-2`
	MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 *Phy_ParametersFR2_maxNumberSRS_PosSpatialRelationsAllServingCells_r16 `optional,ext-2`
}

func (ie *Phy_ParametersFR2) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.PCell_FR2 != nil || ie.Pdsch_RE_MappingFR2_PerSlot != nil || ie.DefaultSpatialRelationPathlossRS_r16 != nil || ie.SpatialRelationUpdateAP_SRS_r16 != nil || ie.MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 != nil
	preambleBits := []bool{hasExtensions, ie.Dummy != nil, ie.Pdsch_RE_MappingFR2_PerSymbol != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Dummy != nil {
		if err = ie.Dummy.Encode(w); err != nil {
			return utils.WrapError("Encode Dummy", err)
		}
	}
	if ie.Pdsch_RE_MappingFR2_PerSymbol != nil {
		if err = ie.Pdsch_RE_MappingFR2_PerSymbol.Encode(w); err != nil {
			return utils.WrapError("Encode Pdsch_RE_MappingFR2_PerSymbol", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.PCell_FR2 != nil || ie.Pdsch_RE_MappingFR2_PerSlot != nil, ie.DefaultSpatialRelationPathlossRS_r16 != nil || ie.SpatialRelationUpdateAP_SRS_r16 != nil || ie.MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap Phy_ParametersFR2", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.PCell_FR2 != nil, ie.Pdsch_RE_MappingFR2_PerSlot != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode PCell_FR2 optional
			if ie.PCell_FR2 != nil {
				if err = ie.PCell_FR2.Encode(extWriter); err != nil {
					return utils.WrapError("Encode PCell_FR2", err)
				}
			}
			// encode Pdsch_RE_MappingFR2_PerSlot optional
			if ie.Pdsch_RE_MappingFR2_PerSlot != nil {
				if err = ie.Pdsch_RE_MappingFR2_PerSlot.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdsch_RE_MappingFR2_PerSlot", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 2
		if extBitmap[1] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{ie.DefaultSpatialRelationPathlossRS_r16 != nil, ie.SpatialRelationUpdateAP_SRS_r16 != nil, ie.MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode DefaultSpatialRelationPathlossRS_r16 optional
			if ie.DefaultSpatialRelationPathlossRS_r16 != nil {
				if err = ie.DefaultSpatialRelationPathlossRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode DefaultSpatialRelationPathlossRS_r16", err)
				}
			}
			// encode SpatialRelationUpdateAP_SRS_r16 optional
			if ie.SpatialRelationUpdateAP_SRS_r16 != nil {
				if err = ie.SpatialRelationUpdateAP_SRS_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SpatialRelationUpdateAP_SRS_r16", err)
				}
			}
			// encode MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 optional
			if ie.MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 != nil {
				if err = ie.MaxNumberSRS_PosSpatialRelationsAllServingCells_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode MaxNumberSRS_PosSpatialRelationsAllServingCells_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}
	}
	return nil
}

func (ie *Phy_ParametersFR2) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var DummyPresent bool
	if DummyPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Pdsch_RE_MappingFR2_PerSymbolPresent bool
	if Pdsch_RE_MappingFR2_PerSymbolPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DummyPresent {
		ie.Dummy = new(Phy_ParametersFR2_dummy)
		if err = ie.Dummy.Decode(r); err != nil {
			return utils.WrapError("Decode Dummy", err)
		}
	}
	if Pdsch_RE_MappingFR2_PerSymbolPresent {
		ie.Pdsch_RE_MappingFR2_PerSymbol = new(Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSymbol)
		if err = ie.Pdsch_RE_MappingFR2_PerSymbol.Decode(r); err != nil {
			return utils.WrapError("Decode Pdsch_RE_MappingFR2_PerSymbol", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 2 bits for 2 extension groups
		extBitmap, err := r.ReadExtBitMap()
		if err != nil {
			return err
		}

		// decode extension group 1
		if len(extBitmap) > 0 && extBitmap[0] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			PCell_FR2Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Pdsch_RE_MappingFR2_PerSlotPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode PCell_FR2 optional
			if PCell_FR2Present {
				ie.PCell_FR2 = new(Phy_ParametersFR2_pCell_FR2)
				if err = ie.PCell_FR2.Decode(extReader); err != nil {
					return utils.WrapError("Decode PCell_FR2", err)
				}
			}
			// decode Pdsch_RE_MappingFR2_PerSlot optional
			if Pdsch_RE_MappingFR2_PerSlotPresent {
				ie.Pdsch_RE_MappingFR2_PerSlot = new(Phy_ParametersFR2_pdsch_RE_MappingFR2_PerSlot)
				if err = ie.Pdsch_RE_MappingFR2_PerSlot.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdsch_RE_MappingFR2_PerSlot", err)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			DefaultSpatialRelationPathlossRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SpatialRelationUpdateAP_SRS_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			MaxNumberSRS_PosSpatialRelationsAllServingCells_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode DefaultSpatialRelationPathlossRS_r16 optional
			if DefaultSpatialRelationPathlossRS_r16Present {
				ie.DefaultSpatialRelationPathlossRS_r16 = new(Phy_ParametersFR2_defaultSpatialRelationPathlossRS_r16)
				if err = ie.DefaultSpatialRelationPathlossRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode DefaultSpatialRelationPathlossRS_r16", err)
				}
			}
			// decode SpatialRelationUpdateAP_SRS_r16 optional
			if SpatialRelationUpdateAP_SRS_r16Present {
				ie.SpatialRelationUpdateAP_SRS_r16 = new(Phy_ParametersFR2_spatialRelationUpdateAP_SRS_r16)
				if err = ie.SpatialRelationUpdateAP_SRS_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode SpatialRelationUpdateAP_SRS_r16", err)
				}
			}
			// decode MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 optional
			if MaxNumberSRS_PosSpatialRelationsAllServingCells_r16Present {
				ie.MaxNumberSRS_PosSpatialRelationsAllServingCells_r16 = new(Phy_ParametersFR2_maxNumberSRS_PosSpatialRelationsAllServingCells_r16)
				if err = ie.MaxNumberSRS_PosSpatialRelationsAllServingCells_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode MaxNumberSRS_PosSpatialRelationsAllServingCells_r16", err)
				}
			}
		}
	}
	return nil
}
