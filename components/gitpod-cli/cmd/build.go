// Copyright (c) 2023 Gitpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

package cmd

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gitpod-io/gitpod/common-go/util"
	"github.com/gitpod-io/gitpod/gitpod-cli/pkg/supervisor"
	"github.com/gitpod-io/gitpod/gitpod-cli/pkg/utils"
	"github.com/gitpod-io/gitpod/supervisor/api"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:    "build",
	Short:  "Builds the workspace image (useful to debug a workspace custom image)",
	Hidden: true,
	Run: func(cmd *cobra.Command, args []string) {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		client, err := supervisor.New(ctx)
		if err != nil {
			utils.LogError(ctx, err, "Could not get workspace info required to build", client)
			return
		}
		defer client.Close()

		wsInfo, err := client.Info.WorkspaceInfo(ctx, &api.WorkspaceInfoRequest{})

		if err != nil {
			log.Fatal(err)
		}

		gitpodConfig, err := util.ParseGitpodConfig(wsInfo.CheckoutLocation)
		fmt.Println(gitpodConfig.Image)
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
