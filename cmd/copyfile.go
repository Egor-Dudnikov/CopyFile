package copyfile

import "os"

func main() {

	in, err := os.Open(os.Args[1])
	out, err := os.Args[2]

}
