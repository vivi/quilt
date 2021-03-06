package db

// The Etcd table contains configuration pertaining to the minion etcd cluster including
// the members and leadership information.
type Etcd struct {
	ID int

	EtcdIPs []string // The set of members in the cluster.

	Leader   bool   // True if this Minion is the leader.
	LeaderIP string // IP address of the current leader, or ""
}

func (e Etcd) String() string {
	return defaultString(e)
}

func (e Etcd) getID() int {
	return e.ID
}

// InsertEtcd creates a new etcd row and inserts it into the database.
func (db Database) InsertEtcd() Etcd {
	result := Etcd{ID: db.nextID()}
	db.insert(result)
	return result
}

// EtcdLeader returns true if the minion is the lead master for the cluster.
func (db Database) EtcdLeader() bool {
	etcds := db.SelectFromEtcd(nil)
	return len(etcds) == 1 && etcds[0].Leader
}

// SelectFromEtcd gets all Etcd rows in the database that satisfy the 'check'.
func (db Database) SelectFromEtcd(check func(Etcd) bool) []Etcd {
	result := []Etcd{}
	for _, row := range db.tables[EtcdTable].rows {
		if check == nil || check(row.(Etcd)) {
			result = append(result, row.(Etcd))
		}
	}
	return result
}

// EtcdLeader returns true if the minion is the lead master for the cluster.
func (conn Conn) EtcdLeader() bool {
	var leader bool
	conn.Transact(func(view Database) error {
		leader = view.EtcdLeader()
		return nil
	})
	return leader
}

// SelectFromEtcd gets all Etcd rows in the database connection that satisfy the
// 'check'.
func (conn Conn) SelectFromEtcd(check func(Etcd) bool) []Etcd {
	var etcdRows []Etcd
	conn.Transact(func(view Database) error {
		etcdRows = view.SelectFromEtcd(check)
		return nil
	})
	return etcdRows
}

func (e Etcd) less(r row) bool {
	return e.ID < r.(Minion).ID
}
