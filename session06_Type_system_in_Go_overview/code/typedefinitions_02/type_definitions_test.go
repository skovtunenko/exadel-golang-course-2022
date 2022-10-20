package typedefinitions_02

type (
	Point struct{ x, y float64 } // Point and struct{ x, y float64 } are different types
	polar Point                  // polar and Point denote different types!!
)

// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct { /* Mutex fields */
}

func (m *Mutex) Lock()   { /* Lock implementation */ }
func (m *Mutex) Unlock() { /* Unlock implementation */ }

// NewMutex has the same composition as Mutex but its method set is empty.
type NewMutex Mutex

// The method set of PtrMutex's underlying type *Mutex remains unchanged,
// but the method set of PtrMutex is empty.
type PtrMutex *Mutex

// The method set of *PrintableMutex contains the methods
// Lock and Unlock bound to its embedded field Mutex.
type PrintableMutex struct {
	Mutex
}

type Block interface {
	BlockSize() int
	Encrypt(src, dst []byte)
	Decrypt(src, dst []byte)
}

// MyBlock is an interface type that has the same method set as Block.
type MyBlock Block
