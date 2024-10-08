<<<<<<< HEAD
package repository

import (
	"fmt"
	"github.com/SudarshanZone/Open_Pos/internal/models"
	"gorm.io/gorm"
)

type FnoPositionRepository struct {
	Db *gorm.DB
}

func (repo *FnoPositionRepository) GetPositionsByClaimMatchAccount(claimMatchAccount string) ([]models.FnoPosition, error) {
	var positions []models.FnoPosition
	query := `
		SELECT
			CASE
				WHEN FCP_PRDCT_TYP = 'F' THEN
					'FUT-' || FCP_UNDRLYNG || '-' || TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY')
				WHEN FCP_PRDCT_TYP = 'O' THEN
					'OPT-' || FCP_UNDRLYNG || '-' || TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY') || '-' ||
					FCP_STRK_PRC || '-' ||
					CASE
						WHEN FCP_OPT_TYP = 'C' THEN 'CE'
						WHEN FCP_OPT_TYP = 'P' THEN 'PE'
						ELSE '*'
					END
				ELSE
					FCP_UNDRLYNG || ' ' || TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY')
			END AS "Contract",
			CASE
				WHEN FCP_OPNPSTN_FLW = 'B' THEN 'BUY'
				WHEN FCP_OPNPSTN_FLW = 'S' THEN 'SELL'
				WHEN FCP_OPNPSTN_FLW = 'N' THEN '*'
				ELSE FCP_OPNPSTN_FLW
			END AS "Position",
			ABS(FCP_OPNPSTN_QTY) AS "TotalQty",
			COALESCE(FCP_AVG_PRC, 0) AS "AvgCostPrice",
			COALESCE(FCP_XCHNG_CD, '') AS "ExchangeCode",
			COALESCE(FCP_IBUY_QTY, 0) AS "BuyQty",
			FCP_CLM_MTCH_ACCNT AS "ClaimMatchAccount",
			FCP_PRDCT_TYP AS "ProductType",
			FCP_INDSTK AS "IndexStock",
			FCP_UNDRLYNG AS "Underlying",
			TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY') AS "ExpiryDate",
			FCP_EXER_TYP AS "ExerciseType",
			FCP_OPT_TYP AS "OptionType",
			COALESCE(FCP_STRK_PRC, 0) AS "StrikePrice",
			FCP_UCC_CD AS "UccCode",
			FCP_OPNPSTN_FLW AS "OpenPstnFlow"
		FROM
			FCP_FO_SPN_CNTRCT_PSTN
		WHERE
			(FCP_OPNPSTN_QTY != 0 OR FCP_IBUY_QTY != 0 OR FCP_ISELL_QTY != 0)
			AND FCP_CLM_MTCH_ACCNT = ?;
	`
	err := repo.Db.Raw(query, claimMatchAccount).Scan(&positions).Error
	if err != nil {
		return nil, fmt.Errorf("error fetching positions: %w", err)
	}
	return positions, nil
=======

package repository

import (
    "fmt"
    "github.com/SudarshanZone/Open_Pos/internal/models"
    "gorm.io/gorm"
)

type FnoPositionRepository struct {
    Db *gorm.DB
}

func (repo *FnoPositionRepository) GetPositionsByClaimMatchAccount(claimMatchAccount string) ([]models.FnoPosition, error) {
    var positions []models.FnoPosition
    query := `
        SELECT
            CASE
                WHEN FCP_PRDCT_TYP = 'F' THEN
                    'FUT-' || FCP_UNDRLYNG || '-' || TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY')
                WHEN FCP_PRDCT_TYP = 'O' THEN
                    'OPT-' || FCP_UNDRLYNG || '-' || TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY') || '-' ||
                    FCP_STRK_PRC || '-' ||
                    CASE
                        WHEN FCP_OPT_TYP = 'C' THEN 'CE'
                        WHEN FCP_OPT_TYP = 'P' THEN 'PE'
                        ELSE '*'
                    END
                ELSE
                    FCP_UNDRLYNG || ' ' || TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY')
            END AS "Contract",
            CASE
                WHEN FCP_OPNPSTN_FLW = 'B' THEN 'BUY'
                WHEN FCP_OPNPSTN_FLW = 'S' THEN 'SELL'
                WHEN FCP_OPNPSTN_FLW = 'N' THEN '*'
                ELSE FCP_OPNPSTN_FLW
            END AS "Position",
            ABS(FCP_OPNPSTN_QTY) AS "TotalQty",
            COALESCE(FCP_AVG_PRC, 0) AS "AvgCostPrice",
            COALESCE(FCP_XCHNG_CD, '') AS "ExchangeCode",
            COALESCE(FCP_IBUY_QTY, 0) AS "BuyQty",
            FCP_CLM_MTCH_ACCNT AS "ClaimMatchAccount",
            FCP_PRDCT_TYP AS "ProductType",
            FCP_INDSTK AS "IndexStock",
            FCP_UNDRLYNG AS "Underlying",
            TO_CHAR(FCP_EXPRY_DT, 'DD-Mon-YYYY') AS "ExpiryDate",
            FCP_EXER_TYP AS "ExerciseType",
            FCP_OPT_TYP AS "OptionType",
            COALESCE(FCP_STRK_PRC, 0) AS "StrikePrice",
            FCP_UCC_CD AS "UccCode",
            FCP_OPNPSTN_FLW AS "OpenPstnFlow"
        FROM
            FCP_FO_SPN_CNTRCT_PSTN
        WHERE
            (FCP_OPNPSTN_QTY != 0 OR FCP_IBUY_QTY != 0 OR FCP_ISELL_QTY != 0)
            AND FCP_CLM_MTCH_ACCNT = ?;
    `
    err := repo.Db.Raw(query, claimMatchAccount).Scan(&positions).Error
    if err != nil {
        return nil, fmt.Errorf("error fetching positions: %w", err)
    }
    return positions, nil
>>>>>>> master
}
