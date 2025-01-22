type EventStore struct {
	db *sql.DB
}

func NewEventStore(dbPath string) (*EventStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
        CREATE TABLE IF NOT EXISTS events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            event_type TEXT NOT NULL,
            model_id INTEGER,
            address TEXT,
            name TEXT,
            version TEXT,
            info TEXT,
            stake INTEGER,
            timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return nil, err
	}

	return &EventStore{db: db}, nil
}

func (s *EventStore) Insert(eventType string, modelId uint64, address string, name string, version string, info string, stake uint64) error {
	_, err := s.db.Exec(`
        INSERT INTO events (event_type, model_id, address, name, version, info, stake)
        VALUES (?, ?, ?, ?, ?, ?, ?)
    `, eventType, modelId, address, name, version, info, stake)
	return err
}

func (s *EventStore) Delete(id int64) error {
	_, err := s.db.Exec("DELETE FROM events WHERE id = ?", id)
	return err
}

func (s *EventStore) Update(id int64, eventType string, modelId uint64, address string, name string, version string, info string, stake uint64) error {
	_, err := s.db.Exec(`
        UPDATE events
        SET event_type = ?, model_id = ?, address = ?, name = ?, version = ?, info = ?, stake = ?
        WHERE id = ?
    `, eventType, modelId, address, name, version, info, stake, id)
	return err
}

func (s *EventStore) GetById(id int64) (*Event, error) {
	row := s.db.QueryRow("SELECT * FROM events WHERE id = ?", id)
	event := &Event{}
	err := row.Scan(&event.Id, &event.EventType, &event.ModelId, &event.Address, &event.Name, &event.Version, &event.Info, &event.Stake, &event.Timestamp)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (s *EventStore) GetAll() ([]*Event, error) {
	rows, err := s.db.Query("SELECT * FROM events")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*Event
	for rows.Next() {
		event := &Event{}
		err := rows.Scan(&event.Id, &event.EventType, &event.ModelId, &event.Address, &event.Name, &event.Version, &event.Info, &event.Stake, &event.Timestamp)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

type Event struct {
	Id        int64
	EventType string
	ModelId   uint64
	Address   string
	Name      string
	Version   string
	Info      string
	Stake     uint64
	Timestamp time.Time
}
