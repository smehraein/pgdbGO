package pgdbGO

import "testing"

func TestConnect(t *testing.T) {
	cases := []struct {
		in PGConnection
	}{
		{PGConnection{"postgres", "", "postgres", "localhost"}},
	}
	for _, c := range cases {
		conn, _ := Connect(c.in)
		err := conn.Ping()
		if err != nil {
			t.Errorf("Error connecting to database: %q", err)
		} else {
			conn.Close()
		}
	}
}
