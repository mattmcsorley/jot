// Copyright © 2018 Matt McSorley <mcsorleymatt@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/mattmcsorley/jot/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// syncCmd represents the sync command
var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync Documents with an external repository",
	Long: `Sync Documents to an external repository. For example:

Sync documents to a remote git repository`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(viper.GetString("add-remote")) > 0 {
			newViper := viper.New()
			newViper.SetConfigName("config") // name of config file (without extension)
			newViper.AddConfigPath("$HOME/.jot")
			err := newViper.ReadInConfig() // Find and read the config file
			if err != nil {                // Handle errors reading the config file
				panic(fmt.Errorf("Fatal error config file: %s \n", err))
			}
			newViper.Set("repository.remote", viper.GetString("add-remote"))
			newViper.WriteConfig()
		}

		sync := internal.NewSync()
		sync.Push()
	},
}

func init() {
	rootCmd.AddCommand(syncCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// syncCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	syncCmd.Flags().String("add-remote", "", "Add a new remote")
	syncCmd.Flags().Bool("list-remotes", false, "List remotes")

	viper.BindPFlag("add-remote", syncCmd.Flags().Lookup("add-remote"))
	viper.BindPFlag("list-remotes", syncCmd.Flags().Lookup("list-remotes"))
}
