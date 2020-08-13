package en

import (
	"testing"
	"time"

	"github.com/bingoohuang/valid"
	. "github.com/go-playground/assert/v2"
	english "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

func TestTranslations(t *testing.T) {
	eng := english.New()
	uni := ut.New(eng, eng)
	trans, _ := uni.GetTranslator("en")

	va := valid.New()

	err := RegisterDefaultTranslations(va, trans)
	Equal(t, err, nil)

	type Inner struct {
		EqCSFieldString  string
		NeCSFieldString  string
		GtCSFieldString  string
		GteCSFieldString string
		LtCSFieldString  string
		LteCSFieldString string
	}

	type Test struct {
		Inner             Inner
		RequiredString    string            `valid:"required"`
		RequiredNumber    int               `valid:"required"`
		RequiredMultiple  []string          `valid:"required"`
		LenString         string            `valid:"len=1"`
		LenNumber         float64           `valid:"len=1113.00"`
		LenMultiple       []string          `valid:"len=7"`
		MinString         string            `valid:"min=1"`
		MinNumber         float64           `valid:"min=1113.00"`
		MinMultiple       []string          `valid:"min=7"`
		MaxString         string            `valid:"max=3"`
		MaxNumber         float64           `valid:"max=1113.00"`
		MaxMultiple       []string          `valid:"max=7"`
		EqString          string            `valid:"eq=3"`
		EqNumber          float64           `valid:"eq=2.33"`
		EqMultiple        []string          `valid:"eq=7"`
		NeString          string            `valid:"ne="`
		NeNumber          float64           `valid:"ne=0.00"`
		NeMultiple        []string          `valid:"ne=0"`
		LtString          string            `valid:"lt=3"`
		LtNumber          float64           `valid:"lt=5.56"`
		LtMultiple        []string          `valid:"lt=2"`
		LtTime            time.Time         `valid:"lt"`
		LteString         string            `valid:"lte=3"`
		LteNumber         float64           `valid:"lte=5.56"`
		LteMultiple       []string          `valid:"lte=2"`
		LteTime           time.Time         `valid:"lte"`
		GtString          string            `valid:"gt=3"`
		GtNumber          float64           `valid:"gt=5.56"`
		GtMultiple        []string          `valid:"gt=2"`
		GtTime            time.Time         `valid:"gt"`
		GteString         string            `valid:"gte=3"`
		GteNumber         float64           `valid:"gte=5.56"`
		GteMultiple       []string          `valid:"gte=2"`
		GteTime           time.Time         `valid:"gte"`
		EqFieldString     string            `valid:"eqfield=MaxString"`
		EqCSFieldString   string            `valid:"eqcsfield=Inner.EqCSFieldString"`
		NeCSFieldString   string            `valid:"necsfield=Inner.NeCSFieldString"`
		GtCSFieldString   string            `valid:"gtcsfield=Inner.GtCSFieldString"`
		GteCSFieldString  string            `valid:"gtecsfield=Inner.GteCSFieldString"`
		LtCSFieldString   string            `valid:"ltcsfield=Inner.LtCSFieldString"`
		LteCSFieldString  string            `valid:"ltecsfield=Inner.LteCSFieldString"`
		NeFieldString     string            `valid:"nefield=EqFieldString"`
		GtFieldString     string            `valid:"gtfield=MaxString"`
		GteFieldString    string            `valid:"gtefield=MaxString"`
		LtFieldString     string            `valid:"ltfield=MaxString"`
		LteFieldString    string            `valid:"ltefield=MaxString"`
		AlphaString       string            `valid:"alpha"`
		AlphanumString    string            `valid:"alphanum"`
		NumericString     string            `valid:"numeric"`
		NumberString      string            `valid:"number"`
		HexadecimalString string            `valid:"hexadecimal"`
		HexColorString    string            `valid:"hexcolor"`
		RGBColorString    string            `valid:"rgb"`
		RGBAColorString   string            `valid:"rgba"`
		HSLColorString    string            `valid:"hsl"`
		HSLAColorString   string            `valid:"hsla"`
		Email             string            `valid:"email"`
		URL               string            `valid:"url"`
		URI               string            `valid:"uri"`
		Base64            string            `valid:"base64"`
		Contains          string            `valid:"contains=purpose"`
		ContainsAny       string            `valid:"containsany=!@#$"`
		Excludes          string            `valid:"excludes=text"`
		ExcludesAll       string            `valid:"excludesall=!@#$"`
		ExcludesRune      string            `valid:"excludesrune=☻"`
		ISBN              string            `valid:"isbn"`
		ISBN10            string            `valid:"isbn10"`
		ISBN13            string            `valid:"isbn13"`
		UUID              string            `valid:"uuid"`
		UUID3             string            `valid:"uuid3"`
		UUID4             string            `valid:"uuid4"`
		UUID5             string            `valid:"uuid5"`
		ASCII             string            `valid:"ascii"`
		PrintableASCII    string            `valid:"printascii"`
		MultiByte         string            `valid:"multibyte"`
		DataURI           string            `valid:"datauri"`
		Latitude          string            `valid:"latitude"`
		Longitude         string            `valid:"longitude"`
		SSN               string            `valid:"ssn"`
		IP                string            `valid:"ip"`
		IPv4              string            `valid:"ipv4"`
		IPv6              string            `valid:"ipv6"`
		CIDR              string            `valid:"cidr"`
		CIDRv4            string            `valid:"cidrv4"`
		CIDRv6            string            `valid:"cidrv6"`
		TCPAddr           string            `valid:"tcp_addr"`
		TCPAddrv4         string            `valid:"tcp4_addr"`
		TCPAddrv6         string            `valid:"tcp6_addr"`
		UDPAddr           string            `valid:"udp_addr"`
		UDPAddrv4         string            `valid:"udp4_addr"`
		UDPAddrv6         string            `valid:"udp6_addr"`
		IPAddr            string            `valid:"ip_addr"`
		IPAddrv4          string            `valid:"ip4_addr"`
		IPAddrv6          string            `valid:"ip6_addr"`
		UinxAddr          string            `valid:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		MAC               string            `valid:"mac"`
		IsColor           string            `valid:"iscolor"`
		StrPtrMinLen      *string           `valid:"min=10"`
		StrPtrMaxLen      *string           `valid:"max=1"`
		StrPtrLen         *string           `valid:"len=2"`
		StrPtrLt          *string           `valid:"lt=1"`
		StrPtrLte         *string           `valid:"lte=1"`
		StrPtrGt          *string           `valid:"gt=10"`
		StrPtrGte         *string           `valid:"gte=10"`
		OneOfString       string            `valid:"oneof=red green"`
		OneOfInt          int               `valid:"oneof=5 63"`
		UniqueSlice       []string          `valid:"unique"`
		UniqueArray       [3]string         `valid:"unique"`
		UniqueMap         map[string]string `valid:"unique"`
		JSONString        string            `valid:"json"`
		LowercaseString   string            `valid:"lowercase"`
		UppercaseString   string            `valid:"uppercase"`
		Datetime          string            `valid:"datetime=2006-01-02"`
	}

	var test Test

	test.Inner.EqCSFieldString = "1234"
	test.Inner.GtCSFieldString = "1234"
	test.Inner.GteCSFieldString = "1234"

	test.MaxString = "1234"
	test.MaxNumber = 2000
	test.MaxMultiple = make([]string, 9)

	test.LtString = "1234"
	test.LtNumber = 6
	test.LtMultiple = make([]string, 3)
	test.LtTime = time.Now().Add(time.Hour * 24)

	test.LteString = "1234"
	test.LteNumber = 6
	test.LteMultiple = make([]string, 3)
	test.LteTime = time.Now().Add(time.Hour * 24)

	test.LtFieldString = "12345"
	test.LteFieldString = "12345"

	test.LtCSFieldString = "1234"
	test.LteCSFieldString = "1234"

	test.AlphaString = "abc3"
	test.AlphanumString = "abc3!"
	test.NumericString = "12E.00"
	test.NumberString = "12E"

	test.Excludes = "this is some test text"
	test.ExcludesAll = "This is Great!"
	test.ExcludesRune = "Love it ☻"

	test.ASCII = "ｶﾀｶﾅ"
	test.PrintableASCII = "ｶﾀｶﾅ"

	test.MultiByte = "1234feerf"

	test.LowercaseString = "ABCDEFG"
	test.UppercaseString = "abcdefg"

	s := "toolong"
	test.StrPtrMaxLen = &s
	test.StrPtrLen = &s

	test.UniqueSlice = []string{"1234", "1234"}
	test.UniqueMap = map[string]string{"key1": "1234", "key2": "1234"}
	test.Datetime = "2008-Feb-01"

	err = va.Struct(test)
	NotEqual(t, err, nil)

	errs, ok := err.(valid.ValidationErrors)
	Equal(t, ok, true)

	tests := []struct {
		ns       string
		expected string
	}{
		{
			ns:       "Test.IsColor",
			expected: "IsColor must be a valid color",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC must contain a valid MAC address",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr must be a resolvable IP address",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 must be a resolvable IPv4 address",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 must be a resolvable IPv6 address",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr must be a valid UDP address",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 must be a valid IPv4 UDP address",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 must be a valid IPv6 UDP address",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr must be a valid TCP address",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 must be a valid IPv4 TCP address",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 must be a valid IPv6 TCP address",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR must contain a valid CIDR notation",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 must contain a valid CIDR notation for an IPv4 address",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 must contain a valid CIDR notation for an IPv6 address",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN must be a valid SSN number",
		},
		{
			ns:       "Test.IP",
			expected: "IP must be a valid IP address",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 must be a valid IPv4 address",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 must be a valid IPv6 address",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI must contain a valid Data URI",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude must contain valid latitude coordinates",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude must contain a valid longitude coordinates",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte must contain multibyte characters",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII must contain only ascii characters",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII must contain only printable ascii characters",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID must be a valid UUID",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 must be a valid version 3 UUID",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 must be a valid version 4 UUID",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 must be a valid version 5 UUID",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN must be a valid ISBN number",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 must be a valid ISBN-10 number",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 must be a valid ISBN-13 number",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes cannot contain the text 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll cannot contain any of the following characters '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune cannot contain the following '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny must contain at least one of the following characters '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains must contain the text 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 must be a valid Base64 string",
		},
		{
			ns:       "Test.Email",
			expected: "Email must be a valid email address",
		},
		{
			ns:       "Test.URL",
			expected: "URL must be a valid URL",
		},
		{
			ns:       "Test.URI",
			expected: "URI must be a valid URI",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString must be a valid RGB color",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString must be a valid RGBA color",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString must be a valid HSL color",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString must be a valid HSLA color",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString must be a valid hexadecimal",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString must be a valid HEX color",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString must be a valid number",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString must be a valid numeric value",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString can only contain alphanumeric characters",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString can only contain alphabetic characters",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString must be less than MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString must be less than or equal to MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString must be greater than MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString must be greater than or equal to MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString cannot be equal to EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString must be less than Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString must be less than or equal to Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString must be greater than Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString must be greater than or equal to Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString cannot be equal to Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString must be equal to Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString must be equal to MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString must be at least 3 characters in length",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber must be 5.56 or greater",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple must contain at least 2 items",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime must be greater than or equal to the current Date & Time",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString must be greater than 3 characters in length",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber must be greater than 5.56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple must contain more than 2 items",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime must be greater than the current Date & Time",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString must be at maximum 3 characters in length",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber must be 5.56 or less",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple must contain at maximum 2 items",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime must be less than or equal to the current Date & Time",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString must be less than 3 characters in length",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber must be less than 5.56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple must contain less than 2 items",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime must be less than the current Date & Time",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString should not be equal to ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber should not be equal to 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple should not be equal to 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString is not equal to 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber is not equal to 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple is not equal to 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString must be a maximum of 3 characters in length",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber must be 1,113.00 or less",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple must contain at maximum 7 items",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString must be at least 1 character in length",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber must be 1,113.00 or greater",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple must contain at least 7 items",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString must be 1 character in length",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber must be equal to 1,113.00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple must contain 7 items",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString is a required field",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber is a required field",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple is a required field",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen must be at least 10 characters in length",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen must be a maximum of 1 character in length",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen must be 2 characters in length",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt must be less than 1 character in length",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte must be at maximum 1 character in length",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt must be greater than 10 characters in length",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte must be at least 10 characters in length",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString must be one of [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt must be one of [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice must contain unique values",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray must contain unique values",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap must contain unique values",
		},
		{
			ns:       "Test.JSONString",
			expected: "JSONString must be a valid json string",
		},
		{
			ns:       "Test.LowercaseString",
			expected: "LowercaseString must be a lowercase string",
		},
		{
			ns:       "Test.UppercaseString",
			expected: "UppercaseString must be an uppercase string",
		},
		{
			ns:       "Test.Datetime",
			expected: "Datetime does not match the 2006-01-02 format",
		},
	}

	for _, tt := range tests {

		var fe valid.FieldError

		for _, e := range errs {
			if tt.ns == e.Namespace() {
				fe = e
				break
			}
		}

		NotEqual(t, fe, nil)
		Equal(t, tt.expected, fe.Translate(trans))
	}
}
