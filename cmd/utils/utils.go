package utils

import (
	"github.com/rupayaproject/rupaya/eth"
	"github.com/rupayaproject/rupaya/eth/downloader"
	"github.com/rupayaproject/rupaya/ethstats"
	"github.com/rupayaproject/rupaya/les"
	"github.com/rupayaproject/rupaya/node"
	"github.com/rupayaproject/rupaya/rupx"
	whisper "github.com/rupayaproject/rupaya/whisper/whisperv6"
)

// RegisterEthService adds an Ethereum client to the stack.
func RegisterEthService(stack *node.Node, cfg *eth.Config) {
	var err error
	if cfg.SyncMode == downloader.LightSync {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			return les.New(ctx, cfg)
		})
	} else {
		err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
			var rupXServ *rupx.RupX
			ctx.Service(&rupXServ)

			fullNode, err := eth.New(ctx, cfg, rupXServ)
			if fullNode != nil && cfg.LightServ > 0 {
				ls, _ := les.NewLesServer(fullNode, cfg)
				fullNode.AddLesServer(ls)
			}
			return fullNode, err
		})
	}
	if err != nil {
		Fatalf("Failed to register the Ethereum service: %v", err)
	}
}

// RegisterShhService configures Whisper and adds it to the given node.
func RegisterShhService(stack *node.Node, cfg *whisper.Config) {
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return whisper.New(cfg), nil
	}); err != nil {
		Fatalf("Failed to register the Whisper service: %v", err)
	}
}

// RegisterEthStatsService configures the Ethereum Stats daemon and adds it to
// th egiven node.
func RegisterEthStatsService(stack *node.Node, url string) {
	if err := stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		// Retrieve both eth and les services
		var ethServ *eth.Ethereum
		ctx.Service(&ethServ)

		var lesServ *les.LightEthereum
		ctx.Service(&lesServ)

		return ethstats.New(url, ethServ, lesServ)
	}); err != nil {
		Fatalf("Failed to register the Ethereum Stats service: %v", err)
	}
}

func RegisterRupXService(stack *node.Node, cfg *rupx.Config) {
	if err := stack.Register(func(n *node.ServiceContext) (node.Service, error) {
		return rupx.New(cfg), nil
	}); err != nil {
		Fatalf("Failed to register the TomoX service: %v", err)
	}
}
