package model

import "github.com/suyono3484/r3d3"

type Helper struct {
}

type nameFilter struct {
	param string
}

func (n *nameFilter) Param() string {
	return n.param
}

type classFilter struct {
	param string
}

func (c *classFilter) Param() string {
	return c.param
}

type statusFilter struct {
	param string
}

func (s *statusFilter) Param() string {
	return s.param
}

func (h Helper) NewNameFilter(name string) r3d3.ListFilter {
	return &nameFilter{
		param: name,
	}
}

func (h Helper) NewClassFilter(class string) r3d3.ListFilter {
	return &classFilter{
		param: class,
	}
}

func (h Helper) NewStatusFilter(status string) r3d3.ListFilter {
	return &statusFilter{
		param: status,
	}
}
