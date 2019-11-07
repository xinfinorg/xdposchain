## XinFin Hybrid Blockchain

XinFin Hybrid Blockchain is an Enterprise ready Blockchain for global trade and finance

Visit: [XinFin.org](https://xinfin.org)
Contribute: [Developer Docs](https://docs.xinfin.org)

## XinFin Network Apothem testnet is community driven project to achieve the following 

- XinFin DPOS (XDPoS) an efficient PBFT consensus that selects 21 random set of validators from MAX 5000 validator Masternodes to achieve a high throughput Energy efficient consensus with instant block finality

- KYC Enforcement on Masternodes for Enterprise Adoption and compliance

- Ability to port/relay limited set of data and transactions from privacy channels to public channel

- Interoperability between applications hosted on Private Blockchains like Corda, Hyperledger, Quorum(JP Morgan) using relayers to XinFin Network

- Customer Centric and consortium driven Governance to equally benefit the validators as well as providing comfort for large scale enterprise applications to be hosted on the Network. This achieves

  - Rapid Upgradability

  - DApps Standardisation for rapid commercialisation

  - Compliance with major global jurisdictions.

### KYC for masternodes

#### OVERVIEW

To add a layer of KYC for masternodes in the current system and a sense of ownership amongst the masternodes hence tying such a cluster of masternodes to physical entity which can held accountable for its actions.

#### Design

We established a bidirectional connection between a candidate and its owner inorder to retrieve a candidate belonging to a specific owner & vice versa.

All the masternodes are recognized by the KYC of their owners and hence are considered as a single verified entity ( for eg. while voting for invalid KYC, only one vote is considered per such cluster )

The contract is very strict in handing out penalty for invalid KYC, it results loss of all funds invested in all of its candidates.

For eg. say A proposes condidates B,C,D by paying for its proposal cost.
If at a later stage if some predecided amount of owners ( investors ) vote that a KYC for a A is invalid then A & all of its candidates (B,C,D) will lose their position & all their funds will be lost ( will remain with contract wallet ).  

### Flow  

To add new masternode

1. Add KYC from a potential owner.

2. Propose a candidate for masternode from that account.

To remove a masternode whose owner has invalid KYC

1. Sign-in from one of the candidates 

2. Vote for invalid kyc funcion

3. When 75% confirmation has been received from the owners ( and not candidates * ), that particular owner against whom the invalid votes have been given will be removed along with all its candidate nodes.

### For developers

### To contribute

Simple create a pull request along with proper reasoning, we'll get back to you.

Our Channels : [Telegram Developer Group](https://t.me/XinFinDevelopers)  or [Public Slack Group](https://launchpass.com/xinfin-public)

