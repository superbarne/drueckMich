package bookmark;

import (
	"path"
	"github.com/watson-developer-cloud/go-sdk/core"
	"github.com/watson-developer-cloud/go-sdk/visualrecognitionv3"
	"github.com/superbarne/drueckMich/api/categoryWvr"
	"gopkg.in/mgo.v2/bson"
)

func AnalyzeImage(url string, bookmark *Bookmark) {
	// Watson Visual Recognition Service instanziieren:
	service, serviceErr := visualrecognitionv3.
		NewVisualRecognitionV3(&visualrecognitionv3.VisualRecognitionV3Options{
			URL:       "https://gateway.watsonplatform.net/visual-recognition/api",
			Version:   "2018-03-19",
			IAMApiKey: "uPQg9r6Om9LxghWPaU9mGhAJABhMZQSlnuRGH_kABUdm", // !!!!!!!!!HIER DEN EIGENEN API-KEY EINTRAGEN!!!!!!!!!!!!!
		})
	if serviceErr != nil {
		panic(serviceErr)
	}

	// Image klassifizieren:

	//--------------------------------------------------------------------------
	// A) Image per URL definieren:
	//--------------------------------------------------------------------------

	// Optionen für die Klassifizierung festlegen:
	classifyOptions := service.NewClassifyOptions()
	classifyOptions.URL = core.StringPtr(url)

	// Schwellwert für den "Verlässlichkeitsscore":
	classifyOptions.Threshold = core.Float32Ptr(0.6)

	classifyOptions.ClassifierIds = []string{"default"}
	//	classifyOptions.ClassifierIds = []string{"default", "food", "explicit"}

	// Ausgabesprache definieren:
	sprache := new(string)
	*sprache = "de"
	classifyOptions.AcceptLanguage = sprache

	// Classify Dienst aufrufen:
	response, responseErr := service.Classify(classifyOptions)
	if responseErr != nil {
		panic(responseErr)
	}

	// Ergebnisdaten aufbereiten:
	classifyResult := service.GetClassifyResult(response)
	classes := classifyResult.Images[0].Classifiers[0].Classes
	imageName := *classifyResult.Images[0].ResolvedURL
	_, imageName = path.Split(imageName) // path.Split NICHT string.Split !!!!!
	classes = classifyResult.Images[0].Classifiers[0].Classes
	for _, wert := range classes {
		categories, _ := categoryWvr.Find(bson.M{"name": wert.ClassName },"createdAt")
		if len(categories) == 1 {
			bookmark.CategoryWvrIds = append(bookmark.CategoryWvrIds, categories[0].ID)
		} else {
			entity := categoryWvr.CategoryWvr{
				ID: bson.NewObjectId(),
				UserId: bookmark.UserId,
				Name: *wert.ClassName,
			}
			categoryWvr.Create(entity)
			bookmark.CategoryWvrIds = append(bookmark.CategoryWvrIds, entity.ID)
		}
	}

}
