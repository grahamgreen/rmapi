package shell

import (
	"errors"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/abiosoft/ishell"
	"github.com/grahamgreen/rmapi/filetree"
	"github.com/grahamgreen/rmapi/model"
)

func findCmd(ctx *ShellCtxt) *ishell.Cmd {
	return &ishell.Cmd{
		Name:      "find",
		Help:      "find files recursively, usage: find dir [regexp]",
		Completer: createDirCompleter(ctx),
		Func: func(c *ishell.Context) {
			if len(c.Args) != 1 && len(c.Args) != 2 {
				c.Err(errors.New("missing arguments; usage find dir [regexp]"))
				return
			}

			start := c.Args[0]

			startNode, err := ctx.api.Filetree().NodeByPath(start, ctx.node)

			if err != nil {
				c.Err(errors.New("start directory doesn't exist"))
				return
			}

			var matchRegexp *regexp.Regexp
			if len(c.Args) == 2 {
				matchRegexp, err = regexp.Compile(c.Args[1])
				if err != nil {
					c.Err(errors.New("failed to compile regexp"))
					return
				}
			}

			filetree.WalkTree(startNode, filetree.FileTreeVistor{
				Visit: func(node *model.Node, path []string) bool {
					var entryType string
					if node.IsDirectory() {
						entryType = "[d] "
					} else {
						entryType = "[f] "
					}
					entryName := entryType + filepath.Join(strings.Join(path, "/"), node.Name())

					if matchRegexp == nil {
						c.Println(entryName)
						return false
					}

					if !matchRegexp.Match([]byte(entryName)) {
						return false
					}

					c.Println(entryName)

					return false
				},
			})
		},
	}
}
