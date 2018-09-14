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
	"os"
	"path/filepath"
	"strings"

	"github.com/mattmcsorley/jot/internal"
	"github.com/mattmcsorley/jot/internal/file"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// journalCmd represents the journal command
var journalCmd = &cobra.Command{
	Use:   "journal",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		journalBaseDir := filepath.Join(viper.GetString("library"), "journal")
		os.MkdirAll(journalBaseDir, os.ModePerm)
		documentService := file.NewDocumentService(journalBaseDir)
		journal := internal.NewJournal(documentService)
		journal.SaveContent(strings.Join(args, " "), viper.GetString("templates.journal"))

	},
}

func init() {
	rootCmd.AddCommand(journalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// journalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// journalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
