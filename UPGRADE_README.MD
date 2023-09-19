# Upgrade Geth EIP's Documentation

Visit: [XinFin.org](https://xinfin.org)
Contribute: [Developer Docs](https://docs.xinfin.org)

#### OVERVIEW

Most of represented features are similar to Ethereum EIP's 

### List of EIP's:

##### Berlin EIPs

* EIP-2565 – lowers ModExp gas cost 
* EIP-2718 – enables easier support for multiple transaction types *

##### London EIP's:

* EIP-1559 – improves the transaction fee market *
* EIP-3198 – returns the BASEFEE from a block *
* EIP-3541 - prevents deploying contracts starting with 0xEF 

##### Paris EIPs:

* EIP-4399 – Supplant DIFFICULTY opcode with PREVRANDAO 

##### Shanghai EIPs

* EIP-3651– Starts the COINBASE address warm
* EIP-3855 – New PUSH0 instruction
* EIP-3860 – Limit and meter initcode

To apply listed eips it is necessary to add into the Genesis the below lines:

```text
"berlinBlock": 5,
"londonBlock": 6,
```

under the line `"constantinopleBlock": 4,` 

(actually you can be modify cmd/puppeth/wizard_genesis.go to setup all hard forks blocks by puppeth)

Notice: eips are marked with * **require to be implemented into the engine**  

***

### Test cases 

#### EIP 3541

Disallow new code starting with the 0xEF byte to be deployed. Code already existing in the account trie starting with 0xEF byte is not affected semantically by this change.

Test:

After block.number == HF_BLOCK new contract creation (via create transaction, CREATE or CREATE2 instructions) results in an exceptional abort if the code's first byte is 0xEF.

https://github.com/ethereum/go-ethereum/blob/65a17c00c7a903f314db2e53381bed19716d50db/core/vm/evm.go#L471

#### EIP 4399

This EIP supplants the semantics of the return value of existing DIFFICULTY (0x44) opcode and renames the opcode to PREVRANDAO (0x44).

Test: 

Beginning with TRANSITION_BLOCK, the DIFFICULTY (0x44) instruction MUST return the value of the mixHash field.

https://github.com/Rock-n-Block/XDPoSChain/blob/d2ce55ec960845bb79ead657aef3a0d24be88d8e/core/types/block.go#L83

#### EIP 3651

The COINBASE address shall be warm at the start of transaction execution, in accordance with the actual cost of reading that account.

#### EIP 3855

Introduce the PUSH0 (0x5f) instruction, which pushes the constant value 0 onto the stack.

#### EIP 3860

Limit and meter initcode

Test:

* Creation transaction with gas limit enough to cover initcode cost
* Creation transaction with gas limit enough to cover intrinsic cost except initcode cost
* CREATE/CREATE2/creation transaction with len(initcode) at MAX_INITCODE_SIZE
* CREATE/CREATE2/creation transaction with len(initcode) at MAX_INITCODE_SIZE+1

https://github.com/ethereum/go-ethereum/blob/65a17c00c7a903f314db2e53381bed19716d50db/core/state_transition.go#L103

#### EIP 2565

Gas cost of the precompiled contract should be more than 200

***

Other eips require implementation into XDPoS engine but can be tested via unittests and with ethash engine
All added eips have unittest which are represented in  https://github.com/ethereum/go-ethereum