package sql

import "uni/dataModels"

type Data struct {
	Students map[string]dataModels.Student
	Masters  map[string]dataModels.Master
	Cuorses  map[string]dataModels.Cuorse
}

var RealData Data

func Init() {

	Students = make(map[string]dataModels.Student)
	Students["5dc57efc-9ff3-48cf-8fe5-9f2e9f0e21da"] = dataModels.Student{
		Id:   "5dc57efc-9ff3-48cf-8fe5-9f2e9f0e21da",
		Name: "hamzah",
	}

	Students["d88f1809-15dc-47ef-8771-d5b7b50a2b67"] = dataModels.Student{
		Id:   "d88f1809-15dc-47ef-8771-d5b7b50a2b67",
		Name: "saeed",
	}

	Masters = make(map[string]dataModels.Master)
	Masters["94c9cdbf-e7d6-4aff-824f-48c57fc46467"] = dataModels.Master{
		Id:   "94c9cdbf-e7d6-4aff-824f-48c57fc46467",
		Name: "khaled",
	}

	Masters["8366fe5a-e57f-4cb3-bb3f-39da1ad6a944"] = dataModels.Master{
		Id:   "8366fe5a-e57f-4cb3-bb3f-39da1ad6a944",
		Name: "nasraddin",
	}

	Cuorses = make(map[string]dataModels.Cuorse)
	Cuorses["090bf6ae-ab44-4448-b99a-a945dbaa404c"] = dataModels.Cuorse{
		Id:   "090bf6ae-ab44-4448-b99a-a945dbaa404c",
		Name: "programming",
	}

	Cuorses["6bd388c7-037a-4c8c-b2dc-b9efb20d26c9"] = dataModels.Cuorse{
		Id:   "6bd388c7-037a-4c8c-b2dc-b9efb20d26c9",
		Name: "english",
	}

	RealData = Data{
		Cuorses:  Cuorses,
		Students: Students,
		Masters:  Masters,
	}

}
