## scalard query staking delegation

Query a delegation based on address and validator address

### Synopsis

Query delegations for an individual delegator on an individual validator.

Example:
$ <appd> query staking delegation scalar1gghjut3ccd8ay0zduzj64hwre2fxs9ld75ru9p scalarvaloper1gghjut3ccd8ay0zduzj64hwre2fxs9ldmqhffj

```
scalard query staking delegation [delegator-addr] [validator-addr] [flags]
```

### Options

```
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for delegation
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")
```

### Options inherited from parent commands

```
      --chain-id string     The network chain ID (default "scalar")
      --home string         directory for config and data (default "$HOME/.scalar")
      --log_format string   The logging format (json|plain) (default "plain")
      --log_level string    The logging level (trace|debug|info|warn|error|fatal|panic) (default "info")
      --trace               print out full stack trace on errors
```

### SEE ALSO

- [scalard query staking](scalard_query_staking.md) - Querying commands for the staking module