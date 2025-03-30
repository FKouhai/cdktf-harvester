package stacks

import (
	"log"
	"os"

	"cdk.tf/go/stack/generated/harvester/harvester/provider"
	"cdk.tf/go/stack/generated/harvester/harvester/virtualmachine"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	yaml "gopkg.in/yaml.v3"
)

type VM_Definition struct {
	CPU             float64   `yaml:"cpu"`
	Name            string    `yaml:"name"`
	Memory          string    `yaml:"memory"`
	Reserved_memory string    `yaml:"reserved_memory,omitempty"`
	Efi             bool      `yaml:"efi,omitempty"`
	Secure_boot     bool      `yaml:"secure_boot,omitempty"`
	Ssh_Keys        []*string `yaml:"ssh_keys,omitempty"`
	Disk            []struct {
		Name       string `yaml:"name"`
		Type       string `yaml:"type"`
		Size       string `yaml:"size"`
		Bus        string `yaml:"bus"`
		Boot_order int    `yaml:"boot_order"`
	} `yaml:"disk"`
	Network_interface []struct {
		Name         string `yaml:"name"`
		Network_name string `yaml:"network_name"`
		Model        string `yaml:"model,omitempty"`
		Type         string `yaml:"type,omitempty"`
	} `yaml:"network_interface"`
}

func (v *VM_Definition) ToVMConfig() *virtualmachine.VirtualmachineConfig {
	var vmConfig virtualmachine.VirtualmachineConfig
	vmConfig.Name = &v.Name
	vmConfig.Cpu = &v.CPU
	vmConfig.Memory = &v.Memory
	if v.Reserved_memory != "" {
		vmConfig.ReservedMemory = &v.Reserved_memory
	}
	vmConfig.Efi = &v.Efi
	vmConfig.SshKeys = &v.Ssh_Keys
	vmConfig.Disk = &v.Disk
	vmConfig.NetworkInterface = &v.Network_interface

	return &vmConfig
}

func (v *VM_Definition) ParseYaml(file string) *VM_Definition {
	b, err := os.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
	}
	err = yaml.Unmarshal(b, &v)
	if err != nil {
		log.Fatalln(err)
	}
	return v

}
func NewStack(scope constructs.Construct, name string, vmConfig *virtualmachine.VirtualmachineConfig) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &name)
	var prov_config provider.HarvesterProviderConfig
	prov_config.Kubecontext = jsii.String("local")
	provider.NewHarvesterProvider(scope, jsii.String(*vmConfig.Name), &prov_config)
	VM(stack, jsii.String(*vmConfig.Name), vmConfig)

	// The code that defines your stack goes here

	return stack

}
func VM(scope constructs.Construct, id *string, vmConfig *virtualmachine.VirtualmachineConfig) virtualmachine.Virtualmachine {
	vm := virtualmachine.NewVirtualmachine(scope, id, vmConfig)

	return vm
}

func ISO(scope constructs.Construct, id string) cdktf.TerraformStack {
	stack := cdktf.NewTerraformStack(scope, &id)

	// The code that defines your stack goes here

	return stack
}
