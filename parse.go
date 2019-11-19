// Package ccpa offers functionality to parse IAB consent strings
// required for compliance with the California Consumer Privacy Act.
package ccpa

import (
	"errors"
	"strconv"
)

// Value is an enum type to represent the possible values of the
// enum fields in the Consent type.
type Value int

const (
	// No represents a 'no' in a given enum field.
	No Value = iota
	// Yes represents a 'yes' in a given enum field.
	Yes
	// NotApplicable means a given enum field does not apply.
	NotApplicable
)

var runeToValue = map[uint8]Value{
	'N': No,
	'Y': Yes,
	'-': NotApplicable,
}

// Consent represents the parsed fields of the consent string.
type Consent struct {
	// Version is the Version of this string specification used to encode the string.
	Version int
	// Explicit specifies if Explicit notice has been provided as required by 1798.115(d)
	// of the CCPA and the opportunity to opt out of the sale of their data
	// pursuant to 1798.120 and 1798.135 of the CCPA.
	Explicit Value
	// OptOut specifies if user opted-out of the sale of his or personal information
	// pursuant to 1798.120 and 1798.135.
	OptOut Value
	// LSPA states wheter a publisher is a signatory to the IAB Limited Service Provider
	// Agreement (LSPA) and the publisher declares that the transaction is covered as a
	// "Covered Opt Out Transaction" or a "Non Opt Out Transaction" as those terms are
	// defined in the Agreement.
	LSPA Value
}

// Parse takes a consent string specified by the IAB 1.0 spec which can be found here:
// https://iabtechlab.com/wp-content/uploads/2019/11/U.S.-Privacy-String-v1.0-IAB-Tech-Lab.pdf
// This string is expected to be passed by publishers (or daisy chained in any redirection)
// via the URL parameter 'us_consent'. It returns a *Consent and nil error upon successful
// parse, or a nil *Consent and error on error.
func Parse(s string) (*Consent, error) {
	if len(s) != 4 {
		return nil, errors.New("consent string should be exactly 4 characters in length")
	}

	var v, err = strconv.Atoi(string(s[0]))
	if err != nil {
		return nil, errors.New("error parsing Specification Version number: " + err.Error())
	}

	var exp, ok = runeToValue[s[1]]
	if !ok {
		return nil, errors.New("unexpected value for Explicit Notice field")
	}

	var opt Value
	opt, ok = runeToValue[s[2]]
	if !ok {
		return nil, errors.New("unexpected value for Opt-Out Sale field")
	}

	var lspa Value
	lspa, ok = runeToValue[s[3]]
	if !ok {
		return nil, errors.New("unexpected value for LSPA field")
	}

	return &Consent{
		Version:  v,
		Explicit: exp,
		OptOut:   opt,
		LSPA:     lspa,
	}, nil
}
