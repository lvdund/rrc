package ies

import (
	"bytes"

	"github.com/lvdund/asn1go/uper"
	"github.com/lvdund/rrc/utils"
)

type NZP_CSI_RS_ResourceSet struct {
	Nzp_CSI_ResourceSetId           NZP_CSI_RS_ResourceSetId             `madatory`
	Nzp_CSI_RS_Resources            []NZP_CSI_RS_ResourceId              `lb:1,ub:maxNrofNZP_CSI_RS_ResourcesPerSet,madatory`
	Repetition                      *NZP_CSI_RS_ResourceSet_repetition   `optional`
	AperiodicTriggeringOffset       *int64                               `lb:0,ub:6,optional`
	Trs_Info                        *NZP_CSI_RS_ResourceSet_trs_Info     `optional`
	AperiodicTriggeringOffset_r16   *int64                               `lb:0,ub:31,optional,ext-1`
	Pdc_Info_r17                    *NZP_CSI_RS_ResourceSet_pdc_Info_r17 `optional,ext-2`
	CmrGroupingAndPairing_r17       *CMRGroupingAndPairing_r17           `optional,ext-2`
	AperiodicTriggeringOffset_r17   *int64                               `lb:0,ub:124,optional,ext-2`
	AperiodicTriggeringOffsetL2_r17 *int64                               `lb:0,ub:31,optional,ext-2`
}

