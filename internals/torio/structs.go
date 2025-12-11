// Package torio
package torio

const (
	P_M = iota
	P_B
)

type room struct {
	resources []resource
}
type player struct {
	resources []resource
	speedMult int
	state     int
}

type resource struct {
	name      string
	amount    int
	regenRate int
}
