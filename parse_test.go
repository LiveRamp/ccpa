package ccpa_test

import (
	"github.com/go-check/check"

	"github.com/LiveRamp/ccpa"
)


type ParseTestSuite struct{}

func (p *ParseTestSuite) TestErrorCases(c *check.C) {
	var tcs = []struct{
		s string
		exp string
	}{
		{
			s: "1YY",
			exp: "consent string should be exactly 4 characters in length",
		},
		{
			s: "A---",
			exp: "error parsing Specification Version number: .*",
		},
		{
			s: "1XYY",
			exp: "unexpected value for Explicit Notice field",
		},
		{
			s: "1YXY",
			exp: "unexpected value for Opt-Out Sale field",
		},
		{
			s: "1YYX",
			exp: "unexpected value for LSPA field",
		},
	}

	for _, tc := range tcs {
		c.Log(tc)

		var con, err = ccpa.Parse(tc.s)
		c.Check(con, check.IsNil)
		c.Check(err, check.ErrorMatches, tc.exp)
	}
}

func (p *ParseTestSuite) TestParse(c *check.C) {
	var tcs = []struct{
		s string
		exp *ccpa.Consent
	}{
		// First 3 are the example strings from the IAB spec.
		// The last test is added for sanity.
		{
			s: "1YYN",
			exp: &ccpa.Consent{
				Version:  1,
				Explicit: ccpa.Yes,
				OptOut:   ccpa.Yes,
				LSPA:     ccpa.No,
			},
		},
		{
			s: "1NYY",
			exp: &ccpa.Consent{
				Version:  1,
				Explicit: ccpa.No,
				OptOut:   ccpa.Yes,
				LSPA:     ccpa.Yes,
			},
		},
		{
			s: "1---",
			exp: &ccpa.Consent{
				Version:  1,
				Explicit: ccpa.NotApplicable,
				OptOut:   ccpa.NotApplicable,
				LSPA:     ccpa.NotApplicable,
			},
		},
		{
			s: "2YN-",
			exp: &ccpa.Consent{
				Version:  2,
				Explicit: ccpa.Yes,
				OptOut:   ccpa.No,
				LSPA:     ccpa.NotApplicable,
			},
		},

	}

	for _, tc := range tcs {
		c.Log(tc)

		var con, err = ccpa.Parse(tc.s)
		c.Check(err, check.IsNil)
		c.Check(con, check.DeepEquals, tc.exp)
	}
}

var _ = check.Suite(&ParseTestSuite{})
