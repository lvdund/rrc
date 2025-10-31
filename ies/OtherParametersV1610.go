package ies

// Other-Parameters-v1610 ::= SEQUENCE
type OtherParametersV1610 struct {
	ResumewithstoredmcgScellsR16 *OtherParametersV1610ResumewithstoredmcgScellsR16
	ResumewithmcgScellconfigR16  *OtherParametersV1610ResumewithmcgScellconfigR16
	ResumewithstoredscgR16       *OtherParametersV1610ResumewithstoredscgR16
	ResumewithscgConfigR16       *OtherParametersV1610ResumewithscgConfigR16
	McgrlfRecoveryviascgR16      *OtherParametersV1610McgrlfRecoveryviascgR16
	OverheatingindforscgR16      *OtherParametersV1610OverheatingindforscgR16
}
