package domain

// SensorType is the dictionary for various sensor types
type SensorType struct {
	Name        string
	Description string
}

func (st SensorType) String() string {
	return st.Name
}
