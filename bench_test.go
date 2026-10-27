package lz4_test

import (
	"bytes"
	"io"
	"testing"

	"hanzo.ai/lz4/v4"
	"hanzo.ai/lz4/v4/internal/lz4block"
)

func BenchmarkCompress(b *testing.B) {
	buf := make([]byte, len(pg1661))
	var c lz4.Compressor

	n, _ := c.CompressBlock(pg1661, buf)

	b.ReportAllocs()
	b.ResetTimer()
	b.ReportMetric(float64(n), "outbytes")

	for i := 0; i < b.N; i++ {
		_, _ = c.CompressBlock(pg1661, buf)
	}
}

func BenchmarkCompressRandom(b *testing.B) {
	buf := make([]byte, lz4.CompressBlockBound(len(random)))
	var c lz4.Compressor

	n, _ := c.CompressBlock(random, buf)

	b.ReportAllocs()
	b.SetBytes(int64(len(random)))
	b.ResetTimer()
	b.ReportMetric(float64(n), "outbytes")

	for i := 0; i < b.N; i++ {
		_, _ = c.CompressBlock(random, buf)
	}
}

func BenchmarkCompressHC(b *testing.B) {
	buf := make([]byte, len(pg1661))
	c := lz4.CompressorHC{Level: 16}

	n, _ := c.CompressBlock(pg1661, buf)

	b.ReportAllocs()
	b.ResetTimer()
	b.ReportMetric(float64(n), "outbytes")

	for i := 0; i < b.N; i++ {
		_, _ = c.CompressBlock(pg1661, buf)
	}
}

func BenchmarkUncompress(b *testing.B) {
	buf := make([]byte, len(pg1661))

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = lz4block.UncompressBlock(pg1661LZ4, buf, nil)
	}
}

func mustLoadFile(f string) []byte {
	b, err := ioutil.ReadFile(f)
	if err != nil {
		panic(err)
	}
	return b
}

var (
	pg1661    = mustLoadFile("testdata/pg1661.txt")
	digits    = mustLoadFile("testdata/e.txt")
	twain     = mustLoadFile("testdata/Mark
