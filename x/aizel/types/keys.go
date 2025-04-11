package types

const (
	// ModuleName defines the module name
	ModuleName = "aizel"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_aizel"

    
)

var (
	ParamsKey = []byte("p_aizel")
)



func KeyPrefix(p string) []byte {
    return []byte(p)
}
