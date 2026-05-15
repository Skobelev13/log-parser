package model

type TopologyNode struct {
	ID string `json:"id"`
}

type Topology struct {
	Nodes []TopologyNode `json:"nodes"`
	Links []Link         `json:"links"`
}
