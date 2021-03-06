package main

import (
	"fmt"

	"github.com/flynn/go-docopt"
	"github.com/theupdateframework/go-tuf"
)

func init() {
	register("gen-key", cmdGenKey, `
usage: tuf gen-key [--expires=<days>] <role>

Generate a new signing key for the given role.

The key will be serialized to JSON and written to the "keys" directory with
filename pattern "ROLE-KEYID.json". The root manifest will also be staged
with the addition of the key's ID to the role's list of key IDs.

Options:
  --expires=<days>   Set the root manifest to expire <days> days from now.
`)
}

func cmdGenKey(args *docopt.Args, repo *tuf.Repo) error {
	role := args.String["<role>"]
	var keyids []string
	var err error
	if arg := args.String["--expires"]; arg != "" {
		expires, err := parseExpires(arg)
		if err != nil {
			return err
		}
		keyids, err = repo.GenKeyWithExpires(role, expires)
	} else {
		keyids, err = repo.GenKey(role)
	}
	if err != nil {
		return err
	}
	for _, id := range keyids {
		fmt.Println("Generated", role, "key with ID", id)
	}
	return nil
}
