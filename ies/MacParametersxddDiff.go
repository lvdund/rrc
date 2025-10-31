package ies

// MAC-ParametersXDD-Diff ::= SEQUENCE
// Extensible
type MacParametersxddDiff struct {
	Skipuplinktxdynamic               *MacParametersxddDiffSkipuplinktxdynamic
	LogicalchannelsrDelaytimer        *MacParametersxddDiffLogicalchannelsrDelaytimer
	LongdrxCycle                      *MacParametersxddDiffLongdrxCycle
	ShortdrxCycle                     *MacParametersxddDiffShortdrxCycle
	MultiplesrConfigurations          *MacParametersxddDiffMultiplesrConfigurations
	Multipleconfiguredgrants          *MacParametersxddDiffMultipleconfiguredgrants
	SecondarydrxGroupR16              *MacParametersxddDiffSecondarydrxGroupR16              `ext`
	EnhancedskipuplinktxdynamicR16    *MacParametersxddDiffEnhancedskipuplinktxdynamicR16    `ext`
	EnhancedskipuplinktxconfiguredR16 *MacParametersxddDiffEnhancedskipuplinktxconfiguredR16 `ext`
}
