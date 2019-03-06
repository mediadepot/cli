package actions

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/jaypipes/ghw"
	"os/exec"
	"github.com/mediadepot/cli/pkg/errors"
)

type InstallAction struct {
}


// This function does the following
// - checks the latest version off CoreOS installer available.
// - ensure that the coreos installer executable is available locally
// - tries to automatically determine a boot disk (defaults to /dev/sda)
// - tries to find all storage disks
// - prompts for Media depot questions (features to install, SSH keys, etc)
// - generates a YAML file compatible with CoreOS
// - asks to format all drives using extfs4
// 		- checks that disk are empty
// - if satisified, will start CoreOS.
//func (e *InstallAction) Start(data map[string]interface{}, dryRun bool) error {
//
//
//
//}

func (e *InstallAction) Validate() (bool, error) {

	fmt.Println("Checking that `coreos-install` binary available")
	_, coreOsErr := exec.LookPath("coreos-install")
	if coreOsErr != nil {
		return false, errors.InvalidArgumentsError("`coreos-install` executable is missing")
	}

	fmt.Println("Checking that `mkfs.ext4` binary available")
	_, extfsErr := exec.LookPath("mkfs.ext4")
	if extfsErr != nil {
		return false, errors.InvalidArgumentsError("`mkfs.ext4` executable is missing")
	}

	return true, nil
}

func (e *InstallAction) QueryBootDisk() error {

	blockInfo, err := ghw.Block()

	disks := blockInfo.Disks

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F5B4 {{ .Name | cyan }} ({{ .SizeBytes | red }} {{ .DriveType |  red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .SizeBytes | red }} {{ .DriveType |  red }})",
		Selected: "\U0001F5B4 {{ .Name | red | cyan }}",
		Details: `
--------- Disk ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Size (Bytes):" | faint }}	{{ .SizeBytes }}
{{ "Drive Type:" | faint }}	{{ .DriveType }}`,
	}

	prompt := promptui.Select {
		Label:     "Boot Disk",
		Items:     disks,
		Templates: templates,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return errors.InvalidArgumentsError(fmt.Sprintf("Prompt failed %v\n", err))
	}

	fmt.Printf("You choose number %d: %s\n", i+1, disks[i].Name)
	return nil
}

func (e *InstallAction) QueryStorageDisks() error {

	blockInfo, err := ghw.Block()

	disks := blockInfo.Disks

	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F5B4 {{ .Name | cyan }} ({{ .SizeBytes | red }} {{ .DriveType |  red }})",
		Inactive: "  {{ .Name | cyan }} ({{ .SizeBytes | red }} {{ .DriveType |  red }})",
		Selected: "\U0001F5B4 {{ .Name | red | cyan }}",
		Details: `
--------- Disk ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Size (Bytes):" | faint }}	{{ .SizeBytes }}
{{ "Drive Type:" | faint }}	{{ .DriveType }}`,
	}

	prompt := promptui.Select {
		Label:     "Storage Disks",
		Items:     disks,
		Templates: templates,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return errors.InvalidArgumentsError(fmt.Sprintf("Prompt failed %v\n", err))
	}

	fmt.Printf("You choose number %d: %s\n", i+1, disks[i].Name)
	return nil
}


func (e *InstallAction) QueryFeatures() error {

}