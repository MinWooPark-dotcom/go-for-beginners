package main

import (
	"bytes"
	"fmt"
	"io"
	"mime"
)

// type WordDecoder struct {
// CharsetReader, if non-nil, defines a function to generate
// charset-conversion readers, converting from the provided
// charset into UTF-8.
// Charsets are always lower-case. utf-8, iso-8859-1 and us-ascii charsets
// are handled by default.
// One of the CharsetReader's result values must be non-nil.
// CharsetReader func(charset string, input io.Reader) (io.Reader, error)
// }

func main() {
	// 객체 생성 방법으로 Go 내장함수 new()를 쓸 수 있다.
	// new()를 사용하면 모든 필드를 Zero value로 초기화하고 객체의 포인터를 리턴.
	// A WordDecoder decodes MIME headers containing RFC 2047 encoded-words.
	dec := new(mime.WordDecoder) // ${<nil>}

	// WordDecoder의 메서드 Decode
	// 인자로 string을 받고 리턴 값으로 string, error
	header, err := dec.Decode("=?utf-8?q?=C2=A1Hola,_se=C3=B1or!?=")
	if err != nil {
		panic(err)
	}
	fmt.Println("header", header) // ¡Hola, señor!

	// CharsetReader 메서드 정의
	dec.CharsetReader = func(charset string, input io.Reader) (io.Reader, error) {
		switch charset {
		case "x-case":
			// Fake character set for example.
			// Real use would integrate with packages such
			// as code.google.com/p/go-charset
			content, err := io.ReadAll(input)
			if err != nil {
				return nil, err
			}
			return bytes.NewReader(bytes.ToUpper(content)), nil
		default:
			return nil, fmt.Errorf("unhandled charset %q", charset)
		}
	}

	// func (*WordDecoder) Decode ¶
	header, err = dec.Decode("=?x-case?q?hello!?=")
	if err != nil {
		panic(err)
	}
	fmt.Println(header) // HELLO!
}
