package main

import (
	"fmt"
	"io"
	"os"
)

type Castom struct {
	write  io.Writer
	copied int
}

func (f *Castom) Write(p []byte) (int, error) {
	f.write.Write(p)
	fmt.Printf("Скопировано %d байт\n", f.copied+len(p))
	f.copied += len(p)
	return len(p), nil
}

func main() {

	in, errin := os.Open(os.Args[1])
	if errin != nil {
		panic(errin)
	}
	defer in.Close()

	out, errout := os.Create(os.Args[2])
	if errout != nil {
		panic(errout)
	}
	defer out.Close()

	CastomWriter := &Castom{
		write:  out,
		copied: 0,
	}

	fmt.Println("Начинается копирование...")

	_, errcopy := io.Copy(CastomWriter, in)
	if errcopy != nil {
		panic(errcopy)
	}

	fmt.Println("Копирование завершено!")
}
