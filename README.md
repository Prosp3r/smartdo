# SMART-TODO

### A simple app built to deploy and interract with an ERC20 Token Smart Contract on ETH compatible testnets including Rinkeby testnet written in Go.

This is a simple smart contract to test out the workings of Go interractions with ERC20 Tokens
Can deploy and manage multiple contracts from this single interface.

### Usage
 $ ./smartdo <command> <options>

list of functionalities and the commands to test them out


1. Create account - Create a new encrypted ethereum compatible wallet
    $ ./smartdo adduser username password
    An account is needed to carry out the different operations below. Encrypted accounts created are temporaily stored in the ./wallet folder.

2. Deploy contract
    This command deploys the pre written ERC20 Toke to the active Testnet or mainnet. Information for all deployed contracts are recorded in files located in ./loadedcontracts folder named according to their given names at the time of deployment.
	$ ./smartdo deploy username password <contractname>

3. Check Wallet Address
    This command returns the wallet address (hex) of the provided username and password provided the account was created on this deployment. Encrypted accounts created are temporaily stored in the ./wallet folder.
    $ ./smartdo mywallet username password
4. Check balance
    $ ./smartdo balance username password
5. Send wei
    Sends wei to the provided address(hex) within the active network(mainnet or testnets)
    $ ./smartdo sendwei username password <recipient_address e.g 0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d> <sendAmount e.g. 200>
6. Mint Token
    This command tells the deployed smart contract to mint tokens and assign them to the given recipient_address.
    $ ./smartdo contract-mint username password <contractName e.g. logi> <recipient_address e.g 0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d> <amountOfTokens e.g. 200000000000000>
7. Transfer Tokens
    This command transfers token from the reserves on the smart contract to the recipient_address. it can only be ran by the token deployer account.
    $ ./smartdo contract-transfer username password <contractName e.g. logi> <recipient_address e.g 0x8be9a9FCA9861b39487C8513C0EfD2D4C697011d> <amountOfTokens e.g. 2000000000000000000>



... A work in progress


