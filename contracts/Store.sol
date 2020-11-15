pragma solidity >0.5.0;

contract Store {

    struct Cart {
        bytes32 key;
        bytes32 value;
    }

    string public version;

    mapping (uint => Cart) public items;

    uint public itemCount;

    constructor() public {}

    function setItem(Cart memory _car) private {
        items[_car.key] = _car.value;
        itemCount++;
    }
}