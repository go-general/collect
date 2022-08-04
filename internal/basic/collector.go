package basic

// Collector is the basic ability of all collects, includes list & set & map.
type Collector interface {
	// IsEmpty checks if collector has no values.
	// return true if empty, false if not.
	IsEmpty() bool

	// Size gets num of values in collector.
	Size() int

	// Clear clears all values in collect, but maybe do nothing with immutable collector.
	Clear()
}
