package coinbasemanager

import (
	"math"

	"github.com/nonsense-project/nonsense/v2/domain/consensus/model"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/model/externalapi"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/constants"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/hashset"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/subnetworks"
	"github.com/nonsense-project/nonsense/v2/domain/consensus/utils/transactionhelper"
	"github.com/nonsense-project/nonsense/v2/infrastructure/db/database"
	"github.com/pkg/errors"
)

type coinbaseManager struct {
	subsidyGenesisReward                    uint64
	preDeflationaryPhaseBaseSubsidy         uint64
	coinbasePayloadScriptPublicKeyMaxLength uint8
	genesisHash                             *externalapi.DomainHash
	deflationaryPhaseDaaScore               uint64
	deflationaryPhaseBaseSubsidy            uint64

	databaseContext     model.DBReader
	dagTraversalManager model.DAGTraversalManager
	ghostdagDataStore   model.GHOSTDAGDataStore
	acceptanceDataStore model.AcceptanceDataStore
	daaBlocksStore      model.DAABlocksStore
	blockStore          model.BlockStore
	pruningStore        model.PruningStore
	blockHeaderStore    model.BlockHeaderStore
}

func (c *coinbaseManager) ExpectedCoinbaseTransaction(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash,
	coinbaseData *externalapi.DomainCoinbaseData) (expectedTransaction *externalapi.DomainTransaction, hasRedReward bool, err error) {

	ghostdagData, err := c.ghostdagDataStore.Get(c.databaseContext, stagingArea, blockHash, true)
	if !database.IsNotFoundError(err) && err != nil {
		return nil, false, err
	}

	// If there's ghostdag data with trusted data we prefer it because we need the original merge set non-pruned merge set.
	if database.IsNotFoundError(err) {
		ghostdagData, err = c.ghostdagDataStore.Get(c.databaseContext, stagingArea, blockHash, false)
		if err != nil {
			return nil, false, err
		}
	}

	acceptanceData, err := c.acceptanceDataStore.Get(c.databaseContext, stagingArea, blockHash)
	if err != nil {
		return nil, false, err
	}

	daaAddedBlocksSet, err := c.daaAddedBlocksSet(stagingArea, blockHash)
	if err != nil {
		return nil, false, err
	}

	txOuts := make([]*externalapi.DomainTransactionOutput, 0, len(ghostdagData.MergeSetBlues()))
	acceptanceDataMap := acceptanceDataFromArrayToMap(acceptanceData)
	for _, blue := range ghostdagData.MergeSetBlues() {
		txOut, hasReward, err := c.coinbaseOutputForBlueBlock(stagingArea, blue, acceptanceDataMap[*blue], daaAddedBlocksSet)
		if err != nil {
			return nil, false, err
		}

		if hasReward {
			txOuts = append(txOuts, txOut)
		}
	}

	txOut, hasRedReward, err := c.coinbaseOutputForRewardFromRedBlocks(
		stagingArea, ghostdagData, acceptanceData, daaAddedBlocksSet, coinbaseData)
	if err != nil {
		return nil, false, err
	}

	if hasRedReward {
		txOuts = append(txOuts, txOut)
	}

	subsidy, err := c.CalcBlockSubsidy(stagingArea, blockHash)
	if err != nil {
		return nil, false, err
	}

	payload, err := c.serializeCoinbasePayload(ghostdagData.BlueScore(), coinbaseData, subsidy)
	if err != nil {
		return nil, false, err
	}

	return &externalapi.DomainTransaction{
		Version:      constants.MaxTransactionVersion,
		Inputs:       []*externalapi.DomainTransactionInput{},
		Outputs:      txOuts,
		LockTime:     0,
		SubnetworkID: subnetworks.SubnetworkIDCoinbase,
		Gas:          0,
		Payload:      payload,
	}, hasRedReward, nil
}

func (c *coinbaseManager) daaAddedBlocksSet(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash) (
	hashset.HashSet, error) {

	daaAddedBlocks, err := c.daaBlocksStore.DAAAddedBlocks(c.databaseContext, stagingArea, blockHash)
	if err != nil {
		return nil, err
	}

	return hashset.NewFromSlice(daaAddedBlocks...), nil
}

