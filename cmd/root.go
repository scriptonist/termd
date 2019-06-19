// Copyright © 2019 NAME HERE <EMAIL ADDRESS>

package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gdamore/tcell"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/rivo/tview"
	"github.com/russross/blackfriday"
	"github.com/scriptonist/termd/internal/console"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "termd",
	Short: "Render markdown on terminal",
	RunE:  run,
}

func run(cmd *cobra.Command, args []string) error {
	var output []byte
	cr := new(console.Console)

	if len(args) == 1 {
		fileb, err := ioutil.ReadFile(args[0])
		if err != nil {
			return err
		}
		output = blackfriday.Run(
			fileb,
			blackfriday.WithRenderer(cr),
			blackfriday.WithExtensions(blackfriday.CommonExtensions),
		)
	} else {
		reader := bufio.NewReader(os.Stdin)
		var input []byte
		buffer := make([]byte, 2<<20)
		for {
			count, err := reader.Read(buffer)
			if count == 0 {
				break
			}
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("Unable to read from pipe :%v", err))
			}
			input = append(input, buffer...)
		}
		output = blackfriday.Run(
			input,
			blackfriday.WithRenderer(cr),
			blackfriday.WithExtensions(blackfriday.CommonExtensions),
		)
	}
	app := tview.NewApplication()
	textView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetChangedFunc(func() {
			app.Draw()
		})
	textView.Write(output)

	textView.SetBorder(true)
	textView.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
	})
	if err := app.SetRoot(textView, true).SetFocus(textView).Run(); err != nil {
		panic(err)
	}

	return nil
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.h-cli.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".h-cli" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".h-cli")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
