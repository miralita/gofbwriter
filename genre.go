package gofbwriter

import (
	"regexp"
	"strings"
)

//Genre - book genre from xsd
type Genre int

//go:generate stringer -type=Genre

const (
	//GenreAccounting genre
	GenreAccounting Genre = iota
	//GenreAdvAnimal genre
	GenreAdvAnimal
	//GenreAdvGeo genre
	GenreAdvGeo
	//GenreAdvHistory genre
	GenreAdvHistory
	//GenreAdvMaritime genre
	GenreAdvMaritime
	//GenreAdvWestern genre
	GenreAdvWestern
	//GenreAdventure genre
	GenreAdventure
	//GenreAntique genre
	GenreAntique
	//GenreAntiqueAnt genre
	GenreAntiqueAnt
	//GenreAntiqueEast genre
	GenreAntiqueEast
	//GenreAntiqueEuropean genre
	GenreAntiqueEuropean
	//GenreAntiqueMyths genre
	GenreAntiqueMyths
	//GenreAntiqueRussian genre
	GenreAntiqueRussian
	//GenreAphorismQuote genre
	GenreAphorismQuote
	//GenreArchitectureBook genre
	GenreArchitectureBook
	//GenreAutoRegulations genre
	GenreAutoRegulations
	//GenreBanking genre
	GenreBanking
	//GenreChildAdv genre
	GenreChildAdv
	//GenreChildDet genre
	GenreChildDet
	//GenreChildEducation genre
	GenreChildEducation
	//GenreChildProse genre
	GenreChildProse
	//GenreChildSf genre
	GenreChildSf
	//GenreChildTale genre
	GenreChildTale
	//GenreChildVerse genre
	GenreChildVerse
	//GenreChildren genre
	GenreChildren
	//GenreCinemaTheatre genre
	GenreCinemaTheatre
	//GenreCityFantasy genre
	GenreCityFantasy
	//GenreCompDb genre
	GenreCompDb
	//GenreCompHard genre
	GenreCompHard
	//GenreCompOsnet genre
	GenreCompOsnet
	//GenreCompProgramming genre
	GenreCompProgramming
	//GenreCompSoft genre
	GenreCompSoft
	//GenreCompWww genre
	GenreCompWww
	//GenreComputers genre
	GenreComputers
	//GenreDesign genre
	GenreDesign
	//GenreDetAction genre
	GenreDetAction
	//GenreDetClassic genre
	GenreDetClassic
	//GenreDetCrime genre
	GenreDetCrime
	//GenreDetEspionage genre
	GenreDetEspionage
	//GenreDetHard genre
	GenreDetHard
	//GenreDetHistory genre
	GenreDetHistory
	//GenreDetIrony genre
	GenreDetIrony
	//GenreDetPolice genre
	GenreDetPolice
	//GenreDetPolitical genre
	GenreDetPolitical
	//GenreDetective genre
	GenreDetective
	//GenreDragonFantasy genre
	GenreDragonFantasy
	//GenreDramaturgy genre
	GenreDramaturgy
	//GenreEconomics genre
	GenreEconomics
	//GenreEssays genre
	GenreEssays
	//GenreFantasyFight genre
	GenreFantasyFight
	//GenreForeignAction genre
	GenreForeignAction
	//GenreForeignAdventure genre
	GenreForeignAdventure
	//GenreForeignAntique genre
	GenreForeignAntique
	//GenreForeignBusiness genre
	GenreForeignBusiness
	//GenreForeignChildren genre
	GenreForeignChildren
	//GenreForeignComp genre
	GenreForeignComp
	//GenreForeignContemporary genre
	GenreForeignContemporary
	//GenreForeignContemporaryLit genre
	GenreForeignContemporaryLit
	//GenreForeignDesc genre
	GenreForeignDesc
	//GenreForeignDetective genre
	GenreForeignDetective
	//GenreForeignDramaturgy genre
	GenreForeignDramaturgy
	//GenreForeignEdu genre
	GenreForeignEdu
	//GenreForeignFantasy genre
	GenreForeignFantasy
	//GenreForeignHome genre
	GenreForeignHome
	//GenreForeignHumor genre
	GenreForeignHumor
	//GenreForeignLanguage genre
	GenreForeignLanguage
	//GenreForeignLove genre
	GenreForeignLove
	//GenreForeignNovel genre
	GenreForeignNovel
	//GenreForeignOther genre
	GenreForeignOther
	//GenreForeignPoetry genre
	GenreForeignPoetry
	//GenreForeignProse genre
	GenreForeignProse
	//GenreForeignPsychology genre
	GenreForeignPsychology
	//GenreForeignPublicism genre
	GenreForeignPublicism
	//GenreForeignReligion genre
	GenreForeignReligion
	//GenreForeignSf genre
	GenreForeignSf
	//GenreGeoGuides genre
	GenreGeoGuides
	//GenreGeographyBook genre
	GenreGeographyBook
	//GenreGlobalEconomy genre
	GenreGlobalEconomy
	//GenreHistoricalFantasy genre
	GenreHistoricalFantasy
	//GenreHome genre
	GenreHome
	//GenreHomeCooking genre
	GenreHomeCooking
	//GenreHomeCrafts genre
	GenreHomeCrafts
	//GenreHomeDiy genre
	GenreHomeDiy
	//GenreHomeEntertain genre
	GenreHomeEntertain
	//GenreHomeGarden genre
	GenreHomeGarden
	//GenreHomeHealth genre
	GenreHomeHealth
	//GenreHomePets genre
	GenreHomePets
	//GenreHomeSex genre
	GenreHomeSex
	//GenreHomeSport genre
	GenreHomeSport
	//GenreHumor genre
	GenreHumor
	//GenreHumorAnecdote genre
	GenreHumorAnecdote
	//GenreHumorFantasy genre
	GenreHumorFantasy
	//GenreHumorProse genre
	GenreHumorProse
	//GenreHumorVerse genre
	GenreHumorVerse
	//GenreIndustries genre
	GenreIndustries
	//GenreJobHunting genre
	GenreJobHunting
	//GenreLiterature18 genre
	GenreLiterature18
	//GenreLiterature19 genre
	GenreLiterature19
	//GenreLiterature20 genre
	GenreLiterature20
	//GenreLoveContemporary genre
	GenreLoveContemporary
	//GenreLoveDetective genre
	GenreLoveDetective
	//GenreLoveErotica genre
	GenreLoveErotica
	//GenreLoveFantasy genre
	GenreLoveFantasy
	//GenreLoveHistory genre
	GenreLoveHistory
	//GenreLoveSf genre
	GenreLoveSf
	//GenreLoveShort genre
	GenreLoveShort
	//GenreMagicianBook genre
	GenreMagicianBook
	//GenreManagement genre
	GenreManagement
	//GenreMarketing genre
	GenreMarketing
	//GenreMilitarySpecial genre
	GenreMilitarySpecial
	//GenreMusicDancing genre
	GenreMusicDancing
	//GenreNarrative genre
	GenreNarrative
	//GenreNewspapers genre
	GenreNewspapers
	//GenreNonfBiography genre
	GenreNonfBiography
	//GenreNonfCriticism genre
	GenreNonfCriticism
	//GenreNonfPublicism genre
	GenreNonfPublicism
	//GenreNonfiction genre
	GenreNonfiction
	//GenreOrgBehavior genre
	GenreOrgBehavior
	//GenrePaperWork genre
	GenrePaperWork
	//GenrePedagogyBook genre
	GenrePedagogyBook
	//GenrePeriodic genre
	GenrePeriodic
	//GenrePersonalFinance genre
	GenrePersonalFinance
	//GenrePoetry genre
	GenrePoetry
	//GenrePopadanec genre
	GenrePopadanec
	//GenrePopularBusiness genre
	GenrePopularBusiness
	//GenreProseClassic genre
	GenreProseClassic
	//GenreProseCounter genre
	GenreProseCounter
	//GenreProseHistory genre
	GenreProseHistory
	//GenreProseMilitary genre
	GenreProseMilitary
	//GenreProseRusClassic genre
	GenreProseRusClassic
	//GenreProseSuClassics genre
	GenreProseSuClassics
	//GenrePsyAlassic genre
	GenrePsyAlassic
	//GenrePsyChilds genre
	GenrePsyChilds
	//GenrePsyGeneric genre
	GenrePsyGeneric
	//GenrePsyPersonal genre
	GenrePsyPersonal
	//GenrePsySexAndFamily genre
	GenrePsySexAndFamily
	//GenrePsySocial genre
	GenrePsySocial
	//GenrePsyTheraphy genre
	GenrePsyTheraphy
	//GenreRealEstate genre
	GenreRealEstate
	//GenreRefDict genre
	GenreRefDict
	//GenreRefEncyc genre
	GenreRefEncyc
	//GenreRefGuide genre
	GenreRefGuide
	//GenreRefRef genre
	GenreRefRef
	//GenreReference genre
	GenreReference
	//GenreReligion genre
	GenreReligion
	//GenreReligionEsoterics genre
	GenreReligionEsoterics
	//GenreReligionRel genre
	GenreReligionRel
	//GenreReligionSelf genre
	GenreReligionSelf
	//GenreRussianContemporary genre
	GenreRussianContemporary
	//GenreRussianFantasy genre
	GenreRussianFantasy
	//GenreSciBiology genre
	GenreSciBiology
	//GenreSciChem genre
	GenreSciChem
	//GenreSciCulture genre
	GenreSciCulture
	//GenreSciHistory genre
	GenreSciHistory
	//GenreSciJuris genre
	GenreSciJuris
	//GenreSciLinguistic genre
	GenreSciLinguistic
	//GenreSciMath genre
	GenreSciMath
	//GenreSciMedicine genre
	GenreSciMedicine
	//GenreSciPhilosophy genre
	GenreSciPhilosophy
	//GenreSciPhys genre
	GenreSciPhys
	//GenreSciPolitics genre
	GenreSciPolitics
	//GenreSciReligion genre
	GenreSciReligion
	//GenreSciTech genre
	GenreSciTech
	//GenreScience genre
	GenreScience
	//GenreSf genre
	GenreSf
	//GenreSfAction genre
	GenreSfAction
	//GenreSfCyberpunk genre
	GenreSfCyberpunk
	//GenreSfDetective genre
	GenreSfDetective
	//GenreSfFantasy genre
	GenreSfFantasy
	//GenreSfHeroic genre
	GenreSfHeroic
	//GenreSfHistory genre
	GenreSfHistory
	//GenreSfHorror genre
	GenreSfHorror
	//GenreSfHumor genre
	GenreSfHumor
	//GenreSfSocial genre
	GenreSfSocial
	//GenreSfSpace genre
	GenreSfSpace
	//GenreShortStory genre
	GenreShortStory
	//GenreSketch genre
	GenreSketch
	//GenreSmallBusiness genre
	GenreSmallBusiness
	//GenreSociologyBook genre
	GenreSociologyBook
	//GenreStock genre
	GenreStock
	//GenreThriller genre
	GenreThriller
	//GenreUpbringingBook genre
	GenreUpbringingBook
	//GenreVampireBook genre
	GenreVampireBook
	//GenreVisualArts genre
	GenreVisualArts
	//GenreUnrecognised genre
	GenreUnrecognised
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")

func toSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}

func (s Genre) toString() string {
	return toSnakeCase(strings.TrimPrefix(s.String(), "Genre"))
}