// coinbaseOutputForBlueBlock calculates the output that should go into the coinbase transaction of blueBlock
// If blueBlock gets no fee - returns nil for txOut
func (c *coinbaseManager) coinbaseOutputForBlueBlock(stagingArea *model.StagingArea,
	blueBlock *externalapi.DomainHash, blockAcceptanceData *externalapi.BlockAcceptanceData,
	mergingBlockDAAAddedBlocksSet hashset.HashSet) (*externalapi.DomainTransactionOutput, bool, error) {

	blockReward, err := c.calcMergedBlockReward(stagingArea, blueBlock, blockAcceptanceData, mergingBlockDAAAddedBlocksSet)
	if err != nil {
		return nil, false, err
	}

	if blockReward == 0 {
		return nil, false, nil
	}

	// the ScriptPublicKey for the coinbase is parsed from the coinbase payload
	_, coinbaseData, _, err := c.ExtractCoinbaseDataBlueScoreAndSubsidy(blockAcceptanceData.TransactionAcceptanceData[0].Transaction)
	if err != nil {
		return nil, false, err
	}

	txOut := &externalapi.DomainTransactionOutput{
		Value:           blockReward,
		ScriptPublicKey: coinbaseData.ScriptPublicKey,
	}

	return txOut, true, nil
}

func (c *coinbaseManager) coinbaseOutputForRewardFromRedBlocks(stagingArea *model.StagingArea,
	ghostdagData *externalapi.BlockGHOSTDAGData, acceptanceData externalapi.AcceptanceData, daaAddedBlocksSet hashset.HashSet,
	coinbaseData *externalapi.DomainCoinbaseData) (*externalapi.DomainTransactionOutput, bool, error) {

	acceptanceDataMap := acceptanceDataFromArrayToMap(acceptanceData)
	totalReward := uint64(0)
	for _, red := range ghostdagData.MergeSetReds() {
		reward, err := c.calcMergedBlockReward(stagingArea, red, acceptanceDataMap[*red], daaAddedBlocksSet)
		if err != nil {
			return nil, false, err
		}

		totalReward += reward
	}

	if totalReward == 0 {
		return nil, false, nil
	}

	return &externalapi.DomainTransactionOutput{
		Value:           totalReward,
		ScriptPublicKey: coinbaseData.ScriptPublicKey,
	}, true, nil
}

func acceptanceDataFromArrayToMap(acceptanceData externalapi.AcceptanceData) map[externalapi.DomainHash]*externalapi.BlockAcceptanceData {
	acceptanceDataMap := make(map[externalapi.DomainHash]*externalapi.BlockAcceptanceData, len(acceptanceData))
	for _, blockAcceptanceData := range acceptanceData {
		acceptanceDataMap[*blockAcceptanceData.BlockHash] = blockAcceptanceData
	}
	return acceptanceDataMap
}

// CalcBlockSubsidy returns the subsidy amount a block at the provided blue score
// should have. This is mainly used for determining how much the coinbase for
// newly generated blocks awards as well as validating the coinbase for blocks
// has the expected value.
func (c *coinbaseManager) CalcBlockSubsidy(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash) (uint64, error) {
	if blockHash.Equal(c.genesisHash) {
		return c.subsidyGenesisReward, nil
	}
	blockDaaScore, err := c.daaBlocksStore.DAAScore(c.databaseContext, stagingArea, blockHash)
	if err != nil {
		return 0, err
	}
	if blockDaaScore < c.deflationaryPhaseDaaScore {
		return c.preDeflationaryPhaseBaseSubsidy, nil
	}

	blockSubsidy := c.calcDeflationaryPeriodBlockSubsidy(blockDaaScore)
	return blockSubsidy, nil
}

