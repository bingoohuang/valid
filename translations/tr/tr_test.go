package tr

import (
	"testing"
	"time"

	"github.com/bingoohuang/valid"
	. "github.com/go-playground/assert/v2"
	turkish "github.com/go-playground/locales/tr"
	ut "github.com/go-playground/universal-translator"
)

func TestTranslations(t *testing.T) {
	tr := turkish.New()
	uni := ut.New(tr, tr)
	trans, _ := uni.GetTranslator("tr")

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
			expected: "IsColor geçerli bir renk olmalıdır",
		},
		{
			ns:       "Test.MAC",
			expected: "MAC geçerli bir MAC adresi içermelidir",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddr çözülebilir bir IP adresi olmalıdır",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4 çözülebilir bir IPv4 adresi olmalıdır",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6 çözülebilir bir IPv6 adresi olmalıdır",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddr geçerli bir UDP adresi olmalıdır",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4 geçerli bir IPv4 UDP adresi olmalıdır",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6 geçerli bir IPv6 UDP adresi olmalıdır",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddr geçerli bir TCP adresi olmalıdır",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4 geçerli bir IPv4 TCP adresi olmalıdır",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6 geçerli bir IPv6 TCP adresi olmalıdır",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDR geçerli bir CIDR gösterimi içermelidir",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4 bir IPv4 adresi için geçerli bir CIDR gösterimi içermelidir",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6 bir IPv6 adresi için geçerli bir CIDR gösterimi içermelidir",
		},
		{
			ns:       "Test.SSN",
			expected: "SSN geçerli bir SSN numarası olmalıdır",
		},
		{
			ns:       "Test.IP",
			expected: "IP geçerli bir IP adresi olmalıdır",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4 geçerli bir IPv4 adresi olmalıdır",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6 geçerli bir IPv6 adresi olmalıdır",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURI geçerli bir Veri URI içermelidir",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitude geçerli bir enlem koordinatı içermelidir",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitude geçerli bir boylam koordinatı içermelidir",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByte çok baytlı karakterler içermelidir",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCII yalnızca ascii karakterler içermelidir",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCII yalnızca yazdırılabilir ascii karakterleri içermelidir",
		},
		{
			ns:       "Test.UUID",
			expected: "UUID geçerli bir UUID olmalıdır",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3 geçerli bir sürüm 3 UUID olmalıdır",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4 geçerli bir sürüm 4 UUID olmalıdır",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5 geçerli bir sürüm 5 UUID olmalıdır",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBN geçerli bir ISBN numarası olmalıdır",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10 geçerli bir ISBN-10 numarası olmalıdır",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13 geçerli bir ISBN-13 numarası olmalıdır",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludes, 'text' metnini içeremez",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAll, '!@#$' karakterlerinden hiçbirini içeremez",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRune, '☻' ifadesini içeremez",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAny, '!@#$' karakterlerinden en az birini içermelidir",
		},
		{
			ns:       "Test.Contains",
			expected: "Contains, 'purpose' metnini içermelidir",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64 geçerli bir Base64 karakter dizesi olmalıdır",
		},
		{
			ns:       "Test.Email",
			expected: "Email geçerli bir e-posta adresi olmalıdır",
		},
		{
			ns:       "Test.URL",
			expected: "URL geçerli bir URL olmalıdır",
		},
		{
			ns:       "Test.URI",
			expected: "URI geçerli bir URI olmalıdır",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorString geçerli bir RGB rengi olmalıdır",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorString geçerli bir RGBA rengi olmalıdır",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorString geçerli bir HSL rengi olmalıdır",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorString geçerli bir HSLA rengi olmalıdır",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalString geçerli bir onaltılık olmalıdır",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorString geçerli bir HEX rengi olmalıdır",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberString geçerli bir sayı olmalıdır",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericString geçerli bir sayısal değer olmalıdır",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumString yalnızca alfanümerik karakterler içerebilir",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaString yalnızca alfabetik karakterler içerebilir",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldString, MaxString değerinden küçük olmalıdır",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldString, MaxString değerinden küçük veya ona eşit olmalıdır",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldString, MaxString değerinden büyük olmalıdır",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldString, MaxString değerinden büyük veya ona eşit olmalıdır",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldString, EqFieldString değerine eşit olmamalıdır",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldString, Inner.LtCSFieldString değerinden küçük olmalıdır",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldString, Inner.LteCSFieldString değerinden küçük veya ona eşit olmalıdır",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldString, Inner.GtCSFieldString değerinden büyük olmalıdır",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldString, Inner.GteCSFieldString değerinden küçük veya ona eşit olmalıdır",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldString, Inner.NeCSFieldString değerine eşit olmamalıdır",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldString, Inner.EqCSFieldString değerine eşit olmalıdır",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldString, MaxString değerine eşit olmalıdır",
		},
		{
			ns:       "Test.GteString",
			expected: "GteString en az 3 karakter uzunluğunda olmalıdır",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumber, 5,56 veya daha büyük olmalıdır",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultiple en az 2 öğe içermelidir",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTime geçerli Tarih ve Saatten büyük veya ona eşit olmalıdır",
		},
		{
			ns:       "Test.GtString",
			expected: "GtString, 3 karakter uzunluğundan fazla olmalıdır",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumber, 5,56 değerinden büyük olmalıdır",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultiple, 2 öğeden daha fazla içermelidir",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTime geçerli Tarih ve Saatten büyük olmalıdır",
		},
		{
			ns:       "Test.LteString",
			expected: "LteString en fazla 3 karakter uzunluğunda olmalıdır",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumber, 5,56 veya daha az olmalıdır",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultiple, maksimum 2 öğe içermelidir",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTime geçerli Tarih ve Saate eşit veya daha küçük olmalıdır",
		},
		{
			ns:       "Test.LtString",
			expected: "LtString, 3 karakter uzunluğundan daha az olmalıdır",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumber, 5,56 değerinden küçük olmalıdır",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultiple, 2 öğeden daha az içermelidir",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTime geçerli Tarih ve Saatten daha az olmalıdır",
		},
		{
			ns:       "Test.NeString",
			expected: "NeString,  değerine eşit olmamalıdır",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumber, 0.00 değerine eşit olmamalıdır",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultiple, 0 değerine eşit olmamalıdır",
		},
		{
			ns:       "Test.EqString",
			expected: "EqString, 3 değerine eşit değil",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumber, 2.33 değerine eşit değil",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultiple, 7 değerine eşit değil",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxString uzunluğu en fazla 3 karakter olmalıdır",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumber, 1.113,00 veya daha az olmalıdır",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultiple maksimum 7 öğe içermelidir",
		},
		{
			ns:       "Test.MinString",
			expected: "MinString en az 1 karakter uzunluğunda olmalıdır",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumber, 1.113,00 veya daha büyük olmalıdır",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultiple en az 7 öğe içermelidir",
		},
		{
			ns:       "Test.LenString",
			expected: "LenString uzunluğu 1 karakter olmalıdır",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumber, 1.113,00 değerine eşit olmalıdır",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultiple, 7 öğe içermelidir",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredString zorunlu bir alandır",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumber zorunlu bir alandır",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultiple zorunlu bir alandır",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLen en az 10 karakter uzunluğunda olmalıdır",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLen uzunluğu en fazla 1 karakter olmalıdır",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLen uzunluğu 2 karakter olmalıdır",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLt, 1 karakter uzunluğundan daha az olmalıdır",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLte en fazla 1 karakter uzunluğunda olmalıdır",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGt, 10 karakter uzunluğundan fazla olmalıdır",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGte en az 10 karakter uzunluğunda olmalıdır",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfString, [red green]'dan biri olmalıdır",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfInt, [5 63]'dan biri olmalıdır",
		},
		{
			ns:       "Test.UniqueSlice",
			expected: "UniqueSlice benzersiz değerler içermelidir",
		},
		{
			ns:       "Test.UniqueArray",
			expected: "UniqueArray benzersiz değerler içermelidir",
		},
		{
			ns:       "Test.UniqueMap",
			expected: "UniqueMap benzersiz değerler içermelidir",
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
