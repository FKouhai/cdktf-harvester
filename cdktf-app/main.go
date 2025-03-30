package main

import (
	"cdk.tf/go/stack/stacks"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

func main() {
	app := cdktf.NewApp(nil)
	var vm stacks.VM_Definition
	// TODO wrap around cobra to accept yaml files as arguments whenever running the binary
	vm.ParseYaml("config.yml")
	b := vm.ToVMConfig()

	stacks.VM(app, "harvester-vm", b)

	app.Synth()
}
