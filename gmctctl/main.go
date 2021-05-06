package main

import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	decodeStr = flag.String("d", "", "the string need to be decode")
	encodeStr = flag.String("e", "", "the string need to be encode")
)

func main() {
	flag.Parse()
	str := *decodeStr
	if str != "" {
		s, err := decode(str)
		if err != nil {
			log.Fatalln(red(err.Error()))
		}

		fmt.Println(green(s))
		return
	}

	str = *encodeStr
	if str != "" {
		s, err := encode([]byte(str))
		if err != nil {
			log.Fatalln(red(err.Error()))
		}

		fmt.Println(green(s))
		return
	}
}

//
func decode(d string) (string, error) {
	hr, err := hex.DecodeString(d)
	if err != nil {
		return "", err
	}

	br := bytes.NewReader(hr)
	zr, err := zlib.NewReader(br)
	if err != nil {
		return "", err
	}

	rbuf, err := ioutil.ReadAll(zr)
	if err != nil {
		return "", err
	}

	return string(rbuf), nil
}

func encode(buf []byte) (string, error) {
	var b bytes.Buffer
	defer b.Reset()

	w := zlib.NewWriter(&b)
	if _, err := w.Write(buf); err != nil {
		return "", err
	}
	w.Close()

	return hex.EncodeToString(b.Bytes()), nil
}

//
func red(str string) string {
	return fmt.Sprintf("\x1b[31m%s\x1b[0m", str)
}

//
func green(str string) string {
	return fmt.Sprintf("\x1b[32m%s\x1b[0m", str)
}
