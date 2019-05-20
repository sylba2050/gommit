package main

import (
	"fmt"
    "os/exec"

	"github.com/manifoldco/promptui"
)

func main() {
	items := []string{
        "feat",
        "fix",
        "update",
        "style",
        "doc",
        "add",
        "delete",
        "refactor",
        "perf",
        "disable",
    }
	index := -1
	var selected string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    "Select Prefix",
			Items:    items,
			AddLabel: "Other",
		}

		index, selected, err = prompt.Run()

		if index == -1 {
			items = append(items, selected)
		}
	}

	if err != nil {
		fmt.Printf("Select failed %v\n", err)
		return
	}

    prompt := promptui.Prompt{
		Label: selected,
	}

	input, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

    err = exec.Command("git", "commit", "-m", "\"" + selected + ": " + input + "\"").Run()
	if err != nil {
		fmt.Printf("commit failed %v\n", err)
		return
	}
}
