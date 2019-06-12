package main

// #include <stdlib.h>
// #include "rocksdb/c.h"
// #cgo CFLAGS: -I./rocksdb/include
// #cgo LDFLAGS: -L ./rocksdb/ -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -ldl
import "C"

import (
	"unsafe"
	"os"
	"fmt"
	"flag"
)

func main() {

	name := flag.String("N", "test.db", "db")
	op := flag.String("O", "r", "operate")
	key := flag.String("K", "default", "key")
	value := flag.String("V", "", "value")
	help := flag.Bool("H", false, "help")
	flag.Usage = func() {
		fmt.Println("Usage : gosocksdb [-N name] -O r/w/d -K default -V test")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	var (
		cErr *C.char
		cName = C.CString(*name)
		cKey = C.CString(*key)
		cValue = C.CString(*value)
	)
	defer func() {
		C.free(unsafe.Pointer(cName))
		C.free(unsafe.Pointer(cKey))
		C.free(unsafe.Pointer(cValue))
	}()
	options := C.rocksdb_options_create()
	C.rocksdb_options_set_create_if_missing(options, 1)
	db := C.rocksdb_open(options, cName, &cErr)

	if cErr != nil {
		defer C.free(unsafe.Pointer(cErr))
		fmt.Println(C.GoString(cErr))
		os.Exit(0)
	}

	switch *op {
	case "r":
		ro := C.rocksdb_readoptions_create()
		var cLength C.size_t
		cValue := C.rocksdb_get(db, ro, cKey, C.size_t(len(*key)), &cLength, &cErr)
		if cErr != nil {
			defer C.free(unsafe.Pointer(cErr))
			fmt.Println(C.GoString(cErr))
		}
		defer C.free(unsafe.Pointer(cValue))
		fmt.Println(string(C.GoBytes(unsafe.Pointer(cValue), C.int(cLength))))

	case "w":
		wo := C.rocksdb_writeoptions_create()
		C.rocksdb_put(db, wo, cKey, C.size_t(len(*key)), cValue, C.size_t(len(*value)), &cErr)
		if cErr != nil {
			defer C.free(unsafe.Pointer(cErr))
			fmt.Println(C.GoString(cErr))
		}
	case "d":
		wo := C.rocksdb_writeoptions_create()
		C.rocksdb_delete(db, wo, cKey, C.size_t(len(*key)), &cErr)
		if cErr != nil {
			defer C.free(unsafe.Pointer(cErr))
			fmt.Println(C.GoString(cErr))
		}
	}
}