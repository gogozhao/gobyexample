package source
import (
	"crypto/sha1"
	"fmt"
	"crypto/md5"
	"encoding/base64"
)

func encodeSHA1() {
	s := "sha1 this string"
	s1 := "cf23df2207d99a74fbe169e3eba035e633b65d94"
	h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Println("source:", s1)
	fmt.Printf("result: %x\n", bs)

}

func encodeMD5() {
	s := "md5 this string"
	s1 := "d77bff3a550c1bf39b78ad2136c5d604"
	m := md5.New()
	m.Write([]byte(s))
	bs := m.Sum(nil)
	fmt.Println(s)
	fmt.Println("source:", s1)
	fmt.Printf("result: %x\n", bs)

}

func encodeBase64() {
	s := "abc123!?$*&()'-=@~"

	fmt.Println("source:", s)

	sEnc := base64.StdEncoding.EncodeToString([]byte(s))
	fmt.Println("encode:", sEnc)

	sDec, _ := base64.StdEncoding.DecodeString(sEnc)
	fmt.Println("decode", string(sDec))

	sEnc = base64.URLEncoding.EncodeToString([]byte(s))
	fmt.Println("encode", sEnc)

	sDec, _ = base64.URLEncoding.DecodeString(sEnc)
	fmt.Println("decode", string(sDec))
}

func Encodes() {
	encodeSHA1()
	fmt.Println("")
	encodeMD5()
	fmt.Println("")
	encodeBase64()
}