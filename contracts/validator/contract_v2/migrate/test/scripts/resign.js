const { walletClient, validator } = require("./constant");

async function run() {
  const res = await walletClient.writeContract({
    ...validator,
    functionName: "resign",
    args: ["0x2C7e9c9C48917a0386AB20791cdfA698B0BC976b"],

  });
  console.log(res);
}

run();
