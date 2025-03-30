package main

import (
	"cdk.tf/go/stack/stacks"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(&cdktf.AppConfig{
		SkipValidation: jsii.Bool(true),
	})
	var vm stacks.VM_Definition
	// TODO wrap around cobra to accept yaml files as arguments whenever running the binary
	vm.ParseYaml("config.yml")
	b := vm.ToVMConfig()

	stacks.NewStack(app, "harvester-vm", b)

	app.Synth()
}
