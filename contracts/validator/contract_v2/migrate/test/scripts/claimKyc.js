const { walletClient, validator } = require("./constant");

async function run() {
  const res = await walletClient.writeContract({
    ...validator,
    functionName: "cliamKYC",
  });
}

run();
