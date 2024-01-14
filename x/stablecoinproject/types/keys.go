package types

const (
	// ModuleName defines the module name
	ModuleName = "stablecoinproject"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stablecoinproject"

    
)

var (
	ParamsKey = []byte("p_stablecoinproject")
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
