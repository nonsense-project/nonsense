package coinbasemanager

import (
	"math"
	"strconv"
	"testing"

	"github.com/nonsense-project/nonsense/v2/domain/consensus/model/externalapi"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/constants"
	"github.com/nonsense-project/nonsense/v2/domain/dagconfig"
)

func TestCalcDeflationaryPeriodBlockSubsidy(t *testing.T) {
	const secondsPerMonth = 2629800
	const secondsPerHalving = secondsPerMonth * 12
	const deflationaryPhaseDaaScore = secondsPerMonth * 6
	const deflationaryPhaseBaseSubsidy = 1_776_465_535
	coinbaseManagerInterface := New(
		nil,
		0,
		0,
		0,
		&externalapi.DomainHash{},
		deflationaryPhaseDaaScore,
		deflationaryPhaseBaseSubsidy,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil)
	coinbaseManagerInstance := coinbaseManagerInterface.(*coinbaseManager)

	tests := []struct {
		name                 string
		blockDaaScore        uint64
		expectedBlockSubsidy uint64
	}{
		{
			name:                 "start of deflationary phase",
			blockDaaScore:        deflationaryPhaseDaaScore,
			expectedBlockSubsidy: deflationaryPhaseBaseSubsidy,
		},
		{
			name:                 "after 1 year",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / 1.4)),
		},
		{
			name:                 "after 2 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*2,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / math.Pow(1.4, 2))),
		},
		{
			name:                 "after 5 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*5,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / math.Pow(1.4, 5))),
		},
		{
			name:                 "after 32 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*32,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / math.Pow(1.4, 32))),
		},
		{
			name:                 "after 64 years",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*64,
			expectedBlockSubsidy: uint64(math.Trunc(deflationaryPhaseBaseSubsidy / math.Pow(1.4, 64))),
		},
		{
			name:                 "just before subsidy depleted",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*63,
			expectedBlockSubsidy: 1,
		},
		{
			name:                 "after subsidy depleted",
			blockDaaScore:        deflationaryPhaseDaaScore + secondsPerHalving*64,
			expectedBlockSubsidy: 0,
		},
	}

	for _, test := range tests {
		blockSubsidy := coinbaseManagerInstance.calcDeflationaryPeriodBlockSubsidy(test.blockDaaScore)
		if blockSubsidy != test.expectedBlockSubsidy {
			t.Errorf("TestCalcDeflationaryPeriodBlockSubsidy: test '%s' failed. Want: %d, got: %d",
				test.name, test.expectedBlockSubsidy, blockSubsidy)
		}
	}
}

func TestMainnetSubsidyScheduleStaysBelowMaximumSupply(t *testing.T) {
	const secondsPerMonth = uint64(2_629_800)

	params := &dagconfig.MainnetParams
	totalSubsidy := params.DeflationaryPhaseDaaScore * params.PreDeflationaryPhaseBaseSubsidy
	for _, monthlySubsidy := range subsidyByDeflationaryMonthTable {
		totalSubsidy += monthlySubsidy * secondsPerMonth
	}

	const expectedTotalSubsidy = uint64(199_999_999_867_608_000)
	if totalSubsidy != expectedTotalSubsidy {
		t.Fatalf("unexpected mainnet subsidy schedule total: got %d, expected %d", totalSubsidy, expectedTotalSubsidy)
	}
	if totalSubsidy > constants.MaxSompi {
		t.Fatalf("mainnet subsidy schedule exceeds maximum supply: got %d, maximum %d", totalSubsidy, constants.MaxSompi)
	}
}

func TestBuildSubsidyTable(t *testing.T) {
	deflationaryPhaseBaseSubsidy := dagconfig.MainnetParams.DeflationaryPhaseBaseSubsidy
	if deflationaryPhaseBaseSubsidy != 1_776_465_535 {
		t.Errorf("TestBuildSubsidyTable: table generation function was not updated to reflect "+
			"the new base subsidy %d. Please fix the constant above and replace subsidyByDeflationaryMonthTable "+
			"in coinbasemanager.go with the printed table", deflationaryPhaseBaseSubsidy)
	}
	coinbaseManagerInterface := New(
		nil,
		0,
		0,
		0,
		&externalapi.DomainHash{},
		0,
		deflationaryPhaseBaseSubsidy,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil,
		nil)
	coinbaseManagerInstance := coinbaseManagerInterface.(*coinbaseManager)

	var subsidyTable []uint64
	for M := uint64(0); ; M++ {
		subsidy := coinbaseManagerInstance.calcDeflationaryPeriodBlockSubsidyFloatCalc(M)
		subsidyTable = append(subsidyTable, subsidy)
		if subsidy == 0 {
			break
		}
	}

	tableStr := "\n{\t"
	for i := 0; i < len(subsidyTable); i++ {
		tableStr += strconv.FormatUint(subsidyTable[i], 10) + ", "
		if (i+1)%25 == 0 {
			tableStr += "\n\t"
		}
	}
	tableStr += "\n}"
	t.Logf("%s", tableStr)
}
