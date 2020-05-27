package stack

import (
	"io"
	"strings"

	corev1alpha1 "github.com/pivotal/kpack/pkg/apis/core/v1alpha1"
	expv1alpha1 "github.com/pivotal/kpack/pkg/apis/experimental/v1alpha1"
	"github.com/spf13/cobra"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/pivotal/build-service-cli/pkg/commands"
	"github.com/pivotal/build-service-cli/pkg/k8s"
)

func NewStatusCommand(clientSetProvider k8s.ClientSetProvider) *cobra.Command {
	var (
		verbose bool
	)

	cmd := &cobra.Command{
		Use:          "status <name>",
		Short:        "Display stack status",
		Long:         `Prints detailed information about the status of a specific stack.`,
		Example:      "tbctl stack status my-stack",
		Args:         cobra.ExactArgs(1),
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cs, err := clientSetProvider.GetClientSet("")
			if err != nil {
				return err
			}

			stack, err := cs.KpackClient.ExperimentalV1alpha1().Stacks().Get(args[0], metav1.GetOptions{})
			if err != nil {
				return err
			}

			return displayStackStatus(cmd.OutOrStdout(), stack, verbose)
		},
	}

	cmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "display mixins")

	return cmd
}

func displayStackStatus(out io.Writer, s *expv1alpha1.Stack, verbose bool) error {
	writer := commands.NewStatusWriter(out)

	items := []string{
		"Status", getStatusText(s),
		"Id", s.Status.Id,
		"Run Image", s.Status.RunImage.LatestImage,
		"Build Image", s.Status.BuildImage.LatestImage,
	}

	if verbose {
		items = append(items, "Mixins", strings.Join(s.Status.Mixins, ", "))
	}

	if err := writer.AddBlock("", items...); err != nil {
		return err
	}

	return writer.Write()
}

func getStatusText(s *expv1alpha1.Stack) string {
	if cond := s.Status.GetCondition(corev1alpha1.ConditionReady); cond != nil {
		if cond.Status == corev1.ConditionTrue {
			return "Ready"
		} else if cond.Status == corev1.ConditionFalse {
			return "Not Ready"
		}
	}
	return "Unknown"
}
