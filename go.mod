module go-cache-benchmark

go 1.16

require (
	github.com/dgraph-io/ristretto v0.0.3
	github.com/dgryski/go-clockpro v0.0.0-20140817124034-edc6d3eeb96e
	github.com/dgryski/go-tinylfu v0.0.0-20210109191853-fba88f4a7f91
	github.com/hashicorp/golang-lru v0.5.1
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pingcap/go-ycsb v0.0.0-20210129115622-04d8656123e4
	github.com/uptrace/go-uptrace v0.0.0-00010101000000-000000000000
)

replace github.com/uptrace/go-uptrace => /home/vmihailenco/Dropbox/box_workspace/beat/go-uptrace
