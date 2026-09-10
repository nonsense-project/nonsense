package flowcontext

import (
	"github.com/nonsense-project/nonsense/v2/domain"
)

// Domain returns the Domain object associated to the flow context.
func (f *FlowContext) Domain() domain.Domain {
	return f.domain
}
