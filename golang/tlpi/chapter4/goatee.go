/*
a tee-like utility written in Golang (go-tee, get it? ha-ha!)
as part the the Linux Programming Interface exercises
athanasios@akostopoulos.com

*/
package main

import "log"
import "os"
import "io"
import "bufio"
import flags "github.com/jessevdk/go-flags"

var opts struct {
	Append bool `short:"a" long:"append" description:"Open output file in append mode"`
}

func main() {
	fileFlags := (os.O_CREATE | os.O_WRONLY)
	args, err := flags.ParseArgs(&opts, os.Args)
	if err != nil {
		log.Fatal(err)
	}
	if len(args) != 2 {
		log.Fatal("only one argument required")
	}
	if opts.Append == true {
		fileFlags = (os.O_APPEND | fileFlags)
	}
	fname := args[1]
	f, err := os.OpenFile(fname, fileFlags, 0666)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// saves a TON of work
	out := io.MultiWriter(os.Stdout, f)
	nBytes, nChunks := int64(0), int64(0)
	r := bufio.NewReader(os.Stdin)
	buf := make([]byte, 0, 4*1024)
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]
		if n == 0 {
			if err == nil {
				continue
			}
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		nChunks++
		nBytes += int64(len(buf))
		if _, err := out.Write(buf); err != nil {
			log.Fatal(err)
		}
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
	}
	log.Println("Bytes:", nBytes, "Chunks:", nChunks)
}