func (ie *NZP_CSI_RS_ResourceSet) Encode(w *uper.UperWriter) error {
	var err error
	hasExtensions := ie.AperiodicTriggeringOffset_r16 != nil || ie.Pdc_Info_r17 != nil || ie.CmrGroupingAndPairing_r17 != nil || ie.AperiodicTriggeringOffset_r17 != nil || ie.AperiodicTriggeringOffsetL2_r17 != nil
	preambleBits := []bool{hasExtensions, ie.Repetition != nil, ie.AperiodicTriggeringOffset != nil, ie.Trs_Info != nil}
	for _, bit := range preambleBits {
		if err = w.WriteBool(bit); err != nil {
			return err
		}
	}
	if err = ie.Nzp_CSI_ResourceSetId.Encode(w); err != nil {
		return utils.WrapError("Encode Nzp_CSI_ResourceSetId", err)
	}
	tmp_Nzp_CSI_RS_Resources := utils.NewSequence[*NZP_CSI_RS_ResourceId]([]*NZP_CSI_RS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourcesPerSet}, false)
	for _, i := range ie.Nzp_CSI_RS_Resources {
		tmp_Nzp_CSI_RS_Resources.Value = append(tmp_Nzp_CSI_RS_Resources.Value, &i)
	}
	if err = tmp_Nzp_CSI_RS_Resources.Encode(w); err != nil {
		return utils.WrapError("Encode Nzp_CSI_RS_Resources", err)
	}
	if ie.Repetition != nil {
		if err = ie.Repetition.Encode(w); err != nil {
			return utils.WrapError("Encode Repetition", err)
		}
	}
	if ie.AperiodicTriggeringOffset != nil {
		if err = w.WriteInteger(*ie.AperiodicTriggeringOffset, &uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
			return utils.WrapError("Encode AperiodicTriggeringOffset", err)
		}
	}
	if ie.Trs_Info != nil {
		if err = ie.Trs_Info.Encode(w); err != nil {
			return utils.WrapError("Encode Trs_Info", err)
		}
	}
	if hasExtensions {
		// Extension bitmap: 2 bits for 2 extension groups
		extBitmap := []bool{ie.AperiodicTriggeringOffset_r16 != nil, ie.Pdc_Info_r17 != nil || ie.CmrGroupingAndPairing_r17 != nil || ie.AperiodicTriggeringOffset_r17 != nil || ie.AperiodicTriggeringOffsetL2_r17 != nil}
		if err := w.WriteExtBitMap(extBitmap); err != nil {
			return utils.WrapError("WriteExtBitMap NZP_CSI_RS_ResourceSet", err)
		}

		// encode extension group 1
		if extBitmap[0] {
			extBuf := new(bytes.Buffer)
			extWriter := uper.NewWriter(extBuf)

			// Write preamble bits for optional fields in extension group 1
			optionals_ext_1 := []bool{ie.AperiodicTriggeringOffset_r16 != nil}
			for _, bit := range optionals_ext_1 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode AperiodicTriggeringOffset_r16 optional
			if ie.AperiodicTriggeringOffset_r16 != nil {
				if err = extWriter.WriteInteger(*ie.AperiodicTriggeringOffset_r16, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode AperiodicTriggeringOffset_r16", err)
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
			optionals_ext_2 := []bool{ie.Pdc_Info_r17 != nil, ie.CmrGroupingAndPairing_r17 != nil, ie.AperiodicTriggeringOffset_r17 != nil, ie.AperiodicTriggeringOffsetL2_r17 != nil}
			for _, bit := range optionals_ext_2 {
				if err := extWriter.WriteBool(bit); err != nil {
					return err
				}
			}

			// encode Pdc_Info_r17 optional
			if ie.Pdc_Info_r17 != nil {
				if err = ie.Pdc_Info_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode Pdc_Info_r17", err)
				}
			}
			// encode CmrGroupingAndPairing_r17 optional
			if ie.CmrGroupingAndPairing_r17 != nil {
				if err = ie.CmrGroupingAndPairing_r17.Encode(extWriter); err != nil {
					return utils.WrapError("Encode CmrGroupingAndPairing_r17", err)
				}
			}
			// encode AperiodicTriggeringOffset_r17 optional
			if ie.AperiodicTriggeringOffset_r17 != nil {
				if err = extWriter.WriteInteger(*ie.AperiodicTriggeringOffset_r17, &uper.Constraint{Lb: 0, Ub: 124}, false); err != nil {
					return utils.WrapError("Encode AperiodicTriggeringOffset_r17", err)
				}
			}
			// encode AperiodicTriggeringOffsetL2_r17 optional
			if ie.AperiodicTriggeringOffsetL2_r17 != nil {
				if err = extWriter.WriteInteger(*ie.AperiodicTriggeringOffsetL2_r17, &uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
					return utils.WrapError("Encode AperiodicTriggeringOffsetL2_r17", err)
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

func (ie *NZP_CSI_RS_ResourceSet) Decode(r *uper.UperReader) error {
	var err error
	var extensionBit bool
	if extensionBit, err = r.ReadBool(); err != nil {
		return err
	}
	var RepetitionPresent bool
	if RepetitionPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var AperiodicTriggeringOffsetPresent bool
	if AperiodicTriggeringOffsetPresent, err = r.ReadBool(); err != nil {
		return err
	}
	var Trs_InfoPresent bool
	if Trs_InfoPresent, err = r.ReadBool(); err != nil {
		return err
	}
	if err = ie.Nzp_CSI_ResourceSetId.Decode(r); err != nil {
		return utils.WrapError("Decode Nzp_CSI_ResourceSetId", err)
	}
	tmp_Nzp_CSI_RS_Resources := utils.NewSequence[*NZP_CSI_RS_ResourceId]([]*NZP_CSI_RS_ResourceId{}, uper.Constraint{Lb: 1, Ub: maxNrofNZP_CSI_RS_ResourcesPerSet}, false)
	fn_Nzp_CSI_RS_Resources := func() *NZP_CSI_RS_ResourceId {
		return new(NZP_CSI_RS_ResourceId)
	}
	if err = tmp_Nzp_CSI_RS_Resources.Decode(r, fn_Nzp_CSI_RS_Resources); err != nil {
		return utils.WrapError("Decode Nzp_CSI_RS_Resources", err)
	}
	ie.Nzp_CSI_RS_Resources = []NZP_CSI_RS_ResourceId{}
	for _, i := range tmp_Nzp_CSI_RS_Resources.Value {
		ie.Nzp_CSI_RS_Resources = append(ie.Nzp_CSI_RS_Resources, *i)
	}
	if RepetitionPresent {
		ie.Repetition = new(NZP_CSI_RS_ResourceSet_repetition)
		if err = ie.Repetition.Decode(r); err != nil {
			return utils.WrapError("Decode Repetition", err)
		}
	}
	if AperiodicTriggeringOffsetPresent {
		var tmp_int_AperiodicTriggeringOffset int64
		if tmp_int_AperiodicTriggeringOffset, err = r.ReadInteger(&uper.Constraint{Lb: 0, Ub: 6}, false); err != nil {
			return utils.WrapError("Decode AperiodicTriggeringOffset", err)
		}
		ie.AperiodicTriggeringOffset = &tmp_int_AperiodicTriggeringOffset
	}
	if Trs_InfoPresent {
		ie.Trs_Info = new(NZP_CSI_RS_ResourceSet_trs_Info)
		if err = ie.Trs_Info.Decode(r); err != nil {
			return utils.WrapError("Decode Trs_Info", err)
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

			AperiodicTriggeringOffset_r16Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode AperiodicTriggeringOffset_r16 optional
			if AperiodicTriggeringOffset_r16Present {
				var tmp_int_AperiodicTriggeringOffset_r16 int64
				if tmp_int_AperiodicTriggeringOffset_r16, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode AperiodicTriggeringOffset_r16", err)
				}
				ie.AperiodicTriggeringOffset_r16 = &tmp_int_AperiodicTriggeringOffset_r16
			}
		}
		// decode extension group 2
		if len(extBitmap) > 1 && extBitmap[1] {
			extBytes, err := r.ReadOpenType()
			if err != nil {
				return err
			}

			extReader := uper.NewReader(bytes.NewReader(extBytes))

			Pdc_Info_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			CmrGroupingAndPairing_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AperiodicTriggeringOffset_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			AperiodicTriggeringOffsetL2_r17Present, err := extReader.ReadBool()
			if err != nil {
				return err
			}
			// decode Pdc_Info_r17 optional
			if Pdc_Info_r17Present {
				ie.Pdc_Info_r17 = new(NZP_CSI_RS_ResourceSet_pdc_Info_r17)
				if err = ie.Pdc_Info_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode Pdc_Info_r17", err)
				}
			}
			// decode CmrGroupingAndPairing_r17 optional
			if CmrGroupingAndPairing_r17Present {
				ie.CmrGroupingAndPairing_r17 = new(CMRGroupingAndPairing_r17)
				if err = ie.CmrGroupingAndPairing_r17.Decode(extReader); err != nil {
					return utils.WrapError("Decode CmrGroupingAndPairing_r17", err)
				}
			}
			// decode AperiodicTriggeringOffset_r17 optional
			if AperiodicTriggeringOffset_r17Present {
				var tmp_int_AperiodicTriggeringOffset_r17 int64
				if tmp_int_AperiodicTriggeringOffset_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 124}, false); err != nil {
					return utils.WrapError("Decode AperiodicTriggeringOffset_r17", err)
				}
				ie.AperiodicTriggeringOffset_r17 = &tmp_int_AperiodicTriggeringOffset_r17
			}
			// decode AperiodicTriggeringOffsetL2_r17 optional
			if AperiodicTriggeringOffsetL2_r17Present {
				var tmp_int_AperiodicTriggeringOffsetL2_r17 int64
				if tmp_int_AperiodicTriggeringOffsetL2_r17, err = extReader.ReadInteger(&uper.Constraint{Lb: 0, Ub: 31}, false); err != nil {
					return utils.WrapError("Decode AperiodicTriggeringOffsetL2_r17", err)
				}
				ie.AperiodicTriggeringOffsetL2_r17 = &tmp_int_AperiodicTriggeringOffsetL2_r17
			}
		}
	}
	return nil
}
