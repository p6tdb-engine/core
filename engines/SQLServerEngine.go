package engines

import "github.com/lhsribas/P6TDB/model"

type SQLServer struct{}

func (s *SQLServer) Engine(db model.Yaml2Go) {}
