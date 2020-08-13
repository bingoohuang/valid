package ja

import (
	"testing"
	"time"

	"github.com/bingoohuang/valid"
	. "github.com/go-playground/assert/v2"
	ja_locale "github.com/go-playground/locales/ja"
	ut "github.com/go-playground/universal-translator"
)

func TestTranslations(t *testing.T) {
	japanese := ja_locale.New()
	uni := ut.New(japanese, japanese)
	trans, _ := uni.GetTranslator("ja")

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

	test.Excludes = "this is some test text"
	test.ExcludesAll = "This is Great!"
	test.ExcludesRune = "Love it ☻"

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
			expected: "IsColorは正しい色でなければなりません",
		},
		{
			ns:       "Test.MAC",
			expected: "MACは正しいMACアドレスを含まなければなりません",
		},
		{
			ns:       "Test.IPAddr",
			expected: "IPAddrは解決可能なIPアドレスでなければなりません",
		},
		{
			ns:       "Test.IPAddrv4",
			expected: "IPAddrv4は解決可能なIPv4アドレスでなければなりません",
		},
		{
			ns:       "Test.IPAddrv6",
			expected: "IPAddrv6は解決可能なIPv6アドレスでなければなりません",
		},
		{
			ns:       "Test.UDPAddr",
			expected: "UDPAddrは正しいUDPアドレスでなければなりません",
		},
		{
			ns:       "Test.UDPAddrv4",
			expected: "UDPAddrv4は正しいIPv4のUDPアドレスでなければなりません",
		},
		{
			ns:       "Test.UDPAddrv6",
			expected: "UDPAddrv6は正しいIPv6のUDPアドレスでなければなりません",
		},
		{
			ns:       "Test.TCPAddr",
			expected: "TCPAddrは正しいTCPアドレスでなければなりません",
		},
		{
			ns:       "Test.TCPAddrv4",
			expected: "TCPAddrv4は正しいIPv4のTCPアドレスでなければなりません",
		},
		{
			ns:       "Test.TCPAddrv6",
			expected: "TCPAddrv6は正しいIPv6のTCPアドレスでなければなりません",
		},
		{
			ns:       "Test.CIDR",
			expected: "CIDRは正しいCIDR表記を含まなければなりません",
		},
		{
			ns:       "Test.CIDRv4",
			expected: "CIDRv4はIPv4アドレスの正しいCIDR表記を含まなければなりません",
		},
		{
			ns:       "Test.CIDRv6",
			expected: "CIDRv6はIPv6アドレスの正しいCIDR表記を含まなければなりません",
		},
		{
			ns:       "Test.SSN",
			expected: "SSNは正しい社会保障番号でなければなりません",
		},
		{
			ns:       "Test.IP",
			expected: "IPは正しいIPアドレスでなければなりません",
		},
		{
			ns:       "Test.IPv4",
			expected: "IPv4は正しいIPv4アドレスでなければなりません",
		},
		{
			ns:       "Test.IPv6",
			expected: "IPv6は正しいIPv6アドレスでなければなりません",
		},
		{
			ns:       "Test.DataURI",
			expected: "DataURIは正しいデータURIを含まなければなりません",
		},
		{
			ns:       "Test.Latitude",
			expected: "Latitudeは正しい緯度の座標を含まなければなりません",
		},
		{
			ns:       "Test.Longitude",
			expected: "Longitudeは正しい経度の座標を含まなければなりません",
		},
		{
			ns:       "Test.MultiByte",
			expected: "MultiByteはマルチバイト文字を含まなければなりません",
		},
		{
			ns:       "Test.ASCII",
			expected: "ASCIIはASCII文字のみを含まなければなりません",
		},
		{
			ns:       "Test.PrintableASCII",
			expected: "PrintableASCIIは印刷可能なASCII文字のみを含まなければなりません",
		},
		{
			ns:       "Test.UUID",
			expected: "UUIDは正しいUUIDでなければなりません",
		},
		{
			ns:       "Test.UUID3",
			expected: "UUID3はバージョンが3の正しいUUIDでなければなりません",
		},
		{
			ns:       "Test.UUID4",
			expected: "UUID4はバージョンが4の正しいUUIDでなければなりません",
		},
		{
			ns:       "Test.UUID5",
			expected: "UUID5はバージョンが4の正しいUUIDでなければなりません",
		},
		{
			ns:       "Test.ISBN",
			expected: "ISBNは正しいISBN番号でなければなりません",
		},
		{
			ns:       "Test.ISBN10",
			expected: "ISBN10は正しいISBN-10番号でなければなりません",
		},
		{
			ns:       "Test.ISBN13",
			expected: "ISBN13は正しいISBN-13番号でなければなりません",
		},
		{
			ns:       "Test.Excludes",
			expected: "Excludesには'text'というテキストを含むことはできません",
		},
		{
			ns:       "Test.ExcludesAll",
			expected: "ExcludesAllには'!@#$'のどれも含めることはできません",
		},
		{
			ns:       "Test.ExcludesRune",
			expected: "ExcludesRuneには'☻'を含めることはできません",
		},
		{
			ns:       "Test.ContainsAny",
			expected: "ContainsAnyは'!@#$'の少なくとも1つを含まなければなりません",
		},
		{
			ns:       "Test.Contains",
			expected: "Containsは'purpose'を含まなければなりません",
		},
		{
			ns:       "Test.Base64",
			expected: "Base64は正しいBase64文字列でなければなりません",
		},
		{
			ns:       "Test.Email",
			expected: "Emailは正しいメールアドレスでなければなりません",
		},
		{
			ns:       "Test.URL",
			expected: "URLは正しいURLでなければなりません",
		},
		{
			ns:       "Test.URI",
			expected: "URIは正しいURIでなければなりません",
		},
		{
			ns:       "Test.RGBColorString",
			expected: "RGBColorStringは正しいRGBカラーコードでなければなりません",
		},
		{
			ns:       "Test.RGBAColorString",
			expected: "RGBAColorStringは正しいRGBAカラーコードでなければなりません",
		},
		{
			ns:       "Test.HSLColorString",
			expected: "HSLColorStringは正しいHSLカラーコードでなければなりません",
		},
		{
			ns:       "Test.HSLAColorString",
			expected: "HSLAColorStringは正しいHSLAカラーコードでなければなりません",
		},
		{
			ns:       "Test.HexadecimalString",
			expected: "HexadecimalStringは正しい16進表記でなければなりません",
		},
		{
			ns:       "Test.HexColorString",
			expected: "HexColorStringは正しいHEXカラーコードでなければなりません",
		},
		{
			ns:       "Test.NumberString",
			expected: "NumberStringは正しい数でなければなりません",
		},
		{
			ns:       "Test.NumericString",
			expected: "NumericStringは正しい数字でなければなりません",
		},
		{
			ns:       "Test.AlphanumString",
			expected: "AlphanumStringはアルファベットと数字のみを含むことができます",
		},
		{
			ns:       "Test.AlphaString",
			expected: "AlphaStringはアルファベットのみを含むことができます",
		},
		{
			ns:       "Test.LtFieldString",
			expected: "LtFieldStringはMaxStringよりも小さくなければなりません",
		},
		{
			ns:       "Test.LteFieldString",
			expected: "LteFieldStringはMaxString以下でなければなりません",
		},
		{
			ns:       "Test.GtFieldString",
			expected: "GtFieldStringはMaxStringよりも大きくなければなりません",
		},
		{
			ns:       "Test.GteFieldString",
			expected: "GteFieldStringはMaxString以上でなければなりません",
		},
		{
			ns:       "Test.NeFieldString",
			expected: "NeFieldStringはEqFieldStringとは異ならなければなりません",
		},
		{
			ns:       "Test.LtCSFieldString",
			expected: "LtCSFieldStringはInner.LtCSFieldStringよりも小さくなければなりません",
		},
		{
			ns:       "Test.LteCSFieldString",
			expected: "LteCSFieldStringはInner.LteCSFieldString以下でなければなりません",
		},
		{
			ns:       "Test.GtCSFieldString",
			expected: "GtCSFieldStringはInner.GtCSFieldStringよりも大きくなければなりません",
		},
		{
			ns:       "Test.GteCSFieldString",
			expected: "GteCSFieldStringはInner.GteCSFieldString以上でなければなりません",
		},
		{
			ns:       "Test.NeCSFieldString",
			expected: "NeCSFieldStringはInner.NeCSFieldStringとは異ならなければなりません",
		},
		{
			ns:       "Test.EqCSFieldString",
			expected: "EqCSFieldStringはInner.EqCSFieldStringと等しくなければなりません",
		},
		{
			ns:       "Test.EqFieldString",
			expected: "EqFieldStringはMaxStringと等しくなければなりません",
		},
		{
			ns:       "Test.GteString",
			expected: "GteStringの長さは少なくとも3文字以上はなければなりません",
		},
		{
			ns:       "Test.GteNumber",
			expected: "GteNumberは5.56かより大きくなければなりません",
		},
		{
			ns:       "Test.GteMultiple",
			expected: "GteMultipleは少なくとも2つの項目を含まなければなりません",
		},
		{
			ns:       "Test.GteTime",
			expected: "GteTimeは現時刻以降でなければなりません",
		},
		{
			ns:       "Test.GtString",
			expected: "GtStringの長さは3文字よりも多くなければなりません",
		},
		{
			ns:       "Test.GtNumber",
			expected: "GtNumberは5.56よりも大きくなければなりません",
		},
		{
			ns:       "Test.GtMultiple",
			expected: "GtMultipleは2つの項目よりも多い項目を含まなければなりません",
		},
		{
			ns:       "Test.GtTime",
			expected: "GtTimeは現時刻よりも後でなければなりません",
		},
		{
			ns:       "Test.LteString",
			expected: "LteStringの長さは最大でも3文字でなければなりません",
		},
		{
			ns:       "Test.LteNumber",
			expected: "LteNumberは5.56かより小さくなければなりません",
		},
		{
			ns:       "Test.LteMultiple",
			expected: "LteMultipleは最大でも2つの項目を含まなければなりません",
		},
		{
			ns:       "Test.LteTime",
			expected: "LteTimeは現時刻以前でなければなりません",
		},
		{
			ns:       "Test.LtString",
			expected: "LtStringの長さは3文字よりも少なくなければなりません",
		},
		{
			ns:       "Test.LtNumber",
			expected: "LtNumberは5.56よりも小さくなければなりません",
		},
		{
			ns:       "Test.LtMultiple",
			expected: "LtMultipleは2つの項目よりも少ない項目を含まなければなりません",
		},
		{
			ns:       "Test.LtTime",
			expected: "LtTimeは現時刻よりも前でなければなりません",
		},
		{
			ns:       "Test.NeString",
			expected: "NeStringはと異ならなければなりません",
		},
		{
			ns:       "Test.NeNumber",
			expected: "NeNumberは0.00と異ならなければなりません",
		},
		{
			ns:       "Test.NeMultiple",
			expected: "NeMultipleの項目の数は0個と異ならなければなりません",
		},
		{
			ns:       "Test.EqString",
			expected: "EqStringは3と等しくありません",
		},
		{
			ns:       "Test.EqNumber",
			expected: "EqNumberは2.33と等しくありません",
		},
		{
			ns:       "Test.EqMultiple",
			expected: "EqMultipleは7と等しくありません",
		},
		{
			ns:       "Test.MaxString",
			expected: "MaxStringの長さは最大でも3文字でなければなりません",
		},
		{
			ns:       "Test.MaxNumber",
			expected: "MaxNumberは1,113.00かより小さくなければなりません",
		},
		{
			ns:       "Test.MaxMultiple",
			expected: "MaxMultipleは最大でも7つの項目を含まなければなりません",
		},
		{
			ns:       "Test.MinString",
			expected: "MinStringの長さは少なくとも1文字はなければなりません",
		},
		{
			ns:       "Test.MinNumber",
			expected: "MinNumberは1,113.00かより大きくなければなりません",
		},
		{
			ns:       "Test.MinMultiple",
			expected: "MinMultipleは少なくとも7つの項目を含まなければなりません",
		},
		{
			ns:       "Test.LenString",
			expected: "LenStringの長さは1文字でなければなりません",
		},
		{
			ns:       "Test.LenNumber",
			expected: "LenNumberは1,113.00と等しくなければなりません",
		},
		{
			ns:       "Test.LenMultiple",
			expected: "LenMultipleは7つの項目を含まなければなりません",
		},
		{
			ns:       "Test.RequiredString",
			expected: "RequiredStringは必須フィールドです",
		},
		{
			ns:       "Test.RequiredNumber",
			expected: "RequiredNumberは必須フィールドです",
		},
		{
			ns:       "Test.RequiredMultiple",
			expected: "RequiredMultipleは必須フィールドです",
		},
		{
			ns:       "Test.StrPtrMinLen",
			expected: "StrPtrMinLenの長さは少なくとも10文字はなければなりません",
		},
		{
			ns:       "Test.StrPtrMaxLen",
			expected: "StrPtrMaxLenの長さは最大でも1文字でなければなりません",
		},
		{
			ns:       "Test.StrPtrLen",
			expected: "StrPtrLenの長さは2文字でなければなりません",
		},
		{
			ns:       "Test.StrPtrLt",
			expected: "StrPtrLtの長さは1文字よりも少なくなければなりません",
		},
		{
			ns:       "Test.StrPtrLte",
			expected: "StrPtrLteの長さは最大でも1文字でなければなりません",
		},
		{
			ns:       "Test.StrPtrGt",
			expected: "StrPtrGtの長さは10文字よりも多くなければなりません",
		},
		{
			ns:       "Test.StrPtrGte",
			expected: "StrPtrGteの長さは少なくとも10文字以上はなければなりません",
		},
		{
			ns:       "Test.OneOfString",
			expected: "OneOfStringは[red green]のうちのいずれかでなければなりません",
		},
		{
			ns:       "Test.OneOfInt",
			expected: "OneOfIntは[5 63]のうちのいずれかでなければなりません",
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
