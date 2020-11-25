package slack

import "flag"

var port = flag.Int("port", 5001, "Server port")

func init() {
	flag.Parse()
}

type server struct {
}
