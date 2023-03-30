package cli

import (
	"path/filepath"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	projectstypes "github.com/lavanet/lava/x/projects/types"
	"github.com/lavanet/lava/x/subscription/types"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var _ = strconv.Itoa(0)

func CmdAddProject() *cobra.Command {
	var disable bool
	cmd := &cobra.Command{
		Use:   "add-project [project-name] [optional: project-description]",
		Short: "Add a new project to a subcsciption",
		Long: `The add-project command allows the subscription owner to create a new project and associate 
		it with its subscription.  Optionally, you can determine the policy of the project using a YAML file
		(see example in cookbook/projects/example_policy.yml). You can also optionally provide a YAML file which 
		consists additional accounts to be added to the project (see example in cookbook/project/example_project_keys.yml).
		Note that in project keys, to define the key type, you should follow the enum described in the top of example_project_keys.yml.
		Finally, you can optionally create a disabled project by using the "--disable" flag.
		Note, after the project is added, its name (a.k.a. index) is 
		changed to "<project_subscription_address>-<original_project_name>".`,
		Example: `required flags: --from <admin-key> (the project's subscription address is also considered admin)
				  optional flags: --policy-file <policy-file-path>, --project-keys-file <project-keys-file-path>, --disable
				  
		lavad tx subscription add-project [project-file-path] --from <admin-key>`,
		Args: cobra.RangeArgs(1, 2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			projectName := args[0]
			projectDescription := ""
			if len(args) == 2 {
				projectDescription = args[1]
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			creator := clientCtx.GetFromAddress().String()

			var policy projectstypes.Policy
			if cmd.Flags().Lookup("policy-file").Changed {
				policyFilePath, err := cmd.Flags().GetString("policy-file")
				if err != nil {
					return err
				}

				err = readYaml(policyFilePath, "Policy", &policy)
				if err != nil {
					return err
				}
			}

			var projectKeys []projectstypes.ProjectKey
			if cmd.Flags().Lookup("project-keys-file").Changed {
				projectKeysFilePath, err := cmd.Flags().GetString("project-keys-file")
				if err != nil {
					return err
				}

				err = readYaml(projectKeysFilePath, "Project-Keys", &projectKeys)
				if err != nil {
					return err
				}
			}

			// keep all the inputs in a single projectData object (used as a container)
			projectData := projectstypes.ProjectData{
				Name:        projectName,
				Description: projectDescription,
				Enabled:     !disable,
				ProjectKeys: projectKeys,
				Policy:      policy,
			}

			msg := types.NewMsgAddProject(
				creator,
				projectData,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	cmd.MarkFlagRequired(flags.FlagFrom)
	cmd.Flags().String("policy-file", "", "policy file path (optional)")
	cmd.Flags().String("project-keys-file", "", "project keys file path (optional)")
	cmd.Flags().BoolVar(&disable, "disable", false, "disables the project (optional)")
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func readYaml(filePath string, primaryKey string, content interface{}) error {
	configPath, configName := filepath.Split(filePath)
	if configPath == "" {
		configPath = "."
	}
	viper.SetConfigName(configName)
	viper.SetConfigType("yml")
	viper.AddConfigPath(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	err = viper.GetViper().UnmarshalKey(primaryKey, content, func(dc *mapstructure.DecoderConfig) {
		dc.ErrorUnset = true
		dc.ErrorUnused = true
	})
	if err != nil {
		return err
	}

	return nil
}
