package dhcpv6

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestParseMessageOptionsWithORO(t *testing.T) {
	buf := []byte{
		0, 6, // ORO option
		0, 2, // length
		0, 3, // IANA Option
		0, 6, // ORO
		0, 2, // length
		0, 4, // IATA
	}

	want := OptionCodes{OptionIANA, OptionIATA}
	var mo MessageOptions
	if err := mo.FromBytes(buf); err != nil {
		t.Errorf("FromBytes = %v", err)
	} else if got := mo.RequestedOptions(); !reflect.DeepEqual(got, want) {
		t.Errorf("ORO = %v, want %v", got, want)
	}
}

func TestOptRequestedOption(t *testing.T) {
	expected := []byte{0, 1, 0, 2}
	var o optRequestedOption
	err := o.FromBytes(expected)
	require.NoError(t, err, "ParseOptRequestedOption() correct options should not error")
}

func TestOptRequestedOptionParseOptRequestedOptionTooShort(t *testing.T) {
	buf := []byte{0, 1, 0}
	var o optRequestedOption
	err := o.FromBytes(buf)
	require.Error(t, err, "A short option should return an error (must be divisible by 2)")
}

func TestOptRequestedOptionString(t *testing.T) {
	buf := []byte{0, 1, 0, 2}
	var o optRequestedOption
	err := o.FromBytes(buf)
	require.NoError(t, err)
	require.Contains(
		t,
		o.String(),
		"Client ID, Server ID",
		"String() should contain the options specified",
	)
	o.OptionCodes = append(o.OptionCodes, 12345)
	require.Contains(
		t,
		o.String(),
		"unknown",
		"String() should contain 'Unknown' for an illegal option",
	)
}
