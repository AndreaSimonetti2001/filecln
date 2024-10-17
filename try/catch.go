package try

import "filecln/logger"

func Catch(e error) {
	if e != nil {
		logger.Error("%v", e)
	}
}
