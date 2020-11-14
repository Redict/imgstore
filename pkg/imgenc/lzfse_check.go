package imgenc

import (
	"fmt"
	"io/ioutil"

	zstd "github.com/valyala/gozstd"
)

func main() {
	v, _ := ioutil.ReadFile("bytes_manager.go")
	c := zstd.CompressLevel(nil, v, 100)
	d, _ := zstd.Decompress(nil, c)
	fmt.Println(c, len(c), len(v))
	for i := 0; i < len(d); i++ {
		fmt.Printf("%c", d[i])
	}
	fmt.Println()
}
