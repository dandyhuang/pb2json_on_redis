package cmd

import (
	"fmt"
	"github.com/dandyhuang/cmd_tools/internal/biz"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd 代表没有调用子命令时的基础命令
var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	// 如果有相关的 action 要执行，请取消下面这行代码的注释
	// Run: func(cmd *cobra.Command, args []string) { },
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./configs/cobra.yaml", "config file (default is $HOME/.cobra.yaml)")
	rootCmd.AddCommand(pb2json)
	rootCmd.AddCommand(json2pb)
}

// Execute 将所有子命令添加到root命令并适当设置标志。会被 main.main() 调用一次。
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var pb2json = &cobra.Command{
	Use:   "pb2json",
	Short: "get redis pb 2 json",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("获取配置文件的mysql.url", viper.GetString(`mysql.url`))
		fmt.Println("获取配置文件的redis.url", viper.GetStringSlice(`redis`))
		fmt.Println("获取配置文件的smtp", viper.GetStringMap("smtp"))
	},
}

var json2pb= &cobra.Command{
	Use:   "json2pb",
	Short: "pb2json into redis",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("获取配置文件的mysql.url", viper.GetString(`mysql.url`))
		fmt.Println("获取配置文件的redis.url", viper.GetStringSlice(`redis`))
		fmt.Println("获取配置文件的smtp", viper.GetStringMap("smtp"))
		jsonFile, err := os.Open( viper.GetString("input_json"))
		if err != nil {
			fmt.Println(err)
		}

		defer jsonFile.Close()

		byteValue, err := ioutil.ReadAll(jsonFile)
		if err != nil {
			fmt.Println(err)
		}
		biz.JsonToPb(viper.GetString("input_proto_file"),viper.GetString("request_message_name"),
			byteValue)
	},
}



func initConfig() {
	// Don't forget to read config either from cfgFile or from home directory!
	fmt.Println("cfgfile:", cfgFile)
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

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".cobra")
	}

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can't read config:", err)
		os.Exit(1)
	}
}
