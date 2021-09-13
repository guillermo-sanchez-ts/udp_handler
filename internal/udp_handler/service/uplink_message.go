package service

import (
	"fmt"
	db "udp_handler/internal/udp_handler/repo/postgres"
)

var err error
var bs db.UplinkMessageService
var UplinkMessage db.UplinkMessage

//CreateUplinkMessage create a UplinkMessage
func UpdateDevice(b *db.UplinkMessage) error {
	//TODO
	return nil
}

func dummy(b *db.UplinkMessage) error {
	fmt.Println("In service")
	bs = b
	err = bs.CreateUplinkMessage()
	if err != nil {
		return err
	}
	return nil
}
