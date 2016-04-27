package shared

import(
	"log"
	"os"
)

func init(){
	initLoggers()
}

func initLoggers(){
	LogV = log.New(os.Stdout, "[VERBOSE]:", log.Ldate | log.Ltime | log.Lshortfile)
	LogD = log.New(os.Stdout, "[DEBUG]:", log.Ldate | log.Ltime | log.Lshortfile)
	LogE = log.New(os.Stderr, "[ERROR]:", log.Ldate | log.Ltime | log.Lshortfile)
	LogW = log.New(os.Stderr, "[WARNING]:", log.Ldate | log.Ltime | log.Lshortfile)
}
