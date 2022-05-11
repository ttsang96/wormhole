const algosdk = require('algosdk');
const TestLib = require('./testlib.ts')
const testLib = new TestLib.TestLib()
const fs = require('fs');
const path = require('path');


import  {
    getAlgoClient,
    submitVAA, 
    submitVAAHdr, 
    simpleSignVAA, 
    getIsTransferCompletedAlgorand,
    parseVAA, 
    CORE_ID,
    TOKEN_BRIDGE_ID,
    attestFromAlgorand,
    AccountToSigner
} from "@certusone/wormhole-sdk/lib/cjs/algorand/Algorand";

import {
    getTempAccounts,
} from "@certusone/wormhole-sdk/lib/cjs/algorand/Helpers";


const guardianKeys = [
    "beFA429d57cD18b7F8A4d91A2da9AB4AF05d0FBe"
]
const guardianPrivKeys = [
    "cfb12303a19cde580bb4dd771639b0d26bc68353645571a8cff516ab2ee113a0"
]


class AlgoTests {
    constructor() {
    }

    async runTests() {
        let seq = Math.floor(new Date().getTime() / 1000.0);

        console.log("test start");
        let client = getAlgoClient();

        let accounts = await getTempAccounts();
        let player = AccountToSigner(accounts[0])

        console.log("attesting some ALGO");
        console.log(await attestFromAlgorand(client, player, 0))
        process.exit(0);

        console.log("test complete");
    }
};

let t = new AlgoTests()
t.runTests()


