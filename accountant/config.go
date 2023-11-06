package accountant

// Config contains configuration for the AccountingBook.
type Config struct {
	TrustedNodesDBPath       string `yaml:"trusted_nodes_db_path"`
	TokensDBPath             string `yaml:"tokens_db_path"`
	TraxsToVerticesMapDBPath string `yaml:"trxs_to_vertices_map_db_path"`
	VerticesDBPath           string `yaml:"vertices_db_path"`
	LoadDAG                  bool   `yaml:"load_dag"`
}
