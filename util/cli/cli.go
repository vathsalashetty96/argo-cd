// Package cmd provides functionally common to various argo CLIs

package cli

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"strings"

	"github.com/google/shlex"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/klog"
	"k8s.io/kubectl/pkg/util/term"
	"sigs.k8s.io/yaml"

	"github.com/vathsalashetty96/argo-cd/common"
	"github.com/vathsalashetty96/argo-cd/util/errors"
	"github.com/vathsalashetty96/argo-cd/util/io"
)

// NewVersionCmd returns a new `version` command to be used as a sub-command to root
func NewVersionCmd(cliName string) *cobra.Command {
	var short bool
	versionCmd := cobra.Command{
		Use:   "version",
		Short: "Print version information",
		Run: func(cmd *cobra.Command, args []string) {
			version := common.GetVersion()
			fmt.Printf("%s: %s\n", cliName, version)
			if short {
				return
			}
			fmt.Printf("  BuildDate: %s\n", version.BuildDate)
			fmt.Printf("  GitCommit: %s\n", version.GitCommit)
			fmt.Printf("  GitTreeState: %s\n", version.GitTreeState)
			if version.GitTag != "" {
				fmt.Printf("  GitTag: %s\n", version.GitTag)
			}
			fmt.Printf("  GoVersion: %s\n", version.GoVersion)
			fmt.Printf("  Compiler: %s\n", version.Compiler)
			fmt.Printf("  Platform: %s\n", version.Platform)
		},
	}
	versionCmd.Flags().BoolVar(&short, "short", false, "print just the version number")
	return &versionCmd
}

// AddKubectlFlagsToCmd adds kubectl like flags to a command and returns the ClientConfig interface
// for retrieving the values.
func AddKubectlFlagsToCmd(cmd *cobra.Command) clientcmd.ClientConfig {
	loadingRules := clientcmd.NewDefaultClientConfigLoadingRules()
	loadingRules.DefaultClientConfig = &clientcmd.DefaultClientConfig
	overrides := clientcmd.ConfigOverrides{}
	kflags := clientcmd.RecommendedConfigOverrideFlags("")
	cmd.PersistentFlags().StringVar(&loadingRules.ExplicitPath, "kubeconfig", "", "Path to a kube config. Only required if out-of-cluster")
	clientcmd.BindOverrideFlags(&overrides, cmd.PersistentFlags(), kflags)
	return clientcmd.NewInteractiveDeferredLoadingClientConfig(loadingRules, &overrides, os.Stdin)
}

// PromptCredentials is a helper to prompt the user for a username and password (unless already supplied)
func PromptCredentials(username, password string) (string, string) {
	return PromptUsername(username), PromptPassword(password)
}

// PromptUsername prompts the user for a username value
func PromptUsername(username string) string {
	return PromptMessage("Username", username)
}

// PromptMessage prompts the user for a value (unless already supplied)
func PromptMessage(message, value string) string {
	for value == "" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(message + ": ")
		valueRaw, err := reader.ReadString('\n')
		errors.CheckError(err)
		value = strings.TrimSpace(valueRaw)
	}
	return value
}

// PromptPassword prompts the user for a password, without local echo. (unless already supplied)
func PromptPassword(password string) string {
	for password == "" {
		fmt.Print("Password: ")
		passwordRaw, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		errors.CheckError(err)
		password = string(passwordRaw)
		fmt.Print("\n")
	}
	return password
}

// AskToProceed prompts the user with a message (typically a yes or no question) and returns whether
// or not they responded in the affirmative or negative.
func AskToProceed(message string) bool {
	for {
		fmt.Print(message)
		reader := bufio.NewReader(os.Stdin)
		proceedRaw, err := reader.ReadString('\n')
		errors.CheckError(err)
		switch strings.ToLower(strings.TrimSpace(proceedRaw)) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		}
	}
}

// ReadAndConfirmPassword is a helper to read and confirm a password from stdin
func ReadAndConfirmPassword() (string, error) {
	for {
		fmt.Print("*** Enter new password: ")
		password, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return "", err
		}
		fmt.Print("\n")
		fmt.Print("*** Confirm new password: ")
		confirmPassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
		if err != nil {
			return "", err
		}
		fmt.Print("\n")
		if string(password) == string(confirmPassword) {
			return string(password), nil
		}
		log.Error("Passwords do not match")
	}
}

// SetLogFormat sets a logrus log format
func SetLogFormat(logFormat string) {
	switch strings.ToLower(logFormat) {
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	case "text":
		if os.Getenv("FORCE_LOG_COLORS") == "1" {
			log.SetFormatter(&log.TextFormatter{ForceColors: true})
		}
	default:
		log.Fatalf("Unknown log format '%s'", logFormat)
	}
}

