package ebpf

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecoderV4(t *testing.T) {
	data := []byte{0xf3, 0xd2, 0x12, 0x0, 0x63, 0x75, 0x72, 0x6c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc7, 0xbd, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb4, 0x5, 0x0, 0x0, 0x25, 0x39, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x0, 0xb, 0x0, 0x0, 0x0, 0xa, 0x0, 0x2, 0xf, 0xac, 0xd9, 0x5, 0xc4, 0x0, 0x50, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	fields := []string{"PID", "Task", "NumSAcks", "SRTT", "RTT", "TotalRetrans", "AdvMSS", "BytesReceived", "SegsIn", "SegsOut", "SAddr", "DAddr", "DPort"}
	expected := `"PID":1233651,"Task":"curl","NumSAcks":0,"SRTT":48583,"RTT":0,"TotalRetrans":0,"AdvMSS":1460,"BytesReceived":14629,"SegsIn":14,"SegsOut":11,"SAddr":"10.0.2.15","DAddr":"172.217.5.196","DPort":80`

	buf := new(bytes.Buffer)
	d := newDecoder(nil, true)
	d.decode(data, fields, buf)

	assert.Contains(t, buf.String(), expected)
}

func TestDecoderV6(t *testing.T) {
	data := []byte{0xdb, 0x77, 0x11, 0x0, 0x66, 0x61, 0x6b, 0x65, 0x68, 0x74, 0x74, 0x70, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x53, 0x1, 0x0, 0x0, 0xe9, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb8, 0xff, 0x0, 0x0, 0x4b, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x6, 0x0, 0x0, 0x0, 0x3, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x1, 0x9b, 0x8e, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	fields := []string{"PID", "Task", "NumSAcks", "SRTT", "RTT", "TotalRetrans", "AdvMSS", "BytesReceived", "SegsIn", "SegsOut", "SAddr", "DAddr", "DPort"}
	expected := `"PID":1144795,"Task":"fakehttp","NumSAcks":0,"SRTT":339,"RTT":233,"TotalRetrans":0,"AdvMSS":65464,"BytesReceived":75,"SegsIn":6,"SegsOut":3,"SAddr":"::1","DAddr":"::1","DPort":39822`

	buf := new(bytes.Buffer)
	d := newDecoder(nil, false)
	d.decode(data, fields, buf)

	assert.Contains(t, buf.String(), expected)
}

func BenchmarkDecoderV4(b *testing.B) {
	data := []byte{0xf3, 0xd2, 0x12, 0x0, 0x63, 0x75, 0x72, 0x6c, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xc7, 0xbd, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb4, 0x5, 0x0, 0x0, 0x25, 0x39, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xe, 0x0, 0x0, 0x0, 0xb, 0x0, 0x0, 0x0, 0xa, 0x0, 0x2, 0xf, 0xac, 0xd9, 0x5, 0xc4, 0x0, 0x50, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	buf := new(bytes.Buffer)
	fields := []string{"PID", "Task", "NumSAcks", "SRTT", "RTT", "TotalRetrans", "AdvMSS", "BytesReceived", "SegsIn", "SegsOut", "SAddr", "DAddr", "DPort"}

	d := newDecoder(nil, true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		d.decode(data, fields, buf)
	}
}

func BenchmarkDecoderV4SixFields(b *testing.B) {
	data := []byte{0xa, 0x0, 0x2, 0xf, 0xac, 0xd9, 0x5, 0xce, 0x10, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4a, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x7, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	buf := new(bytes.Buffer)
	fields := []string{"SAddr", "DAddr", "BytesReceived", "BytesSent", "RTT", "TotalRetrans"}

	d := newDecoder(nil, true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		d.decode(data, fields, buf)
	}
}

func BenchmarkDecoderV4NineFields(b *testing.B) {
	data := []byte{0xa, 0x0, 0x2, 0xf, 0xac, 0xd9, 0x5, 0xce, 0x10, 0x2, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x4a, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x8, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0xb4, 0x5, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x50, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}
	buf := new(bytes.Buffer)
	fields := []string{"SAddr", "DAddr", "BytesReceived", "BytesSent", "RTT", "PacketsOut", "AdvMSS", "TotalRetrans", "DPort"}

	d := newDecoder(nil, true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		d.decode(data, fields, buf)
	}
}
