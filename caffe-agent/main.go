package main

import (
	"fmt"
	"os"

	"github.com/rai-project/caffe"
	_ "github.com/rai-project/caffe/predict"
	cmd "github.com/rai-project/dlframework/framework/cmd/server"
	"github.com/rai-project/tracer"
)

func main() {

	rootCmd, err := cmd.NewRootCommand(caffe.FrameworkManifest)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	defer tracer.Close()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
