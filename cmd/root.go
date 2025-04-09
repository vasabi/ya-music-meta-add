package cmd

import (
	"context"
	"github.com/TheonAegor/go-framework/pkg/config"
	"github.com/TheonAegor/go-framework/pkg/config/envConfig"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"ya-music-meta-add/cmd/metadata"
	"ya-music-meta-add/internal"
)

var (
	cfgFile  string
	logLevel logrus.Level

	rootCmd = &cobra.Command{
		Use:   "yama",
		Short: "Allow to upd metadata and sort downloaded mp3 ya music files",
		Long:  `.`,
	}
	log = logrus.New()
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig, initHandler)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.PersistentFlags().Uint32VarP((*uint32)(&logLevel), "verbosity", "v", 5, "verbosity level")

	log.Level = logLevel

	err := viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	if err != nil {
		panic(err)
	}

	rootCmd.AddCommand(metadata.Metadata)
}

func initConfig() {
	viper.AutomaticEnv()

	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		root, err := os.Getwd()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra.yaml" (without extension).
		viper.AddConfigPath(root)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}

	ctx := context.Background()

	if err := viper.ReadInConfig(); err == nil {
		logrus.Infof("Using config file path: %s", viper.ConfigFileUsed())
	}

	if err := viper.Unmarshal(&internal.GlobalConfig); err != nil {
		logrus.Fatalf("cannot unmarshal config: %s", err.Error())
	}

	opts := append([]config.Config(nil),
		config.NewConfig(config.Struct(&internal.GlobalConfig)),
		envConfig.NewConfig(config.Struct(&internal.GlobalConfig)),
	)

	if err := config.Load(ctx, opts...); err != nil {
		logrus.Fatalf("failed to load config: %+v", err)
	}

	logrus.Debugf("config: %+v", internal.GlobalConfig)
}

func initHandler() {
	ctx := context.Background()

	if ctx == nil {
		ctx = context.Background()
	}
}
