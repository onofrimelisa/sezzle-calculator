// backend/internal/domain/registry.go
package domain

type RegistryEntry struct {
	Type        OperationType
	Name        string
	Description string
	Category    string
	Binary      BinaryOperation
	Unary       UnaryOperation
}

type Registry struct {
	operations map[string]RegistryEntry
}

func NewRegistry() *Registry {
	return &Registry{operations: make(map[string]RegistryEntry)}
}

func (r *Registry) RegisterBinary(op BinaryOperation) {
	r.operations[op.Name()] = RegistryEntry{
		Type: Binary, Name: op.Name(), Description: op.Description(),
		Category: op.Category(), Binary: op,
	}
}

func (r *Registry) RegisterUnary(op UnaryOperation) {
	r.operations[op.Name()] = RegistryEntry{
		Type: Unary, Name: op.Name(), Description: op.Description(),
		Category: op.Category(), Unary: op,
	}
}

func (r *Registry) Get(name string) (RegistryEntry, error) {
	entry, ok := r.operations[name]
	if !ok {
		return RegistryEntry{}, ErrInvalidOperation
	}
	return entry, nil
}

func (r *Registry) List() []RegistryEntry {
	entries := make([]RegistryEntry, 0, len(r.operations))
	for _, entry := range r.operations {
		entries = append(entries, entry)
	}
	return entries
}
