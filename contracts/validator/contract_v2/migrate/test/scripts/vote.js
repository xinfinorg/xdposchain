const { walletClient, validator } = require("./constant");

async function run() {
  const res = await walletClient.writeContract({
    ...validator,
    functionName: "propose",
    args: ["0x2C7e9c9C48917a0386AB20791cdfA698B0BC976b"],
    value: "10000000000000000000000000",
  });

  await walletClient.writeContract({
    ...validator,
    functionName: "vote",
    args: ["0x2C7e9c9C48917a0386AB20791cdfA698B0BC976b"],
    value: "10000000000000000000000000",
  });

  await walletClient.writeContract({
    ...validator,
    functionName: "unvote",
    args: [
      "0x2C7e9c9C48917a0386AB20791cdfA698B0BC976b",
      "10000000000000000000000000",
    ],
  });
}

run();
