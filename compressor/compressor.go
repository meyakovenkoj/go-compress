package compressor

import (
	"bufio"
	"io"

	"compress/gzip"

	"github.com/andybalholm/brotli"
	"github.com/golang/snappy"
	"github.com/klauspost/compress/zstd"
	"github.com/larzconwell/bzip2"
	"github.com/meyakovenkoj/go-compress/predict"
	"github.com/pierrec/lz4"
	"github.com/ulikunitz/xz"
)

type BufWriteCloser struct {
	*bufio.Writer
}

func (mwc *BufWriteCloser) Close() error {
	return nil
}

func Compress(uncompressedStream io.Writer, algorithm predict.EncodeType) (compressedStream io.WriteCloser) {
	var writer io.WriteCloser
	cfg := predict.ReadConfig()
	switch algorithm {
	case predict.Lz4:
		writer = lz4.NewWriter(uncompressedStream)
	case predict.Brotli:
		writer = brotli.NewWriterLevel(uncompressedStream, cfg.Brotli.Level)
	case predict.Bzip2:
		writer = bzip2.NewWriter(uncompressedStream)
	case predict.Lzma:
		writer, _ = xz.NewWriter(uncompressedStream)
	case predict.Gzip:
		writer = gzip.NewWriter(uncompressedStream)
	case predict.Snappy:
		writer = snappy.NewWriter(uncompressedStream)
	case predict.Zstd:
		writer, _ = zstd.NewWriter(uncompressedStream)
	case predict.None:
		bw := bufio.NewWriter(uncompressedStream)
		writer = &BufWriteCloser{bw}
	}

	return writer
}
