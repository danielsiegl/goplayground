package main

import (
	"flag"
	"fmt"
	"net/url"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	project := flag.String("project", "", "Project in owner/repo format")
	branch := flag.String("branch", "", "Branch name")
	flag.Parse()

	if *project == "" || *branch == "" {
		fmt.Println("Usage: codex -project <owner/repo> -branch <branch>")
		os.Exit(1)
	}

	workspaceURL := buildURL(*project, *branch)
	if err := openBrowser(workspaceURL); err != nil {
		fmt.Printf("Failed to open browser: %v\n", err)
		os.Exit(1)
	}
}

func buildURL(project, branch string) string {
	segments := strings.Split(project, "/")
	for i, s := range segments {
		segments[i] = url.PathEscape(s)
	}
	p := strings.Join(segments, "/")
	b := url.QueryEscape(branch)
	return fmt.Sprintf("https://chatgpt.com/codex/%s?branch=%s", p, b)
}

func openBrowser(url string) error {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", url)
	case "darwin":
		cmd = exec.Command("open", url)
	default:
		return fmt.Errorf("unsupported platform")
	}
	return cmd.Start()
}
