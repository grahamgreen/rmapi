package shell

import (
	"fmt"

	"github.com/abiosoft/ishell"
)

func nukeCmd(ctx *ShellCtxt) *ishell.Cmd {
	return &ishell.Cmd{
		Name:      "nuke",
		Help:      "deletes everything",
		Completer: createEntryCompleter(ctx),
		Func: func(c *ishell.Context) {
			fmt.Print("Are you sure, this will DELETE EVERYTHING! type [YES]:")
			var response string
			_, err := fmt.Scanln(&response)
			if err != nil {
				return
			}
			if response != "YES" {
				return
			}
			fmt.Println("Nuking")
			err = ctx.api.Nuke()

			if err != nil {
				c.Err(fmt.Errorf("failed to delete entry: %v", err))
				return
			}
			ctx.api.Filetree().Clear()
		},
	}
}
