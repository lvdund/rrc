package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type SlotFormatIndicator struct {
	Sfi_RNTI                              RNTI_Value                      `madatory`
	Dci_PayloadSize                       int64                           `lb:1,ub:maxSFI_DCI_PayloadSize,madatory`
	SlotFormatCombToAddModList            []SlotFormatCombinationsPerCell `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional`
	SlotFormatCombToReleaseList           []ServCellIndex                 `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional`
	AvailableRB_SetsToAddModList_r16      []AvailableRB_SetsPerCell_r16   `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional,ext-1`
	AvailableRB_SetsToReleaseList_r16     []ServCellIndex                 `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional,ext-1`
	SwitchTriggerToAddModList_r16         []SearchSpaceSwitchTrigger_r16  `lb:1,ub:4,optional,ext-1`
	SwitchTriggerToReleaseList_r16        []ServCellIndex                 `lb:1,ub:4,optional,ext-1`
	Co_DurationsPerCellToAddModList_r16   []CO_DurationsPerCell_r16       `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional,ext-1`
	Co_DurationsPerCellToReleaseList_r16  []ServCellIndex                 `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional,ext-1`
	SwitchTriggerToAddModListSizeExt_r16  []SearchSpaceSwitchTrigger_r16  `lb:1,ub:maxNrofAggregatedCellsPerCellGroupMinus4_r16,optional,ext-2`
	SwitchTriggerToReleaseListSizeExt_r16 []ServCellIndex                 `lb:1,ub:maxNrofAggregatedCellsPerCellGroupMinus4_r16,optional,ext-2`
	Co_DurationsPerCellToAddModList_r17   []CO_DurationsPerCell_r17       `lb:1,ub:maxNrofAggregatedCellsPerCellGroup,optional,ext-3`
}

