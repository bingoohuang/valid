package es

import (
	"testing"
	"time"

	"github.com/bingoohuang/valid"
	. "github.com/go-playground/assert/v2"
	spanish "github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
)

func TestTranslations(t *testing.T) {
	spa := spanish.New()
	uni := ut.New(spa, spa)
	trans, _ := uni.GetTranslator("es")

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
			expected: "IsColor debe ser un color válido",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC debe contener una dirección MAC válida",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr debe ser una dirección IP resoluble",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 debe ser una dirección IPv4 resoluble",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 debe ser una dirección IPv6 resoluble",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr debe ser una dirección UDP válida",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 debe ser una dirección IPv4 UDP válida",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 debe ser una dirección IPv6 UDP válida",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr debe ser una dirección TCP válida",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 debe ser una dirección IPv4 TCP válida",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 debe ser una dirección IPv6 TCP válida",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR debe contener una anotación válida del CIDR",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 debe contener una notación CIDR válida para una dirección IPv4",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 debe contener una notación CIDR válida para una dirección IPv6",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN debe ser un número válido de SSN",
		},
		{
			ns:       "Test.IP",
			expected: "IP debe ser una dirección IP válida",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 debe ser una dirección IPv4 válida",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 debe ser una dirección IPv6 válida",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI debe contener un URI de datos válido",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude debe contener coordenadas de latitud válidas",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude debe contener unas coordenadas de longitud válidas",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte debe contener caracteres multibyte",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII debe contener sólo caracteres ascii",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII debe contener sólo caracteres ASCII imprimibles",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID debe ser un UUID válido",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 debe ser una versión válida 3 UUID",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 debe ser una versión válida 4 UUID",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 debe ser una versión válida 5 UUID",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN debe ser un número ISBN válido",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 debe ser un número ISBN-10 válido",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 debe ser un número ISBN-13 válido",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes no puede contener el texto 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll no puede contener ninguno de los siguientes caracteres '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune no puede contener lo siguiente '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny debe contener al menos uno de los siguientes caracteres '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains debe contener el texto 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 debe ser una cadena de Base64 válida",
		},
		{
			ns:       "Test.Email",
			expected: "Email debe ser una dirección de correo electrónico válida",
		},
		{
			ns:       "Test.URL",
			expected: "URL debe ser un URL válido",
		},
		{
			ns:       "Test.URI",
			expected: "URI debe ser una URI válida",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString debe ser un color RGB válido",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString debe ser un color RGBA válido",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString debe ser un color HSL válido",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString debe ser un color HSL válido",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString debe ser un hexadecimal válido",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString debe ser un color HEX válido",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString debe ser un número válido",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString debe ser un valor numérico válido",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString sólo puede contener caracteres alfanuméricos",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString sólo puede contener caracteres alfabéticos",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString debe ser menor que MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString debe ser menor o igual a MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString debe ser mayor que MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString debe ser mayor o igual a MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString no puede ser igual a EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString debe ser menor que Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString debe ser menor o igual a Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString debe ser mayor que Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString debe ser mayor o igual a Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString no puede ser igual a Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString debe ser igual a Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString debe ser igual a MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString debe tener al menos 3 caracteres de longitud",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber debe ser 5,56 o mayor",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple debe contener al menos 2 elementos",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime debe ser después o durante la fecha y hora actuales",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString debe ser mayor que 3 caracteres en longitud",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber debe ser mayor que 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple debe contener más de 2 elementos",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime debe ser después de la fecha y hora actual",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString debe tener un máximo de 3 caracteres de longitud",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber debe ser 5,56 o menos",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple debe contener como máximo 2 elementos",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime debe ser antes o durante la fecha y hora actual",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString debe tener menos de 3 caracteres de longitud",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber debe ser menos de 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple debe contener menos de 2 elementos",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime debe ser antes de la fecha y hora actual",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString no debería ser igual a ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber no debería ser igual a 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple no debería ser igual a 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString no es igual a 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber no es igual a 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple no es igual a 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString debe tener un máximo de 3 caracteres de longitud",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber debe ser 1.113,00 o menos",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple debe contener como máximo 7 elementos",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString debe tener al menos 1 carácter de longitud",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber debe ser 1.113,00 o más",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple debe contener al menos 7 elementos",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString debe tener 1 carácter de longitud",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber debe ser igual a 1.113,00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple debe contener 7 elementos",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString es un campo requerido",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber es un campo requerido",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple es un campo requerido",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen debe tener al menos 10 caracteres de longitud",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen debe tener un máximo de 1 carácter de longitud",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen debe tener 2 caracteres de longitud",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt debe tener menos de 1 carácter de longitud",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte debe tener un máximo de 1 carácter de longitud",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt debe ser mayor que 10 caracteres en longitud",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte debe tener al menos 10 caracteres de longitud",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString debe ser uno de [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt debe ser uno de [5 63]",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice debe contener valores únicos",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray debe contener valores únicos",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap debe contener valores únicos",
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
