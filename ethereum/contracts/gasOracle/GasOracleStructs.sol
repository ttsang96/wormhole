// contracts/Structs.sol
// SPDX-License-Identifier: Apache 2

pragma solidity ^0.8.0;

contract GasOracleStructs {
    
    struct UpgradeContract {
        // Governance Header
        // module: "GasOracle" left-padded
        bytes32 module;
        // governance action: 2
        uint8 action;
        // governance paket chain id
        uint16 chainId;
        // Address of the new contract
        bytes32 newContract;
    }
}
