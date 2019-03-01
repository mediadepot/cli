package actions

import (
	"fmt"
	"os/exec"

	"github.com/jaypipes/ghw"
	"github.com/manifoldco/promptui"
)

type InstallAction struct {
}


// This function does the following
// - checks the latest version off CoreOS installer available.
// - ensure that the coreos installer executable is available locally
// - tries to automatically determine a boot disk (defaults to /dev/sda)
// - tries to find all storage disks
// - asks to format all drives using extfs4
// 		- checks that disk are empty
// - prompts for Media depot questions (features to install, SSH keys, etc)
// - generates a YAML file compatible with CoreOS
// - if satisified, will start CoreOS.
func (e *InstallAction) Start(data map[string]interface{}, dryRun bool) error {

	e.Validate()



}

func (e *InstallAction) Validate() (bool, error) {
	errorMsgs := []string{}
	valid := true

	fmt.Println("Checking that `coreos-install` binary available")
	_, lookErr := exec.LookPath("coreos-install")
	if lookErr != nil {
		valid = false
		errorMsgs = append(errorMsgs, "`coreos-install` executable is missing")
	}


	_, lookErr := exec.LookPath("mkfs.ext4")
	if lookErr != nil {
		valid = false
		errorMsgs = append(errorMsgs, "`mkfs.ext4` executable is missing")
	}

	return (valid, errorMsgs)
}