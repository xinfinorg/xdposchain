const { createPublicClient, http, createWalletClient } = require("viem");
const { privateKeyToAccount } = require("viem/accounts");
const validatorABI = require("../abi/validatorABI.json");
require("dotenv").config();
const privateKey = process.env.PRIVATE_KEY;
const rpcUrl = process.env.RPC_URL;
const xdc = {
  id: 551,
  name: "XDC Devnet",
  network: "XDC Devnet",
  nativeCurrency: {
    decimals: 18,
    name: "XDC",
    symbol: "XDC",
  },
  rpcUrls: {
    public: { http: [rpcUrl] },
    default: { http: [rpcUrl] },
  },
};

const account = privateKeyToAccount(privateKey);

const publicClient = createPublicClient({
  chain: xdc,
  transport: http(rpcUrl),
});

const walletClient = createWalletClient({
  chain: xdc,
  transport: http(rpcUrl),
  account,
});

const validator = {
  address: "0x0000000000000000000000000000000000000088",
  abi: validatorABI,
};

module.exports = {
  publicClient,
  walletClient,
  validator,
};
