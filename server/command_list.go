package main

import (
	"encoding/json"
	"fmt"

	"github.com/mattermost/mattermost-server/v5/model"
	"github.com/mattermost/mattermost-server/v5/plugin"
)

func executeList(p *Plugin, c *plugin.Context, cmdArgs *model.CommandArgs, args ...string) *model.CommandResponse {
	keys, appErr := p.API.KVList(0, 10000)
	if appErr != nil {
		return p.responsef(cmdArgs, "Error listing keys. err=%v", appErr)
	}

	if len(keys) == 0 || keys[0] == "null" {
		return p.responsef(cmdArgs, "No keys found.")
	}

	b, err := json.MarshalIndent(keys, "", "  ")
	if err != nil {
		return p.responsef(cmdArgs, "Error marshaling keys. err=%v", err)
	}

	res := fmt.Sprintf("```\n%s\n```", string(b))
	return p.responsef(cmdArgs, res)
}
