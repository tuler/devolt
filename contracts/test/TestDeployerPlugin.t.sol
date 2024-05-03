//SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import {Test} from "forge-std/Test.sol";
import {console} from "forge-std/console.sol";
import {DeployerPlugin} from "@contracts/proxy/DeployerPlugin.sol";
import {Volt} from "@contracts/token/ERC20/Volt.sol";

contract TestDeployerPlugin is Test {
    DeployerPlugin deployerPlugin;
    Volt volt;

    address application = address(1);

    function setUp() public {
        deployerPlugin = (new DeployerPlugin){
            salt: bytes32(abi.encode(1596))
        }();
    }

    function testDeployAnyContract() public {
        bytes memory bytecode = type(Volt).creationCode;
        vm.prank(application);
        address addr = deployerPlugin.deployAnyContract(
            abi.encodePacked(bytecode, abi.encode(application))
        );
        assertTrue(addr != address(0));        
    }
}