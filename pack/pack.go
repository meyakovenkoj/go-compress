package pack

// package main

import (
	"archive/tar"
	"io"
	"log"
	"os"
	"path/filepath"
)

func Pack(src string, buf io.Writer) (packedStream *tar.Writer) {
	// tar > buf
	tarStream := tar.NewWriter(buf)

	// walk through every file in the folder
	filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
		// generate tar header
		header, err := tar.FileInfoHeader(fi, file)
		if err != nil {
			return err
		}

		// must provide real name
		// (see https://golang.org/src/archive/tar/common.go?#L626)
		header.Name = filepath.ToSlash(file)

		// write header
		if err := tarStream.WriteHeader(header); err != nil {
			return err
		}
		// if not a dir, write file content
		if !fi.IsDir() {
			data, err := os.Open(file)
			if err != nil {
				return err
			}
			if _, err := io.Copy(tarStream, data); err != nil {
				return err
			}
		}
		return nil
	})

	return tarStream
}

func Unpack(dstPath string, packedStream io.Reader) {
	tarReader := tar.NewReader(packedStream)

	for true {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Unpack: Next() failed: %s", err.Error())
		}

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.Mkdir(filepath.Join(dstPath, header.Name), 0755); err != nil {
				log.Fatalf("Unpack: Mkdir() failed: %s", err.Error())
			}
		case tar.TypeReg:
			outFile, err := os.Create(filepath.Join(dstPath, header.Name))
			if err != nil {
				log.Fatalf("Unpack: Create() failed: %s", err.Error())
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				log.Fatalf("Unpack: Copy() failed: %s", err.Error())
			}
			outFile.Close()

		default:
			log.Fatalf(
				"Unpack: unknown type: %b in %s",
				header.Typeflag,
				header.Name)
		}

	}
}

// func main() {
// 	f, err := os.OpenFile("/tmp/123.tar", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	var mode predict.EncodeType
// 	mode = 1
// 	compressed := compressor.Compress(f, mode)

// 	packed := Pack("./main.go", compressed)
// 	packed.Close()
// 	compressed.Close()

// 	os.Mkdir("temp", os.ModePerm)
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	outf, err := os.Open("/tmp/123.tar")
// 	if err != nil {
// 		panic(err)
// 	}
// 	decompressed := compressor.Decompress(outf, mode)
// 	Unpack("temp", decompressed)

// }
