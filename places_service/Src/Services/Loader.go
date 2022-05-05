package Services

import "mpasGo/Database"

var GoogleServiceInstance = NewGoogleMaps(ProviderConfig{
	Key: "AIzaSyB7QSklJ7Be2n0DjtsiIukAeQu4922KfeI",
})

var PhotonServiceInstance = NewPhoton(PhotonConfig{
	BaseUrl: "http://localhost:2322",
})

var MongoServiceInstance = NewMongo(MongoConfig{
	RepositoryInstance: Database.RepositoryInstance,
})