func (c *coinbaseManager) calcDeflationaryPeriodBlockSubsidy(blockDaaScore uint64) uint64 {
	// We define a year as 365.25 days and a month as 365.25 / 12 = 30.4375
	// secondsPerMonth = 30.4375 * 24 * 60 * 60
	const secondsPerMonth = 2629800
	// Note that this calculation implicitly assumes that block per second = 1 (by assuming daa score diff is in second units).
	monthsSinceDeflationaryPhaseStarted := (blockDaaScore - c.deflationaryPhaseDaaScore) / secondsPerMonth
	// Return the pre-calculated value from subsidy-per-month table
	return c.getDeflationaryPeriodBlockSubsidyFromTable(monthsSinceDeflationaryPhaseStarted)
}

/*
This table was pre-calculated by calling `calcDeflationaryPeriodBlockSubsidyFloatCalc` for all months until reaching 0 subsidy.
To regenerate this table, run `TestBuildSubsidyTable` in coinbasemanager_test.go (note the `deflationaryPhaseBaseSubsidy` therein)
*/
var subsidyByDeflationaryMonthTable = []uint64{
	1776465535, 1727346442, 1679585488, 1633145119, 1587988821, 1544081091, 1501387405, 1459874195, 1419508821, 1380259546, 1342095509, 1304986703, 1268903953, 1233818887, 1199703920, 1166532227, 1134277729, 1102915065, 1072419575, 1042767282, 1013934872, 985899675, 958639649, 932133359, 906359966,
	881299205, 856931371, 833237305, 810198378, 787796475, 766013982, 744833773, 724239194, 704214054, 684742606, 665809542, 647399976, 629499432, 612093836, 595169504, 578713127, 562711767, 547152844, 532024123, 517313710, 503010038, 489101861, 475578244, 462428554, 449642451,
	437209883, 425121074, 413366519, 401936977, 390823460, 380017231, 369509793, 359292884, 349358472, 339698746, 330306110, 321173179, 312292773, 303657910, 295261799, 287097840, 279159614, 271440879, 263935566, 256637774, 249541766, 242641961, 235932935, 229409414, 223066267,
	216898507, 210901285, 205069886, 199399724, 193886342, 188525404, 183312696, 178244118, 173315686, 168523525, 163863867, 159333047, 154927505, 150643775, 146478490, 142428374, 138490244, 134661003, 130937640, 127317227, 123796919, 120373946, 117045619, 113809319, 110662503,
	107602696, 104627493, 101734553, 98921603, 96186430, 93526885, 90940876, 88426370, 85981390, 83604013, 81292371, 79044645, 76859069, 74733923, 72667538, 70658288, 68704593, 66804918, 64957769, 63161693, 61415279, 59717152, 58065979, 56460461, 54899335,
	53381373, 51905384, 50470205, 49074709, 47717798, 46398406, 45115495, 43868056, 42655109, 41475699, 40328900, 39213810, 38129552, 37075274, 36050146, 35053364, 34084142, 33141718, 32225353, 31334326, 30467935, 29625499, 28806357, 28009864, 27235394,
	26482338, 25750104, 25038117, 24345815, 23672656, 23018109, 22381661, 21762810, 21161071, 20575969, 20007046, 19453853, 18915956, 18392932, 17884369, 17389868, 16909040, 16441507, 15986901, 15544864, 15115050, 14697121, 14290747, 13895609, 13511397,
	13137808, 12774549, 12421334, 12077885, 11743933, 11419215, 11103474, 10796464, 10497943, 10207676, 9925435, 9650998, 9384149, 9124678, 8872381, 8627061, 8388524, 8156582, 7931053, 7711760, 7498531, 7291197, 7089596, 6893570, 6702963,
	6517627, 6337415, 6162186, 5991802, 5826130, 5665038, 5508400, 5356093, 5207998, 5063997, 4923978, 4787831, 4655448, 4526725, 4401561, 4279859, 4161521, 4046455, 3934571, 3825781, 3719998, 3617141, 3517127, 3419879, 3325320,
	3233375, 3143972, 3057042, 2972515, 2890325, 2810408, 2732700, 2657141, 2583672, 2512234, 2442770, 2375228, 2309553, 2245694, 2183601, 2123225, 2064518, 2007434, 1951929, 1897958, 1845480, 1794452, 1744836, 1696591, 1649681,
	1604067, 1559715, 1516589, 1474655, 1433881, 1394235, 1355684, 1318200, 1281752, 1246311, 1211851, 1178343, 1145762, 1114082, 1083278, 1053325, 1024201, 995882, 968346, 941571, 915537, 890222, 865608, 841674, 818401,
	795773, 773770, 752375, 731572, 711344, 691675, 672551, 653955, 635873, 618291, 601195, 584572, 568409, 552692, 537411, 522551, 508103, 494054, 480393, 467110, 454195, 441636, 429425, 417551, 406006,
	394780, 383865, 373251, 362930, 352895, 343138, 333650, 324425, 315454, 306732, 298251, 290004, 281986, 274189, 266608, 259236, 252068, 245098, 238321, 231732, 225324, 219094, 213036, 207146, 201418,
	195849, 190434, 185168, 180048, 175070, 170229, 165523, 160946, 156496, 152169, 147961, 143870, 139892, 136024, 132263, 128606, 125050, 121592, 118230, 114961, 111782, 108692, 105686, 102764, 99923,
	97160, 94473, 91861, 89321, 86851, 84450, 82115, 79844, 77637, 75490, 73403, 71373, 69400, 67481, 65615, 63801, 62037, 60321, 58653, 57032, 55455, 53921, 52430, 50981, 49571,
	48200, 46868, 45572, 44312, 43087, 41895, 40737, 39610, 38515, 37450, 36415, 35408, 34429, 33477, 32551, 31651, 30776, 29925, 29098, 28293, 27511, 26750, 26010, 25291, 24592,
	23912, 23251, 22608, 21983, 21375, 20784, 20209, 19650, 19107, 18579, 18065, 17565, 17080, 16607, 16148, 15702, 15268, 14845, 14435, 14036, 13648, 13270, 12903, 12547, 12200,
	11862, 11534, 11215, 10905, 10604, 10311, 10025, 9748, 9479, 9217, 8962, 8714, 8473, 8239, 8011, 7789, 7574, 7365, 7161, 6963, 6770, 6583, 6401, 6224, 6052,
	5885, 5722, 5564, 5410, 5260, 5115, 4973, 4836, 4702, 4572, 4446, 4323, 4203, 4087, 3974, 3864, 3757, 3653, 3552, 3454, 3358, 3266, 3175, 3087, 3002,
	2919, 2838, 2760, 2684, 2609, 2537, 2467, 2399, 2332, 2268, 2205, 2144, 2085, 2027, 1971, 1917, 1864, 1812, 1762, 1713, 1666, 1620, 1575, 1531, 1489,
	1448, 1408, 1369, 1331, 1294, 1258, 1224, 1190, 1157, 1125, 1094, 1063, 1034, 1005, 978, 951, 924, 899, 874, 850, 826, 803, 781, 759, 738,
	718, 698, 679, 660, 642, 624, 607, 590, 574, 558, 542, 527, 513, 499, 485, 471, 458, 446, 433, 421, 410, 398, 387, 377, 366,
	356, 346, 337, 327, 318, 309, 301, 292, 284, 276, 269, 261, 254, 247, 240, 234, 227, 221, 215, 209, 203, 197, 192, 187, 181,
	176, 171, 167, 162, 158, 153, 149, 145, 141, 137, 133, 129, 126, 122, 119, 116, 112, 109, 106, 103, 100, 98, 95, 92, 90,
	87, 85, 82, 80, 78, 76, 74, 72, 70, 68, 66, 64, 62, 60, 59, 57, 56, 54, 52, 51, 50, 48, 47, 46, 44,
	43, 42, 41, 40, 38, 37, 36, 35, 34, 33, 32, 31, 31, 30, 29, 28, 27, 27, 26, 25, 24, 24, 23, 22, 22,
	21, 20, 20, 19, 19, 18, 18, 17, 17, 16, 16, 15, 15, 14, 14, 14, 13, 13, 13, 12, 12, 11, 11, 11, 11,
	10, 10, 10, 9, 9, 9, 9, 8, 8, 8, 8, 7, 7, 7, 7, 7, 6, 6, 6, 6, 6, 5, 5, 5, 5,
	5, 5, 5, 4, 4, 4, 4, 4, 4, 4, 4, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
	1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0,
}

