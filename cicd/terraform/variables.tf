locals {
    /**
    Load the nodes data from s3
    Below is the the format the config needs to follow:
    {{Name of the node, in a pattern of 'xdc'+ number. i.e xdc50}}: {
      pk: {{Value of the node private key}},
      ... any other configuration we want to pass.
    }    
    Note: No `n` is allowed in the node name
    **/
    predefinedNodesConfig = jsondecode(data.aws_s3_object.xdc_node_config.body)
    envs = { for tuple in regexall("(.*)=(.*)", file(".env")) : tuple[0] => tuple[1] }
    logLevel = local.envs["log_level"]

    # regions = [
    #   {
    #     "name": "us-east-2", // Ohio
    #     "start": local.envs["us_east_2_start"],
    #     "end": local.envs["us_east_2_end"],
    #   }
    # ]

    # keyNames = {
    #   for r in local.regions : 
    #     r.name => [for i in range(r.start, r.end+1) : "xdc${i}"]
    # }

    # nodeKeys = {
    #   for r in local.regions : 
    #     r.name => { for i in local.keyNames[r.name]: i => local.predefinedNodesConfig[i] }
    # }
    
    rpcTestnetNodeKeys = { "testnet-rpc1": local.predefinedNodesConfig["testnet-rpc1"]} // we hardcode the rpc to a single node for now
    rpcMainnetNodeKeys = { "mainnet-rpc1": local.predefinedNodesConfig["mainnet-rpc1"]} // we hardcode the rpc to a single node for now
}
