## scalard query ibc channel packet-receipt

Query a packet receipt

### Synopsis

Query a packet receipt

```
scalard query ibc channel packet-receipt [port-id] [channel-id] [sequence] [flags]
```

### Examples

```
<appd> query ibc channel packet-receipt [port-id] [channel-id] [sequence]
```

### Options

```
      --height int      Use a specific height to query state at (this can error if the node is pruning state)
  -h, --help            help for packet-receipt
      --node string     <host>:<port> to Tendermint RPC interface for this chain (default "tcp://localhost:26657")
  -o, --output string   Output format (text|json) (default "text")
      --prove           show proofs for the query results (default true)
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

- [scalard query ibc channel](scalard_query_ibc_channel.md) - IBC channel query subcommands
