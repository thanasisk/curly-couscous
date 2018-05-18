/*
Proggie to demonstrate the need for atomicity in append operations
as part the the Linux Programming Interface exercises
athanasios@akostopoulos.com

*/
package main

import "log"
import "os"
import "math/rand"
import flags "github.com/jessevdk/go-flags"

var opts struct {
	Size     int64 `short:"s" long:"size" description:"number of bytes to write"`
	Xclusive bool  `short:"x" long:"xclusive" description:"screw atomic append"`
}

func main() {
	file_flags := (os.O_APPEND | os.O_WRONLY | os.O_CREATE)
	regular_file_flags := (os.O_WRONLY | os.O_CREATE)
	args, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		log.Fatal(err)
	}
	if len(args) != 2 {
		log.Print(os.Args)
		log.Fatal("only one argument required - the filename")
	}
	if opts.Xclusive == true {
		file_flags = regular_file_flags
	}
	sz_buf := 1024
	random_bytes := make([]byte, sz_buf)
	rand.Read(random_bytes)
	var nChunks int64
	nChunks = opts.Size / int64(sz_buf)
	sz_leftover := opts.Size - (int64(sz_buf) * nChunks)
	leftovers := make([]byte, sz_leftover)
	fname := args[1]
	f, err := os.OpenFile(fname, file_flags, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// the real deal starts here
	if opts.Xclusive == true {
		if _, err := f.Seek(0, os.SEEK_END); err != nil {
			log.Fatal(err)
		}
	}
	for i := 0; int64(i) < nChunks; i++ {
		if _, err := f.Write(random_bytes); err != nil {
			log.Fatal(err)
		}
	}
	if _, err := f.Write(leftovers); err != nil {
		log.Fatal(err)
	}
}
