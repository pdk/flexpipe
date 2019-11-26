package flexpipe

// RecordFilter is a function that takes a Record and returns true/false to
// indicate if a certain condition is met.
type RecordFilter func(Record) bool

// RecordTransform is a function that takes a Record as input, and produces a
// potentially new or changed Record as output.
type RecordTransform func(Record) Record

// RecordMultiplier is a function that takes a single Record as input, and
// produces 0, 1 or more Records as output.
type RecordMultiplier func(Record) []Record
