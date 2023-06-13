package compressor

import (
	"bufio"
	"compress/bzip2"
	"compress/gzip"
	"io"

	"github.com/andybalholm/brotli"
	"github.com/golang/snappy"

	"github.com/klauspost/compress/zstd"
	"github.com/meyakovenkoj/go-compress/predict"
	"github.com/pierrec/lz4"
	"github.com/xi2/xz"
)

func Decompress(compressedStream io.Reader, algorithm predict.EncodeType) (uncompressedStream io.Reader) {
	var reader io.Reader
	switch algorithm {
	case predict.Lz4:
		reader = lz4.NewReader(compressedStream)
	case predict.Brotli:
		reader = brotli.NewReader(compressedStream)
	case predict.Bzip2:
		reader = bzip2.NewReader(compressedStream)
	case predict.Lzma:
		reader, _ = xz.NewReader(compressedStream, 0)
	case predict.Gzip:
		reader, _ = gzip.NewReader(compressedStream)
	case predict.Snappy:
		reader = snappy.NewReader(compressedStream)
	case predict.Zstd:
		reader, _ = zstd.NewReader(compressedStream)
	case predict.None:
		reader = bufio.NewReader(compressedStream)
	}

	return reader
}
