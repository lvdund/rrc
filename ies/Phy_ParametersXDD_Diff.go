package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type Phy_ParametersXDD_Diff struct {
	DynamicSFI                      *Phy_ParametersXDD_Diff_dynamicSFI                      `optional`
	TwoPUCCH_F0_2_ConsecSymbols     *Phy_ParametersXDD_Diff_twoPUCCH_F0_2_ConsecSymbols     `optional`
	TwoDifferentTPC_Loop_PUSCH      *Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUSCH      `optional`
	TwoDifferentTPC_Loop_PUCCH      *Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUCCH      `optional`
	Dl_SchedulingOffset_PDSCH_TypeA *Phy_ParametersXDD_Diff_dl_SchedulingOffset_PDSCH_TypeA `optional,ext-1`
	Dl_SchedulingOffset_PDSCH_TypeB *Phy_ParametersXDD_Diff_dl_SchedulingOffset_PDSCH_TypeB `optional,ext-1`
	Ul_SchedulingOffset             *Phy_ParametersXDD_Diff_ul_SchedulingOffset             `optional,ext-1`
}

func (ie *Phy_ParametersXDD_Diff) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.Dl_SchedulingOffset_PDSCH_TypeA != nil || ie.Dl_SchedulingOffset_PDSCH_TypeB != nil || ie.Ul_SchedulingOffset != nil
	preambleBits := []bool{hasExtensions, ie.DynamicSFI != nil, ie.TwoPUCCH_F0_2_ConsecSymbols != nil, ie.TwoDifferentTPC_Loop_PUSCH != nil, ie.TwoDifferentTPC_Loop_PUCCH != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.DynamicSFI != nil {
		if err = ie.DynamicSFI.Encode(w); err != nil {
			return utils.WrapError("Encode DynamicSFI", err)
		}
	}
	if ie.TwoPUCCH_F0_2_ConsecSymbols != nil {
		if err = ie.TwoPUCCH_F0_2_ConsecSymbols.Encode(w); err != nil {
			return utils.WrapError("Encode TwoPUCCH_F0_2_ConsecSymbols", err)
		}
	}
	if ie.TwoDifferentTPC_Loop_PUSCH != nil {
		if err = ie.TwoDifferentTPC_Loop_PUSCH.Encode(w); err != nil {
			return utils.WrapError("Encode TwoDifferentTPC_Loop_PUSCH", err)
		}
	}
	if ie.TwoDifferentTPC_Loop_PUCCH != nil {
		if err = ie.TwoDifferentTPC_Loop_PUCCH.Encode(w); err != nil {
			return utils.WrapError("Encode TwoDifferentTPC_Loop_PUCCH", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 1 bits for 1 extension groups
		extBitmap := []bool{ie.Dl_SchedulingOffset_PDSCH_TypeA != nil || ie.Dl_SchedulingOffset_PDSCH_TypeB != nil || ie.Ul_SchedulingOffset != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap Phy_ParametersXDD_Diff", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.Dl_SchedulingOffset_PDSCH_TypeA != nil, ie.Dl_SchedulingOffset_PDSCH_TypeB != nil, ie.Ul_SchedulingOffset != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Dl_SchedulingOffset_PDSCH_TypeA optional
			if ie.Dl_SchedulingOffset_PDSCH_TypeA != nil {
				if err = ie.Dl_SchedulingOffset_PDSCH_TypeA.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_SchedulingOffset_PDSCH_TypeA", err)
				}
			}
			// encode Dl_SchedulingOffset_PDSCH_TypeB optional
			if ie.Dl_SchedulingOffset_PDSCH_TypeB != nil {
				if err = ie.Dl_SchedulingOffset_PDSCH_TypeB.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Dl_SchedulingOffset_PDSCH_TypeB", err)
				}
			}
			// encode Ul_SchedulingOffset optional
			if ie.Ul_SchedulingOffset != nil {
				if err = ie.Ul_SchedulingOffset.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Ul_SchedulingOffset", err)
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

func (ie *Phy_ParametersXDD_Diff) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var DynamicSFIPresent bool
	if DynamicSFIPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoPUCCH_F0_2_ConsecSymbolsPresent bool
	if TwoPUCCH_F0_2_ConsecSymbolsPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoDifferentTPC_Loop_PUSCHPresent bool
	if TwoDifferentTPC_Loop_PUSCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var TwoDifferentTPC_Loop_PUCCHPresent bool
	if TwoDifferentTPC_Loop_PUCCHPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if DynamicSFIPresent {
		ie.DynamicSFI = new(Phy_ParametersXDD_Diff_dynamicSFI)
		if err = ie.DynamicSFI.Decode(r); err != nil {
			return utils.WrapError("Decode DynamicSFI", err)
		}
	}
	if TwoPUCCH_F0_2_ConsecSymbolsPresent {
		ie.TwoPUCCH_F0_2_ConsecSymbols = new(Phy_ParametersXDD_Diff_twoPUCCH_F0_2_ConsecSymbols)
		if err = ie.TwoPUCCH_F0_2_ConsecSymbols.Decode(r); err != nil {
			return utils.WrapError("Decode TwoPUCCH_F0_2_ConsecSymbols", err)
		}
	}
	if TwoDifferentTPC_Loop_PUSCHPresent {
		ie.TwoDifferentTPC_Loop_PUSCH = new(Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUSCH)
		if err = ie.TwoDifferentTPC_Loop_PUSCH.Decode(r); err != nil {
			return utils.WrapError("Decode TwoDifferentTPC_Loop_PUSCH", err)
		}
	}
	if TwoDifferentTPC_Loop_PUCCHPresent {
		ie.TwoDifferentTPC_Loop_PUCCH = new(Phy_ParametersXDD_Diff_twoDifferentTPC_Loop_PUCCH)
		if err = ie.TwoDifferentTPC_Loop_PUCCH.Decode(r); err != nil {
			return utils.WrapError("Decode TwoDifferentTPC_Loop_PUCCH", err)
		}
	}

	if extensionBit {
		// Read extension bitmap: 1 bits for 1 extension groups
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

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Dl_SchedulingOffset_PDSCH_TypeAPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Dl_SchedulingOffset_PDSCH_TypeBPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Ul_SchedulingOffsetPresent, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Dl_SchedulingOffset_PDSCH_TypeA optional
			if Dl_SchedulingOffset_PDSCH_TypeAPresent {
				ie.Dl_SchedulingOffset_PDSCH_TypeA = new(Phy_ParametersXDD_Diff_dl_SchedulingOffset_PDSCH_TypeA)
				if err = ie.Dl_SchedulingOffset_PDSCH_TypeA.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_SchedulingOffset_PDSCH_TypeA", err)
				}
			}
			// decode Dl_SchedulingOffset_PDSCH_TypeB optional
			if Dl_SchedulingOffset_PDSCH_TypeBPresent {
				ie.Dl_SchedulingOffset_PDSCH_TypeB = new(Phy_ParametersXDD_Diff_dl_SchedulingOffset_PDSCH_TypeB)
				if err = ie.Dl_SchedulingOffset_PDSCH_TypeB.Decode(extReader); err != nil {
					return utils.WrapError("Decode Dl_SchedulingOffset_PDSCH_TypeB", err)
				}
			}
			// decode Ul_SchedulingOffset optional
			if Ul_SchedulingOffsetPresent {
				ie.Ul_SchedulingOffset = new(Phy_ParametersXDD_Diff_ul_SchedulingOffset)
				if err = ie.Ul_SchedulingOffset.Decode(extReader); err != nil {
					return utils.WrapError("Decode Ul_SchedulingOffset", err)
				}
			}
		}
	}
	return nil
}