func (ie *SlotFormatIndicator) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := len(ie.AvailableRB_SetsToAddModList_r16) > 0 || len(ie.AvailableRB_SetsToReleaseList_r16) > 0 || len(ie.SwitchTriggerToAddModList_r16) > 0 || len(ie.SwitchTriggerToReleaseList_r16) > 0 || len(ie.Co_DurationsPerCellToAddModList_r16) > 0 || len(ie.Co_DurationsPerCellToReleaseList_r16) > 0 || len(ie.SwitchTriggerToAddModListSizeExt_r16) > 0 || len(ie.SwitchTriggerToReleaseListSizeExt_r16) > 0 || len(ie.Co_DurationsPerCellToAddModList_r17) > 0
	preambleBits := []bool{hasExtensions, len(ie.SlotFormatCombToAddModList) > 0, len(ie.SlotFormatCombToReleaseList) > 0}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Sfi_RNTI.Encode(w); err != nil {
		return utils.WrapError("Encode Sfi_RNTI", err)
	}
	if err = w.WriteInteger(ie.Dci_PayloadSize, &aper.Constraint{Lb: 1, Ub: maxSFI_DCI_PayloadSize}, false); err != nil {
		return utils.WrapError("WriteInteger Dci_PayloadSize", err)
	}
	if len(ie.SlotFormatCombToAddModList) > 0 {
		tmp_SlotFormatCombToAddModList := utils.NewSequence[*SlotFormatCombinationsPerCell]([]*SlotFormatCombinationsPerCell{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
		for _, i := range ie.SlotFormatCombToAddModList {
			tmp_SlotFormatCombToAddModList.Value = append(tmp_SlotFormatCombToAddModList.Value, &i)
		}
		if err = tmp_SlotFormatCombToAddModList.Encode(w); err != nil {
			return utils.WrapError("Encode SlotFormatCombToAddModList", err)
		}
	}
	if len(ie.SlotFormatCombToReleaseList) > 0 {
		tmp_SlotFormatCombToReleaseList := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
		for _, i := range ie.SlotFormatCombToReleaseList {
			tmp_SlotFormatCombToReleaseList.Value = append(tmp_SlotFormatCombToReleaseList.Value, &i)
		}
		if err = tmp_SlotFormatCombToReleaseList.Encode(w); err != nil {
			return utils.WrapError("Encode SlotFormatCombToReleaseList", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{len(ie.AvailableRB_SetsToAddModList_r16) > 0 || len(ie.AvailableRB_SetsToReleaseList_r16) > 0 || len(ie.SwitchTriggerToAddModList_r16) > 0 || len(ie.SwitchTriggerToReleaseList_r16) > 0 || len(ie.Co_DurationsPerCellToAddModList_r16) > 0 || len(ie.Co_DurationsPerCellToReleaseList_r16) > 0, len(ie.SwitchTriggerToAddModListSizeExt_r16) > 0 || len(ie.SwitchTriggerToReleaseListSizeExt_r16) > 0, len(ie.Co_DurationsPerCellToAddModList_r17) > 0}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap SlotFormatIndicator", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{len(ie.AvailableRB_SetsToAddModList_r16) > 0, len(ie.AvailableRB_SetsToReleaseList_r16) > 0, len(ie.SwitchTriggerToAddModList_r16) > 0, len(ie.SwitchTriggerToReleaseList_r16) > 0, len(ie.Co_DurationsPerCellToAddModList_r16) > 0, len(ie.Co_DurationsPerCellToReleaseList_r16) > 0}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode AvailableRB_SetsToAddModList_r16 optional
			if len(ie.AvailableRB_SetsToAddModList_r16) > 0 {
				tmp_AvailableRB_SetsToAddModList_r16 := utils.NewSequence[*AvailableRB_SetsPerCell_r16]([]*AvailableRB_SetsPerCell_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				for _, i := range ie.AvailableRB_SetsToAddModList_r16 {
					tmp_AvailableRB_SetsToAddModList_r16.Value = append(tmp_AvailableRB_SetsToAddModList_r16.Value, &i)
				}
				if err = tmp_AvailableRB_SetsToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AvailableRB_SetsToAddModList_r16", err)
				}
			}
			// encode AvailableRB_SetsToReleaseList_r16 optional
			if len(ie.AvailableRB_SetsToReleaseList_r16) > 0 {
				tmp_AvailableRB_SetsToReleaseList_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				for _, i := range ie.AvailableRB_SetsToReleaseList_r16 {
					tmp_AvailableRB_SetsToReleaseList_r16.Value = append(tmp_AvailableRB_SetsToReleaseList_r16.Value, &i)
				}
				if err = tmp_AvailableRB_SetsToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode AvailableRB_SetsToReleaseList_r16", err)
				}
			}
			// encode SwitchTriggerToAddModList_r16 optional
			if len(ie.SwitchTriggerToAddModList_r16) > 0 {
				tmp_SwitchTriggerToAddModList_r16 := utils.NewSequence[*SearchSpaceSwitchTrigger_r16]([]*SearchSpaceSwitchTrigger_r16{}, aper.Constraint{Lb: 1, Ub: 4}, false)
				for _, i := range ie.SwitchTriggerToAddModList_r16 {
					tmp_SwitchTriggerToAddModList_r16.Value = append(tmp_SwitchTriggerToAddModList_r16.Value, &i)
				}
				if err = tmp_SwitchTriggerToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SwitchTriggerToAddModList_r16", err)
				}
			}
			// encode SwitchTriggerToReleaseList_r16 optional
			if len(ie.SwitchTriggerToReleaseList_r16) > 0 {
				tmp_SwitchTriggerToReleaseList_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: 4}, false)
				for _, i := range ie.SwitchTriggerToReleaseList_r16 {
					tmp_SwitchTriggerToReleaseList_r16.Value = append(tmp_SwitchTriggerToReleaseList_r16.Value, &i)
				}
				if err = tmp_SwitchTriggerToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SwitchTriggerToReleaseList_r16", err)
				}
			}
			// encode Co_DurationsPerCellToAddModList_r16 optional
			if len(ie.Co_DurationsPerCellToAddModList_r16) > 0 {
				tmp_Co_DurationsPerCellToAddModList_r16 := utils.NewSequence[*CO_DurationsPerCell_r16]([]*CO_DurationsPerCell_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				for _, i := range ie.Co_DurationsPerCellToAddModList_r16 {
					tmp_Co_DurationsPerCellToAddModList_r16.Value = append(tmp_Co_DurationsPerCellToAddModList_r16.Value, &i)
				}
				if err = tmp_Co_DurationsPerCellToAddModList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Co_DurationsPerCellToAddModList_r16", err)
				}
			}
			// encode Co_DurationsPerCellToReleaseList_r16 optional
			if len(ie.Co_DurationsPerCellToReleaseList_r16) > 0 {
				tmp_Co_DurationsPerCellToReleaseList_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				for _, i := range ie.Co_DurationsPerCellToReleaseList_r16 {
					tmp_Co_DurationsPerCellToReleaseList_r16.Value = append(tmp_Co_DurationsPerCellToReleaseList_r16.Value, &i)
				}
				if err = tmp_Co_DurationsPerCellToReleaseList_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Co_DurationsPerCellToReleaseList_r16", err)
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
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 2
			optionals_ext_2 := []bool{len(ie.SwitchTriggerToAddModListSizeExt_r16) > 0, len(ie.SwitchTriggerToReleaseListSizeExt_r16) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode SwitchTriggerToAddModListSizeExt_r16 optional
			if len(ie.SwitchTriggerToAddModListSizeExt_r16) > 0 {
				tmp_SwitchTriggerToAddModListSizeExt_r16 := utils.NewSequence[*SearchSpaceSwitchTrigger_r16]([]*SearchSpaceSwitchTrigger_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroupMinus4_r16}, false)
				for _, i := range ie.SwitchTriggerToAddModListSizeExt_r16 {
					tmp_SwitchTriggerToAddModListSizeExt_r16.Value = append(tmp_SwitchTriggerToAddModListSizeExt_r16.Value, &i)
				}
				if err = tmp_SwitchTriggerToAddModListSizeExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SwitchTriggerToAddModListSizeExt_r16", err)
				}
			}
			// encode SwitchTriggerToReleaseListSizeExt_r16 optional
			if len(ie.SwitchTriggerToReleaseListSizeExt_r16) > 0 {
				tmp_SwitchTriggerToReleaseListSizeExt_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroupMinus4_r16}, false)
				for _, i := range ie.SwitchTriggerToReleaseListSizeExt_r16 {
					tmp_SwitchTriggerToReleaseListSizeExt_r16.Value = append(tmp_SwitchTriggerToReleaseListSizeExt_r16.Value, &i)
				}
				if err = tmp_SwitchTriggerToReleaseListSizeExt_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode SwitchTriggerToReleaseListSizeExt_r16", err)
				}
			}

			if err := extWriter.Close(); err != nil {
				return err
			}

			if err := w.WriteOpenType(extBuf.Bytes()); err != nil {
				return err
			}
		}

		// encode extension group 3
		if extBitmap[2] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 3
			optionals_ext_3 := []bool{len(ie.Co_DurationsPerCellToAddModList_r17) > 0}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Co_DurationsPerCellToAddModList_r17 optional
			if len(ie.Co_DurationsPerCellToAddModList_r17) > 0 {
				tmp_Co_DurationsPerCellToAddModList_r17 := utils.NewSequence[*CO_DurationsPerCell_r17]([]*CO_DurationsPerCell_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				for _, i := range ie.Co_DurationsPerCellToAddModList_r17 {
					tmp_Co_DurationsPerCellToAddModList_r17.Value = append(tmp_Co_DurationsPerCellToAddModList_r17.Value, &i)
				}
				if err = tmp_Co_DurationsPerCellToAddModList_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Co_DurationsPerCellToAddModList_r17", err)
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

func (ie *SlotFormatIndicator) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var SlotFormatCombToAddModListPresent bool
	if SlotFormatCombToAddModListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var SlotFormatCombToReleaseListPresent bool
	if SlotFormatCombToReleaseListPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Sfi_RNTI.Decode(r); err != nil {
		return utils.WrapError("Decode Sfi_RNTI", err)
	}
	var tmp_int_Dci_PayloadSize int64
	if tmp_int_Dci_PayloadSize, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: maxSFI_DCI_PayloadSize}, false); err != nil {
		return utils.WrapError("ReadInteger Dci_PayloadSize", err)
	}
	ie.Dci_PayloadSize = tmp_int_Dci_PayloadSize
	if SlotFormatCombToAddModListPresent {
		tmp_SlotFormatCombToAddModList := utils.NewSequence[*SlotFormatCombinationsPerCell]([]*SlotFormatCombinationsPerCell{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
		fn_SlotFormatCombToAddModList := func() *SlotFormatCombinationsPerCell {
			return new(SlotFormatCombinationsPerCell)
		}
		if err = tmp_SlotFormatCombToAddModList.Decode(r, fn_SlotFormatCombToAddModList); err != nil {
			return utils.WrapError("Decode SlotFormatCombToAddModList", err)
		}
		ie.SlotFormatCombToAddModList = []SlotFormatCombinationsPerCell{}
		for _, i := range tmp_SlotFormatCombToAddModList.Value {
			ie.SlotFormatCombToAddModList = append(ie.SlotFormatCombToAddModList, *i)
		}
	}
	if SlotFormatCombToReleaseListPresent {
		tmp_SlotFormatCombToReleaseList := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
		fn_SlotFormatCombToReleaseList := func() *ServCellIndex {
			return new(ServCellIndex)
		}
		if err = tmp_SlotFormatCombToReleaseList.Decode(r, fn_SlotFormatCombToReleaseList); err != nil {
			return utils.WrapError("Decode SlotFormatCombToReleaseList", err)
		}
		ie.SlotFormatCombToReleaseList = []ServCellIndex{}
		for _, i := range tmp_SlotFormatCombToReleaseList.Value {
			ie.SlotFormatCombToReleaseList = append(ie.SlotFormatCombToReleaseList, *i)
		}
	}

	if extensionBit {
		// Read extension bitmap: 3 bits for 3 extension groups
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

			AvailableRB_SetsToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AvailableRB_SetsToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SwitchTriggerToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SwitchTriggerToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Co_DurationsPerCellToAddModList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			Co_DurationsPerCellToReleaseList_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode AvailableRB_SetsToAddModList_r16 optional
			if AvailableRB_SetsToAddModList_r16Present {
				tmp_AvailableRB_SetsToAddModList_r16 := utils.NewSequence[*AvailableRB_SetsPerCell_r16]([]*AvailableRB_SetsPerCell_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				fn_AvailableRB_SetsToAddModList_r16 := func() *AvailableRB_SetsPerCell_r16 {
					return new(AvailableRB_SetsPerCell_r16)
				}
				if err = tmp_AvailableRB_SetsToAddModList_r16.Decode(extReader, fn_AvailableRB_SetsToAddModList_r16); err != nil {
					return utils.WrapError("Decode AvailableRB_SetsToAddModList_r16", err)
				}
				ie.AvailableRB_SetsToAddModList_r16 = []AvailableRB_SetsPerCell_r16{}
				for _, i := range tmp_AvailableRB_SetsToAddModList_r16.Value {
					ie.AvailableRB_SetsToAddModList_r16 = append(ie.AvailableRB_SetsToAddModList_r16, *i)
				}
			}
			// decode AvailableRB_SetsToReleaseList_r16 optional
			if AvailableRB_SetsToReleaseList_r16Present {
				tmp_AvailableRB_SetsToReleaseList_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				fn_AvailableRB_SetsToReleaseList_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_AvailableRB_SetsToReleaseList_r16.Decode(extReader, fn_AvailableRB_SetsToReleaseList_r16); err != nil {
					return utils.WrapError("Decode AvailableRB_SetsToReleaseList_r16", err)
				}
				ie.AvailableRB_SetsToReleaseList_r16 = []ServCellIndex{}
				for _, i := range tmp_AvailableRB_SetsToReleaseList_r16.Value {
					ie.AvailableRB_SetsToReleaseList_r16 = append(ie.AvailableRB_SetsToReleaseList_r16, *i)
				}
			}
			// decode SwitchTriggerToAddModList_r16 optional
			if SwitchTriggerToAddModList_r16Present {
				tmp_SwitchTriggerToAddModList_r16 := utils.NewSequence[*SearchSpaceSwitchTrigger_r16]([]*SearchSpaceSwitchTrigger_r16{}, aper.Constraint{Lb: 1, Ub: 4}, false)
				fn_SwitchTriggerToAddModList_r16 := func() *SearchSpaceSwitchTrigger_r16 {
					return new(SearchSpaceSwitchTrigger_r16)
				}
				if err = tmp_SwitchTriggerToAddModList_r16.Decode(extReader, fn_SwitchTriggerToAddModList_r16); err != nil {
					return utils.WrapError("Decode SwitchTriggerToAddModList_r16", err)
				}
				ie.SwitchTriggerToAddModList_r16 = []SearchSpaceSwitchTrigger_r16{}
				for _, i := range tmp_SwitchTriggerToAddModList_r16.Value {
					ie.SwitchTriggerToAddModList_r16 = append(ie.SwitchTriggerToAddModList_r16, *i)
				}
			}
			// decode SwitchTriggerToReleaseList_r16 optional
			if SwitchTriggerToReleaseList_r16Present {
				tmp_SwitchTriggerToReleaseList_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: 4}, false)
				fn_SwitchTriggerToReleaseList_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_SwitchTriggerToReleaseList_r16.Decode(extReader, fn_SwitchTriggerToReleaseList_r16); err != nil {
					return utils.WrapError("Decode SwitchTriggerToReleaseList_r16", err)
				}
				ie.SwitchTriggerToReleaseList_r16 = []ServCellIndex{}
				for _, i := range tmp_SwitchTriggerToReleaseList_r16.Value {
					ie.SwitchTriggerToReleaseList_r16 = append(ie.SwitchTriggerToReleaseList_r16, *i)
				}
			}
			// decode Co_DurationsPerCellToAddModList_r16 optional
			if Co_DurationsPerCellToAddModList_r16Present {
				tmp_Co_DurationsPerCellToAddModList_r16 := utils.NewSequence[*CO_DurationsPerCell_r16]([]*CO_DurationsPerCell_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				fn_Co_DurationsPerCellToAddModList_r16 := func() *CO_DurationsPerCell_r16 {
					return new(CO_DurationsPerCell_r16)
				}
				if err = tmp_Co_DurationsPerCellToAddModList_r16.Decode(extReader, fn_Co_DurationsPerCellToAddModList_r16); err != nil {
					return utils.WrapError("Decode Co_DurationsPerCellToAddModList_r16", err)
				}
				ie.Co_DurationsPerCellToAddModList_r16 = []CO_DurationsPerCell_r16{}
				for _, i := range tmp_Co_DurationsPerCellToAddModList_r16.Value {
					ie.Co_DurationsPerCellToAddModList_r16 = append(ie.Co_DurationsPerCellToAddModList_r16, *i)
				}
			}
			// decode Co_DurationsPerCellToReleaseList_r16 optional
			if Co_DurationsPerCellToReleaseList_r16Present {
				tmp_Co_DurationsPerCellToReleaseList_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				fn_Co_DurationsPerCellToReleaseList_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_Co_DurationsPerCellToReleaseList_r16.Decode(extReader, fn_Co_DurationsPerCellToReleaseList_r16); err != nil {
					return utils.WrapError("Decode Co_DurationsPerCellToReleaseList_r16", err)
				}
				ie.Co_DurationsPerCellToReleaseList_r16 = []ServCellIndex{}
				for _, i := range tmp_Co_DurationsPerCellToReleaseList_r16.Value {
					ie.Co_DurationsPerCellToReleaseList_r16 = append(ie.Co_DurationsPerCellToReleaseList_r16, *i)
				}
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			SwitchTriggerToAddModListSizeExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			SwitchTriggerToReleaseListSizeExt_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode SwitchTriggerToAddModListSizeExt_r16 optional
			if SwitchTriggerToAddModListSizeExt_r16Present {
				tmp_SwitchTriggerToAddModListSizeExt_r16 := utils.NewSequence[*SearchSpaceSwitchTrigger_r16]([]*SearchSpaceSwitchTrigger_r16{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroupMinus4_r16}, false)
				fn_SwitchTriggerToAddModListSizeExt_r16 := func() *SearchSpaceSwitchTrigger_r16 {
					return new(SearchSpaceSwitchTrigger_r16)
				}
				if err = tmp_SwitchTriggerToAddModListSizeExt_r16.Decode(extReader, fn_SwitchTriggerToAddModListSizeExt_r16); err != nil {
					return utils.WrapError("Decode SwitchTriggerToAddModListSizeExt_r16", err)
				}
				ie.SwitchTriggerToAddModListSizeExt_r16 = []SearchSpaceSwitchTrigger_r16{}
				for _, i := range tmp_SwitchTriggerToAddModListSizeExt_r16.Value {
					ie.SwitchTriggerToAddModListSizeExt_r16 = append(ie.SwitchTriggerToAddModListSizeExt_r16, *i)
				}
			}
			// decode SwitchTriggerToReleaseListSizeExt_r16 optional
			if SwitchTriggerToReleaseListSizeExt_r16Present {
				tmp_SwitchTriggerToReleaseListSizeExt_r16 := utils.NewSequence[*ServCellIndex]([]*ServCellIndex{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroupMinus4_r16}, false)
				fn_SwitchTriggerToReleaseListSizeExt_r16 := func() *ServCellIndex {
					return new(ServCellIndex)
				}
				if err = tmp_SwitchTriggerToReleaseListSizeExt_r16.Decode(extReader, fn_SwitchTriggerToReleaseListSizeExt_r16); err != nil {
					return utils.WrapError("Decode SwitchTriggerToReleaseListSizeExt_r16", err)
				}
				ie.SwitchTriggerToReleaseListSizeExt_r16 = []ServCellIndex{}
				for _, i := range tmp_SwitchTriggerToReleaseListSizeExt_r16.Value {
					ie.SwitchTriggerToReleaseListSizeExt_r16 = append(ie.SwitchTriggerToReleaseListSizeExt_r16, *i)
				}
			}
		}
		// decode extension group 3
		if len(extBitmap) > 2 && extBitmap[2] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := aper.NewReader(bytes.NewReader(extBytes))

			Co_DurationsPerCellToAddModList_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Co_DurationsPerCellToAddModList_r17 optional
			if Co_DurationsPerCellToAddModList_r17Present {
				tmp_Co_DurationsPerCellToAddModList_r17 := utils.NewSequence[*CO_DurationsPerCell_r17]([]*CO_DurationsPerCell_r17{}, aper.Constraint{Lb: 1, Ub: maxNrofAggregatedCellsPerCellGroup}, false)
				fn_Co_DurationsPerCellToAddModList_r17 := func() *CO_DurationsPerCell_r17 {
					return new(CO_DurationsPerCell_r17)
				}
				if err = tmp_Co_DurationsPerCellToAddModList_r17.Decode(extReader, fn_Co_DurationsPerCellToAddModList_r17); err != nil {
					return utils.WrapError("Decode Co_DurationsPerCellToAddModList_r17", err)
				}
				ie.Co_DurationsPerCellToAddModList_r17 = []CO_DurationsPerCell_r17{}
				for _, i := range tmp_Co_DurationsPerCellToAddModList_r17.Value {
					ie.Co_DurationsPerCellToAddModList_r17 = append(ie.Co_DurationsPerCellToAddModList_r17, *i)
				}
			}
		}
	}
	return nil
}
