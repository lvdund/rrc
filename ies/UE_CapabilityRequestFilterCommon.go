package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/aper"
	"github.com/lvdund/rrc/utils"
)

type UE_CapabilityRequestFilterCommon struct {
	Mrdc_Request                 *UE_CapabilityRequestFilterCommon_mrdc_Request                 `optional`
	CodebookTypeRequest_r16      *UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16      `optional,ext-1`
	UplinkTxSwitchRequest_r16    *UE_CapabilityRequestFilterCommon_uplinkTxSwitchRequest_r16    `optional,ext-1`
	RequestedCellGrouping_r16    []CellGrouping_r16                                             `lb:1,ub:maxCellGroupings_r16,optional,ext-2`
	FallbackGroupFiveRequest_r17 *UE_CapabilityRequestFilterCommon_fallbackGroupFiveRequest_r17 `optional,ext-3`
}

func (ie *UE_CapabilityRequestFilterCommon) Encode(w *aper.AperWriter) error {
	var err error
	hasExtensions := ie.CodebookTypeRequest_r16 != nil || ie.UplinkTxSwitchRequest_r16 != nil || len(ie.RequestedCellGrouping_r16) > 0 || ie.FallbackGroupFiveRequest_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Mrdc_Request != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if ie.Mrdc_Request != nil {
		if err = ie.Mrdc_Request.Encode(w); err != nil {
			return utils.WrapError("Encode Mrdc_Request", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 3 bits for 3 extension groups
		extBitmap := []bool{ie.CodebookTypeRequest_r16 != nil || ie.UplinkTxSwitchRequest_r16 != nil, len(ie.RequestedCellGrouping_r16) > 0, ie.FallbackGroupFiveRequest_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap UE_CapabilityRequestFilterCommon", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := aper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.CodebookTypeRequest_r16 != nil, ie.UplinkTxSwitchRequest_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode CodebookTypeRequest_r16 optional
			if ie.CodebookTypeRequest_r16 != nil {
				if err = ie.CodebookTypeRequest_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CodebookTypeRequest_r16", err)
				}
			}
			// encode UplinkTxSwitchRequest_r16 optional
			if ie.UplinkTxSwitchRequest_r16 != nil {
				if err = ie.UplinkTxSwitchRequest_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode UplinkTxSwitchRequest_r16", err)
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
			optionals_ext_2 := []bool{len(ie.RequestedCellGrouping_r16) > 0}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode RequestedCellGrouping_r16 optional
			if len(ie.RequestedCellGrouping_r16) > 0 {
				tmp_RequestedCellGrouping_r16 := utils.NewSequence[*CellGrouping_r16]([]*CellGrouping_r16{}, aper.Constraint{Lb: 1, Ub: maxCellGroupings_r16}, false)
				for _, i := range ie.RequestedCellGrouping_r16 {
					tmp_RequestedCellGrouping_r16.Value = append(tmp_RequestedCellGrouping_r16.Value, &i)
				}
				if err = tmp_RequestedCellGrouping_r16.Encode(extWriter); err != nil {
					return utils.WrapError("Encode RequestedCellGrouping_r16", err)
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
			optionals_ext_3 := []bool{ie.FallbackGroupFiveRequest_r17 != nil}
			for _, bit := range optionals_ext_3 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode FallbackGroupFiveRequest_r17 optional
			if ie.FallbackGroupFiveRequest_r17 != nil {
				if err = ie.FallbackGroupFiveRequest_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode FallbackGroupFiveRequest_r17", err)
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

func (ie *UE_CapabilityRequestFilterCommon) Decode(r *aper.AperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var Mrdc_RequestPresent bool
	if Mrdc_RequestPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if Mrdc_RequestPresent {
		ie.Mrdc_Request = new(UE_CapabilityRequestFilterCommon_mrdc_Request)
		if err = ie.Mrdc_Request.Decode(r); err != nil {
			return utils.WrapError("Decode Mrdc_Request", err)
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

			CodebookTypeRequest_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			UplinkTxSwitchRequest_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode CodebookTypeRequest_r16 optional
			if CodebookTypeRequest_r16Present {
				ie.CodebookTypeRequest_r16 = new(UE_CapabilityRequestFilterCommon_codebookTypeRequest_r16)
				if err = ie.CodebookTypeRequest_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode CodebookTypeRequest_r16", err)
				}
			}
			// decode UplinkTxSwitchRequest_r16 optional
			if UplinkTxSwitchRequest_r16Present {
				ie.UplinkTxSwitchRequest_r16 = new(UE_CapabilityRequestFilterCommon_uplinkTxSwitchRequest_r16)
				if err = ie.UplinkTxSwitchRequest_r16.Decode(extReader); err != nil {
					return utils.WrapError("Decode UplinkTxSwitchRequest_r16", err)
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

			RequestedCellGrouping_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode RequestedCellGrouping_r16 optional
			if RequestedCellGrouping_r16Present {
				tmp_RequestedCellGrouping_r16 := utils.NewSequence[*CellGrouping_r16]([]*CellGrouping_r16{}, aper.Constraint{Lb: 1, Ub: maxCellGroupings_r16}, false)
				fn_RequestedCellGrouping_r16 := func() *CellGrouping_r16 {
					return new(CellGrouping_r16)
				}
				if err = tmp_RequestedCellGrouping_r16.Decode(extReader, fn_RequestedCellGrouping_r16); err != nil {
					return utils.WrapError("Decode RequestedCellGrouping_r16", err)
				}
				ie.RequestedCellGrouping_r16 = []CellGrouping_r16{}
				for _, i := range tmp_RequestedCellGrouping_r16.Value {
					ie.RequestedCellGrouping_r16 = append(ie.RequestedCellGrouping_r16, *i)
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

			FallbackGroupFiveRequest_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode FallbackGroupFiveRequest_r17 optional
			if FallbackGroupFiveRequest_r17Present {
				ie.FallbackGroupFiveRequest_r17 = new(UE_CapabilityRequestFilterCommon_fallbackGroupFiveRequest_r17)
				if err = ie.FallbackGroupFiveRequest_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode FallbackGroupFiveRequest_r17", err)
				}
			}
		}
	}
	return nil
}
