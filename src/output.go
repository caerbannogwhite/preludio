package preludiocore

import (
	"github.com/caerbannogwhite/enchanter/dataframe"
	"github.com/caerbannogwhite/enchanter/meta"
)

// Output is what a run hands back: the log, and one dataframe per result.
// A result that is a bare series comes back as a one-column dataframe.
// Rendering belongs to the caller, through enchanter's DataFrame.Table.
type Output struct {
	Log  []meta.LogEnty
	Data []dataframe.DataFrame
}
