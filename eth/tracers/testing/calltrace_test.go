package testing

import (
	"encoding/json"
	"math/big"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"unicode"

	"github.com/XinFinOrg/XDPoSChain/common"
	"github.com/XinFinOrg/XDPoSChain/common/hexutil"
	"github.com/XinFinOrg/XDPoSChain/common/math"
	"github.com/XinFinOrg/XDPoSChain/core"
	"github.com/XinFinOrg/XDPoSChain/core/rawdb"
	"github.com/XinFinOrg/XDPoSChain/core/types"
	"github.com/XinFinOrg/XDPoSChain/core/vm"
	"github.com/XinFinOrg/XDPoSChain/eth/tracers"
	"github.com/XinFinOrg/XDPoSChain/rlp"
	"github.com/XinFinOrg/XDPoSChain/tests"

	// Force-load the native, to trigger registration
	_ "github.com/XinFinOrg/XDPoSChain/eth/tracers/native"
)

type callContext struct {
	Number     math.HexOrDecimal64   `json:"number"`
	Difficulty *math.HexOrDecimal256 `json:"difficulty"`
	Time       math.HexOrDecimal64   `json:"timestamp"`
	GasLimit   math.HexOrDecimal64   `json:"gasLimit"`
	Miner      common.Address        `json:"miner"`
}

// callTrace is the result of a callTracer run.
type callTrace struct {
	Type    string          `json:"type"`
	From    common.Address  `json:"from"`
	To      common.Address  `json:"to"`
	Input   hexutil.Bytes   `json:"input"`
	Output  hexutil.Bytes   `json:"output"`
	Gas     *hexutil.Uint64 `json:"gas,omitempty"`
	GasUsed *hexutil.Uint64 `json:"gasUsed,omitempty"`
	Value   *hexutil.Big    `json:"value,omitempty"`
	Error   string          `json:"error,omitempty"`
	Calls   []callTrace     `json:"calls,omitempty"`
}

// callTracerTest defines a single test to check the call tracer against.
type callTracerTest struct {
	Genesis      *core.Genesis   `json:"genesis"`
	Context      *callContext    `json:"context"`
	Input        string          `json:"input"`
	TracerConfig json.RawMessage `json:"tracerConfig"`
	Result       *callTrace      `json:"result"`
}

// Iterates over all the input-output datasets in the tracer test harness and
// runs the JavaScript tracers against them.
func TestCallTracerLegacy(t *testing.T) {
	testCallTracer("callTracerLegacy", "call_tracer_legacy", t)
}

func TestCallTracer(t *testing.T) {
	testCallTracer("callTracer", "call_tracer", t)
}

func testCallTracer(tracerName string, dirPath string, t *testing.T) {
	files, err := os.ReadDir(filepath.Join("..", "testdata", dirPath))
	if err != nil {
		t.Fatalf("failed to retrieve tracer test suite: %v", err)
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		file := file // capture range variable
		t.Run(camel(strings.TrimSuffix(file.Name(), ".json")), func(t *testing.T) {
			t.Parallel()

			var (
				test = new(callTracerTest)
				tx   = new(types.Transaction)
			)
			// Call tracer test found, read if from disk
			if blob, err := os.ReadFile(filepath.Join("..", "testdata", dirPath, file.Name())); err != nil {
				t.Fatalf("failed to read testcase: %v", err)
			} else if err := json.Unmarshal(blob, test); err != nil {
				t.Fatalf("failed to parse testcase: %v", err)
			}
			if err := rlp.DecodeBytes(common.FromHex(test.Input), tx); err != nil {
				t.Fatalf("failed to parse testcase input: %v", err)
			}
			// Configure a blockchain with the given prestate
			var (
				signer    = types.MakeSigner(test.Genesis.Config, new(big.Int).SetUint64(uint64(test.Context.Number)))
				origin, _ = signer.Sender(tx)
				txContext = vm.TxContext{
					Origin:   origin,
					GasPrice: tx.GasPrice(),
				}
				context = vm.BlockContext{
					CanTransfer: core.CanTransfer,
					Transfer:    core.Transfer,
					Coinbase:    test.Context.Miner,
					BlockNumber: new(big.Int).SetUint64(uint64(test.Context.Number)),
					Time:        new(big.Int).SetUint64(uint64(test.Context.Time)),
					Difficulty:  (*big.Int)(test.Context.Difficulty),
					GasLimit:    uint64(test.Context.GasLimit),
				}
				statedb = tests.MakePreState(rawdb.NewMemoryDatabase(), test.Genesis.Alloc)
			)
			tracer, err := tracers.New(tracerName, new(tracers.Context), test.TracerConfig)
			if err != nil {
				t.Fatalf("failed to create call tracer: %v", err)
			}
			evm := vm.NewEVM(context, txContext, statedb, nil, test.Genesis.Config, vm.Config{Tracer: tracer})
			msg, err := tx.AsMessage(signer, nil, nil, nil)
			if err != nil {
				t.Fatalf("failed to prepare transaction for tracing: %v", err)
			}
			st := core.NewStateTransition(evm, msg, new(core.GasPool).AddGas(tx.Gas()))
			if _, err = st.TransitionDb(common.Address{}); err != nil {
				t.Fatalf("failed to execute transaction: %v", err)
			}
			// Retrieve the trace result and compare against the etalon
			res, err := tracer.GetResult()
			if err != nil {
				t.Fatalf("failed to retrieve trace result: %v", err)
			}
			ret := new(callTrace)
			if err := json.Unmarshal(res, ret); err != nil {
				t.Fatalf("failed to unmarshal trace result: %v", err)
			}

			if !jsonEqual(ret, test.Result) {
				// uncomment this for easier debugging
				//have, _ := json.MarshalIndent(ret, "", " ")
				//want, _ := json.MarshalIndent(test.Result, "", " ")
				//t.Fatalf("trace mismatch: \nhave %+v\nwant %+v", string(have), string(want))
				t.Fatalf("trace mismatch: \nhave %+v\nwant %+v", ret, test.Result)
			}
		})
	}
}

