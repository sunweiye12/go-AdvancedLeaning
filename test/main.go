package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	tenantId := strings.TrimSpace("0")
	fmt.Println(tenantId)
	fmt.Print("第一次获取:", tenantId)
	tenantId1, err := strconv.ParseInt(tenantId, 10, 64)

	if err != nil {
		fmt.Println("报错")
	}
	fmt.Println(tenantId1)
	fmt.Printf("tenantId1: %T", tenantId1)
}

//func containKey(target string, arr []string) bool {
//	for _, e := range arr {
//		if target == strings.TrimSpace(e) {
//			return true
//		}
//	}
//	return false
//}
//
//func compress(in []byte) ([]byte, error) {
//	r := bytes.NewReader(in)
//	w := &bytes.Buffer{}
//	zw := lz4.NewWriter(w)
//	_, err := io.Copy(zw, r)
//	if err != nil {
//		return nil, err
//	}
//	// Closing is *very* important
//	if err := zw.Close(); err != nil {
//		return nil, err
//	}
//	return w.Bytes(), nil
//}
//
//func decompress(in []byte) ([]byte, error) {
//	r := bytes.NewReader(in)
//	w := &bytes.Buffer{}
//
//	zr := lz4.NewReader(r)
//	_, err := io.Copy(w, zr)
//	if err != nil {
//		return nil, err
//	}
//	return w.Bytes(), nil
//}
