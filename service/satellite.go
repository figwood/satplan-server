package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"satplan/common"
	"satplan/dao/db"
	"satplan/entity"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

func AddSatellite(newSat *entity.NewSatDTO) (int, error) {
	satName := strings.TrimSpace(newSat.Name)
	tleDetails := strings.Split(newSat.Tle, "\n")
	if len(tleDetails) < 3 {
		return 0, errors.New("bad format of tle")
	}
	tle := entity.TleData{
		Line0: tleDetails[0],
		Line1: tleDetails[1],
		Line2: tleDetails[2],
	}

	if len(satName) == 0 {
		satName = strings.TrimSpace(tle.Line0)
	}

	line1Details := strings.Split(tle.Line1, " ")
	noardId := line1Details[1]

	newSatToDB := entity.Satellite{
		Name:     satName,
		NoardId:  noardId,
		HexColor: getRandomColor(),
	}
	err := db.CreateSatellite(&newSatToDB)

	if err != nil {
		return 0, err
	}
	err = db.CreateTle(&entity.Tle{
		SatNoardId: noardId,
		Time:       common.GetUtcNowTimeStampSec(),
		Line1:      tle.Line1,
		Line2:      tle.Line2,
	})

	return newSatToDB.Id, err
}

func getRandomColor() string {
	min := 0
	max := 255
	r := rand.Intn(max-min) + min
	g := rand.Intn(max-min) + min
	b := rand.Intn(max-min) + min
	return fmt.Sprintf("#%x%x%x", r, g, b)
}

func GetSatTree() *[]entity.SatItem {
	allSats := GetAllSatellites()
	satTree := []entity.SatItem{}
	for _, s := range *allSats {
		if s.SenItems != nil && len(*s.SenItems) != 0 {
			satTree = append(satTree, s)
		}
	}
	return &satTree
}

func GetAllSatellites() *[]entity.SatItem {
	satellites := db.FindAllSatellites()
	mapSat := map[string]entity.SatItem{}
	for _, sat := range *satellites {
		mapSat[sat.NoardId] = entity.SatItem{
			Id:       sat.Id,
			Name:     sat.Name,
			NoardId:  sat.NoardId,
			HexColor: sat.HexColor,
			SenItems: &[]entity.SenItem{},
		}
	}
	sensors := db.FindAllSensors()

	satItems := []entity.SatItem{}

	for _, sen := range *sensors {
		if _, ok := mapSat[sen.SatNoardId]; !ok {
			continue
		}

		*(mapSat[sen.SatNoardId].SenItems) = append(*(mapSat[sen.SatNoardId].SenItems),
			entity.SenItem{
				Id:             sen.Id,
				Name:           sen.Name,
				Resolution:     sen.Resolution,
				Width:          sen.Width,
				RightSideAngle: sen.RightSideAngle,
				LeftSideAngle:  sen.LeftSideAngle,
				ObserveAngle:   sen.ObserveAngle,
				InitAngle:      sen.InitAngle,
				HexColor:       sen.HexColor,
			})
	}
	for _, m := range mapSat {
		satItems = append(satItems, m)
	}
	//sort by sat name asc
	sort.Slice(satItems, func(i, j int) bool {
		return strings.Compare(satItems[i].Name, satItems[j].Name) <= 0
	})
	return &satItems
}

func GetSatelliteById(satId string) (*entity.Satellite, error) {
	return db.FindSatelliteByNoardId(satId)
}

func UpdateSatellite(id int, satDTO *entity.SatDTO) error {
	satInDB, err := db.FindSatelliteById(id)
	if err != nil || satInDB.Id == 0 {
		return errors.New("error finding satellite")
	}
	satInDB.Name = satDTO.SatName
	satInDB.HexColor = satDTO.HexColor
	return db.SaveSatellite(satInDB)
}

func DeleteSatelliteById(satId int) error {
	return db.DeleteSatelliteById(satId)
}

func UpdateTles() error {
	tleTxt, err := getNewTles()
	if err != nil {
		return err
	}
	tleDetails := strings.Split(tleTxt, "\n")
	tles := []entity.Tle{}
	for i := 0; i < len(tleDetails)-3; i += 3 {
		//get noard id
		line1Details := strings.Split(tleDetails[i+1], " ")
		noardId := line1Details[1]
		sat, _ := db.FindSatelliteByNoardId(noardId)
		if sat == nil || sat.Id == 0 {
			db.CreateSatellite(&entity.Satellite{
				Name:     strings.TrimSpace(tleDetails[i]),
				NoardId:  noardId,
				HexColor: getRandomColor(),
			})
		}
		tles = append(tles, entity.Tle{
			SatNoardId: noardId,
			Time:       common.GetUtcNowTimeStampSec(),
			Line1:      tleDetails[i+1],
			Line2:      tleDetails[i+2],
		})
	}
	db.DeleteTles()
	db.BatCreateTle(&tles)

	//recaculate
	return RecalPath()
}

func getNewTles() (string, error) {
	url := "http://celestrak.com/NORAD/elements/resource.txt"
	/*
		SCD 1
		1 22490U 93009B   21202.41204679  .00000203  00000-0  53852-5 0  9991
		2 22490  24.9705  95.0049 0042752 280.5291 240.5150 14.44613837501644
		TECHSAT 1B (GO-32)
		1 25397U 98043D   21202.84242513 -.00000023  00000-0  95154-5 0  9991
		2 25397  98.7723 148.3036 0001524  31.1555 328.9713 14.23711884196605
	*/

	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(body), nil
}

func RecalPath() error {
	curPath, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}
	dataFolder := common.GetEnvValue("DATA_FOLDER", "../data")
	output, err := exec.Command(fmt.Sprintf("%s/%s", curPath, "calpath"), dataFolder).Output()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	log.Info(output)
	return nil
}
