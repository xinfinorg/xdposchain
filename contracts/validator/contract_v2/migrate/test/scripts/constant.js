const { createPublicClient, http, createWalletClient } = require("viem");
const { privateKeyToAccount } = require("viem/accounts");
const validatorABI = require("../abi/validatorABI.json");
require("dotenv").config();
const m1PrivateKey = process.env.M1_PRIVATE_KEY;
const m2PrivateKey = process.env.M2_PRIVATE_KEY;
const m3PrivateKey = process.env.M3_PRIVATE_KEY;
const m4PrivateKey = process.env.M4_PRIVATE_KEY;
const m5PrivateKey = process.env.M5_PRIVATE_KEY;

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

const masternode1 = privateKeyToAccount(m1PrivateKey);
const masternode2 = privateKeyToAccount(m2PrivateKey);
const masternode3 = privateKeyToAccount(m3PrivateKey);
const masternode4 = privateKeyToAccount(m4PrivateKey);
const masternode5 = privateKeyToAccount(m5PrivateKey);

const publicClient = createPublicClient({
  chain: xdc,
  transport: http(rpcUrl),
});

const walletClient = createWalletClient({
  chain: xdc,
  transport: http(rpcUrl),
});

const validator = {
  address: "0x0000000000000000000000000000000000000088",
  abi: validatorABI,
};

function expect(condition, message) {
  if (!condition) {
    throw new Error(message || "Assertion failed");
  }
}

module.exports = {
  publicClient,
  walletClient,
  validator,
  expect,
  masternode1,
  masternode2,
  masternode3,
  masternode4,
  masternode5,
};
