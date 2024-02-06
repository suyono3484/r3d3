package model

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/suyono3484/r3d3"
)

type Model struct {
	db *sqlx.DB
}

const (
	listQuery     = "SELECT id, name, status FROM spacecraft"
	getQuery      = "SELECT id, name, class crew, image, value, status FROM spacecraft WHERE id = $1"
	armamentQuery = "SELECT a.name, d.qty FROM armament_detail d JOIN armament a ON a.id = d.armament_id JOIN spacecraft s ON s.id = d.spacecraft_id WHERE s.id = $1"
	name          = "name"
	class         = "class"
	status        = "status "
)

type listRow struct {
	IdCraft     int64  `db:"id"`
	NameCraft   string `db:"name"`
	StatusCraft string `db:"status"`
}

type armamentDB struct {
	TitleArmament string `db:"title"`
	QtyArmament   int    `db:"qty"`
}

type spaceRow struct {
	IdCraft       int64           `db:"id"`
	NameCraft     string          `db:"name"`
	StatusCraft   string          `db:"status"`
	ClassCraft    string          `db:"class"`
	CrewCraft     uint64          `db:"crew"`
	ImageURLCraft string          `db:"image"`
	ValueCraft    float64         `db:"value"`
	ArmamentCraft []r3d3.Armament `db:"-"`
}

func NewModel(dsn string) (*Model, error) {
	var (
		err error
	)
	m := &Model{}
	if m.db, err = sqlx.Connect("mysql", dsn); err != nil {
		return nil, err
	}

	return m, nil
}

func (m *Model) Get(id int64) (r3d3.SpaceCraft, error) {
	var (
		r    *sqlx.Row
		err  error
		rows *sqlx.Rows
	)

	result := &spaceRow{}
	r = m.db.QueryRowx(getQuery, id)
	err = r.StructScan(result)
	if err != nil {
		return nil, err
	}

	listArmament := make([]r3d3.Armament, 0)
	if rows, err = m.db.Queryx(armamentQuery, id); err != nil {
		return nil, err
	}

	for rows.Next() {
		v := &armamentDB{}
		if err = rows.StructScan(v); err != nil {
			return nil, err
		}
		listArmament = append(listArmament, v)
	}

	result.ArmamentCraft = listArmament

	return result, nil
}

func (m *Model) List(filter ...r3d3.ListFilter) ([]r3d3.SpaceCraftInList, error) {
	filterMap := make(map[string]map[string]any)
	for _, f := range filter {
		switch v := f.(type) {
		case *nameFilter:
			m, ok := filterMap[name]
			if ok {
				m[v.param] = nil
			} else {
				m = make(map[string]any)
				m[v.param] = nil
			}
			filterMap[name] = m
		case *classFilter:
			m, ok := filterMap[class]
			if ok {
				m[v.param] = nil
			} else {
				m = make(map[string]any)
				m[v.param] = nil
			}
			filterMap[class] = m
		case *statusFilter:
			m, ok := filterMap[status]
			if ok {
				m[v.param] = nil
			} else {
				m = make(map[string]any)
				m[v.param] = nil
			}
			filterMap[status] = m
		}
	}

	q := listQuery
	if len(filterMap) > 0 {
		q += " WHERE"
		fk := true
		for k, v := range filterMap {
			if fk {
				fk = false
			} else {
				q += " AND "
			}
			q += fmt.Sprintf(" %s = ", k)
			if len(v) > 1 {
				first := true
				q += fmt.Sprintf("ANY(")
				for vf := range v {
					if first {
						q += fmt.Sprintf("'%s'", vf)
						first = false
					} else {
						q += fmt.Sprintf(",'%s'", vf)
					}
				}
				q += ")"
			} else {
				for vf := range v {
					q += fmt.Sprintf("'%s'", vf)
				}
			}
		}
	}

	var (
		err error
		r   *sqlx.Rows
	)
	if r, err = m.db.Queryx(q); err != nil {
		return nil, err
	}

	result := make([]r3d3.SpaceCraftInList, 0)
	for r.Next() {
		v := &listRow{}
		if err = r.StructScan(v); err != nil {
			return nil, err
		}
		result = append(result, v)
	}

	return result, nil
}

func (l *listRow) ID() int64 {
	return l.IdCraft
}

func (l *listRow) Name() string {
	return l.NameCraft
}

func (l *listRow) Status() string {
	return l.StatusCraft
}

func (s *spaceRow) ID() int64 {
	return s.IdCraft
}

func (s *spaceRow) Name() string {
	return s.NameCraft
}

func (s *spaceRow) Class() string {
	return s.ClassCraft
}

func (s *spaceRow) Crew() uint64 {
	return s.CrewCraft
}

func (s *spaceRow) ImageURL() string {
	return s.ImageURLCraft
}

func (s *spaceRow) Value() float64 {
	return s.ValueCraft
}

func (s *spaceRow) Status() string {
	return s.StatusCraft
}

func (s *spaceRow) Armament() []r3d3.Armament {
	return s.ArmamentCraft
}

func (a *armamentDB) Title() string {
	return a.TitleArmament
}

func (a *armamentDB) Qty() int {
	return a.QtyArmament
}
