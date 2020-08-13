package ru

import (
	"log"
	//"github.com/rustery/validator"
	"testing"
	"time"

	"github.com/bingoohuang/valid"
	. "github.com/go-playground/assert/v2"
	russian "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
)

func TestTranslations(t *testing.T) {
	ru := russian.New()
	uni := ut.New(ru, ru)
	trans, _ := uni.GetTranslator("ru")

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

	s := "toolong"
	test.StrPtrMaxLen = &s
	test.StrPtrLen = &s

	test.UniqueSlice = []string{"1234", "1234"}
	test.UniqueMap = map[string]string{"key1": "1234", "key2": "1234"}

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
			expected: "IsColor должен быть цветом",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC должен содержать MAC адрес",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr должен быть распознаваемым IP адресом",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 должен быть распознаваемым IPv4 адресом",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 должен быть распознаваемым IPv6 адресом",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr должен быть UDP адресом",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 должен быть IPv4 UDP адресом",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 должен быть IPv6 UDP адресом",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr должен быть TCP адресом",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 должен быть IPv4 TCP адресом",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 должен быть IPv6 TCP адресом",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR должен содержать CIDR обозначения",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 должен содержать CIDR обозначения для IPv4 адреса",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 должен содержать CIDR обозначения для IPv6 адреса",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN должен быть SSN номером",
		},
		{
			ns:       "Test.IP",
			expected: "IP должен быть IP адресом",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 должен быть IPv4 адресом",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 должен быть IPv6 адресом",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI должен содержать Data URI",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude должен содержать координаты широты",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude должен содержать координаты долготы",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte должен содержать мультибайтные символы",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII должен содержать только ascii символы",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII должен содержать только доступные для печати ascii символы",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID должен быть UUID",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 должен быть UUID 3 версии",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 должен быть UUID 4 версии",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 должен быть UUID 5 версии",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN должен быть ISBN номером",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 должен быть ISBN-10 номером",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 должен быть ISBN-13 номером",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes не должен содержать текст 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll не должен содержать символы '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune не должен содержать '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny должен содержать минимум один из символов '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains должен содержать текст 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 должен быть Base64 строкой",
		},
		{
			ns:       "Test.Email",
			expected: "Email должен быть email адресом",
		},
		{
			ns:       "Test.URL",
			expected: "URL должен быть URL",
		},
		{
			ns:       "Test.URI",
			expected: "URI должен быть URI",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString должен быть RGB цветом",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString должен быть RGBA цветом",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString должен быть HSL цветом",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString должен быть HSLA цветом",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString должен быть шестнадцатеричной строкой",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString должен быть HEX цветом",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString должен быть цифрой",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString должен быть цифровым значением",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString должен содержать только буквы и цифры",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString должен содержать только буквы",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString должен быть менее MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString должен быть менее или равен MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString должен быть больше MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString должен быть больше или равен MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString не должен быть равен EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString должен быть менее Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString должен быть менее или равен Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString должен быть больше Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString должен быть больше или равен Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString не должен быть равен Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString должен быть равен Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString должен быть равен MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString должен содержать минимум 3 символы",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber должен быть больше или равно 5.56",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple должен содержать минимум 2 элементы",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime должна быть позже или равна текущему моменту",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString должен быть длиннее 3 символы",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber должен быть больше 5.56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple должен содержать более 2 элементы",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime должна быть позже текущего момента",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString должен содержать максимум 3 символы",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber должен быть менее или равен 5.56",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple должен содержать максимум 2 элементы",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime must be less than or equal to the current Date & Time",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString должен иметь менее 3 символы",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber должен быть менее 5.56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple должен содержать менее 2 элементы",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime must be less than the current Date & Time",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString должен быть не равен ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber должен быть не равен 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple должен быть не равен 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString не равен 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber не равен 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple не равен 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString должен содержать максимум 3 символы",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber должен быть меньше или равно 1,113.00",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple должен содержать максимум 7 элементы",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString должен содержать минимум 1 символ",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber должен быть больше или равно 1,113.00",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple должен содержать минимум 7 элементы",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString должен быть длиной в 1 символ",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber должен быть равен 1,113.00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple должен содержать 7 элементы",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString обязательное поле",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber обязательное поле",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple обязательное поле",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen должен содержать минимум 10 символы",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen должен содержать максимум 1 символ",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen должен быть длиной в 2 символы",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt должен иметь менее 1 символ",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte должен содержать максимум 1 символ",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt должен быть длиннее 10 символы",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte должен содержать минимум 10 символы",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString должен быть одним из [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt должен быть одним из [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice должен содержать уникальные значения",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray должен содержать уникальные значения",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap должен содержать уникальные значения",
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

		log.Println(fe)

		NotEqual(t, fe, nil)
		Equal(t, tt.expected, fe.Translate(trans))
	}
}
