package mms

import (
	"encoding/json"
	"net/http"
	"project/config"
	"project/internal/alpha2"
	"project/internal/model"
	"project/pkg/logging"
	"sort"
)

func CheckMMSInfo(cfg *config.Config, logger *logging.Logger) ([]model.MMSDataModel, error) {

	var mmsSliceSum []model.MMSDataModel
	var mmsSliceRes []model.MMSDataModel

	codeA2, err := alpha2.CountryCodeAlpha2(cfg)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	prov, err := alpha2.GetProviders(cfg, "sms")
	if err != nil {
		logger.Errorf("error read file providers: %v", err)
	}

	resp, err := http.Get(cfg.MMSHost + ":" + cfg.MMSPort)
	if err != nil {
		return nil, err
	}

	logger.Infof("Status code MMS: %v", resp.StatusCode)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {

		err = json.NewDecoder(resp.Body).Decode(&mmsSliceSum)
		if err != nil {
			return nil, err
		}

		for _, m := range mmsSliceSum {
			for _, p := range prov {
				if m.Provider == p {
					for _, c := range codeA2 {
						if m.Country == c.Alpha2 {
							m.Country = c.Country
							mmsSliceRes = append(mmsSliceRes, m)
						}
					}
				}
			}
		}

		sort.SliceStable(mmsSliceRes, func(i, j int) bool {
			return mmsSliceRes[i].Provider < mmsSliceRes[j].Provider
		})
	}

	return mmsSliceRes, nil

}

func SortMMSInfo(mmsInfo []model.MMSDataModel) ([]model.MMSDataModel, error) {

	var first []model.MMSDataModel
	for _, v := range mmsInfo {
		first = append(first, v)
	}

	sort.SliceStable(first, func(i, j int) bool {
		return first[i].Country < first[j].Country
	})

	return first, nil
}
