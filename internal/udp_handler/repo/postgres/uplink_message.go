package postgres

import (
	"udp_handler/pkg/models"
)

//Uplink use from models
type UplinkMessage models.UplinkMessage

//BreedService use from models
type UplinkMessageService models.UplinkMessageService

//CreateUplink create a breed
func UpdateDevice(message UplinkMessage) error {

	// TO DO
	// fmt.Println("In Repo")

	// stmt, err := client.DbClient.Prepare("INSERT INTO public.Uplink (breed_name,category_id) VALUES ($1,$2);")
	// if err != nil {
	// 	return err
	// }
	// //closing the statement to prevent memory leaks
	// defer stmt.Close()
	// _, err = stmt.Exec(b.BreedName, b.CategoryID)
	// if err != nil {
	// 	return err
	// }
	// return nil
	return nil
}
