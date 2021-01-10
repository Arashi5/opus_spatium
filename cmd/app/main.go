package main

import "work_space/playground/engine/gc"

func main() {
	/* пакеты */
	//imports.SimpleImportModule()

	/* потоки */
	//stream.StreamIn()
	//stream.StreamOut()
	//stream.StreamError()

	/* logger */
	//logger.Log()
	//logger.LogFatal()
	//logger.LogPanic()
	//logger.CustomLog()

	/* error */
	//err.ReturnError()
	//err.Error()

	/* Garbage Collection*/
	gc.GCExample()
}