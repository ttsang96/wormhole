// This file contains the token config to be used in the mainnet environment.
//
// This file was generated: Mon Jul 11 2022 14:30:49 GMT+0000 (Coordinated Universal Time) using a min notional of 1000000

package governor

func tokenList() []tokenConfigEntry {
	return []tokenConfigEntry{
		tokenConfigEntry{chain: 1, addr: "c6fa7af3bedbad3a3d65f36aabc97431b1bbe4c2d2f6e0e47ca60203452f5d61", symbol: "USDC", coinGeckoId: "usd-coin", decimals: 6, price: 1.002},            // Addr: EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v, Notional: 6164049
		tokenConfigEntry{chain: 1, addr: "ce010e60afedb22717bd63192f54145a3f965a33bb82d2c7029eb2ce1e208264", symbol: "USDT", coinGeckoId: "tether", decimals: 6, price: 1.002},              // Addr: Es9vMFrzaCERmJfrF4H2FYD4KCoNkY11McCe8BenwNYB, Notional: 1317546
		tokenConfigEntry{chain: 1, addr: "069b8857feab8184fb687f634618c035dac439dc1aeb3b5598a0f00000000001", symbol: "SOL", coinGeckoId: "wrapped-solana", decimals: 8, price: 34.94},       // Addr: So11111111111111111111111111111111111111112, Notional: 4145006
		tokenConfigEntry{chain: 2, addr: "00000000000000000000000018aaa7115705e8be94bffebde57af9bfc265b998", symbol: "AUDIO", coinGeckoId: "audius", decimals: 8, price: 0.331554},          // Addr: 0x18aaa7115705e8be94bffebde57af9bfc265b998, Notional: 2514516
		tokenConfigEntry{chain: 2, addr: "0000000000000000000000001f9840a85d5af5bf1d1762f925bdaddc4201f984", symbol: "UNI", coinGeckoId: "uniswap", decimals: 8, price: 6.07},               // Addr: 0x1f9840a85d5af5bf1d1762f925bdaddc4201f984, Notional: 3811326
		tokenConfigEntry{chain: 2, addr: "00000000000000000000000027702a26126e0b3702af63ee09ac4d1a084ef628", symbol: "ALEPH", coinGeckoId: "aleph", decimals: 8, price: 0.261251},           // Addr: 0x27702a26126e0b3702af63ee09ac4d1a084ef628, Notional: 4087650
		tokenConfigEntry{chain: 2, addr: "0000000000000000000000002e95cea14dd384429eb3c4331b776c4cfbb6fcd9", symbol: "THN", coinGeckoId: "throne", decimals: 8, price: 0.00806137},          // Addr: 0x2e95cea14dd384429eb3c4331b776c4cfbb6fcd9, Notional: 2904441
		tokenConfigEntry{chain: 2, addr: "000000000000000000000000476c5e26a75bd202a9683ffd34359c0cc15be0ff", symbol: "SRM", coinGeckoId: "serum", decimals: 6, price: 0.990139},             // Addr: 0x476c5e26a75bd202a9683ffd34359c0cc15be0ff, Notional: 3777684
		tokenConfigEntry{chain: 2, addr: "00000000000000000000000050d1c9771902476076ecfc8b2a83ad6b9355a4c9", symbol: "FTX Token", coinGeckoId: "ftx-token", decimals: 8, price: 25.28},      // Addr: 0x50d1c9771902476076ecfc8b2a83ad6b9355a4c9, Notional: 76413178
		tokenConfigEntry{chain: 2, addr: "000000000000000000000000514910771af9ca656af840dff83e8264ecf986ca", symbol: "LINK", coinGeckoId: "chainlink", decimals: 8, price: 6.2},             // Addr: 0x514910771af9ca656af840dff83e8264ecf986ca, Notional: 3885156
		tokenConfigEntry{chain: 2, addr: "0000000000000000000000005a98fcbea516cf06857215779fd812ca3bef1b32", symbol: "LDO", coinGeckoId: "lido-dao", decimals: 8, price: 0.626729},          // Addr: 0x5a98fcbea516cf06857215779fd812ca3bef1b32, Notional: 1540284
		tokenConfigEntry{chain: 2, addr: "0000000000000000000000005ab6a4f46ce182356b6fa2661ed8ebcafce995ad", symbol: "SPRT", coinGeckoId: "sportium", decimals: 8, price: 0.531235},         // Addr: 0x5ab6a4f46ce182356b6fa2661ed8ebcafce995ad, Notional: 11301355
		tokenConfigEntry{chain: 2, addr: "0000000000000000000000006b175474e89094c44da98b954eedeac495271d0f", symbol: "DAI", coinGeckoId: "dai", decimals: 8, price: 1.001},                  // Addr: 0x6b175474e89094c44da98b954eedeac495271d0f, Notional: 1392561
		tokenConfigEntry{chain: 2, addr: "0000000000000000000000006b3595068778dd592e39a122f4f5a5cf09c90fe2", symbol: "SUSHI", coinGeckoId: "sushi", decimals: 8, price: 1.22},               // Addr: 0x6b3595068778dd592e39a122f4f5a5cf09c90fe2, Notional: 4375991
		tokenConfigEntry{chain: 2, addr: "000000000000000000000000707f9118e33a9b8998bea41dd0d46f38bb963fc8", symbol: "bETH", coinGeckoId: "anchor-beth-token", decimals: 8, price: 1809.71}, // Addr: 0x707f9118e33a9b8998bea41dd0d46f38bb963fc8, Notional: 9572136
		tokenConfigEntry{chain: 2, addr: "000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48", symbol: "USDC", coinGeckoId: "usd-coin", decimals: 6, price: 1.002},            // Addr: 0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48, Notional: 15554680
		tokenConfigEntry{chain: 2, addr: "000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2", symbol: "WETH", coinGeckoId: "weth", decimals: 8, price: 1141.11},              // Addr: 0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2, Notional: 112845457
		tokenConfigEntry{chain: 2, addr: "000000000000000000000000dac17f958d2ee523a2206206994597c13d831ec7", symbol: "USDT", coinGeckoId: "tether", decimals: 6, price: 1.002},              // Addr: 0xdac17f958d2ee523a2206206994597c13d831ec7, Notional: 12368526
		tokenConfigEntry{chain: 2, addr: "000000000000000000000000e831f96a7a1dce1aa2eb760b1e296c6a74caa9d5", symbol: "NEXM", coinGeckoId: "nexum", decimals: 8, price: 0.30185},             // Addr: 0xe831f96a7a1dce1aa2eb760b1e296c6a74caa9d5, Notional: 60980205
		tokenConfigEntry{chain: 2, addr: "000000000000000000000000f1f955016ecbcd7321c7266bccfb96c68ea5e49b", symbol: "RLY", coinGeckoId: "rally-2", decimals: 8, price: 0.04335401},         // Addr: 0xf1f955016ecbcd7321c7266bccfb96c68ea5e49b, Notional: 1503154
		tokenConfigEntry{chain: 3, addr: "0000000000000000000000008f5cd460d57ac54e111646fc569179144c7f0c28", symbol: "PLY", coinGeckoId: "playnity", decimals: 6, price: 0.01504583},        // Addr: terra13awdgcx40tz5uygkgm79dytez3x87rpg4uhnvu, Notional: 1632611
		tokenConfigEntry{chain: 3, addr: "0000000000000000000000002c71557d2edfedd8330e52be500058a014d329e7", symbol: "BTL", coinGeckoId: "bitlocus", decimals: 6, price: 0.00665618},        // Addr: terra193c42lfwmlkasvcw22l9qqzc5q2dx208tkd7wl, Notional: 3257337
		tokenConfigEntry{chain: 3, addr: "010000000000000000000000000000000000000000000000000000756c756e61", symbol: "LUNA", coinGeckoId: "terra-luna", decimals: 6, price: 0.00010766},     // Addr: uluna, Notional: 16829932
		tokenConfigEntry{chain: 3, addr: "0100000000000000000000000000000000000000000000000000000075757364", symbol: "UST", coinGeckoId: "terrausd", decimals: 6, price: 0.04981399},        // Addr: uusd, Notional: 11547433
		tokenConfigEntry{chain: 4, addr: "00000000000000000000000055d398326f99059ff775485246999027b3197955", symbol: "USDT", coinGeckoId: "tether", decimals: 8, price: 1.002},              // Addr: 0x55d398326f99059ff775485246999027b3197955, Notional: 4471874
		tokenConfigEntry{chain: 4, addr: "000000000000000000000000bb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c", symbol: "WBNB", coinGeckoId: "wbnb", decimals: 8, price: 230.91},               // Addr: 0xbb4cdb9cbd36b01bd1cbaebf2de08d9173bc095c, Notional: 6947299
		tokenConfigEntry{chain: 4, addr: "000000000000000000000000e9e7cea3dedca5984780bafc599bd69add087d56", symbol: "BUSD", coinGeckoId: "binance-usd", decimals: 8, price: 1.002},         // Addr: 0xe9e7cea3dedca5984780bafc599bd69add087d56, Notional: 5948329
		tokenConfigEntry{chain: 4, addr: "000000000000000000000000fafd4cb703b25cb22f43d017e7e0d75febc26743", symbol: "WEYU", coinGeckoId: "weyu", decimals: 8, price: 0.00143735},           // Addr: 0xfafd4cb703b25cb22f43d017e7e0d75febc26743, Notional: 5913640
		tokenConfigEntry{chain: 5, addr: "0000000000000000000000000d500b1d8e8ef31e21c99d1db9a6444d3adf1270", symbol: "WMATIC", coinGeckoId: "wmatic", decimals: 8, price: 0.568687},         // Addr: 0x0d500b1d8e8ef31e21c99d1db9a6444d3adf1270, Notional: 3410326
		tokenConfigEntry{chain: 5, addr: "0000000000000000000000002791bca1f2de4661ed88a30c99a7a9449aa84174", symbol: "USDC", coinGeckoId: "usd-coin", decimals: 6, price: 1.002},            // Addr: 0x2791bca1f2de4661ed88a30c99a7a9449aa84174, Notional: 9635898
		tokenConfigEntry{chain: 5, addr: "000000000000000000000000c2132d05d31c914a87c6611c10748aeb04b58e8f", symbol: "USDT", coinGeckoId: "tether", decimals: 6, price: 1.002},              // Addr: 0xc2132d05d31c914a87c6611c10748aeb04b58e8f, Notional: 3347561
		tokenConfigEntry{chain: 6, addr: "000000000000000000000000b31f66aa3c1e785363f0875a1b74e27b85fd66c7", symbol: "WAVAX", coinGeckoId: "wrapped-avax", decimals: 8, price: 18.21},       // Addr: 0xb31f66aa3c1e785363f0875a1b74e27b85fd66c7, Notional: 4663611
		tokenConfigEntry{chain: 6, addr: "000000000000000000000000b97ef9ef8734c71904d8002f8b6bc66dd9c48a6e", symbol: "USDC", coinGeckoId: "usd-coin", decimals: 6, price: 1.002},            // Addr: 0xb97ef9ef8734c71904d8002f8b6bc66dd9c48a6e, Notional: 3701774
	}
}
