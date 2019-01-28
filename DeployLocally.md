// 1. create a folder to store XinFin-Test-Net data on your machine
$ export DATA_DIR=/path/to/your/data/folder
$ mkdir -p $DATA_DIR/XDC

// 2. download our genesis file
$ export GENESIS_PATH=$DATA_DIR/genesis.json
$ curl -L https://github.com/XinFinOrg/XDPoS-TestNet-Apothem/raw/master/genesis/testnet.json -o $GENESIS_PATH

// 3. init the chain from genesis
$ XDC init $GENESIS_PATH --datadir $DATA_DIR

// 4. get a test account. Create a new one if you don't have any:
$ export KEYSTORE_DIR=keystore
$ touch $DATA_DIR/password && echo 'your-password' > $DATA_DIR/password
$ XDC account new \
      --datadir $DATA_DIR \
      --keystore $KEYSTORE_DIR \
      --password $DATA_DIR/password

// if you already have a test account, import it now
$ XDC  account import ./private_key \
      --datadir $DATA_DIR \
      --keystore $KEYSTORE_DIR \
      --password $DATA_DIR/password

// get the account
$ account=$(
  XDC account list --datadir $DATA_DIR  --keystore $KEYSTORE_DIR \
  2> /dev/null \
  | head -n 1 \
  | cut -d"{" -f 2 | cut -d"}" -f 1
)

// 5. prepare the bootnodes list
$ export BOOTNODES="enode://4d3c2cc0ce7135c1778c6f1cfda623ab44b4b6db55289543d48ecfde7d7111fd420c42174a9f2fea511a04cf6eac4ec69b4456bfaaae0e5bd236107d3172b013@52.221.28.223:30301,enode://298780104303fcdb37a84c5702ebd9ec660971629f68a933fd91f7350c54eea0e294b0857f1fd2e8dba2869fcc36b83e6de553c386cf4ff26f19672955d9f312@13.251.101.216:30301,enode://46dba3a8721c589bede3c134d755eb1a38ae7c5a4c69249b8317c55adc8d46a369f98b06514ecec4b4ff150712085176818d18f59a9e6311a52dbe68cff5b2ae@13.250.94.232:30301"

// 6. Start up XDC now
$ export NAME=YOUR_NODE_NAME
$ XDC \
  --verbosity 4 \
  --datadir $DATA_DIR \
  --keystore $KEYSTORE_DIR \
  --identity $NAME \
  --password $DATA_DIR \
  --networkid 89 \
  --port 30303 \
  --rpc \
  --rpccorsdomain "*" \
  --rpcaddr 0.0.0.0 \
  --rpcport 8545 \
  --rpcvhosts "*" \
  --ws \
  --wsaddr 0.0.0.0 \
  --wsport 8546 \
  --wsorigins "*" \
  --mine \
  --gasprice "1" \
  --targetgaslimit "420000000"
