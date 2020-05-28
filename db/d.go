package db

func (m MongoDB) DelOne(payload interface{}) error {
	ct := m.MakeContext(10)
	_, err := col.DeleteOne(ct, payload)
	return err
}
