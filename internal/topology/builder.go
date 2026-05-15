
package topology

import "log-parser/internal/model"

func Build(links []model.Link) model.Topology {
	seen := make(map[string]bool)

	var nodes []model.TopologyNode

	for _, link := range links {
		if !seen[link.Switch] {
			nodes = append(nodes, model.TopologyNode{ID: link.Switch})
			seen[link.Switch] = true
		}

		if !seen[link.Peer] {
			nodes = append(nodes, model.TopologyNode{ID: link.Peer})
			seen[link.Peer] = true
		}
	}

	return model.Topology{
		Nodes: nodes,
		Links: links,
	}
}