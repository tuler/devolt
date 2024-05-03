// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Script} from "forge-std/Script.sol";
import {console} from "forge-std/console.sol";
import {Volt} from "@contracts/token/ERC20/Volt.sol";

contract Bytecode is Script {
    function run() external view {
        bytes memory bytecode = type(Volt).creationCode;
        console.log("Bytecode of Volt Contract:");
        console.logBytes(bytecode);
    }
}