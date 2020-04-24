package clusterbuilder

import (
	"fmt"

	"github.com/pivotal/kpack/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func NewDeleteCommand(kpackClient versioned.Interface) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "delete <name>",
		Short:   "Delete a clusterbuilder",
		Long:    "Delete a clusterbuilder from the cluster.",
		Example: "tbctl clusterbuilder delete my-clusterbuilder",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := kpackClient.ExperimentalV1alpha1().CustomClusterBuilders().Delete(args[0], &metav1.DeleteOptions{})
			if err != nil {
				return err
			}

			_, err = fmt.Fprintf(cmd.OutOrStdout(), "\"%s\" deleted\n", args[0])
			return err
		},
		SilenceUsage: true,
	}

	return cmd
}