func (c *coinbaseManager) getDeflationaryPeriodBlockSubsidyFromTable(month uint64) uint64 {
	if month >= uint64(len(subsidyByDeflationaryMonthTable)) {
		month = uint64(len(subsidyByDeflationaryMonthTable) - 1)
	}
	return subsidyByDeflationaryMonthTable[month]
}

func (c *coinbaseManager) calcDeflationaryPeriodBlockSubsidyFloatCalc(month uint64) uint64 {
	baseSubsidy := c.deflationaryPhaseBaseSubsidy
	subsidy := float64(baseSubsidy) / math.Pow(1.4, float64(month)/12)
	return uint64(subsidy)
}

func (c *coinbaseManager) calcMergedBlockReward(stagingArea *model.StagingArea, blockHash *externalapi.DomainHash,
	blockAcceptanceData *externalapi.BlockAcceptanceData, mergingBlockDAAAddedBlocksSet hashset.HashSet) (uint64, error) {

	if !blockHash.Equal(blockAcceptanceData.BlockHash) {
		return 0, errors.Errorf("blockAcceptanceData.BlockHash is expected to be %s but got %s",
			blockHash, blockAcceptanceData.BlockHash)
	}

	if !mergingBlockDAAAddedBlocksSet.Contains(blockHash) {
		return 0, nil
	}

	totalFees := uint64(0)
	for _, txAcceptanceData := range blockAcceptanceData.TransactionAcceptanceData {
		if txAcceptanceData.IsAccepted {
			totalFees += txAcceptanceData.Fee
		}
	}

	block, err := c.blockStore.Block(c.databaseContext, stagingArea, blockHash)
	if err != nil {
		return 0, err
	}

	_, _, subsidy, err := c.ExtractCoinbaseDataBlueScoreAndSubsidy(block.Transactions[transactionhelper.CoinbaseTransactionIndex])
	if err != nil {
		return 0, err
	}

	return subsidy + totalFees, nil
}

