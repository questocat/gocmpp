//go:build go1.18
// +build go1.18

package cmpp_test

import (
	"testing"

	cmpp "github.com/bigwhite/gocmpp"
)

func FuzzSubmit(f *testing.F) {
	data := []byte{
		0x00, 0x00, 0x00, 0xbd, 0x00, 0x00, 0x00, 0x04, 0x00, 0x00, 0x00, 0x17, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x01, 0x01, 0x01, 0x01, 0x74, 0x65, 0x73, 0x74, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x02, 0x31, 0x33, 0x35, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x39, 0x36, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x08, 0x39, 0x30, 0x30, 0x30, 0x30,
		0x31, 0x30, 0x32, 0x31, 0x30, 0x00, 0x00, 0x00, 0x00, 0x31, 0x35, 0x31, 0x31, 0x30, 0x35, 0x31,
		0x33, 0x31, 0x35, 0x35, 0x35, 0x31, 0x30, 0x31, 0x2b, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x39, 0x30, 0x30, 0x30, 0x30,
		0x31, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x01, 0x31, 0x33, 0x35, 0x30, 0x30, 0x30, 0x30, 0x32, 0x36, 0x39, 0x36, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x1e, 0x6d, 0x4b, 0x8b, 0xd5, 0x00, 0x67, 0x00, 0x6f, 0x00,
		0x63, 0x00, 0x6d, 0x00, 0x70, 0x00, 0x70, 0x00, 0x20, 0x00, 0x73, 0x00, 0x75, 0x00, 0x62, 0x00,
		0x6d, 0x00, 0x69, 0x00, 0x74, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}
	f.Add(data[8:])
	f.Fuzz(func(t *testing.T, data []byte) {
		p := &cmpp.Cmpp2SubmitReqPkt{}
		_ = p.Unpack(data)
	})
}