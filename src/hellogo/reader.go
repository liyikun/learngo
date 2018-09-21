package main

import (
	"io"
	"strings"
	"os"
)

type MyReader struct {

}
//
//func (mr MyReader) NewReader(c string) MyReader {
//	mr.content = c
//	return mr
//}


func (r MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (int, error) {
	n, err := reader.r.Read(b)
	for i:=0; i<n; i++{
		switch {
		case b[i] >= 'A' && b[i] <'N' :
			b[i] += 13
		case b[i]>='N' && b[i] <='Z':
			b[i] -=13
		case b[i] >= 'a' && b[i] <'n' :
			b[i] += 13
		case b[i]>='n' && b[i] <='z':
			b[i] -=13
		}
	}
	return n, err
}



func main()  {

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	//
	//r := strings.NewReader("Hello, Reader!")
	//
	//b := make([]byte, 8)
	//
	//fmt.Println(len("aaaaaaa"))
	//for {
	//	n, err := r.Read(b)
	//	fmt.Printf("n =%v err = %v b = %v\n", n, err, b)
	//	fmt.Printf("b[:n] = %q\n", b[:n])
	//	if err == io.EOF {
	//		break
	//	}
	//}
}