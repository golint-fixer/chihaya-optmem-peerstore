package optmem

import (
	"fmt"
	"log"
)

func logf(format string, v ...interface{}) {
	log.Printf("optmem: "+format, v)
}

func logln(v ...interface{}) {
	log.Println("optmem: " + fmt.Sprint(v...))
}
