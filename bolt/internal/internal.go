package internal

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/ruprict/wtf"
)

//go:generate protoc --go_out=. internal.proto

func MarshalDial(d *wtf.Dial) ([]byte, error) {
	return proto.Marshal(&Dial{
		ID:      proto.Int64(int64(d.ID)),
		UserID:  proto.Int64(int64(d.UserID)),
		Name:    proto.String(d.Name),
		Level:   proto.Float64(d.Level),
		ModTime: proto.Int64(d.ModTime.UnixNano()),
	})
}

func UnmarshalDial(data []byte, d *wtf.Dial) error {
	var pb Dial

	if err := proto.Unmarshal(data, &pb); err != nil {
		return err
	}

	d.ID = wtf.DialID(pb.GetID())
	d.UserID = wtf.UserID(pb.GetUserID())
	d.Name = pb.GetName()
	d.Level = pb.GetLevel()
	d.ModTime = time.Unix(0, pb.GetModTime()).UTC()

	return nil

}