// New instantiates a new CoinbaseManager
func New(
	databaseContext model.DBReader,

	subsidyGenesisReward uint64,
	preDeflationaryPhaseBaseSubsidy uint64,
	coinbasePayloadScriptPublicKeyMaxLength uint8,
	genesisHash *externalapi.DomainHash,
	deflationaryPhaseDaaScore uint64,
	deflationaryPhaseBaseSubsidy uint64,

	dagTraversalManager model.DAGTraversalManager,
	ghostdagDataStore model.GHOSTDAGDataStore,
	acceptanceDataStore model.AcceptanceDataStore,
	daaBlocksStore model.DAABlocksStore,
	blockStore model.BlockStore,
	pruningStore model.PruningStore,
	blockHeaderStore model.BlockHeaderStore) model.CoinbaseManager {

	return &coinbaseManager{
		databaseContext: databaseContext,

		subsidyGenesisReward:                    subsidyGenesisReward,
		preDeflationaryPhaseBaseSubsidy:         preDeflationaryPhaseBaseSubsidy,
		coinbasePayloadScriptPublicKeyMaxLength: coinbasePayloadScriptPublicKeyMaxLength,
		genesisHash:                             genesisHash,
		deflationaryPhaseDaaScore:               deflationaryPhaseDaaScore,
		deflationaryPhaseBaseSubsidy:            deflationaryPhaseBaseSubsidy,

		dagTraversalManager: dagTraversalManager,
		ghostdagDataStore:   ghostdagDataStore,
		acceptanceDataStore: acceptanceDataStore,
		daaBlocksStore:      daaBlocksStore,
		blockStore:          blockStore,
		pruningStore:        pruningStore,
		blockHeaderStore:    blockHeaderStore,
	}
}
