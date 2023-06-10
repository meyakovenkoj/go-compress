package typing

import (
	//"github.com/gabriel-vasile/mimetype"
	"fmt"
	"os"

	"github.com/h2non/filetype"
	"github.com/h2non/filetype/types"
)

func Match(file *os.File) (kind types.Type) {
	header := make([]byte, 261)
	n, _ := file.Read(header)
	if n == 0 {
		fmt.Println("get eof")
		return
	}
	kind, _ = filetype.Match(header)
	return
}

func main() {
	// buf, _ := ioutil.ReadFile("sample.jpg")
	file, _ := os.Open("movie.mp4")

	// We only have to pass the file header = first 261 bytes
	kind := Match(file)
	// kind, _ := filetype.Match(buf)
	if kind == filetype.Unknown {
		fmt.Println("Unknown file type")
		return
	}

	fmt.Printf("File type: %s. MIME: %s\n", kind.Extension, kind.MIME.Value)
}

// func get_type(filename):
//     type, enc = mimetypes.guess_type(filename)
//     if type is not None and type.startswith('text/'):
//         type = 'text'
//     match type:
//         case 'text':
//             return Type.TXT
//         case 'application/json':
//             return Type.JSN
//         case 'application/octet-stream':
//             return Type.BIN
//         case None:
//             return Type.BIN
//         case _:
//             return Type.UKN
