package ies

// CA-ParametersNR-v1610-interFreqDAPS-r16 ::= SEQUENCE
type CaParametersnrV1610InterfreqdapsR16 struct {
	InterfreqasyncdapsR16                       *CaParametersnrV1610InterfreqdapsR16InterfreqasyncdapsR16
	InterfreqdiffscsDapsR16                     *CaParametersnrV1610InterfreqdapsR16InterfreqdiffscsDapsR16
	InterfreqmultiulTransmissiondapsR16         *CaParametersnrV1610InterfreqdapsR16InterfreqmultiulTransmissiondapsR16
	InterfreqsemistaticpowersharingdapsMode1R16 *CaParametersnrV1610InterfreqdapsR16InterfreqsemistaticpowersharingdapsMode1R16
	InterfreqsemistaticpowersharingdapsMode2R16 *CaParametersnrV1610InterfreqdapsR16InterfreqsemistaticpowersharingdapsMode2R16
	InterfreqdynamicpowersharingdapsR16         *CaParametersnrV1610InterfreqdapsR16InterfreqdynamicpowersharingdapsR16
	InterfrequlTranscancellationdapsR16         *CaParametersnrV1610InterfreqdapsR16InterfrequlTranscancellationdapsR16
}
