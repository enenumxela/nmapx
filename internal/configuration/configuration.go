package configuration

import (
	"fmt"
)

const (
	VERSION string = "1.0.0"
)

var (
	BANNER string = fmt.Sprintf(`
_ __  _ __ ___   __ _ _ ____  __
| '_ \| '_ `+"`"+` _ \ / _`+"`"+` | '_ \ \/ /
| | | | | | | | | (_| | |_) >  < 
|_| |_|_| |_| |_|\__,_| .__/_/\_\
                      |_| v%s`, VERSION)
)
