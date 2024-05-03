// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {DeployerPlugin} from "@contracts/proxy/DeployerPlugin.sol";

contract DeployContracts is Script {
    function run() external {
        bytes32 _salt = bytes32(abi.encode(1596));
        vm.startBroadcast();
        DeployerPlugin proxy = new DeployerPlugin{salt: _salt}();
        vm.stopBroadcast();
        console.log(
            "DeployerPlugin address:",
            address(proxy),
            "at network:",
            block.chainid
        );
    }
}