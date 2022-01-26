//SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.0 <0.9.0;


import "../.deps/npm/@openzeppelin/contracts/token/ERC20.sol";
import "../.deps/npm/@openzeppelin/contracts/access/Ownable.sol";

/// @custom:security-contact prosper@samedayshop.com

contract Todo is ERC20, Ownable {
    // address owner;

    // Task[] tasks;
    // struct Task{
    //     string content;
    //     bool status;
    // }


    constructor() ERC20("SWIPE", "MTK") {
        // owner = msg.sender;
        _mint(msg.sender, 10000 * 10 ** decimals());   
    }

    function mint(address to, uint256 amount) public onlyOwner {
        _mint(to, amount);
    }

    // modifier isOwner(){
    //     require(owner == msg.sender);
    //     _;
    // }

    // function add(string memory _content) public isOwner {
    //     tasks.push(Task(_content, false));
    // }

    // function get(uint _id) public isOwner view returns(Task memory) {
    //     return tasks[_id];
    // }

    // function list() public isOwner view returns(Task[] memory){
    //     return tasks;
    // } 

    // function update(uint _id, string memory _content) public {

    //     tasks[_id].content = _content;
    // }

    // function toggle(uint _id) public isOwner {
    //     tasks[_id].status = !tasks[_id].status;
    // }

    // function remove(uint _id) public isOwner {
    //     for (uint i = _id; i < tasks.length-1; i++){
    //         tasks[i] = tasks[i+1];
    //     }
    //     tasks.pop();
    // }
}