package db

func (m MongoDB) InOne(payload interface{}) error {
	ct := m.MakeContext(10)
	_, err := col.InsertOne(ct, payload)
	return err
}

func (m MongoDB) InMany(payload []interface{}) error {
	ct := m.MakeContext(25)
	_, err := col.InsertMany(ct, payload)
	return err
}
