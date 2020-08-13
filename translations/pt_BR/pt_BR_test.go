package pt_BR

import (
	"testing"
	"time"

	"github.com/bingoohuang/valid"
	. "github.com/go-playground/assert/v2"
	brazilian_portuguese "github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
)

func TestTranslations(t *testing.T) {
	ptbr := brazilian_portuguese.New()
	uni := ut.New(ptbr, ptbr)
	trans, _ := uni.GetTranslator("pt_BR")

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
		RequiredString    string    `valid:"required"`
		RequiredNumber    int       `valid:"required"`
		RequiredMultiple  []string  `valid:"required"`
		LenString         string    `valid:"len=1"`
		LenNumber         float64   `valid:"len=1113.00"`
		LenMultiple       []string  `valid:"len=7"`
		MinString         string    `valid:"min=1"`
		MinNumber         float64   `valid:"min=1113.00"`
		MinMultiple       []string  `valid:"min=7"`
		MaxString         string    `valid:"max=3"`
		MaxNumber         float64   `valid:"max=1113.00"`
		MaxMultiple       []string  `valid:"max=7"`
		EqString          string    `valid:"eq=3"`
		EqNumber          float64   `valid:"eq=2.33"`
		EqMultiple        []string  `valid:"eq=7"`
		NeString          string    `valid:"ne="`
		NeNumber          float64   `valid:"ne=0.00"`
		NeMultiple        []string  `valid:"ne=0"`
		LtString          string    `valid:"lt=3"`
		LtNumber          float64   `valid:"lt=5.56"`
		LtMultiple        []string  `valid:"lt=2"`
		LtTime            time.Time `valid:"lt"`
		LteString         string    `valid:"lte=3"`
		LteNumber         float64   `valid:"lte=5.56"`
		LteMultiple       []string  `valid:"lte=2"`
		LteTime           time.Time `valid:"lte"`
		GtString          string    `valid:"gt=3"`
		GtNumber          float64   `valid:"gt=5.56"`
		GtMultiple        []string  `valid:"gt=2"`
		GtTime            time.Time `valid:"gt"`
		GteString         string    `valid:"gte=3"`
		GteNumber         float64   `valid:"gte=5.56"`
		GteMultiple       []string  `valid:"gte=2"`
		GteTime           time.Time `valid:"gte"`
		EqFieldString     string    `valid:"eqfield=MaxString"`
		EqCSFieldString   string    `valid:"eqcsfield=Inner.EqCSFieldString"`
		NeCSFieldString   string    `valid:"necsfield=Inner.NeCSFieldString"`
		GtCSFieldString   string    `valid:"gtcsfield=Inner.GtCSFieldString"`
		GteCSFieldString  string    `valid:"gtecsfield=Inner.GteCSFieldString"`
		LtCSFieldString   string    `valid:"ltcsfield=Inner.LtCSFieldString"`
		LteCSFieldString  string    `valid:"ltecsfield=Inner.LteCSFieldString"`
		NeFieldString     string    `valid:"nefield=EqFieldString"`
		GtFieldString     string    `valid:"gtfield=MaxString"`
		GteFieldString    string    `valid:"gtefield=MaxString"`
		LtFieldString     string    `valid:"ltfield=MaxString"`
		LteFieldString    string    `valid:"ltefield=MaxString"`
		AlphaString       string    `valid:"alpha"`
		AlphanumString    string    `valid:"alphanum"`
		NumericString     string    `valid:"numeric"`
		NumberString      string    `valid:"number"`
		HexadecimalString string    `valid:"hexadecimal"`
		HexColorString    string    `valid:"hexcolor"`
		RGBColorString    string    `valid:"rgb"`
		RGBAColorString   string    `valid:"rgba"`
		HSLColorString    string    `valid:"hsl"`
		HSLAColorString   string    `valid:"hsla"`
		Email             string    `valid:"email"`
		URL               string    `valid:"url"`
		URI               string    `valid:"uri"`
		Base64            string    `valid:"base64"`
		Contains          string    `valid:"contains=purpose"`
		ContainsAny       string    `valid:"containsany=!@#$"`
		Excludes          string    `valid:"excludes=text"`
		ExcludesAll       string    `valid:"excludesall=!@#$"`
		ExcludesRune      string    `valid:"excludesrune=☻"`
		ISBN              string    `valid:"isbn"`
		ISBN10            string    `valid:"isbn10"`
		ISBN13            string    `valid:"isbn13"`
		UUID              string    `valid:"uuid"`
		UUID3             string    `valid:"uuid3"`
		UUID4             string    `valid:"uuid4"`
		UUID5             string    `valid:"uuid5"`
		ASCII             string    `valid:"ascii"`
		PrintableASCII    string    `valid:"printascii"`
		MultiByte         string    `valid:"multibyte"`
		DataURI           string    `valid:"datauri"`
		Latitude          string    `valid:"latitude"`
		Longitude         string    `valid:"longitude"`
		SSN               string    `valid:"ssn"`
		IP                string    `valid:"ip"`
		IPv4              string    `valid:"ipv4"`
		IPv6              string    `valid:"ipv6"`
		CIDR              string    `valid:"cidr"`
		CIDRv4            string    `valid:"cidrv4"`
		CIDRv6            string    `valid:"cidrv6"`
		TCPAddr           string    `valid:"tcp_addr"`
		TCPAddrv4         string    `valid:"tcp4_addr"`
		TCPAddrv6         string    `valid:"tcp6_addr"`
		UDPAddr           string    `valid:"udp_addr"`
		UDPAddrv4         string    `valid:"udp4_addr"`
		UDPAddrv6         string    `valid:"udp6_addr"`
		IPAddr            string    `valid:"ip_addr"`
		IPAddrv4          string    `valid:"ip4_addr"`
		IPAddrv6          string    `valid:"ip6_addr"`
		UinxAddr          string    `valid:"unix_addr"` // can't fail from within Go's net package currently, but maybe in the future
		MAC               string    `valid:"mac"`
		IsColor           string    `valid:"iscolor"`
		StrPtrMinLen      *string   `valid:"min=10"`
		StrPtrMaxLen      *string   `valid:"max=1"`
		StrPtrLen         *string   `valid:"len=2"`
		StrPtrLt          *string   `valid:"lt=1"`
		StrPtrLte         *string   `valid:"lte=1"`
		StrPtrGt          *string   `valid:"gt=10"`
		StrPtrGte         *string   `valid:"gte=10"`
		OneOfString       string    `valid:"oneof=red green"`
		OneOfInt          int       `valid:"oneof=5 63"`
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

	test.Excludes = "este é um texto de teste"
	test.ExcludesAll = "Isso é Ótimo!"
	test.ExcludesRune = "Amo isso ☻"

	test.ASCII = "ｶﾀｶﾅ"
	test.PrintableASCII = "ｶﾀｶﾅ"

	test.MultiByte = "1234feerf"

	s := "toolong"
	test.StrPtrMaxLen = &s
	test.StrPtrLen = &s

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
			expected: "IsColor deve ser uma cor válida",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC deve conter um endereço MAC válido",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr deve ser um endereço IP resolvível",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 deve ser um endereço IPv4 resolvível",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 deve ser um endereço IPv6 resolvível",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr deve ser um endereço UDP válido",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 deve ser um endereço IPv4 UDP válido",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 deve ser um endereço IPv6 UDP válido",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr deve ser um endereço TCP válido",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 deve ser um endereço IPv4 TCP válido",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 deve ser um endereço IPv6 TCP válido",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR deve conter uma notação CIDR válida",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 deve conter uma notação CIDR válida para um endereço IPv4",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 deve conter uma notação CIDR válida para um endereço IPv6",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN deve ser um número SSN válido",
		},
		{
			ns:       "Test.IP",
			expected: "IP deve ser um endereço de IP válido",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 deve ser um endereço IPv4 válido",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 deve ser um endereço IPv6 válido",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI deve conter um URI data válido",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude deve conter uma coordenada de latitude válida",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude deve conter uma coordenada de longitude válida",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte deve conter caracteres multibyte",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII deve conter apenas caracteres ascii",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII deve conter apenas caracteres ascii imprimíveis",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID deve ser um UUID válido",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 deve ser um UUID versão 3 válido",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 deve ser um UUID versão 4 válido",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 deve ser um UUID versão 5 válido",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN deve ser um número ISBN válido",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 deve ser um número ISBN-10 válido",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 deve ser um número ISBN-13 válido",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes não deve conter o texto 'text'",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll não deve conter nenhum dos caracteres '!@#$'",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune não deve conter '☻'",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny deve conter pelo menos um dos caracteres '!@#$'",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains deve conter o texto 'purpose'",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 deve ser uma string Base64 válida",
		},
		{
			ns:       "Test.Email",
			expected: "Email deve ser um endereço de e-mail válido",
		},
		{
			ns:       "Test.URL",
			expected: "URL deve ser uma URL válida",
		},
		{
			ns:       "Test.URI",
			expected: "URI deve ser uma URI válida",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString deve ser uma cor RGB válida",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString deve ser uma cor RGBA válida",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString deve ser uma cor HSL válida",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString deve ser uma cor HSLA válida",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString deve ser um hexadecimal válido",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString deve ser uma cor HEX válida",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString deve ser um número válido",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString deve ser um valor numérico válido",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString deve conter caracteres alfanuméricos",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString deve conter caracteres alfabéticos",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString deve ser menor que MaxString",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString deve ser menor ou igual a MaxString",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString deve ser maior do que MaxString",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString deve ser maior ou igual a MaxString",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString não deve ser igual a EqFieldString",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString deve ser menor que Inner.LtCSFieldString",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString deve ser menor ou igual a Inner.LteCSFieldString",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString deve ser maior do que Inner.GtCSFieldString",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString deve ser maior ou igual a Inner.GteCSFieldString",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString não deve ser igual a Inner.NeCSFieldString",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString deve ser igual a Inner.EqCSFieldString",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString deve ser igual a MaxString",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString deve ter pelo menos 3 caracteres",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber deve ser 5,56 ou superior",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple deve conter pelo menos 2 itens",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime deve ser maior ou igual à Data e Hora atual",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString deve ter mais de 3 caracteres",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber deve ser maior do que 5,56",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple deve conter mais de 2 itens",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime deve ser maior que a Data e Hora atual",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString deve ter no máximo 3 caracteres",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber deve ser 5,56 ou menor",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple deve conter no máximo 2 itens",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime deve ser menor ou igual à Data e Hora atual",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString deve ter menos de 3 caracteres",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber deve ser menor que 5,56",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple deve conter menos de 2 itens",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime deve ser inferior à Data e Hora atual",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString não deve ser igual a ",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber não deve ser igual a 0.00",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple não deve ser igual a 0",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString não é igual a 3",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber não é igual a 2.33",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple não é igual a 7",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString deve ter no máximo 3 caracteres",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber deve ser 1.113,00 ou menor",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple deve conter no máximo 7 itens",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString deve ter pelo menos 1 caractere",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber deve ser 1.113,00 ou superior",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple deve conter pelo menos 7 itens",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString deve ter 1 caractere",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber deve ser igual a 1.113,00",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple deve conter 7 itens",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString é um campo requerido",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber é um campo requerido",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple é um campo requerido",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen deve ter pelo menos 10 caracteres",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen deve ter no máximo 1 caractere",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen deve ter 2 caracteres",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt deve ter menos de 1 caractere",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte deve ter no máximo 1 caractere",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt deve ter mais de 10 caracteres",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte deve ter pelo menos 10 caracteres",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString deve ser um de [red green]",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt deve ser um de [5 63]",
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
