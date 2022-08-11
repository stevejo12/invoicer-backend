package controller

import (
	"encoding/json"
	"invoicer/main/config"
	"invoicer/main/models"
	"net/http"
)

func GetPhoneCodeList(w http.ResponseWriter, r *http.Request) {
	var allPhoneCodeList []interface{}
	var response models.JsonResponse

	data, err := config.DB.Query("SELECT * FROM phonecodelist")

	if err != nil {
		w.WriteHeader(500)
		response = models.JsonResponse{
			Code:    500,
			Data:    allPhoneCodeList,
			Message: "Something wrong with the database query to retrieve data",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	defer data.Close()

	for data.Next() {
		var phoneCodeData models.PhoneCodeList
		if err := data.Scan(&phoneCodeData.ID, &phoneCodeData.Name, &phoneCodeData.PhoneCode); err != nil {
			w.WriteHeader(500)
			response = models.JsonResponse{
				Code:    500,
				Data:    allPhoneCodeList,
				Message: "Something wrong while scanning data",
			}
			json.NewEncoder(w).Encode(response)
			return
		}

		allPhoneCodeList = append(allPhoneCodeList, phoneCodeData)
	}

	if data.Err() != nil {
		w.WriteHeader(500)
		response = models.JsonResponse{
			Code:    500,
			Data:    allPhoneCodeList,
			Message: "Something wrong with the data retrieved from database",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response = models.JsonResponse{
		Code: 200,
		Data: allPhoneCodeList,
	}

	json.NewEncoder(w).Encode(response)
}
