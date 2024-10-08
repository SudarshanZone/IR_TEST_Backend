package models
<<<<<<< HEAD

type FnoPosition struct {
	Contract            string  `gorm:"column:FFO_CONTRACT"`
	Position            string  `gorm:"column:FFO_PSTN"`
	TotalQty            int     `gorm:"column:FFO_QTY"`
	AvgCostPrice        float64 `gorm:"column:FFO_AVG_PRC"`
	ExchangeCode        string  `gorm:"column:FCP_XCHNG_CD"`
	BuyQty              int     `gorm:"column:FCP_IBUY_QTY"`
	ClaimMatchAccount   string  `gorm:"column:FCP_CLM_MTCH_ACCNT"`
	ProductType         string  `gorm:"column:FCP_PRDCT_TYP"`
	IndexStock          string  `gorm:"column:FCP_INDSTK"`
	Underlying          string  `gorm:"column:FCP_UNDRLYNG"`
	ExpiryDate          string  `gorm:"column:FCP_EXPRY_DT"`
	ExerciseType        string  `gorm:"column:FCP_EXER_TYP"`
	OptionType          string  `gorm:"column:FCP_OPT_TYP"`
	StrikePrice         float64 `gorm:"column:FCP_STRK_PRC"`
	UCCCode             string  `gorm:"column:FCP_UCC_CD"`
}

=======
>>>>>>> master
type OrderDetails struct {
    ContractDescriptor string  `gorm:"column:ContractDescriptor"`
    VTCDate            string  `gorm:"column:VTCDate"`
    BuySell            string  `gorm:"column:BuySell"`
    Quantity           int32   `gorm:"column:Quantity"`
    Status             string  `gorm:"column:Status"`
    OrderPrice         float32 `gorm:"column:OrderPrice"`
    Open               string  `gorm:"column:Open"`
}