// SetLogLevel parses and sets a logrus log level
func SetLogLevel(logLevel string) {
	level, err := log.ParseLevel(logLevel)
	errors.CheckError(err)
	log.SetLevel(level)
}

// SetGLogLevel set the glog level for the k8s go-client
func SetGLogLevel(glogLevel int) {
	klog.InitFlags(nil)
	_ = flag.Set("logtostderr", "true")
	_ = flag.Set("v", strconv.Itoa(glogLevel))
}

func writeToTempFile(pattern string, data []byte) string {
	f, err := ioutil.TempFile("", pattern)
	errors.CheckError(err)
	defer io.Close(f)
	_, err = f.Write(data)
	errors.CheckError(err)
	return f.Name()
}

func stripComments(input []byte) []byte {
	var stripped []byte
	lines := bytes.Split(input, []byte("\n"))
	for i, line := range lines {
		if bytes.HasPrefix(bytes.TrimSpace(line), []byte("#")) {
			continue
		}
		stripped = append(stripped, line...)
		if i < len(lines)-1 {
			stripped = append(stripped, '\n')
		}
	}
	return stripped
}

const (
	defaultEditor  = "vi"
	editorEnv      = "EDITOR"
	commentsHeader = `# Please edit the object below. Lines beginning with a '#' will be ignored,
# and an empty file will abort the edit. If an error occurs while saving this file will be
# reopened with the relevant failures."
`
)

func setComments(input []byte, comments string) []byte {
	input = stripComments(input)
	var commentLines []string
	for _, line := range strings.Split(comments, "\n") {
		if line != "" {
			commentLines = append(commentLines, "# "+line)
		}
	}
	parts := []string{commentsHeader}
	if len(commentLines) > 0 {
		parts = append(parts, strings.Join(commentLines, "\n"))
	}
	parts = append(parts, string(input))
	return []byte(strings.Join(parts, "\n"))
}

// InteractiveEdit launches an interactive editor
func InteractiveEdit(filePattern string, data []byte, save func(input []byte) error) {
	var editor string
	var editorArgs []string
	if overrideEditor := os.Getenv(editorEnv); overrideEditor == "" {
		editor = defaultEditor
	} else {
		parts := strings.Fields(overrideEditor)
		editor = parts[0]
		editorArgs = parts[1:]
	}

	errorComment := ""
	for {
		data = setComments(data, errorComment)
		tempFile := writeToTempFile(filePattern, data)
		cmd := exec.Command(editor, append(editorArgs, tempFile)...)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		err := (term.TTY{In: os.Stdin, TryDev: true}).Safe(cmd.Run)
		errors.CheckError(err)

		updated, err := ioutil.ReadFile(tempFile)
		errors.CheckError(err)
		if string(updated) == "" || string(updated) == string(data) {
			errors.CheckError(fmt.Errorf("Edit cancelled, no valid changes were saved."))
			break
		} else {
			data = stripComments(updated)
		}

		err = save(data)
		if err == nil {
			break
		}
		errorComment = err.Error()
	}
}

// PrintDiff prints a diff between two unstructured objects to stdout using an external diff utility
// Honors the diff utility set in the KUBECTL_EXTERNAL_DIFF environment variable
func PrintDiff(name string, live *unstructured.Unstructured, target *unstructured.Unstructured) error {
	tempDir, err := ioutil.TempDir("", "argocd-diff")
	if err != nil {
		return err
	}
	targetFile := path.Join(tempDir, name)
	targetData := []byte("")
	if target != nil {
		targetData, err = yaml.Marshal(target)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(targetFile, targetData, 0644)
	if err != nil {
		return err
	}
	liveFile := path.Join(tempDir, fmt.Sprintf("%s-live.yaml", name))
	liveData := []byte("")
	if live != nil {
		liveData, err = yaml.Marshal(live)
		if err != nil {
			return err
		}
	}
	err = ioutil.WriteFile(liveFile, liveData, 0644)
	if err != nil {
		return err
	}
	cmdBinary := "diff"
	var args []string
	if envDiff := os.Getenv("KUBECTL_EXTERNAL_DIFF"); envDiff != "" {
		parts, err := shlex.Split(envDiff)
		if err != nil {
			return err
		}
		cmdBinary = parts[0]
		args = parts[1:]
	}
	cmd := exec.Command(cmdBinary, append(args, liveFile, targetFile)...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