// jsonEqual is similar to reflect.DeepEqual, but does a 'bounce' via json prior to
// comparison
func jsonEqual(x, y interface{}) bool {
	xTrace := new(callTrace)
	yTrace := new(callTrace)
	if xj, err := json.Marshal(x); err == nil {
		json.Unmarshal(xj, xTrace)
	} else {
		return false
	}
	if yj, err := json.Marshal(y); err == nil {
		json.Unmarshal(yj, yTrace)
	} else {
		return false
	}
	return reflect.DeepEqual(xTrace, yTrace)
}

// camel converts a snake cased input string into a camel cased output.
func camel(str string) string {
	pieces := strings.Split(str, "_")
	for i := 1; i < len(pieces); i++ {
		pieces[i] = string(unicode.ToUpper(rune(pieces[i][0]))) + pieces[i][1:]
	}
	return strings.Join(pieces, "")
}
func BenchmarkTracers(b *testing.B) {
	files, err := os.ReadDir(filepath.Join("..", "testdata", "call_tracer"))
	if err != nil {
		b.Fatalf("failed to retrieve tracer test suite: %v", err)
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		file := file // capture range variable
		b.Run(camel(strings.TrimSuffix(file.Name(), ".json")), func(b *testing.B) {
			blob, err := os.ReadFile(filepath.Join("..", "testdata", "call_tracer", file.Name()))
			if err != nil {
				b.Fatalf("failed to read testcase: %v", err)
			}
			test := new(callTracerTest)
			if err := json.Unmarshal(blob, test); err != nil {
				b.Fatalf("failed to parse testcase: %v", err)
			}
			benchTracer("callTracer", test, b)
		})
	}
}

func benchTracer(tracerName string, test *callTracerTest, b *testing.B) {
	// Configure a blockchain with the given prestate
	tx := new(types.Transaction)
	if err := rlp.DecodeBytes(common.FromHex(test.Input), tx); err != nil {
		b.Fatalf("failed to parse testcase input: %v", err)
	}
	signer := types.MakeSigner(test.Genesis.Config, new(big.Int).SetUint64(uint64(test.Context.Number)))
	msg, err := tx.AsMessage(signer, nil, nil, nil)
	if err != nil {
		b.Fatalf("failed to prepare transaction for tracing: %v", err)
	}
	origin, _ := signer.Sender(tx)
	txContext := vm.TxContext{
		Origin:   origin,
		GasPrice: tx.GasPrice(),
	}
	context := vm.BlockContext{
		CanTransfer: core.CanTransfer,
		Transfer:    core.Transfer,
		Coinbase:    test.Context.Miner,
		BlockNumber: new(big.Int).SetUint64(uint64(test.Context.Number)),
		Time:        new(big.Int).SetUint64(uint64(test.Context.Time)),
		Difficulty:  (*big.Int)(test.Context.Difficulty),
		GasLimit:    uint64(test.Context.GasLimit),
	}
	statedb := tests.MakePreState(rawdb.NewMemoryDatabase(), test.Genesis.Alloc)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		tracer, err := tracers.New(tracerName, new(tracers.Context), nil)
		if err != nil {
			b.Fatalf("failed to create call tracer: %v", err)
		}
		evm := vm.NewEVM(context, txContext, statedb, nil, test.Genesis.Config, vm.Config{Tracer: tracer})
		snap := statedb.Snapshot()
		st := core.NewStateTransition(evm, msg, new(core.GasPool).AddGas(tx.Gas()))
		if _, err = st.TransitionDb(common.Address{}); err != nil {
			b.Fatalf("failed to execute transaction: %v", err)
		}
		if _, err = tracer.GetResult(); err != nil {
			b.Fatal(err)
		}
		statedb.RevertToSnapshot(snap)
	}
}

