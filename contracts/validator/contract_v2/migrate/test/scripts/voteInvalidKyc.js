const { walletClient, validator } = require("./constant");

async function run() {
  owner = "";
  await walletClient.writeContract({
    ...validator,
    functionName: "voteInvalidKYC",
    args: [owner],
  });
}

run();
