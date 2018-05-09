package errorhandler

import "log"

// HandleError encapsulates a try catch block with a nice log output
func HandleError(attemptingToDoWhat string, e *error, failOnError bool) {
	if *e != nil {
		log.Println("Error", attemptingToDoWhat, "(", *e, ")")
		if failOnError {
			log.Panic("Terminating application because of previous errors")
		}
	}
}