type contractTracerTest struct {
	Genesis      *core.Genesis   `json:"genesis"`
	Context      *callContext    `json:"context"`
	Input        string          `json:"input"`
	TracerConfig json.RawMessage `json:"tracerConfig"`
	Result       []string        `json:"result"`
}

func testContractTracer(tracerName string, dirPath string, t *testing.T) {
	files, err := os.ReadDir(filepath.Join("..", "testdata", dirPath))
	if err != nil {
		t.Fatalf("failed to retrieve tracer test suite: %v", err)
	}
	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".json") {
			continue
		}
		file := file // capture range variable
		t.Run(camel(strings.TrimSuffix(file.Name(), ".json")), func(t *testing.T) {
			t.Parallel()

			var (
				test = new(contractTracerTest)
				tx   = new(types.Transaction)
			)
			// Call tracer test found, read if from disk
			if blob, err := os.ReadFile(filepath.Join("..", "testdata", dirPath, file.Name())); err != nil {
				t.Fatalf("failed to read testcase: %v", err)
			} else if err := json.Unmarshal(blob, test); err != nil {
				t.Fatalf("failed to parse testcase: %v", err)
			}
			if err := rlp.DecodeBytes(common.FromHex(test.Input), tx); err != nil {
				t.Fatalf("failed to parse testcase input: %v", err)
			}
			// Configure a blockchain with the given prestate
			var (
				signer    = types.MakeSigner(test.Genesis.Config, new(big.Int).SetUint64(uint64(test.Context.Number)))
				origin, _ = signer.Sender(tx)
				txContext = vm.TxContext{
					Origin:   origin,
					GasPrice: tx.GasPrice(),
				}
				context = vm.BlockContext{
					CanTransfer: core.CanTransfer,
					Transfer:    core.Transfer,
					Coinbase:    test.Context.Miner,
					BlockNumber: new(big.Int).SetUint64(uint64(test.Context.Number)),
					Time:        new(big.Int).SetUint64(uint64(test.Context.Time)),
					Difficulty:  (*big.Int)(test.Context.Difficulty),
					GasLimit:    uint64(test.Context.GasLimit),
				}
				statedb = tests.MakePreState(rawdb.NewMemoryDatabase(), test.Genesis.Alloc)
			)
			tracer, err := tracers.New(tracerName, new(tracers.Context), test.TracerConfig)
			if err != nil {
				t.Fatalf("failed to create call tracer: %v", err)
			}
			evm := vm.NewEVM(context, txContext, statedb, nil, test.Genesis.Config, vm.Config{Tracer: tracer})
			msg, err := tx.AsMessage(signer, nil, nil, nil)
			if err != nil {
				t.Fatalf("failed to prepare transaction for tracing: %v", err)
			}
			st := core.NewStateTransition(evm, msg, new(core.GasPool).AddGas(tx.Gas()))
			if _, err = st.TransitionDb(common.Address{}); err != nil {
				t.Fatalf("failed to execute transaction: %v", err)
			}
			// Retrieve the trace result and compare against the etalon
			res, err := tracer.GetResult()
			if err != nil {
				t.Fatalf("failed to retrieve trace result: %v", err)
			}
			ret := new([]string)
			if err := json.Unmarshal(res, ret); err != nil {
				t.Fatalf("failed to unmarshal trace result: %v", err)
			}

			if !reflect.DeepEqual(*ret, test.Result) {
				// uncomment this for easier debugging
				//have, _ := json.MarshalIndent(ret, "", " ")
				//want, _ := json.MarshalIndent(test.Result, "", " ")
				//t.Fatalf("trace mismatch: \nhave %+v\nwant %+v", string(have), string(want))
				t.Fatalf("trace mismatch: \nhave %+v\nwant %+v", *ret, test.Result)
			}
		})
	}
}

func TestContractTracer(t *testing.T) {
	testContractTracer("contractTracer", "contract_tracer", t)
}
