package keys_test

import (
	"github.com/boostxinnovation/boostx-partner-sdk-go/boostx/client"
	"github.com/boostxinnovation/boostx-partner-sdk-go/boostx/handlers"
	"github.com/boostxinnovation/boostx-partner-sdk-go/boostx/keys"
)

// Compile-time proof MemoryKeyStore satisfies both key-store interfaces. These
// live in an external test package because handlers imports keys (for
// ErrUnknownPartner), so an internal-test import of handlers would cycle.
var (
	_ handlers.KeyStore = (*keys.MemoryKeyStore)(nil)
	_ client.KeyStore   = (*keys.MemoryKeyStore)(nil)
)
