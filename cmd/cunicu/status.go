package main

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/stv0g/cunicu/pkg/config"
	"github.com/stv0g/cunicu/pkg/crypto"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"

	rpcproto "github.com/stv0g/cunicu/pkg/proto/rpc"
)

var (
	indent bool

	statusCmd = &cobra.Command{
		Use:               "status [interface-name [peer-public-key]]",
		Short:             "Show current status of the cunīcu daemon, its interfaces and peers",
		Aliases:           []string{"show"},
		Run:               status,
		Args:              cobra.RangeArgs(0, 2),
		ValidArgsFunction: interfaceValidArgs,
	}
)

func init() {
	pf := statusCmd.PersistentFlags()
	pf.VarP(&format, "format", "f", "Output `format` (one of: human, json)")
	pf.BoolVarP(&indent, "indent", "i", true, "Format and indent JSON ouput")

	if err := statusCmd.RegisterFlagCompletionFunc("format", cobra.FixedCompletions([]string{"human", "json"}, cobra.ShellCompDirectiveNoFileComp)); err != nil {
		panic(err)
	}

	addClientCommand(rootCmd, statusCmd)
}

func status(cmd *cobra.Command, args []string) {
	p := &rpcproto.GetStatusParams{}

	if len(args) > 0 {
		p.Interface = args[0]
		if len(args) > 1 {
			pk, err := crypto.ParseKey(args[1])
			if err != nil {
				logger.Fatal("Invalid public key", zap.Error(err))
			}

			p.Peer = pk.Bytes()
		}
	}

	sts, err := rpcClient.GetStatus(context.Background(), p)
	if err != nil {
		logger.Fatal("Failed to retrieve status from daemon", zap.Error(err))
	}

	switch format {
	case config.OutputFormatJSON:
		mo := protojson.MarshalOptions{
			AllowPartial:    true,
			UseProtoNames:   true,
			EmitUnpopulated: false,
		}

		if indent {
			mo.Multiline = true
			mo.Indent = "  "
		}

		buf, err := mo.Marshal(sts)
		if err != nil {
			logger.Fatal("Failed to marshal", zap.Error(err))
		}

		if _, err = stdout.Write(buf); err != nil {
			logger.Fatal("Failed to write to stdout", zap.Error(err))
		}

	case config.OutputFormatHuman:
		if err := sts.Dump(stdout, verbosityLevel); err != nil {
			logger.Fatal("Failed to write to stdout", zap.Error(err))
		}
	}
}
