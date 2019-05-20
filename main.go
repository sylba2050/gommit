package main

import (
	"fmt"
	"strings"
    "os/exec"

	"github.com/manifoldco/promptui"

    "./settings"
)

func main() {
	prefixes := settings.GetSettings()
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "{{ .Name | red }} {{ (.Description) | faint }}",
		Inactive: "{{ .Name | cyan }} {{ (.Description)| faint }}",
		Selected: "{{ .Name | cyan }}",
	}

	searcher := func(input string, index int) bool {
		pepper := prefixes[index]
		name := strings.Replace(strings.ToLower(pepper.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	selectPrompt := promptui.Select{
		Label:     "Select Prefix",
		Items:     prefixes,
		Templates: templates,
		Size:      4,
		Searcher:  searcher,
	}

	i, _, err := selectPrompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

    prompt := promptui.Prompt{
		Label: prefixes[i].Name,
	}

	input, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

    out, err := exec.Command("git", "commit", "-m", prefixes[i].Name + ": " + input).Output()
	if err != nil {
		fmt.Printf("commit failed %v\n", err)
	}
    fmt.Printf("%s", out)
}
